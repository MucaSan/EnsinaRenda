package service

import (
	"context"
	"ensina-renda/domain/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var JwtSecretKey = os.Getenv("JWT_SECRET_KEY")

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (s *JwtService) GerarJWT(ctx context.Context, usuario *model.Usuario) (string, error) {
	claims := jwt.MapClaims{
		"email":     usuario.Email,
		"nome":      usuario.Nome,
		"user_uuid": usuario.Id.String(),
		"exp":       time.Now().Add(time.Hour * 24 * 3).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Internal, "falha ao gerar o token")
	}

	return tokenString, nil
}

func (s *JwtService) DecodificarUUID(ctx context.Context, tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "método de assinatura inesperado: %v", token.Header["alg"])
		}

		return []byte(JwtSecretKey), nil

	}, jwt.WithStrictDecoding())

	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "token inválido ou expirado: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_uuid"].(string), nil
	}

	return "", status.Errorf(codes.Unauthenticated, "token inválido")
}
