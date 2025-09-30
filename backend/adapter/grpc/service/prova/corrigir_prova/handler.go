package corrigir_prova

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
	"strconv"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.CorrigirProvaRequest,
) (*pb.CorrigirProvaResponse, error) {

	converter := NewCorrigirProvaConverter(in)

	provaRespondida, err := converter.ToDomain(ctx)

	provaBase, err := container.ProvaController().CorrigirProva(ctx, idModulo, provaRespondida)
	if err != nil {
		return RespostaErro(err)
	}

	err = container.ProvaController().GerarProva(ctx, provaBase)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.CorrigirProvaResponse{
		Mensagem: "prova gerada com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.CorrigirProvaResponse, error) {
	return &pb.CorrigirProvaResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
