package corrigir_prova

import (
	"context"
	"encoding/json"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/domain/model"
	"errors"
	"fmt"
)

type CorrigirProvaConverter struct {
	base any
}

func NewCorrigirProvaConverter(base any) *CorrigirProvaConverter {
	return &CorrigirProvaConverter{
		base: base,
	}
}

func (c *CorrigirProvaConverter) ToDomain(ctx context.Context) (*model.ProvaRespondida, error) {
	corrigirProvaRequest, ok := c.base.(*pb.CorrigirProvaRequest)
	if !ok {
		return nil, errors.New("nao foi possivel converter base para corrigir_prova_request")
	}

	var provaRespondida *model.ProvaRespondida

	err := json.Unmarshal([]byte(corrigirProvaRequest.ProvaRespondida), &provaRespondida)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(`JSON nao esta no formato correto, deve seguir o formato: {"titulo_prova":"Teste", "questoes_respondidas": [{"numQuestao":1, "enunciado": "Teste", "opcoes":["opcao1", "opcao2", "opcao3", "opcao4", "resposta_usuario":0]}]}, erro: %s`, err.Error()))
	}

	return provaRespondida, nil
}
