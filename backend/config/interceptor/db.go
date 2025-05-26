package interceptor

import (
	"context"
	"ensina-renda/config/database"
	"time"

	"google.golang.org/grpc"
)

/*
	Função de inicialização do banco de dados.
	Basicamente, assim que o client (frontend) enviar uma requisição para o backend, o interceptor irá "interceptar" a requisição, e embedar
	a conexão com o banco, para os repositórios poderem utilizar, antes do controllers iniciarem o uso efetivo do repositório.
*/

func DatabaseUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	/*
		1. Inicia o banco de dados com a abstração do GORM (*gorm.DB)
		2. Inicia a conversão do modelo ORM para a conexão direta com o banco (tipo *sql.DB)
		3. Inicia a função "defer" que garante que mesmo com crash do interceptor ou fim do programa, a conexão com o banco será fechada.
	*/
	gormDb, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	db, err := gormDb.DB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	/*
		1. Inicia uma variável "timeout" com 10 segundos.
		2. Atribui ao valor do ctx a duração de 10 segundos.
		3. Garante que após 10 segundos, a requisição gRPC será cancelada.
	*/
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Embeda um valor de par-chave no contexto, para os repositórios acessarem.
	ctx = context.WithValue(ctx, database.DbContextKey, gormDb)

	// Passa o processo para o controller responsável e trata os possíveis erros.
	resp, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
