package controller

import (
	"context"
	"ensina-renda/domain/model"
	"ensina-renda/repository/iface"
)

type UsuarioController struct {
	usuarioRepository iface.UsuarioRepository
}

func NewUsuarioController(usuarioRepository iface.UsuarioRepository) *UsuarioController {
	return &UsuarioController{
		usuarioRepository: usuarioRepository,
	}
}

func (uc *UsuarioController) CadastrarUsuario(ctx context.Context) (*model.Usuario, error) {
	return nil, nil
}
