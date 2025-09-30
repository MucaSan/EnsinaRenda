package domain

import (
	"context"
	"ensina-renda/domain/model"
)

type AgenteProfessor interface {
	GerarProva(ctx context.Context, provaBase string) (string, error)
	CorrigirProva(ctx context.Context, questaoRespondida *model.ProvaRespondida) ([]string, error)
}
