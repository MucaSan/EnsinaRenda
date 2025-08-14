package model

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UsuarioModulo struct {
	IDUsuario uuid.UUID `gorm:"type:uuid;primaryKey;column:id_usuario"`
	IDModulo  int       `gorm:"primaryKey;column:id_modulo"`
	Concluido bool      `gorm:"default:false;column:concluido"`
}

func (ua *UsuarioModulo) IsValid() error {
	validate := validator.New()

	return validate.Struct(ua)
}

// TableName sobrescreve o nome da tabela padrão do GORM
func (UsuarioModulo) TableName() string {
	return "usuario_modulo"
}
