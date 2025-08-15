package controller

import (
	"context"
	ifaceController "ensina-renda/controller/iface"
	"ensina-renda/domain/model"
	"ensina-renda/repository/iface"
	"fmt"

	"github.com/google/uuid"
)

type ModuloController struct {
	moduloRepository iface.ModuloRepository
}

func NewModuloController(moduloRepository iface.ModuloRepository) ifaceController.ModuloController {
	return &ModuloController{
		moduloRepository: moduloRepository,
	}
}

func (uc *ModuloController) CompletarModulo(ctx context.Context, usuarioModulo *model.UsuarioModulo) error {
	if err := uc.moduloRepository.CompletarModulo(ctx, usuarioModulo.IDModulo, usuarioModulo.IDUsuario); err != nil {
		return fmt.Errorf("erro ao processar usuario_modulo para completar modulo: %v", err)
	}

	return nil
}

func (uc *ModuloController) GetUsuarioModulo(ctx context.Context, idAula int, idUsuario uuid.UUID) (*model.UsuarioModulo, error) {
	usuarioModulo, err := uc.moduloRepository.GetUsuarioModulo(ctx, idAula, idUsuario)
	if err != nil {
		return nil, fmt.Errorf("erro ao processar usuario_modulo para usuario: %v", err)
	}

	if usuarioModulo == nil || usuarioModulo.IDUsuario == uuid.Nil {
		return nil, fmt.Errorf("nao existe registro para o usuario com o ID fornecido")
	}

	return usuarioModulo, nil
}
