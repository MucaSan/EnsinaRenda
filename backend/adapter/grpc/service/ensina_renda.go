package service

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/aula/concluir_aula"
	"ensina-renda/adapter/grpc/service/container"
	"ensina-renda/adapter/grpc/service/usuario/cadastrar_aluno"
	"ensina-renda/adapter/grpc/service/usuario/verificar_aluno"
)

type EnsinaRendaService struct {
	pb.EnsinaRendaServiceServer
	container container.EnsinaRendaContainerInterface
}

func NewEnsinaRendaService(container container.EnsinaRendaContainerInterface) *EnsinaRendaService {
	return &EnsinaRendaService{
		container: container,
	}
}

func (s *EnsinaRendaService) CadastrarAluno(ctx context.Context, req *pb.CadastrarAlunoRequest) (*pb.CadastrarAlunoResponse, error) {
	return cadastrar_aluno.Handle(ctx, s.container, req)
}

func (s *EnsinaRendaService) VerificarAluno(ctx context.Context, req *pb.VerificarAlunoRequest) (*pb.VerificarAlunoResponse, error) {
	return verificar_aluno.Handle(ctx, s.container, req)
}

func (s *EnsinaRendaService) ConcluirAula(ctx context.Context, req *pb.ConcluirAulaRequest) (*pb.ConcluirAulaResponse, error) {
	return concluir_aula.Handle(ctx, s.container, req)
}
