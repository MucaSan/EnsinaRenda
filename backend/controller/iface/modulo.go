package ifaceController

import (
	"context"
	"ensina-renda/domain/model"

	"github.com/google/uuid"
)

type ModuloController interface {
	CompletarModulo(ctx context.Context, usuarioModulo *model.UsuarioModulo) error
	GetUsuarioModulo(ctx context.Context, idAula int, idUsuario uuid.UUID) (*model.UsuarioModulo, error)
}
