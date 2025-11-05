package model

import (
	"github.com/google/uuid"
)

type UsuarioModuloAula struct {
	IDModulo        int       `gorm:"primaryKey;column:id_modulo"`
	IDUsuario       uuid.UUID `gorm:"type:uuid;primaryKey;column:id_usuario"`
	IDAula          int       `gorm:"primaryKey;column:id_aula"`
	AulaConcluida   bool      `gorm:"default:false;column:aula_concluida"`
	ModuloConcluida bool      `gorm:"default:false;column:modulo_concluido"`
}
