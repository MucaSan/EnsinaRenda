package database

import (
	"context"
	context_config "ensina-renda/config/context"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Par de chave de acesso no contexto do banco.
const DbContextKey context_config.ContextKey = "db"

// InitDB initializes the database connection using GORM
func InitDB() (*gorm.DB, error) {
	// String de conexão no banco de dados
	dsn := "host=localhost port=5432 user=postgres dbname=ensina_renda sslmode=disable"

	novoLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Tempo de SQL para classificar como lento
			LogLevel:      logger.Info, // Nivel de log
			Colorful:      true,        // Ativar cores
		},
	)

	// Abrir a ORM do GORM
	gormDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: novoLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDb, err := gormDb.DB()
	if err != nil {
		return nil, err
	}

	// Configura timeout para o banco
	sqlDb.SetConnMaxLifetime(time.Minute * 5)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(5)

	return gormDb, nil
}

// A função GetDB retorna a conexão atual do banco pelo contexto passado
func GetDB(ctx context.Context) *gorm.DB {
	// Realizando assert no valor da interface{} do ctx.Value e verificando se é *gorm.DB
	if db, ok := ctx.Value(DbContextKey).(*gorm.DB); ok {
		return db
	}
	return nil
}
