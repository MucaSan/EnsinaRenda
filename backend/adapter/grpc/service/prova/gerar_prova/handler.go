package gerar_prova

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.GerarProvaRequest,
) (*pb.GerarProvaResponse, error) {

	idModulo := int(in.IdModulo)

	provaBase, err := container.ProvaController().GetProvaBase(ctx, idModulo)
	if err != nil {
		return RespostaErro(err)
	}

	err = container.ProvaController().GerarProva(ctx, provaBase)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.GerarProvaResponse{
		Mensagem: "prova gerada com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.GerarProvaResponse, error) {
	return &pb.GerarProvaResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
