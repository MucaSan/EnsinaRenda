package model

import "github.com/google/uuid"

type CorrecaoProva struct {
	IDModulo        int       `json:"id_modulo" gorm:"primaryKey;column:id_modulo"`
	IDUsuario       uuid.UUID `json:"id_usuario" gorm:"primaryKey;column:id_usuario"`
	ConteudoAnalise string    `json:"conteudo_analise" gorm:"column:conteudo_analise"`
}
