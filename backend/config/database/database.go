package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type contextKey string

// Inicialização de variáveis de contexto do banco, para evitar colisões de tipos.
const DbContextKey contextKey = "db"

//Função responsável por inicializar o banco de dados pelo interceptor.

func InitDB() (*sql.DB, error) {
	// Inicializa a string de conexão com o banco de dados relacional postgres.
	connStr := "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable"
	postgresDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return postgresDB, nil
}

// Função responsável por ser chamada pelos repositories para operações com o banco.

func GetDB(ctx context.Context) *sql.DB {
	// Retorna a conexão com o banco de dados pela par de chaves configurado no interceptor.
	if db, ok := ctx.Value(DbContextKey).(*sql.DB); ok {
		return db
	}
	return nil
}
