package model

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UsuarioAula struct {
	IDUsuario uuid.UUID `gorm:"type:uuid;primaryKey;column:id_usuario"`
	IDAula    int       `gorm:"primaryKey;column:id_aula"`
	Concluido bool      `gorm:"default:false;column:concluido"`
}

func (ua *UsuarioAula) IsValid() error {
	validate := validator.New()

	return validate.Struct(ua)
}

// TableName sobrescreve o nome da tabela padrão do GORM
func (UsuarioAula) TableName() string {
	return "usuario_aula"
}
