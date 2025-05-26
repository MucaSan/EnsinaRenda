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

func (s *EnsinaRendaService) CadastrarAluno(ctx context.Context, req *pb.CadastrarAlunoRequest) (*pb.CadastrarAlunoResponse, error) {
	return nil, nil
}
