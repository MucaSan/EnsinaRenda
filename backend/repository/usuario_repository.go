package repository

import (
	"context"
	"ensina-renda/config/auth"
	"ensina-renda/config/database"
	"ensina-renda/domain/model"

	"github.com/google/uuid"
)

type UsuarioRepository struct {
}

func NewUsuarioRepository() *UsuarioRepository {
	return &UsuarioRepository{}
}

func (r *UsuarioRepository) CriarUsuario(ctx context.Context, usuario *model.Usuario) error {
	err := database.GetDB(ctx).Create(usuario).Error
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

	if usuario.Id != uuid.Nil {
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

	if usuario.Id != uuid.Nil {
		return true, nil
	}

	return false, nil
}

func (r *UsuarioRepository) VerificarEmailUsuario(ctx context.Context, hash_email string) (bool, error) {
	var usuario *model.Usuario
	err := database.GetDB(ctx).Table("usuario").Where("email = ?", hash_email).Find(&usuario).Error
	if err != nil {
		return false, err
	}

	if usuario.Id != uuid.Nil {
		return true, nil
	}

	return false, nil
}

func (r *UsuarioRepository) VerificarSenhaUsuario(ctx context.Context, hash_senha string) (bool, error) {
	var usuario *model.Usuario
	err := database.GetDB(ctx).Table("usuario").Where("senha = ?", hash_senha).Find(&usuario).Error
	if err != nil {
		return false, err
	}

	if usuario.Id != uuid.Nil {
		return true, nil
	}

	return false, nil
}

func (r *UsuarioRepository) GetUsuario(ctx context.Context, email, senha string) (*model.Usuario, error) {
	var usuario *model.Usuario

	err := database.GetDB(ctx).Table("usuario").
		Where("email  = ?", email).
		Where("senha = ?", senha).Find(&usuario).Error
	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepository) GetUsuarioPeloId(ctx context.Context) (*model.Usuario, error) {
	var usuario *model.Usuario
	err := database.GetDB(ctx).Table("usuario").
		Where("id = ?", auth.GetUserUuidPeloContext(ctx)).Find(&usuario).Error
	if err != nil {
		return nil, err
	}

	return usuario, nil
}
