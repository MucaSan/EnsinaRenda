package domain

import (
	"context"
	"ensina-renda/domain/model"
)

type JwtServiceInterface interface {
	GerarJWT(ctx context.Context, usuario *model.Usuario) (string, error)
	DecodificarUUID(ctx context.Context, tokenString string) (string, error)
}
