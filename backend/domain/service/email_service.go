package domain

type EmailService interface {
	EnviarEmail(email, token string) error
}
