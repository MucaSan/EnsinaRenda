package service

import (
	domain "ensina-renda/domain/service"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService struct {
	client *sendgrid.Client
}

func NewEmailService() domain.EmailService {
	return &EmailService{
		client: sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY")),
	}
}

func (s *EmailService) EnviarEmail(email, token string) error {
	from := mail.NewEmail("Equipe Ensina Renda", "nao-responda@ensinararenda.com.br")

	subject := "Redefinição de Senha - EnsinaRenda"

	to := mail.NewEmail("Usuário do EnsinaRenda", email)

	resetLink := fmt.Sprintf("https://ensinararenda.com.br/reset?token=%s", token)

	plainTextContent := fmt.Sprintf("Olá! Para redefinir sua senha, clique no link: %s", resetLink)

	htmlContent := fmt.Sprintf(`
		<strong>Olá!</strong>
		<p>Recebemos uma solicitação para redefinir sua senha.</p>
		<p>Clique no link abaixo para criar uma nova senha:</p>
		<a href="%s" target="_blank">Redefinir Minha Senha</a>
		<p>Se você não solicitou isso, pode ignorar este e-mail.</p>
	`, resetLink)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	resposta, err := s.client.Send(message)
	if err != nil {
		return err
	}

	fmt.Printf("Resposta do SendGrid - Status Code: %d\n", resposta.StatusCode)
	fmt.Printf("Resposta do SendGrid - Body: %s\n", resposta.Body)
	fmt.Printf("Resposta do SendGrid - Headers: %v\n", resposta.Headers)

	// O status 202 NÃO garante a entrega, apenas que o SendGrid aceitou o trabalho.
	if resposta.StatusCode >= 200 && resposta.StatusCode < 300 {
		fmt.Println("E-mail aceito para envio pelo SendGrid.")
		return nil
	}

	// Se o status for 4xx ou 5xx, algo está errado com o *pedido*
	fmt.Printf("ERRO: O SendGrid rejeitou o pedido com status %d\n", resposta.StatusCode)
	return fmt.Errorf("sendgrid rejeitou com status: %d", resposta.StatusCode)
}
