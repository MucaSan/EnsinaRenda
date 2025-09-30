package model

import "encoding/json"

type ProvaRespondida struct {
	TituloProva         string              `json:"titulo_prova"`
	QuestoesRespondidas []QuestaoRespondida `json:"questoes_respondidas"`
}

func (cp *ProvaRespondida) FormatarParaJSONString() (string, error) {
	jsonBytes, err := json.Marshal(cp)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
