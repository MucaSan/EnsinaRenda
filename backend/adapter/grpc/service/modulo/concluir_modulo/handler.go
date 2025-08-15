package concluir_modulo

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.ConcluirModuloRequest,
) (*pb.ConcluirModuloResponse, error) {

	usuarioModuloConverter := NewUsuarioModuloConverter(in)

	usuarioModulo, err := usuarioModuloConverter.ToDomain(ctx)
	if err != nil {
		return RespostaErro(err)
	}

	usuarioModulo, err = container.ModuloController().GetUsuarioModulo(ctx, usuarioModulo.IDModulo, usuarioModulo.IDUsuario)
	if err != nil {
		return RespostaErro(err)
	}

	err = container.ModuloController().CompletarModulo(ctx, usuarioModulo)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.ConcluirModuloResponse{
		Mensagem: "modulo concluido com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.ConcluirModuloResponse, error) {
	return &pb.ConcluirModuloResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
