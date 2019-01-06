package main

//+build wireinject

import (
	"context"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	ServiceContainer = wire.NewSet(NewDatabase)
)

func NewDatabase() *sqlx.DB {
	return sqlx.MustConnect("postgres", "host=localhost port=5432 user=ben password=password dbname=tapx sslmode=disable")
}

func initializeDB(ctx context.Context) (*sqlx.DB, error) {
	wire.Build(ServiceContainer)
	return &sqlx.DB{}, nil
}
