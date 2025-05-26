package repository

import (
	"context"
	"ensina-renda/config/database"
	"ensina-renda/domain/model"

	"github.com/google/uuid"
)

type UsuarioRepository struct {
}

func (r *UsuarioRepository) CriarUsuario(ctx context.Context, usuario *model.Usuario) error {
	err := database.GetDB(ctx).Save(usuario).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UsuarioRepository) VerificarEmail(ctx context.Context, email string) (bool, error) {
	var usuario *model.Usuario

	err := database.GetDB(ctx).Table("usuario").Where("email = ?", email).Find(&usuario).Error
	if err != nil {
		return false, err
	}

	if usuario != nil {
		return true, nil
	}

	return false, nil
}

func (r *UsuarioRepository) VerificarUsuarioCadastrado(ctx context.Context, id uuid.UUID) (bool, error) {
	var usuario *model.Usuario

	err := database.GetDB(ctx).Table("usuario").Where("id = ?", id).Find(&usuario).Error
	if err != nil {
		return false, err
	}

	if usuario != nil {
		return true, nil
	}

	return false, nil
}
