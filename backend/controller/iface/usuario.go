package ifaceController

import (
	"context"
	"ensina-renda/domain/model"
)

type UsuarioController interface {
	CadastrarUsuario(ctx context.Context, usuario *model.Usuario) error
}
