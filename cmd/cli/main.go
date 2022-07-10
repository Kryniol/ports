package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/Kryniol/ports/domain"
	"github.com/Kryniol/ports/infrastructure"
)

var inputPath = flag.String("path", "", "Path to the JSON input file")

func main() {
	flag.Parse()

	if *inputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	runApp(ctx, logger)
}

func runApp(ctx context.Context, l *zap.Logger) {
	repo := infrastructure.NewInMemoryPortRepository()
	reader := infrastructure.NewJSONReader(*inputPath, l)
	svc := domain.NewPortService(reader, repo)
	err := svc.SavePorts(ctx)
	if err != nil {
		l.Fatal("couldn't save ports", zap.Error(err))
	}

	l.Info("ports have been saved")
}
