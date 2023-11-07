package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/zerothion/imageboard/internal/delivery"
	"github.com/zerothion/imageboard/internal/delivery/rest"
	"github.com/zerothion/imageboard/internal/repo"
	"github.com/zerothion/imageboard/internal/repo/postgres"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("Failed to load .env file", "err", err)
	}

	var db_url = os.Getenv("POSTGRES_URL")
	if db_url == "" {
		slog.Error("Env `POSTGRES_URL` is not set!")
		os.Exit(1)
	}

	db, err := pgxpool.New(context.Background(), db_url)
	if err != nil {
		slog.Error("Failed to connect to database!", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping(context.Background())
	if err != nil {
		slog.Error("Failed to ping database!", "err", err)
		os.Exit(1)
	}

	s := delivery.NewHTTP(delivery.Repos{
		UserRepo: postgres.NewUserRepo(db),
	})
	rest.AddUserHandlers(s)
	s.ServeDefault()
}
