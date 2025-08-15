package cadastrar_aluno

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.CadastrarAlunoRequest,
) (*pb.CadastrarAlunoResponse, error) {

	usuarioConverter := NewUsuarioConverter(in)

	usuario, err := usuarioConverter.ToDomain(ctx)
	if err != nil {
		return RespostaErro(err)
	}

	err = container.UsuarioController().CadastrarUsuario(ctx, usuario)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.CadastrarAlunoResponse{
		Mensagem: "aluno cadastrado com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.CadastrarAlunoResponse, error) {
	return &pb.CadastrarAlunoResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
