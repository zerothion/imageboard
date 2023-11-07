package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/zerothion/imageboard/internal/delivery"
	"github.com/zerothion/imageboard/internal/repo"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("Failed to load .env file", "err", err)
	}

	s := delivery.NewHTTP(delivery.Repos{
		UserRepo: repo.NewUserRepoStub(),
	})
	s.ServeDefault()
}
