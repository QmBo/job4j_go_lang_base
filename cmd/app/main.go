package main

import (
	"context"
	"log"

	"job4j.ru/go-lang-base/internal/config"
	"job4j.ru/go-lang-base/internal/db"
	"job4j.ru/go-lang-base/internal/repository"
	"job4j.ru/go-lang-base/internal/tracker"
	"job4j.ru/go-lang-base/internal/trackerstore"
)

func main() {
	ctx := context.Background()

	cfg := db.Config{
		Host:     config.Env("DB_HOST", "localhost"),
		Port:     config.EnvInt("DB_PORT", 5432),
		User:     config.Env("DB_USER", "postgres"),
		Password: config.Env("DB_PASSWORD", "password"),
		DBName:   config.Env("DB_NAME", "local"),
		SSLMode:  config.Env("DB_SSLMODE", "disable"),
		Schema:   config.Env("DB_SCHEMA", "tracker"),
	}

	pool, err := db.NewPool(ctx, cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := repository.NewRepoPg(pool)

	ui := trackerstore.UI{
		Cotext: ctx,
		In:     tracker.ConsoleInput{},
		Out:    tracker.ConsoleOutput{},
		Store:  repo,
	}

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
