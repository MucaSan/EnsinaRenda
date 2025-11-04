package atualizar_senha

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.AtualizarSenhaRequest,
) (*pb.AtualizarSenhaResponse, error) {

	usuario, err := container.UsuarioController().BuscarUsuarioPeloJWT(ctx, in.Token)
	if err != nil {
		return RespostaErro(err)
	}

	err = container.UsuarioController().AtualizarSenha(ctx, usuario, in.Senha)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.AtualizarSenhaResponse{
		Mensagem: "senha atualizada com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.AtualizarSenhaResponse, error) {
	return &pb.AtualizarSenhaResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
