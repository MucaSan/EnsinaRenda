package verificar_aluno

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/domain/model"
	"errors"
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
	verificarAlunoRequest, ok := uc.base.(*pb.VerificarAlunoRequest)
	if !ok {
		return nil, errors.New("nao foi possivel converter base para verificar_aluno_request")
	}

	modelUsuario := &model.Usuario{
		Email: verificarAlunoRequest.Email,
		Senha: verificarAlunoRequest.Senha,
	}

	return modelUsuario, nil

}
