package http

import (
	"bytes"
	"context"
	"encoding/json"
	"ensina-renda/domain/model"
	domain "ensina-renda/domain/service"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

type AgenteProfessor struct {
}

func NewAgenteProfessor() domain.AgenteProfessor {
	return &AgenteProfessor{}
}

// Estrutura para a requisição que será enviada à API do ChatGPT
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Estrutura para as mensagens dentro da requisição
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Estrutura para a resposta que será recebida da API
type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

// Nova estrutura para o conteúdo JSON dentro da string
type ContentResponse struct {
	IdModulo int             `json:"id_modulo"`
	Conteudo json.RawMessage `json:"conteudo_prova"`
}

var chaveApi = os.Getenv("OPENAI_API_KEY")
var url = "https://api.openai.com/v1/chat/completions"

const (
	promptAgenteGerarProva = "Você é um agente especializado em educação financeira em renda fixa. Sua tarefa principal é customizar uma prova fornecida em formato JSON. Mantenha o formato JSON idêntico ao original, mas obrigatóriamente modifique o conteúdo das questões para torná-las mais desafiadoras. Transforme perguntas genéricas em questões mais complexas, e mude questões muito técnicas para situações-problema envolventes. O título da prova deve ser mantido ou levemente adaptado para refletir a nova dificuldade. A quantidade de questões não podem se alterar. Sua resposta deve ser APENAS o JSON da prova modificada, sem nenhum texto adicional."

	promptAgenteAnalisarProva = `Você é um professor de finanças. Seu objetivo é analisar as respostas de um aluno para uma prova em JSON e gerar um feedback detalhado. A prova base já contém as respostas corretas. Compare as respostas do aluno com as corretas: Se a resposta estiver errada, explique o erro e sugira um ponto de melhoria. Mantenha um tom de apoio. Se a resposta estiver correta, dê uma mensagem de incentivo e adicione uma curiosidade breve sobre o tema. A sua resposta deve conter somente o texto do feedback, e nada mais.`
)

func (a *AgenteProfessor) GerarProva(ctx context.Context, provaBase string) (string, error) {
	return realizarRequisicao(promptAgenteGerarProva, provaBase)

}

func (a *AgenteProfessor) CorrigirProva(ctx context.Context, questaoRespondida *model.ProvaRespondida) ([]string, error) {
	semaforo := make(chan struct{}, 5)
	var grupoEspera sync.WaitGroup
	var mutex sync.Mutex
	var canalErros = make(chan error, len(questaoRespondida.QuestoesRespondidas))

	var feedbacks []string

	for _, questao := range questaoRespondida.QuestoesRespondidas {
		grupoEspera.Add(1)
		semaforo <- struct{}{}
		go func() {
			defer grupoEspera.Done()
			defer func() { <-semaforo }()

			feedback, err := realizarRequisicaoPorQuestao(promptAgenteAnalisarProva, questao)
			if err != nil {
				canalErros <- err
			}

			mutex.Lock()
			feedbacks = append(feedbacks, feedback)
			mutex.Unlock()
		}()
	}

	grupoEspera.Wait()
	close(semaforo)
	close(canalErros)

	for err := range canalErros {
		return nil, err
	}

	return feedbacks, nil
}

func realizarRequisicao(promptAgente, promptAnalise string) (string, error) {
	requestBody := ChatRequest{
		Model: "gpt-3.5-turbo-16k",
		Messages: []Message{
			{
				Role:    "system",
				Content: promptAgente,
			},
			{
				Role:    "user",
				Content: promptAnalise,
			},
		},
	}

	// Converte a estrutura da requisição em um JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return "", err
	}

	// Cria a requisição HTTP POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Erro ao criar a requisição:", err)
		return "", err
	}

	// Adiciona os headers necessários
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+chaveApi)

	// Envia a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Erro ao enviar a requisição: %w", err)
	}
	defer resp.Body.Close()

	// Lê a resposta, mesmo que seja um erro HTTP
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Erro ao ler a resposta: %w", err)
	}

	// Adicionado: Log da resposta RAW para depuração
	fmt.Println("Status da resposta:", resp.Status)
	fmt.Println("Corpo da resposta (RAW):", string(body))

	// Trata o erro de forma mais robusta, incluindo o corpo da resposta de erro
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erro na requisição. Status: %s. Corpo da resposta: %s", resp.Status, string(body))
	}

	// Decodifica a resposta JSON
	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		return "", fmt.Errorf("Erro ao decodificar JSON: %w. Corpo da resposta: %s", err, string(body))
	}

	if len(chatResponse.Choices) > 0 {
		jsonContentString := chatResponse.Choices[0].Message.Content

		var contentResponse ContentResponse
		err = json.Unmarshal([]byte(jsonContentString), &contentResponse)
		if err != nil {
			return "", fmt.Errorf("Erro ao decodificar a string de conteúdo JSON: %w", err)
		}

		conteudoGerado, err := contentResponse.Conteudo.MarshalJSON()
		if err != nil {
			return "", fmt.Errorf("Erro ao decodificar a string de conteúdo JSON: %w", err)
		}
		return string(conteudoGerado), nil
	}

	return "", nil

}

func realizarRequisicaoPorQuestao(promptAgente string, questaoRespondida model.QuestaoRespondida) (string, error) {
	requestBody := ChatRequest{
		Model: "gpt-3.5-turbo-16k",
		Messages: []Message{
			{
				Role:    "system",
				Content: promptAgente,
			},
			{
				Role: "user",
				Content: fmt.Sprintf(
					"Enunciado: %s Opcoes: %s Resposta do Usuario: %d",
					questaoRespondida.Enunciado,
					strings.Join(questaoRespondida.Opcoes, " , "),
					questaoRespondida.RespostaUsuario,
				),
			},
		},
	}

	// Converte a estrutura da requisição em um JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return "", err
	}

	// Cria a requisição HTTP POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Erro ao criar a requisição:", err)
		return "", err
	}

	// Adiciona os headers necessários
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+chaveApi)

	// Envia a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Erro ao enviar a requisição: %w", err)
	}
	defer resp.Body.Close()

	// Lê a resposta, mesmo que seja um erro HTTP
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Erro ao ler a resposta: %w", err)
	}

	// Adicionado: Log da resposta RAW para depuração
	fmt.Println("Status da resposta:", resp.Status)
	fmt.Println("Corpo da resposta (RAW):", string(body))

	// Trata o erro de forma mais robusta, incluindo o corpo da resposta de erro
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erro na requisição. Status: %s. Corpo da resposta: %s", resp.Status, string(body))
	}

	// Decodifica a resposta JSON
	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		return "", fmt.Errorf("Erro ao decodificar JSON: %w. Corpo da resposta: %s", err, string(body))
	}

	if len(chatResponse.Choices) > 0 {
		jsonContentString := chatResponse.Choices[0].Message.Content

		return string(jsonContentString), nil
	}

	return "", nil
}
