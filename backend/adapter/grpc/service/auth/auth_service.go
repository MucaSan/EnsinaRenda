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

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

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
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", status.Errorf(codes.Internal, "falha ao gerar o token")
	}

	return tokenString, nil
}
