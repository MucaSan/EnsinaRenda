package repository

import (
	"context"
	"ensina-renda/config/database"
	"ensina-renda/domain/model"

	"github.com/google/uuid"
)

type ModuloRepository struct {
}

func NewModuloRepository() *ModuloRepository {
	return &ModuloRepository{}
}

func (r *ModuloRepository) CompletarModulo(ctx context.Context, idModulo int, idUsuario uuid.UUID) error {
	err := database.GetDB(ctx).Table("usuario_modulo").
		Where("idModulo = ? ", idModulo).
		Where("id_usuario = ?", idUsuario).
		Updates(map[string]interface{}{"concluido": true}).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ModuloRepository) GetUsuarioModulo(ctx context.Context, idModulo int, idUsuario uuid.UUID) (*model.UsuarioModulo, error) {
	var usuarioModulo *model.UsuarioModulo

	err := database.GetDB(ctx).
		Table("usuario_modulo").
		Where("id_modulo  = ?", idModulo).
		Where("id_usuario = ?", idUsuario).
		Find(&usuarioModulo).Error

	if err != nil {
		return nil, err
	}

	return usuarioModulo, nil
}
