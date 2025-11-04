package service

import (
	"crypto/sha256"
	domain "ensina-renda/domain/service"
	"fmt"
)

type HashService struct {
}

func NewHashService() domain.HashService {
	return &HashService{}
}

func (s *HashService) GerarHashSHA256(valor string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(valor)))
}
