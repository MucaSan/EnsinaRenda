package get_prova_gerada

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.GetProvaGeradaRequest,
) (*pb.GetProvaGeradaResponse, error) {

	provaUsuario, err := container.ProvaController().GetProvaUsuario(ctx, in.IdModulo)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.GetProvaGeradaResponse{
		ProvaGerada: provaUsuario.ConteudoGerado,
		Mensagem:    "busca de prova gerada com sucesso!",
		Sucesso:     true,
	}, nil
}

func RespostaErro(err error) (*pb.GetProvaGeradaResponse, error) {
	return &pb.GetProvaGeradaResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
