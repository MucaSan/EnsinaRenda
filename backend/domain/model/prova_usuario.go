package model

import (
	"time"
)

type ProvaUsuario struct {
	IDModulo       int        `gorm:"column:id_modulo;primaryKey"`
	IDUsuario      string     `gorm:"column:id_usuario;primaryKey"`
	ConteudoGerado string     `gorm:"column:conteudo_gerado;type:jsonb"`
	GeradoEm       time.Time  `gorm:"column:gerado_em"`
	AtualizadoEm   *time.Time `gorm:"column:atualizado_em"`
}

func (u ProvaUsuario) TableName() string {
	return "provas_usuario"
}
