package corrigir_prova

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.CorrigirProvaRequest,
) (*pb.CorrigirProvaResponse, error) {

	converter := NewCorrigirProvaConverter(in)

	provaRespondida, err := converter.ToDomain(ctx)
	if err != nil {
		return nil, err
	}

	err = container.ProvaController().CorrigirProva(ctx, in.IdModulo, provaRespondida)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.CorrigirProvaResponse{
		Mensagem: "prova corrigida com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.CorrigirProvaResponse, error) {
	return &pb.CorrigirProvaResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
