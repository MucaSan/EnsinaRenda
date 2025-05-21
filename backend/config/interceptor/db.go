package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func DatabaseUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Pre-processing

	// Call the handler
	return handler(ctx, req)

	// Post-processing
	log.Println("After handling:", info.FullMethod)
	return resp, err
}
