package model

import (
	"encoding/json"
)

type ProvaBase struct {
	IdModulo      int    `json:"id_modulo"`
	ConteudoProva string `json:"conteudo_prova"`
}

func (pb *ProvaBase) TableName() string {
	return "provas_base"
}

func (pb *ProvaBase) FormatarParaJSONString() (string, error) {
	jsonBytes, err := json.Marshal(pb)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
