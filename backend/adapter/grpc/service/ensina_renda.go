package service

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
)

type EnsinaRendaService struct {
	pb.EnsinaRendaServiceServer
}

func NewEnsinaRendaService() *EnsinaRendaService {
	return &EnsinaRendaService{}
}

func (s *EnsinaRendaService) CriarUsuario(ctx context.Context, req *pb.CriarUsuarioRequest) (*pb.CriarUsuarioResponse, error) {
	return nil, nil
}
