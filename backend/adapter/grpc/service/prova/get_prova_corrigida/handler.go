package get_prova_corrigida

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.GetProvaCorrigidaRequest,
) (*pb.GetProvaCorrigidaResponse, error) {

	correcaoProva, err := container.ProvaController().GetCorrecaoProva(ctx, in.IdModulo)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.GetProvaCorrigidaResponse{
		ProvaCorrigida: correcaoProva.ConteudoAnalise,
		Mensagem:       "busca de prova corrigida com sucesso!",
		Sucesso:        true,
	}, nil
}

func RespostaErro(err error) (*pb.GetProvaCorrigidaResponse, error) {
	return &pb.GetProvaCorrigidaResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
