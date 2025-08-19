package realizar_login

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.RealizarLoginRequest,
) (*pb.RealizarLoginResponse, error) {

	usuarioConverter := NewUsuarioConverter(in)

	usuario, err := usuarioConverter.ToDomain(ctx)
	if err != nil {
		return RespostaErro(err)
	}

	usuario, err = container.UsuarioController().GetUsuario(ctx, usuario)
	if err != nil {
		return RespostaErro(err)
	}

	token, err := container.UsuarioController().RealizarLogin(ctx, usuario)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.RealizarLoginResponse{
		Token:    token,
		Mensagem: "usuario logado com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.RealizarLoginResponse, error) {
	return &pb.RealizarLoginResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
