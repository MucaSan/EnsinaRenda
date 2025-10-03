package ifaceController

import (
	"context"
	"ensina-renda/domain/model"
)

type ProvaController interface {
	GetProvaBase(ctx context.Context, idModulo string) (*model.ProvaBase, error)
	GerarProva(ctx context.Context, provaBase *model.ProvaBase) error
	GetProvaUsuario(ctx context.Context, idModulo string) (*model.ProvaUsuario, error)
	CorrigirProva(ctx context.Context, idModulo string, provaRespondida *model.ProvaRespondida) error
}
