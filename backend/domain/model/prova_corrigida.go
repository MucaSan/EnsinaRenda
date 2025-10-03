package model

import "encoding/json"

type ProvaCorrigida struct {
	TituloProva        string             `json:"titulo_prova"`
	QuestoesCorrigidas []QuestaoCorrigida `json:"questoes_corrigidas"`
}

func (pc *ProvaCorrigida) FormatarParaJSONString() (string, error) {
	jsonBytes, err := json.Marshal(pc)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
