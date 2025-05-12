package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ndraaa15/foreglyc-server/bootstrap"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	app := bootstrap.New()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signalChan
		logrus.Info("Received signal shutdown ",
			logrus.Fields{
				"signal": sig,
				"pid":    os.Getpid(),
			},
		)
		logrus.Info("Shutting down gracefully...")
		app.Shutdown(ctx)
		os.Exit(0)
	}()

	app.Run()
}
