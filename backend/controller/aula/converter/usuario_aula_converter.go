package converter

import (
	"context"
	"ensina-renda/domain/model"
)

type UsuarioAulaConverterInterface interface {
	ToDomain(ctx context.Context) (*model.UsuarioAula, error)
}
