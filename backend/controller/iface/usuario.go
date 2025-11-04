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
	GetUsuarioPeloIdDoContexto(ctx context.Context) (*model.Usuario, error)
	GetUsuarioPeloEmail(ctx context.Context, email string) (*model.Usuario, error)
	AtualizarSenha(ctx context.Context, usuario *model.Usuario, senha string) error
	GetUsuarioPeloId(ctx context.Context, id string) (*model.Usuario, error)
	ProvisionarUsuarioModulos(ctx context.Context, usuario *model.Usuario, modulos []*model.Modulo) error
	ProvisionarUsuarioAulas(ctx context.Context, usuario *model.Usuario, aulas []*model.Aula) error
	CriptografarEmail(ctx context.Context, email string) string
	EnviarEmail(ctx context.Context, email, token string) error
	GerarToken(ctx context.Context, usuario *model.Usuario) (string, error)
	BuscarUsuarioPeloJWT(ctx context.Context, token string) (*model.Usuario, error)
}
