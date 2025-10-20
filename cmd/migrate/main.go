package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v5/stdlib"

	_ "cinelist/infrastructure/database/migrations"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Uso: go run ./cmd/migrate <comando>\nExemplo: go run ./cmd/migrate up")
	}
	command := os.Args[1]

	godotenv.Load()
	dbURL := os.Getenv("DATABASE_CONNECTION_STRING")
	if dbURL == "" {
		log.Fatal("ERRO: Variável de ambiente DATABASE_CONNECTION_STRING não definida.\nExemplo: export DATABASE_CONNECTION_STRING=\"postgres://usuario:senha@localhost:5432/nome_db?sslmode=disable\"")
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Falha ao definir o dialeto: %v", err)
	}

	db, err := sql.Open("pgx", dbURL) // "pgx" é o nome do driver
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Falha ao pingar o banco: %v", err)
	}

	migrationsDir := "infrastructure/database/migrations"

	if err := goose.RunContext(context.Background(), command, db, migrationsDir, os.Args[2:]...); err != nil {
		log.Fatalf("Comando 'goose %s' falhou: %v", command, err)
	}

	log.Printf("Comando 'goose %s' concluído com sucesso.", command)
}
