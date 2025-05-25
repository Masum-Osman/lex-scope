package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
	"gorm.io/gorm/logger"
)

func main() {
	app := fx.New(
		config.Module,
		logger.Module,
		di.Module,
		fx.Invoke(run),
	)

	app.Start(context.Background())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	app.Stop(context.Background())
}

func run(lc fx.Lifecycle, logger logger.Logger) {
	lc.Append((fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Lex Scope service started...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Lex Scope service stopped.")
			return nil
		},
	}))
}
