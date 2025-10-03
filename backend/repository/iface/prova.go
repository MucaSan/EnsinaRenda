package iface

import (
	"context"
	"ensina-renda/domain/model"
)

type ProvaRepository interface {
	GetProvaBase(ctx context.Context, id_modulo string) (*model.ProvaBase, error)
	SalvarProva(ctx context.Context, provaBase *model.ProvaUsuario) error
	GetProvaUsuario(ctx context.Context, id_modulo string) (*model.ProvaUsuario, error)
	SalvarCorrecaoProva(ctx context.Context, correcaoProva *model.CorrecaoProva) error
	GetCorrecaoProva(ctx context.Context, idModulo string) (*model.CorrecaoProva, error)
}
