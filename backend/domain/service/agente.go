package domain

import (
	"context"
)

type AgenteProfessor interface {
	GerarProva(ctx context.Context, provaBase string) (string, error)
	CorrigirProva(ctx context.Context, provaGerada string) (string, error)
}
