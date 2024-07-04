package app

import (
	grpcapp "gameng/internal/app/grpc"
	"gameng/internal/config"
	wsmock "gameng/internal/service/mock"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, cfg *config.Config) *App {
	const op = "app.New"

	log = log.With(slog.String("op", op))

	// log.Info("initializing storage")
	//storage, err := postgres.New(&cfg.Postgres)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info("storage initialized")
	//
	//log.Info("initializing cache")
	//cache, err := redis.New(&cfg.Redis, storage)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info("cache initialized")

	log.Info("initializing grpc server")
	grpcServer := grpcapp.New(log, cfg.GRPC.Port, wsmock.New())
	log.Info("grpc server initialized")

	return &App{
		GRPCSrv: grpcServer,
	}
}
