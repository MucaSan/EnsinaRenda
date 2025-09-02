package ifaceController

import (
	"context"
	"ensina-renda/domain/model"
)

type UsuarioController interface {
	CadastrarUsuario(ctx context.Context, usuario *model.Usuario) error
	VerificarCredenciaisUsuario(ctx context.Context, usuario *model.Usuario) (bool, error)
	GetUsuario(ctx context.Context, usuario *model.Usuario) (*model.Usuario, error)
	RealizarLogin(ctx context.Context, usuario *model.Usuario) (string, error)
	GetUsuarioPeloId(ctx context.Context) (*model.Usuario, error)
}
