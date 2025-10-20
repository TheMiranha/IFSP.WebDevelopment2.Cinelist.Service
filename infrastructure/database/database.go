package infrastructure_database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	errorLoadingEnv := godotenv.Load()

	if errorLoadingEnv != nil {
		panic(errorLoadingEnv)
	}

	psqlInfo := os.Getenv("DATABASE_CONNECTION_STRING")

	conn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	return conn
}
