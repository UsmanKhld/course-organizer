package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := fmt.Sprintf("postgres://postgres:postgres@localhost:5432/course_organizer")
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	DB = pool
	fmt.Println("✅ Connected to Postgres")
}
