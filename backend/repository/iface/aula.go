package iface

import (
	"context"
	"ensina-renda/domain/model"

	"github.com/google/uuid"
)

type AulaRepository interface {
	CompletarAula(ctx context.Context, id_aula int, id_usuario uuid.UUID) error
	GetUsuarioAula(ctx context.Context, id_aula int, id_usuario uuid.UUID) (*model.UsuarioAula, error)
	ListarUsuarioModuloAulas(ctx context.Context, id_usuario uuid.UUID) ([]*model.UsuarioModuloAula, error)
}
