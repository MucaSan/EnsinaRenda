package controller

import (
	"context"
	jwtService "ensina-renda/adapter/grpc/service/auth"
	servicosUsuario "ensina-renda/adapter/grpc/service/email"

	"ensina-renda/domain/model"
	domain "ensina-renda/domain/service"
	"ensina-renda/repository/iface"
	"fmt"
	"time"
)

type UsuarioController struct {
	usuarioRepository iface.UsuarioRepository
	jwtService        domain.JwtServiceInterface
	emailService      domain.EmailService
	hashService       domain.HashService
}

func NewUsuarioController(usuarioRepository iface.UsuarioRepository) *UsuarioController {
	return &UsuarioController{
		usuarioRepository: usuarioRepository,
		jwtService:        jwtService.NewJwtService(),
		emailService:      servicosUsuario.NewEmailService(),
		hashService:       servicosUsuario.NewHashService(),
	}
}

func (uc *UsuarioController) CadastrarUsuario(ctx context.Context, usuario *model.Usuario) error {
	email := usuario.Email

	temUsuario, err := uc.usuarioRepository.VerificarEmail(ctx, email)
	if err != nil {
		return err
	}

	if temUsuario {
		return fmt.Errorf("ja existe um usuario com o email %s no sistema", email)
	}

	if err = uc.usuarioRepository.CriarUsuario(ctx, usuario); err != nil {
		return err
	}

	usuarioCadastrado, err := uc.usuarioRepository.VerificarUsuarioCadastrado(ctx, usuario.Id)
	if err != nil {
		return err
	}

	if !usuarioCadastrado {
		return fmt.Errorf("o usuario nao foi cadastrado com sucesso")
	}

	return nil
}

func (uc *UsuarioController) VerificarCredenciaisUsuario(ctx context.Context, usuario *model.Usuario) (bool, error) {
	emailCerto, err := uc.usuarioRepository.VerificarEmailUsuario(ctx, usuario.Email)
	if err != nil {
		return false, err
	}

	senhaCerta, err := uc.usuarioRepository.VerificarSenhaUsuario(ctx, usuario.Senha)
	if err != nil {
		return false, err
	}

	if !emailCerto {
		return false, nil
	}

	if !senhaCerta {
		return false, nil
	}

	return true, nil
}

func (uc *UsuarioController) GetUsuario(ctx context.Context, usuario *model.Usuario) (*model.Usuario, error) {
	usuario, err := uc.usuarioRepository.GetUsuario(ctx, usuario.Email, usuario.Senha)
	if err != nil {
		return nil, fmt.Errorf("houve um erro ao tentar procurar o usuario: %s", err.Error())
	}

	return usuario, nil
}

func (uc *UsuarioController) RealizarLogin(ctx context.Context, usuario *model.Usuario) (string, error) {
	token, err := uc.jwtService.GerarJWT(ctx, usuario)
	if err != nil {
		return "", fmt.Errorf("houve um erro ao tentar gerar o jwt da sessao: %s", err.Error())
	}

	return token, nil
}

func (uc *UsuarioController) GetUsuarioPeloId(ctx context.Context, id string) (*model.Usuario, error) {
	usuario, err := uc.usuarioRepository.GetUsuarioPeloId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("houve um erro ao tentar buscar o usuario pelo id: %s", err.Error())
	}

	return usuario, nil
}

func (uc *UsuarioController) GetUsuarioPeloIdDoContexto(ctx context.Context) (*model.Usuario, error) {
	usuario, err := uc.usuarioRepository.GetUsuarioPeloIdDoContexto(ctx)
	if err != nil {
		return nil, fmt.Errorf("houve um erro ao tentar procurar o usuario: %s", err.Error())
	}

	return usuario, nil
}

func (uc *UsuarioController) GetUsuarioPeloEmail(ctx context.Context, email string) (*model.Usuario, error) {
	usuario, err := uc.usuarioRepository.GetUsuarioPeloEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("houve um erro ao tentar procurar o usuario pelo email: %s", err.Error())
	}

	return usuario, nil
}

func (uc *UsuarioController) AtualizarSenha(ctx context.Context, usuario *model.Usuario, senha string) error {
	usuario.Senha = senha
	horarioAtual := time.Now()
	usuario.AtualizadoEm = &horarioAtual

	err := uc.usuarioRepository.AtualizarUsuario(ctx, usuario)
	if err != nil {
		return fmt.Errorf("houve um erro ao tentar atualizar a senha do usuario: %s", err.Error())
	}

	return nil
}

func (uc *UsuarioController) ProvisionarUsuarioModulos(
	ctx context.Context,
	usuario *model.Usuario,
	modulos []*model.Modulo,
) error {
	usuarioModulos := make([]*model.UsuarioModulo, 0)

	for _, modulo := range modulos {
		usuarioModulos = append(usuarioModulos, &model.UsuarioModulo{
			IDUsuario: usuario.Id,
			IDModulo:  modulo.ID,
			Concluido: false,
		})
	}

	err := uc.usuarioRepository.ProvisionarUsuarioModulos(ctx, usuarioModulos)
	if err != nil {
		return fmt.Errorf("houve um erro ao tentar provisionar usuario modulos: %s", err.Error())
	}

	return nil
}

func (uc *UsuarioController) ProvisionarUsuarioAulas(
	ctx context.Context,
	usuario *model.Usuario,
	aulas []*model.Aula,
) error {
	usuarioAulas := make([]*model.UsuarioAula, 0)

	for _, aula := range aulas {
		usuarioAulas = append(usuarioAulas, &model.UsuarioAula{
			IDUsuario: usuario.Id,
			IDAula:    aula.ID,
			Concluido: false,
		})
	}

	err := uc.usuarioRepository.ProvisionarUsuarioAulas(ctx, usuarioAulas)
	if err != nil {
		return fmt.Errorf("houve um erro ao tentar provisionar usuario aulas: %s", err.Error())
	}

	return nil
}

func (uc *UsuarioController) CriptografarEmail(ctx context.Context, email string) string {
	return uc.hashService.GerarHashSHA256(email)
}

func (uc *UsuarioController) EnviarEmail(ctx context.Context, email, token string) error {
	err := uc.emailService.EnviarEmail(email, token)
	if err != nil {
		return fmt.Errorf("houve um erro ao tentar enviar email: %s", err.Error())
	}

	return nil
}

func (uc *UsuarioController) GerarToken(ctx context.Context, usuario *model.Usuario) (string, error) {
	token, err := uc.jwtService.GerarJWT(ctx, usuario)
	if err != nil {
		return "", fmt.Errorf(
			"houve um erro ao tentar gerar o jwt de resetar a senha: %s",
			err.Error(),
		)
	}

	return token, nil
}

func (uc *UsuarioController) BuscarUsuarioPeloJWT(ctx context.Context, token string) (*model.Usuario, error) {
	idUsuario, err := uc.jwtService.DecodificarUUID(ctx, token)
	if err != nil {
		return nil, fmt.Errorf(
			"nao foi possivel decodificar token JWT: %s",
			err.Error(),
		)
	}

	usuario, err := uc.usuarioRepository.GetUsuarioPeloId(ctx, idUsuario)
	if err != nil {
		return nil, fmt.Errorf(
			"nao foi possivel buscar usuario por por meio do ID do token: %s",
			err.Error(),
		)
	}

	return usuario, nil
}
