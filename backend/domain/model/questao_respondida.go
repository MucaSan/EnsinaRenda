package model

type QuestaoRespondida struct {
	Questao
	RespostaUsuario int `json:"resposta_usuario"`
}
