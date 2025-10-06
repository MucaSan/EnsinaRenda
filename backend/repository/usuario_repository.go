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
		Where("senha = ?", senha).First(&usuario).Error
	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepository) GetUsuarioPeloIdDoContexto(ctx context.Context) (*model.Usuario, error) {
	var usuario *model.Usuario
	err := database.GetDB(ctx).Table("usuario").
		Where("id = ?", auth.GetUserUuidPeloContext(ctx)).First(&usuario).Error
	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepository) GetUsuarioPeloEmail(ctx context.Context, email string) (*model.Usuario, error) {
	var usuario *model.Usuario

	err := database.GetDB(ctx).Table("usuario").
		Where("email = ?", email).First(&usuario).Error
	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepository) AtualizarUsuario(ctx context.Context, usuario *model.Usuario) error {
	err := database.GetDB(ctx).Table("usuario").Save(usuario).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UsuarioRepository) GetUsuarioPeloId(ctx context.Context, id string) (*model.Usuario, error) {
	var usuario *model.Usuario

	err := database.GetDB(ctx).Table("usuario").
		Where("id = ?", id).First(&usuario).Error
	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepository) ProvisionarUsuarioModulos(ctx context.Context, usuarioModulos []*model.UsuarioModulo) error {
	err := database.GetDB(ctx).Table("usuario_modulo").Create(usuarioModulos).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UsuarioRepository) ProvisionarUsuarioAulas(ctx context.Context, usuarioAulas []*model.UsuarioAula) error {
	err := database.GetDB(ctx).Table("usuario_aula").Create(usuarioAulas).Error
	if err != nil {
		return err
	}

	return nil
}
