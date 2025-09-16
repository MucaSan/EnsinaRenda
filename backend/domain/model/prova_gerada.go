package model

import "encoding/json"

type ProvaGerada struct {
	TituloProva string    `json:"titulo_prova"`
	Questoes    []Questao `json:"questoes"`
}

func (cp *ProvaGerada) FormatarParaJSONString() (string, error) {
	jsonBytes, err := json.Marshal(cp)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
