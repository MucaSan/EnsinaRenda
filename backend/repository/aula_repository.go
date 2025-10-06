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

func (r *AulaRepository) ListarUsuarioModuloAulas(ctx context.Context, id_usuario uuid.UUID) ([]*model.UsuarioModuloAula, error) {
	var moduloAulas []*model.UsuarioModuloAula

	err := database.GetDB(ctx).
		Table("usuario_aula").
		Select(`
					modulo_aula.id_modulo AS id_modulo, 
					usuario_aula.id_usuario AS id_usuario,
					usuario_aula.id_aula AS id_aula, 
					usuario_aula.concluido AS concluido
					`,
		).
		Joins("INNER JOIN modulo_aula ON usuario_aula.id_aula = modulo_aula.id_aula").
		Where("usuario_aula.id_usuario = ?", id_usuario).
		Find(&moduloAulas).Error
	if err != nil {
		return nil, err
	}

	return moduloAulas, nil
}

func (r *AulaRepository) ListarAulas(ctx context.Context) ([]*model.Aula, error) {
	var aulas []*model.Aula

	err := database.GetDB(ctx).Table("aula").Find(&aulas).Error
	if err != nil {
		return nil, err
	}

	return aulas, nil
}
