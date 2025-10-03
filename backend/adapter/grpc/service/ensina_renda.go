package service

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/aula/concluir_aula"
	"ensina-renda/adapter/grpc/service/aula/listar_modulo_aula"
	"ensina-renda/adapter/grpc/service/auth/realizar_login"
	"ensina-renda/adapter/grpc/service/container"
	"ensina-renda/adapter/grpc/service/modulo/concluir_modulo"
	"ensina-renda/adapter/grpc/service/prova/corrigir_prova"
	"ensina-renda/adapter/grpc/service/prova/gerar_prova"
	"ensina-renda/adapter/grpc/service/prova/get_prova_corrigida"
	"ensina-renda/adapter/grpc/service/prova/get_prova_gerada"
	"ensina-renda/adapter/grpc/service/usuario/atualizar_senha"
	"ensina-renda/adapter/grpc/service/usuario/cadastrar_aluno"
	"ensina-renda/adapter/grpc/service/usuario/get_usuario_email"
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

func (s *EnsinaRendaService) ConcluirModulo(ctx context.Context, req *pb.ConcluirModuloRequest) (*pb.ConcluirModuloResponse, error) {
	return concluir_modulo.Handle(ctx, s.container, req)
}

func (s *EnsinaRendaService) RealizarLogin(ctx context.Context, req *pb.RealizarLoginRequest) (*pb.RealizarLoginResponse, error) {
	return realizar_login.Handle(ctx, s.container, req)
}

func (s *EnsinaRendaService) ListarModuloAulas(ctx context.Context, req *pb.ListarModuloAulasRequest) (*pb.ListarModuloAulasResponse, error) {
	return listar_modulo_aula.Handle(ctx, s.container, req)
}

func (s *EnsinaRendaService) AtualizarSenha(ctx context.Context, req *pb.AtualizarSenhaRequest) (*pb.AtualizarSenhaResponse, error) {
	return atualizar_senha.Handle(ctx, s.container, req)
}

func (s *EnsinaRendaService) GetUsuarioPeloEmail(ctx context.Context, req *pb.GetUsuarioPeloEmailRequest) (*pb.GetUsuarioPeloEmailResponse, error) {
	return get_usuario_email.Handle(ctx, s.container, req)
}

func (s *EnsinaRendaService) GerarProva(ctx context.Context, req *pb.GerarProvaRequest) (*pb.GerarProvaResponse, error) {
	return gerar_prova.Handle(ctx, s.container, req)
}

func (s *EnsinaRendaService) CorrigirProva(ctx context.Context, req *pb.CorrigirProvaRequest) (*pb.CorrigirProvaResponse, error) {
	return corrigir_prova.Handle(ctx, s.container, req)
}

func (s *EnsinaRendaService) GetProvaGerada(ctx context.Context, req *pb.GetProvaGeradaRequest) (*pb.GetProvaGeradaResponse, error) {
	return get_prova_gerada.Handle(ctx, s.container, req)
}

func (s *EnsinaRendaService) GetProvaCorrigida(ctx context.Context, req *pb.GetProvaCorrigidaRequest) (*pb.GetProvaCorrigidaResponse, error) {
	return get_prova_corrigida.Handle(ctx, s.container, req)
}
