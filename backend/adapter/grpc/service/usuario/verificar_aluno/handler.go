package verificar_aluno

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.VerificarAlunoRequest,
) (*pb.VerificarAlunoResponse, error) {

	usuarioConverter := NewUsuarioConverter(in)

	usuario, err := usuarioConverter.ToDomain(ctx)
	if err != nil {
		return RespostaErro(err)
	}

	verificadoComSucesso, err := container.UsuarioController().VerificarCredenciaisUsuario(ctx, usuario)
	if err != nil {
		return RespostaErro(err)
	}

	if !verificadoComSucesso {
		return &pb.VerificarAlunoResponse{
			Mensagem: "aluno nao existe, ou credenciais estão erradas!",
			Sucesso:  true,
		}, nil
	}

	return &pb.VerificarAlunoResponse{
		Mensagem: "aluno logado com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.VerificarAlunoResponse, error) {
	return &pb.VerificarAlunoResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, err
}
