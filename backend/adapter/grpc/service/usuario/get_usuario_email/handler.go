package get_usuario_email

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
	dto "ensina-renda/adapter/grpc/service/usuario"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.GetUsuarioPeloEmailRequest,
) (*pb.GetUsuarioPeloEmailResponse, error) {

	usuario, err := container.UsuarioController().GetUsuarioPeloEmail(ctx, in.Email)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.GetUsuarioPeloEmailResponse{
		Usuario:  dto.ConverterUsuarioParaPb(usuario),
		Mensagem: "busca de usuario pelo emails com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.GetUsuarioPeloEmailResponse, error) {
	return &pb.GetUsuarioPeloEmailResponse{
		Usuario:  nil,
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
