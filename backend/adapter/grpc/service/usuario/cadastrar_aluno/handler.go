package cadastrar_aluno

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service/container"
	"ensina-renda/config/database"
	"ensina-renda/domain/model"
	"log"
	"sync"
	"time"
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

	providenciarCursoAoAluno(container, usuario)

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

func providenciarCursoAoAluno(container container.EnsinaRendaContainerInterface, usuario *model.Usuario) {
	go func() {
		gormDb, err := database.InitDB()
		if err != nil {
			return
		}

		db, err := gormDb.DB()
		if err != nil {
			return
		}
		defer db.Close()

		timeout := 10 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		ctx = context.WithValue(ctx, database.DbContextKey, gormDb)

		grupoEspera := &sync.WaitGroup{}
		var aulas []*model.Aula
		var modulos []*model.Modulo

		grupoEspera.Add(1)
		go func() {
			defer grupoEspera.Done()

			aulas, err = container.AulaController().ListarAulas(ctx)
			if err != nil {
				log.Println(err)
				return
			}
		}()

		grupoEspera.Add(1)
		go func() {
			defer grupoEspera.Done()

			modulos, err = container.ModuloController().ListarModulos(ctx)
			if err != nil {
				log.Println(err)
				return
			}
		}()

		grupoEspera.Wait()

		grupoEspera.Add(1)
		go func() {
			defer grupoEspera.Done()

			err = container.UsuarioController().ProvisionarUsuarioAulas(ctx, usuario, aulas)
			if err != nil {
				log.Println(err)
				return
			}
		}()

		grupoEspera.Add(1)
		go func() {
			defer grupoEspera.Done()

			err = container.UsuarioController().ProvisionarUsuarioModulos(ctx, usuario, modulos)
			if err != nil {
				log.Println(err)
				return
			}
		}()

		grupoEspera.Wait()
	}()
}
