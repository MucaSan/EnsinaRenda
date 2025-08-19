package realizar_login

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/domain/model"
	"errors"
)

type UsuarioConverter struct {
	base any
}

func NewUsuarioConverter(base any) *UsuarioConverter {
	return &UsuarioConverter{
		base: base,
	}
}

func (uc *UsuarioConverter) ToDomain(ctx context.Context) (*model.Usuario, error) {
	realizarLoginRequest, ok := uc.base.(*pb.RealizarLoginRequest)
	if !ok {
		return nil, errors.New("nao foi possivel converter base para realizar_login_request")
	}

	modelUsuario := &model.Usuario{
		Email: realizarLoginRequest.Email,
		Senha: realizarLoginRequest.Senha,
	}

	return modelUsuario, nil

}
