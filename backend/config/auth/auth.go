package auth

type contextKey string

// Par de chave e acesso aos dados de autenticacao do usuario
const UserUuidContextKey contextKey = "user_uuid"
const EmailContextKey contextKey = "email"
