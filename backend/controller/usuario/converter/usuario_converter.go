package converter

import (
	"context"
	"ensina-renda/domain/model"
)

type UsuarioConverterInterface interface {
	ToDomain(ctx context.Context, base any) (*model.Usuario, error)
}
