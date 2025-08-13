package iface

import (
	"context"
	"ensina-renda/domain/model"

	"github.com/google/uuid"
)

type UsuarioRepository interface {
	CriarUsuario(ctx context.Context, usuario *model.Usuario) error
	VerificarEmail(ctx context.Context, email string) (bool, error)
	VerificarUsuarioCadastrado(ctx context.Context, id uuid.UUID) (bool, error)
	VerificarEmailUsuario(ctx context.Context, hash_email string) (bool, error)
	VerificarSenhaUsuario(ctx context.Context, hash_senha string) (bool, error)
	GetUsuario(ctx context.Context, id uuid.UUID) (*model.Usuario, error)
}
