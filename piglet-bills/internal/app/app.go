package app

import (
	"fmt"
	"log/slog"
	pgmigration "piglet-bills-service/internal/app/postgres"

	"piglet-bills-service/internal/services/accounting"
	"piglet-bills-service/internal/storage/psql"

	grpcapp "piglet-bills-service/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcServer string,
	grpcPort string,
	migrationPath string,
	dbUser string,
	dbPassword string,
	dbHost string,
	dbPort string,
	dbName string,
) *App {
	if err := pgmigration.RunMigration(
		"file://"+migrationPath,
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			dbUser, dbPassword, dbHost, dbPort, dbName),
	); err != nil {
		panic(err)
	}

	storage, err := psql.New(
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)
	if err != nil {
		panic(err)
	}

	accountingService := accounting.New(log, storage)

	grpcApp := grpcapp.New(log, accountingService, grpcServer, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
