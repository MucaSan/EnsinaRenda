package ifaceController

import (
	"context"
	"ensina-renda/domain/model"
)

type ProvaController interface {
	GetProvaBase(ctx context.Context, idModulo int) (*model.ProvaBase, error)
	GerarProva(ctx context.Context, provaBase *model.ProvaBase) error
	GetProvaUsuario(ctx context.Context, idModulo int) (*model.ProvaUsuario, error)
}
