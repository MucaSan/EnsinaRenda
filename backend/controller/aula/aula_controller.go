package controller

import (
	"context"
	ifaceController "ensina-renda/controller/iface"
	"ensina-renda/domain/model"
	"ensina-renda/repository/iface"
	"fmt"

	"github.com/google/uuid"
)

type AulaController struct {
	aulaRepository iface.AulaRepository
}

func NewAulaController(aulaRepository iface.AulaRepository) ifaceController.AulaController {
	return &AulaController{
		aulaRepository: aulaRepository,
	}
}

func (uc *AulaController) CompletarAula(ctx context.Context, usuarioAula *model.UsuarioAula) error {
	if err := uc.aulaRepository.CompletarAula(ctx, usuarioAula.IDAula, usuarioAula.IDUsuario); err != nil {
		return fmt.Errorf("erro ao processar usuario_aula para completar aula")
	}

	return nil
}

func (uc *AulaController) GetUsuarioAula(ctx context.Context, idAula int, idUsuario uuid.UUID) (*model.UsuarioAula, error) {
	usuarioAula, err := uc.aulaRepository.GetUsuarioAula(ctx, idAula, idUsuario)
	if err != nil {
		return nil, fmt.Errorf("erro ao processar usuario_aula para usuario %v", err)
	}

	return usuarioAula, nil
}
