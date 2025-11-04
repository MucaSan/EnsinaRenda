package enviar_reset_senha

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
)

func Handle(
	ctx context.Context,
	container container.EnsinaRendaContainerInterface,
	in *pb.EnviarResetSenhaRequest,
) (*pb.EnviarResetSenhaResponse, error) {
	email := in.Email

	emailHasheado := container.UsuarioController().CriptografarEmail(ctx, email)

	usuario, err := container.UsuarioController().GetUsuarioPeloEmail(ctx, emailHasheado)
	if err != nil {
		return RespostaErro(err)
	}

	token, err := container.UsuarioController().GerarToken(ctx, usuario)
	if err != nil {
		return RespostaErro(err)
	}

	err = container.UsuarioController().EnviarEmail(ctx, email, token)
	if err != nil {
		return RespostaErro(err)
	}

	return &pb.EnviarResetSenhaResponse{
		Mensagem: "e-mail enviado com sucesso!",
		Sucesso:  true,
	}, nil
}

func RespostaErro(err error) (*pb.EnviarResetSenhaResponse, error) {
	return &pb.EnviarResetSenhaResponse{
		Mensagem: err.Error(),
		Sucesso:  false,
	}, nil
}
