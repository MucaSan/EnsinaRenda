package iface

import (
	"context"
	"ensina-renda/domain/model"

	"github.com/google/uuid"
)

type ModuloRepository interface {
	GetUsuarioModulo(ctx context.Context, idModulo int, idUsuario uuid.UUID) (*model.UsuarioModulo, error)
	CompletarModulo(ctx context.Context, idModulo int, idUsuario uuid.UUID) error
	ListarModulos(ctx context.Context) ([]*model.Modulo, error)
}
