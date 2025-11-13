package media_final

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.MediaFinalRequest,
) (*pb.MediaFinalResponse, error) {
	mediaFinal, porcentagemFinal, err := container.ProvaController().ObterResultadoFinal(ctx)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.MediaFinalResponse{
		ResultadoFinal: &pb.ResultadoFinal{
			MediaFinal:       mediaFinal,
			PorcentagemMedia: porcentagemFinal,
		},
		Mensagem: "resultado obtido com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.MediaFinalResponse, error) {
	return &pb.MediaFinalResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
