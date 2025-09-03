package interceptor

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	service "ensina-renda/adapter/grpc/service/auth"
	"ensina-renda/config/auth"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	exigeAutenticar := true
	switch req.(type) {
	case *pb.CadastrarAlunoRequest:
		exigeAutenticar = false
	case *pb.RealizarLoginRequest:
		exigeAutenticar = false
	case *pb.GetUsuarioPeloEmailRequest:
		exigeAutenticar = false
	case *pb.AtualizarSenhaRequest:
		exigeAutenticar = false
	default:
	}

	if !exigeAutenticar {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "os meta dados nao foram providos ao servidor")
	}

	auths := md.Get("authorization")
	if auths == nil {
		return nil, status.Errorf(codes.Unauthenticated, "o header de authorization esta vazio")
	}

	bearerToken := auths[0]

	prefixoBearer := "Bearer "

	if !strings.HasPrefix(bearerToken, prefixoBearer) {
		return nil, status.Errorf(codes.Unauthenticated, "o header de authorization nao contem o prefixo bearer")
	}

	tokenString := strings.TrimPrefix(bearerToken, prefixoBearer)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "método de assinatura inesperado: %v", token.Header["alg"])
		}
		return []byte(service.JwtSecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("falha na validação do token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "token inválido ou claims não encontrados")
	}

	if _, ok = claims["email"]; !ok {
		return nil, status.Errorf(codes.Unauthenticated, "o token criptografado nao contem o campo email")
	}

	if _, ok = claims["nome"]; !ok {
		return nil, status.Errorf(codes.Unauthenticated, "o token criptografado nao contem o campo nome")
	}

	if _, ok = claims["user_uuid"]; !ok {
		return nil, status.Errorf(codes.Unauthenticated, "o token criptografado nao contem o campo user_uuid")
	}

	if _, ok := claims["exp"]; !ok {
		return nil, status.Errorf(codes.Unauthenticated, "o token criptografado nao contem a expiracao")
	}

	horarioAtual := float64(time.Now().Unix())
	horarioExpiracao := claims["exp"].(float64)

	if horarioAtual >= horarioExpiracao {
		return nil, status.Errorf(codes.Unauthenticated, "o token criptografado expirou, tempo limite de 3 dias")
	}

	ctx = context.WithValue(ctx, auth.UserUuidContextKey, claims["user_uuid"].(string))
	ctx = context.WithValue(ctx, auth.EmailContextKey, claims["email"].(string))

	return handler(ctx, req)
}
