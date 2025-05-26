package model

import (
	"errors"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Usuario struct {
	Id             uuid.UUID  `validate:"required" gorm:"column:id"`
	Nome           string     `validate:"required,min=3,max=255" gorm:"column: nome"`
	Email          string     `validate:"required,email,max=64" gorm:"column: email"`
	Senha          string     `validate:"required,min=4,max=64" gorm:"column:senha"`
	DataNascimento string     `validate:"required,datetime=2006-01-02" gorm:"column:data_nascimento"`
	CriadoEm       *time.Time `validate:"required" gorm:"column:criado_em"`
	AtualizadoEm   *time.Time `gorm:"column:atualizado_em"`
	DeletadoEm     *time.Time `gorm:"column:deletado_em"`
}

func (u *Usuario) IsValid() error {
	validate := validator.New()

	if u.AtualizadoEm.Before(*u.CriadoEm) {
		return errors.New("o campo de atualizacao nao pode ser antes do campo de criacao")
	}

	if u.DeletadoEm.Before(*u.CriadoEm) {
		return errors.New("o campo de delecao nao pode ser antes do campo de delecao")
	}

	if u.DeletadoEm.Before(*u.AtualizadoEm) {
		return errors.New("o campo de delecao nao pode ser antes do campo de atualizacao")
	}

	return validate.Struct(u)
}

func (u *Usuario) TableName() string {
	return "usuario"
}
