package concluir_aula

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.ConcluirAulaRequest,
) (*pb.ConcluirAulaResponse, error) {

	usuarioAulaConverter := NewUsuarioAulaConverter(in)

	usuarioAula, err := usuarioAulaConverter.ToDomain(ctx)
	if err != nil {
		return RespostaErro(err)
	}

	usuarioAula, err = container.AulaController().GetUsuarioAula(ctx, usuarioAula.IDAula, usuarioAula.IDUsuario)
	if err != nil {
		return RespostaErro(err)
	}

	err = container.AulaController().CompletarAula(ctx, usuarioAula)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.ConcluirAulaResponse{
		Mensagem: "aula concluida com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.ConcluirAulaResponse, error) {
	return &pb.ConcluirAulaResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, err
}
