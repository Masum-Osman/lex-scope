package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Masum-Osman/lex-scope/pkg/config"
	"go.uber.org/fx"
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
