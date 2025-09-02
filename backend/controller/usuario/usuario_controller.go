package controller

import (
	"context"
	service "ensina-renda/adapter/grpc/service/auth"
	"ensina-renda/domain/model"
	domain "ensina-renda/domain/service"
	"ensina-renda/repository/iface"
	"fmt"
)

type UsuarioController struct {
	usuarioRepository iface.UsuarioRepository
	jwtService        domain.JwtServiceInterface
}

func NewUsuarioController(usuarioRepository iface.UsuarioRepository) *UsuarioController {
	return &UsuarioController{
		usuarioRepository: usuarioRepository,
		jwtService:        service.NewJwtService(),
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

func (uc *UsuarioController) GetUsuarioPeloId(ctx context.Context) (*model.Usuario, error) {
	usuario, err := uc.usuarioRepository.GetUsuarioPeloId(ctx)
	if err != nil {
		return nil, fmt.Errorf("houve um erro ao tentar procurar o usuario: %s", err.Error())
	}

	return usuario, nil
}
