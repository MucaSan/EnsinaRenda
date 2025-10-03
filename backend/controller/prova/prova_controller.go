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
	"strconv"
	"time"

	"github.com/google/uuid"
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

func (pc *ProvaController) CorrigirProva(ctx context.Context, idModulo string, provaRespondida *model.ProvaRespondida) error {
	provaUsuario, err := pc.provaRepository.GetProvaUsuario(ctx, idModulo)
	if err != nil {
		return fmt.Errorf("erro ao pegar prova do usuario do modulo fornecido: %v", err)
	}

	provaGerada, err := provaUsuario.FormatarConteudoParaProva()
	if err != nil {
		return fmt.Errorf("erro ao formatar prova do usuario do conteudo: %v", err)
	}

	err = adicionarRespostasCorretas(provaGerada, provaRespondida)
	if err != nil {
		return err
	}

	provaCorrigida, err := pc.agenteProfessor.CorrigirProva(ctx, provaRespondida)
	if err != nil {
		return fmt.Errorf("erro ao corrigir prova respondida: %v", err)
	}

	correcaoProva, err := gerarCorrecaoProva(ctx, provaCorrigida, idModulo)
	if err != nil {
		return err
	}

	err = pc.provaRepository.SalvarCorrecaoProva(ctx, correcaoProva)
	if err != nil {
		return fmt.Errorf("erro ao salvar prova corrigida: %v", err)
	}

	return nil
}

func adicionarRespostasCorretas(provaGerada *model.ProvaGerada, provaRespondida *model.ProvaRespondida) error {
	if provaGerada == nil {
		return fmt.Errorf("prova gerada esta vazia, gere outra prova")
	}

	if provaRespondida == nil {
		return fmt.Errorf("prova respondida esta vazia")
	}

	quantidadeQuestoesGeradas := len(provaGerada.Questoes)
	quantidadeQuestoesRespondidas := len(provaRespondida.QuestoesRespondidas)

	if quantidadeQuestoesGeradas != quantidadeQuestoesRespondidas {
		return fmt.Errorf("provas estao com tamanhos diferentes, houve algum erro no processamento das provas")
	}

	for i, questaoGerada := range provaGerada.Questoes {
		provaRespondida.QuestoesRespondidas[i].RespostaCorreta = questaoGerada.RespostaCorreta
	}

	return nil
}

func gerarCorrecaoProva(
	ctx context.Context,
	provaCorrigida *model.ProvaCorrigida,
	idModulo string,
) (*model.CorrecaoProva, error) {
	conteudoCorrecaoProva, err := provaCorrigida.FormatarParaJSONString()
	if err != nil {
		return nil, fmt.Errorf("erro ao formatar para JSON: %v", err)
	}

	usuarioUuid, err := uuid.Parse(auth.GetUserUuidPeloContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar uuid do usuario: %v", err)
	}

	numeroIdModulo, err := strconv.Atoi(idModulo)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar numero id modulo: %v", err)
	}

	return &model.CorrecaoProva{
		IDUsuario:       usuarioUuid,
		IDModulo:        numeroIdModulo,
		ConteudoAnalise: conteudoCorrecaoProva,
	}, nil
}
