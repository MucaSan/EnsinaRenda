package controller

import (
	"context"
	"encoding/json"
	"ensina-renda/config/auth"
	ifaceController "ensina-renda/controller/iface"
	"ensina-renda/domain/model"
	domain "ensina-renda/domain/service"
	"ensina-renda/repository/iface"
	"fmt"
	"time"
)

type ProvaController struct {
	provaRepository iface.ProvaRepository
	agenteProfessor domain.AgenteProfessor
}

func NewProvaController(provaRepository iface.ProvaRepository, agente domain.AgenteProfessor) ifaceController.ProvaController {
	return &ProvaController{
		provaRepository: provaRepository,
		agenteProfessor: agente,
	}
}

func (pc *ProvaController) GetProvaBase(ctx context.Context, idModulo string) (*model.ProvaBase, error) {
	provaBase, err := pc.provaRepository.GetProvaBase(ctx, idModulo)

	if err != nil {
		return nil, fmt.Errorf("erro ao processar usuario_modulo para completar modulo: %v", err)
	}

	return provaBase, nil
}

func (pc *ProvaController) GerarProva(ctx context.Context, provaBase *model.ProvaBase) error {
	conteudoProvaBase, err := provaBase.FormatarParaJSONString()
	if err != nil {
		return fmt.Errorf("erro ao formatar prova para JSON: %v", err)
	}

	objetoJsonComEscape, err := pc.agenteProfessor.GerarProva(ctx, conteudoProvaBase)
	if err != nil {
		return fmt.Errorf("erro ao gerar prova com agente professor: %v", err)
	}

	var objetoJsonSemEscape string

	err = json.Unmarshal([]byte(objetoJsonComEscape), &objetoJsonSemEscape)
	if err != nil {
		return fmt.Errorf("erro ao remover escape do json: %v", err)
	}

	conteudoProvaGeradaEmBytes := []byte(objetoJsonSemEscape)

	var provaGerada *model.ProvaGerada

	err = json.Unmarshal(conteudoProvaGeradaEmBytes, &provaGerada)
	if err != nil {
		return fmt.Errorf("erro ao parsear prova gerada para estrutura JSON: %v", err)
	}

	conteudoProvaString, err := provaGerada.FormatarParaJSONString()
	if err != nil {
		return fmt.Errorf("erro ao formatar prova gerada para JSON: %v", err)
	}

	tempoAtual := time.Now()
	provaUsuario := &model.ProvaUsuario{
		IDModulo:       provaBase.IdModulo,
		IDUsuario:      auth.GetUserUuidPeloContext(ctx),
		ConteudoGerado: conteudoProvaString,
		GeradoEm:       tempoAtual,
		AtualizadoEm:   &tempoAtual,
	}

	err = pc.provaRepository.SalvarProva(ctx, provaUsuario)
	if err != nil {
		return fmt.Errorf("erro ao salvar prova: %v", err)
	}

	return nil
}

func (pc *ProvaController) GetProvaUsuario(ctx context.Context, idModulo string) (*model.ProvaUsuario, error) {
	provaUsuario, err := pc.provaRepository.GetProvaUsuario(ctx, idModulo)
	if err != nil {
		return nil, err
	}

	return provaUsuario, nil
}
