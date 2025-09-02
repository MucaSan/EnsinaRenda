package auth

import "context"

type ContextKey string

// Par de chave e acesso aos dados de autenticacao do usuario
const UserUuidContextKey ContextKey = "user_uuid"
const EmailContextKey ContextKey = "email"

func GetUserUuidPeloContext(ctx context.Context) string {
	return ctx.Value(UserUuidContextKey).(string)
}

func GetEmailPeloContext(ctx context.Context) ContextKey {
	return ctx.Value(EmailContextKey).(ContextKey)
}
