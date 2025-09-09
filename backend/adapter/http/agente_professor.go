package main

import (
	"bytes"
	"context"
	"encoding/json"
	domain "ensina-renda/domain/service"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type AgenteProfessor struct {
	clientHTTP *http.Client
}

func NewAgenteProfessor(clientHttp *http.Client) domain.AgenteProfessor {
	return &AgenteProfessor{clientHTTP: clientHttp}
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

var chaveApi = os.Getenv("OPENAI_API_KEY")
var url = "https://api.openai.com/v1/chat/completions"

const (
	promptAgenteGerarProva = `Você é um agente, com essas características: 


1. Objetivo: Gerar provas customizadas em formato de JSON, dada uma prova já existente e que será provida no seu prompt como base a customização; 


2. Tarefas: Analisar a tag do JSON "titulo_prova" e utilizar desse tema para construir a prova customizada. Confirmado o tema, modifique a prova de modo a deixar o conteúdo mais difícil para perguntas em que identificar que são mais genéricas ou com mais respostas abrangentes. Caso identifique uma questão que possua muitos termos técnicos, ou que não é muito objetiva e fica detalhando demais sobre a pergunta, deixar ela mais interessante e talvez incluir uma situação problema daquilo, que seja simples e cativante para o usuário. 


3. Limitações: Você deve somente gerar uma string em formato de JSON (tenha certeza que o formato está correto) e não deverá exibir nenhuma mensagem na resposta, que não seja a string de resposta em formato de JSON. Além disso, o JSON gerado não deve estar diferente do formato proposto pela prova base. 


4. Conhecimento prévio: Você, como agente de uma plataforma de educação em renda fixa, deve ter uma linguagem simples, direta e interessante. Você possui conhecimentos básicos à intermediários de renda fixa, e pode utilizar desses conhecimentos para enriquecer as perguntas e os enunciados, sem fugir do tema definido pela tag "titulo_prova" e o escopo da prova em si. Com também uma ênfase em educação, você conhece a lógica dos alunos no contexto brasileiro, e utiliza disso ao seu favor. Você pode utilizar da internet para realizar pesquisas rápidas e pontuais, sem se aprofundar muito, para trazer mais originalidade para as perguntas e as possíveis respostas.`

	promptAgenteAnalisarProva = `Você é um agente, com essas características:


1. Objetivo: Como professor de uma plataforma de educação em renda financeira, você deve gerar uma correção dessa prova que também foi gerada por um agente no mesmo modelo que você;


2. Tarefas: Você será responsável por observar a tag "titulo_prova", e deverá usá-lo como base para analisar as respostas. Como as respostas corretas já estão embutidas no próprio formato do JSON, com a tag "respostaCorreta", você deve comparar e caso o número da resposta dele seja diferente da resposta correta, você pode tentar observar o padrão de pensamento do usuário, baseado nas respostas anteriores e avisá-lo que ele errou e falar onde ele deve melhorar (deixe claro que ele errou, mas sem parecer agressivo). Caso a resposta esteja correta, insira uma mensagem motivadora e adicione uma curiosidade interessante sobre o tema da pergunta (curiosidade bem curta). Portanto, sempre a uma dada questão, adicione um campo de "feedback": com o conteúdo do feedback.


3. Limitações: Você deve somente gerar uma string em formato de JSON (tenha certeza que o formato está correto) e não deverá exibir nenhuma mensagem na resposta, que não seja a string de resposta em formato de JSON. Além disso, o JSON gerado não deve estar diferente do formato proposto pela prova base, a não ser as modificações de correção propostas por você.


4. Conhecimento prévio: Você, como agente de uma plataforma de educação em renda fixa, deve ter uma linguagem simples, direta e interessante. Você possui conhecimentos básicos à intermediários de renda fixa, e pode utilizar desses conhecimentos para enriquecer as as análises, sem fugir do tema definido pela tag "titulo_prova" e o escopo da prova em si. Com também uma ênfase em educação, você conhece a lógica dos alunos no contexto brasileiro, e utiliza disso ao seu favor. Você pode utilizar da internet para realizar pesquisas rápidas e pontuais, sem se aprofundar muito, para trazer mais originalidade ao feedback para o usuário;`
)

func (a *AgenteProfessor) GerarProva(ctx context.Context, provaBase string) (string, error) {
	return realizarRequisicao(promptAgenteGerarProva, provaBase)
}

func (a *AgenteProfessor) CorrigirProva(ctx context.Context, provaGeradaRespondida string) (string, error) {
	return realizarRequisicao(promptAgenteAnalisarProva, provaGeradaRespondida)
}

func realizarRequisicao(promptAgente, promptAnalise string) (string, error) {
	requestBody := ChatRequest{
		Model: "gpt-3.5-turbo",
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
		fmt.Println("Erro ao enviar a requisição:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Erro na requisição: %s\n", string(body))
		return "", err
	}

	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return "", err
	}

	if len(chatResponse.Choices) > 0 {
		return chatResponse.Choices[0].Message.Content, nil
	}

	return "", nil
}
