package iface

import (
	"context"
	"ensina-renda/domain/model"
)

type ProvaRepository interface {
	GetProvaBase(ctx context.Context, id_modulo int) (*model.ProvaBase, error)
	SalvarProva(ctx context.Context, provaBase *model.ProvaUsuario) error
	GetProvaUsuario(ctx context.Context, id_modulo int) (*model.ProvaUsuario, error)
}
