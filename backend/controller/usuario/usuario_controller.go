package controller

import (
	"context"
	"ensina-renda/domain/model"
	"ensina-renda/repository/iface"
	"fmt"
)

type UsuarioController struct {
	usuarioRepository iface.UsuarioRepository
}

func NewUsuarioController(usuarioRepository iface.UsuarioRepository) *UsuarioController {
	return &UsuarioController{
		usuarioRepository: usuarioRepository,
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
