package main

import (
	"apex_roulette/application_service/service"
	"apex_roulette/config"
	"apex_roulette/infrastructure/logger"
	"apex_roulette/presentation/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	config := config.New()
	l := logger.NewLogger(os.Stdout)

	l.StructLog(service.LogLevelInfo, "Starting bot...")

	// 全てのギルドに対して動作させる場合は空文字を指定
	guild := ""
	d := server.New(config, guild, l)

	go d.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	defer close(quit)

	<-quit
	l.StructLog(service.LogLevelInfo, "Stopping bot...")
	d.Shutdown()
}
