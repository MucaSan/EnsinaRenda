package main

import (
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/adapter/grpc/service"
	"ensina-renda/config/interceptor"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

/*
	Esse é o arquivo principal do projeto, responsável por iniciar o servidor gRPC e registrar os interceptors.
*/

func main() {
	// Faz com que o processo comece a ouvir através do protocolo TCP-IP na porta 50051.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("\nCouldn't create listener for localhost port :50051. \nError found: %s", err.Error())
	}

	// Registra, respectivamente, o servidor gRPC (com os interceptors nele) e o serviço com a API gRPC.
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.DatabaseUnaryInterceptor),
	)
	ensinaRendaService := service.NewEnsinaRendaService()
	pb.RegisterEnsinaRendaServiceServer(grpcServer, ensinaRendaService)

	//Disponibiliza o servidor para ser consumido.
	log.Println("Server started on port 50051.")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("\nCouldn't serve with provided listener.\nError found: %s", err.Error())
	}
}
