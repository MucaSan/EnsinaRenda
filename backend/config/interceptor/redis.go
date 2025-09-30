package interceptor

import (
	"context"
	"ensina-renda/config/redis"

	"google.golang.org/grpc"
)

/*
	Função de inicialização do banco de dados.
	Basicamente, assim que o client (frontend) enviar uma requisição para o backend, o interceptor irá "interceptar" a requisição, e embedar
	a conexão com o banco, para os repositórios poderem utilizar, antes do controllers iniciarem o uso efetivo do repositório.
*/

func RedisUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	clientRedis := redis.InitRedis()

	// Embeda um valor de par-chave no contexto, para os serviços utilizarem o redis.
	ctx = context.WithValue(ctx, redis.RedisContextKey, clientRedis)

	// Passa o processo para o controller responsável e trata os possíveis erros.
	return handler(ctx, req)
}
