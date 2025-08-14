package converter

import (
	"context"
	"ensina-renda/domain/model"
)

type UsuarioModuloConverterInterface interface {
	ToDomain(ctx context.Context) (*model.UsuarioModulo, error)
}
