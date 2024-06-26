package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	transactionsgrpc "piglet-transactions-service/internal/grpc/transactions"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	server     string
	port       string
}

func New(
	log *slog.Logger,
	transactionsService transactionsgrpc.Transactions,
	categoryService transactionsgrpc.Categories,
	billService transactionsgrpc.Bills,
	grpcBillsCli *grpc.ClientConn,
	server string,
	port string,
) *App {
	gRPCServer := grpc.NewServer()

	transactionsgrpc.Register(gRPCServer, grpcBillsCli, transactionsService, categoryService, billService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		server:     server,
		port:       port,
	}
}

func (a *App) MustStart() {
	if err := a.startServer(); err != nil {
		panic(err)
	}
}

func (a *App) startServer() error {
	const op = "piglet-transactions | grpcapp.startServer"

	log := a.log.With(
		slog.String("op", op),
		slog.String("server", a.server),
		slog.String("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.server, a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running and listening", slog.String("addr", l.Addr().String()))

	if err = a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "piglet-transactions | grpcapp.Stop"

	a.log.With(slog.String("op", op), slog.String("server", a.server)).
		Info("stopping grpc server", slog.String("port", a.port))

	a.gRPCServer.GracefulStop()
}
