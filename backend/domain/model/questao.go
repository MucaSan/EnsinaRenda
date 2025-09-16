package model

type Questao struct {
	NumQuestao      int      `json:"numQuestao"`
	Enunciado       string   `json:"enunciado"`
	Opcoes          []string `json:"opcoes"`
	RespostaCorreta int      `json:"resposta_correta"`
}
