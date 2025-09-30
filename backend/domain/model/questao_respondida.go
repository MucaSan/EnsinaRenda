package model

type QuestaoRespondida struct {
	NumQuestao      int      `json:"numQuestao"`
	Enunciado       string   `json:"enunciado"`
	Opcoes          []string `json:"opcoes"`
	RespostaUsuario int      `json:"resposta_usuario"`
}
