package model

type QuestaoCorrigida struct {
	Questao
	RespostaUsuario int    `json:"resposta_usuario"`
	Feedback        string `json:"feedback"`
}
