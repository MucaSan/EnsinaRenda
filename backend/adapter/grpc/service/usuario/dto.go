package dto

import (
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/domain/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConverterUsuarioParaPb(usuario *model.Usuario) *pb.Usuario {
	return &pb.Usuario{
		Id:       usuario.Id.String(),
		Email:    usuario.Email,
		Nome:     usuario.Nome,
		CriadoEm: timestamppb.New(usuario.CriadoEm),
	}
}
