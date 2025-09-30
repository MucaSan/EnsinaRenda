package auth

import (
	"context"
	context_config "ensina-renda/config/context"
)

// Par de chave e acesso aos dados de autenticacao do usuario
const UserUuidContextKey context_config.ContextKey = "user_uuid"
const EmailContextKey context_config.ContextKey = "email"

func GetUserUuidPeloContext(ctx context.Context) string {
	return ctx.Value(UserUuidContextKey).(string)
}

func GetEmailPeloContext(ctx context.Context) string {
	return ctx.Value(EmailContextKey).(string)
}
