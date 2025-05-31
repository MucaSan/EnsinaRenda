package cadastrar_aluno

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/domain/model"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type UsuarioConverter struct {
	base any
}

func NewUsuarioConverter(base any) *UsuarioConverter {
	return &UsuarioConverter{
		base: base,
	}
}

func (uc *UsuarioConverter) ToDomain(ctx context.Context) (*model.Usuario, error) {
	cadastrarAlunoRequest, ok := uc.base.(*pb.CadastrarAlunoRequest)
	if !ok {
		return nil, errors.New("nao foi possivel converter base para cadastrar_aluno_request")
	}

	err := validarCadastrarAlunoRequest(cadastrarAlunoRequest)
	if err != nil {
		return nil, err
	}

	dataNascimento, _ := time.Parse("02/01/2006", cadastrarAlunoRequest.DataNascimento)
	modelUsuario := &model.Usuario{
		Id:             uuid.New(),
		Nome:           cadastrarAlunoRequest.Nome,
		Email:          cadastrarAlunoRequest.Email,
		Senha:          cadastrarAlunoRequest.Senha,
		DataNascimento: dataNascimento,
		CriadoEm:       time.Now(),
	}

	if err = modelUsuario.IsValid(); err != nil {
		return nil, err
	}

	return modelUsuario, nil

}

func validarCadastrarAlunoRequest(cadastrarAlunoRequest *pb.CadastrarAlunoRequest) error {
	relacaoNomeValorCampo := map[string]string{
		"nome":            cadastrarAlunoRequest.Nome,
		"email":           cadastrarAlunoRequest.Email,
		"senha":           cadastrarAlunoRequest.Senha,
		"data_nascimento": cadastrarAlunoRequest.DataNascimento,
	}

	if err := validarCamposObrigatorios(relacaoNomeValorCampo); err != nil {
		return err
	}

	if !IsValidSHA256(relacaoNomeValorCampo["email"]) {
		return errors.New("o campo de email nao e um hash SHA-256")
	}

	if !IsValidSHA256(relacaoNomeValorCampo["senha"]) {
		return errors.New("o campo de senha nao e um hash SHA-256")
	}

	tamanhoNome := len(relacaoNomeValorCampo["nome"])

	if tamanhoNome > 255 || tamanhoNome < 3 {
		return errors.New("o campo de nome nao pode ser maior que 255 caracteres ou menor que 3 caracteres")
	}

	if err := validarDataNascimento(relacaoNomeValorCampo["data_nascimento"]); err != nil {
		return err
	}

	return nil
}

func validarCamposObrigatorios(relacaoNomeValorCampo map[string]string) error {
	for nomeCampo, valorCampo := range relacaoNomeValorCampo {
		if valorCampo == "" {
			return fmt.Errorf("o campo %s e obrigatorio, preencha-o", nomeCampo)
		}
	}

	return nil
}

func IsValidSHA256(hash string) bool {
	// SHA-256 regex: 64 caracteres hexadecimais
	sha256Regex := regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
	return sha256Regex.MatchString(hash)
}

func validarDataNascimento(data_nascimento string) error {
	// Define o layout para DD/MM/YYYY
	layout := "02/01/2006" // Referencia do Go: 02=DD, 01=MM, 2006=YYYY

	// Transforma a string para o formato time.Time, para verificacao
	_, err := time.Parse(layout, data_nascimento)
	if err != nil {
		return errors.New("o campo data_nascimento esta em um formato invalido, deve seguir o padrao brasileiro DD/MM/YYYY (ex: 02/01/2006)")
	}

	return nil
}
