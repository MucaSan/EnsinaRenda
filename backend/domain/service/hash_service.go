package domain

type HashService interface {
	GerarHashSHA256(valor string) string
}
