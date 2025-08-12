package repository

import (
	"context"
	"ensina-renda/config/database"
	"ensina-renda/domain/model"

	"github.com/google/uuid"
)

type AulaRepository struct {
}

func NewAulaRepository() *AulaRepository {
	return &AulaRepository{}
}

func (r *AulaRepository) CompletarAula(ctx context.Context, id_aula int, id_usuario uuid.UUID) error {
	err := database.GetDB(ctx).Table("usuario_aula").
		Where("id_aula = ? ", id_aula).
		Where("id_usuario = ?", id_usuario).
		Updates(map[string]interface{}{"concluido": true}).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (r *AulaRepository) GetUsuarioAula(ctx context.Context, id_aula int, id_usuario uuid.UUID) (*model.UsuarioAula, error) {
	var usuarioAula *model.UsuarioAula

	err := database.GetDB(ctx).
		Table("usuario_aula").
		Where("id_aula  = ?", id_aula).
		Where("id_usuario = ?", id_usuario).
		Find(&usuarioAula).Error

	if err != nil {
		return nil, err
	}

	return usuarioAula, nil
}
