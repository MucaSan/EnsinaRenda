package main

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service"
	"ensina-renda/adapter/grpc/service/container"
	"ensina-renda/config/interceptor"
	controller "ensina-renda/controller/usuario"
	"ensina-renda/repository"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
	Esse é o arquivo principal do projeto, responsável por iniciar paralelamente em uma única função o proxy reverso HTTP e o servidor gRPC.
*/

func main() {
	// Inicia um canal de erros com tamanho 2
	canalErro := make(chan error, 2)

	// Inicialização da variável de grupo de espera
	var wg sync.WaitGroup
	// Adiciona duas go rotinas para o processo de paralelismo
	wg.Add(2)

	// Roda o servidor gRPC em paralelo
	go func() {
		defer wg.Done()

		log.Println("Subindo servidor gRPC em localhost:9090")
		err := subirServidorGRPC()
		if err != nil {
			canalErro <- err
		}
	}()

	// Roda o proxy reverso HTTP em paralelo
	go func() {
		defer wg.Done()

		log.Println("Subindo proxy reverso em localhost:8081")
		err := subirProxyReversoHTTP()
		if err != nil {
			canalErro <- err
		}
	}()

	// Caso ocorra um erro, só entrará nos logs de parada caso os dois falhem.
	wg.Wait()

	err := <-canalErro
	if err != nil {
		log.Printf("Erro encontrado: %v \n", err.Error())
		return
	}

	log.Printf("Os servidores tiveram uma conclusão inesperada.")

}

// Sobe o servidor gRPC para processar as requisições do frontend e interagir diretamente com o banco de dados.
func subirServidorGRPC() error {
	// Faz com que o processo comece a ouvir através do protocolo TCP-IP na porta 9090.
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("\nNao foi possivel criar um listener para o localhost na :9090. \nErro encontrado: %s", err.Error())
	}

	// Inicializa os repositórios para serem utilizados pelos controllers.
	usuarioRepository := repository.NewUsuarioRepository()

	// Container conterá todos os controllers do sistema, sendo utilizados pelos handlers
	container := container.NewEnsinaRendaContainer(controller.NewUsuarioController(usuarioRepository))

	// Registra, respectivamente, o servidor gRPC (com os interceptors nele) e o serviço com a API gRPC.
	servidorGrpc := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.DatabaseUnaryInterceptor),
	)
	ensinaRendaService := service.NewEnsinaRendaService(container)
	pb.RegisterEnsinaRendaServiceServer(servidorGrpc, ensinaRendaService)

	// Inicializa o servidor gRPC
	return servidorGrpc.Serve(lis)
}

// Sobe o proxy reverso, que serve como uma abstração para a camada do servidor gRPC, basicamente, transformando requisições HTTP para gRPC.
func subirProxyReversoHTTP() error {
	// Porta de entrada para o endpoint de registro do proxy reverso HTTP.
	endpointservidorGrpc := flag.String("endpoint-servidor-grpc", "localhost:9090", "endpoint de entrada para o servidor gRPC")

	// Cria o canal de comunicação do contexto, que interpreta a requisição do frontend e embeda isso na variável a fim de manuntenção das ações necessárias para o backend.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//Registra o servidor MUX, com a comunicação TSL/SSL inativada.
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterEnsinaRendaServiceHandlerFromEndpoint(ctx, mux, *endpointservidorGrpc, opts)
	if err != nil {
		return err
	}

	// Inicializa o servidor do proxy reverso com a configuração do MUX na porta 8081 (evitar subir o frontend e o backend na mesma porta)
	return http.ListenAndServe(":8081", mux)
}
