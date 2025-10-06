package ifaceController

import (
	"context"
	"ensina-renda/domain/model"

	"github.com/google/uuid"
)

type AulaController interface {
	CompletarAula(ctx context.Context, usuarioAula *model.UsuarioAula) error
	GetUsuarioAula(ctx context.Context, idAula int, idUsuario uuid.UUID) (*model.UsuarioAula, error)
	ListarUsuarioAulaModulo(ctx context.Context, idUsuario uuid.UUID) (map[int][]*model.UsuarioModuloAula, error)
	ListarAulas(ctx context.Context) ([]*model.Aula, error)
}
