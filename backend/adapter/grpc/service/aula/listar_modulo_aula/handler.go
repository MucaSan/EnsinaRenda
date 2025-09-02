package listar_modulo_aula

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
	"ensina-renda/config/auth"

	"github.com/google/uuid"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.ListarModuloAulasRequest,
) (*pb.ListarModuloAulasResponse, error) {
	usuario, err := container.UsuarioController().GetUsuarioPeloId(ctx)
	if err != nil {
		return RespostaErro(err)
	}

	if usuario == nil || usuario.Id == uuid.Nil {
		return &pb.ListarModuloAulasResponse{
			ModuloAulas: nil,
			Mensagem:    "usuario nao existe",
			Sucesso:     false,
		}, nil
	}

	stringUsuarioUuid := auth.GetUserUuidPeloContext(ctx)

	usuarioUuid, err := uuid.Parse(stringUsuarioUuid)
	if err != nil {
		return &pb.ListarModuloAulasResponse{
			ModuloAulas: nil,
			Mensagem:    "usuario uuid nao e valido!",
			Sucesso:     false,
		}, nil
	}

	moduloAulas, err := container.AulaController().ListarUsuarioAulaModulo(ctx, usuarioUuid)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.ListarModuloAulasResponse{
		ModuloAulas: ConverterMapaParaPb(moduloAulas),
		Mensagem:    "status das aulas de cada modulo para o usuario listadas com sucesso!",
		Sucesso:     true,
	}, nil
}

func RespostaErro(err error) (*pb.ListarModuloAulasResponse, error) {
	return &pb.ListarModuloAulasResponse{
		ModuloAulas: nil,
		Mensagem:    err.Error(),
		Sucesso:     false,
	}, nil
}
