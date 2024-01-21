package app

import (
	"Test-Task2/internal/config"
	"Test-Task2/internal/register_layers"
	"Test-Task2/pkg/api/jwt"
	"Test-Task2/pkg/database/inmemory"
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	ctx context.Context
	log *zap.Logger
	db  inmemory.Storage
}

func NewApp(ctx context.Context, log *zap.Logger) *App {
	return &App{
		ctx: ctx,
		log: log,
	}
}

func (app *App) Run() error {
	cfg := config.GetConfig()

	j := jwt.NewJWT(
		cfg.JWT.AccessToken,
		cfg.JWT.SecretKey,
	)

	db := inmemory.NewInMemory()

	server := grpc.NewServer()

	gRepo := register_layers.NewGRepository(db, app.log)
	gUsecase := register_layers.NewGUsecase(gRepo)
	gHandler := register_layers.NewGDelivery(gUsecase, j)
	gHandler.RegisterRoutes(server)

	listener, err := net.Listen("tcp", fmt.Sprint(cfg.GRPC.Host+":"+cfg.GRPC.Port))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server starts to work")

	if err = server.Serve(listener); err != nil {
		return err
	}

	return nil
}
