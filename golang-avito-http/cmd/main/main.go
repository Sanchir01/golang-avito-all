package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sanchir01/golang-avito/internal/app"
	httpserver "github.com/Sanchir01/golang-avito/internal/server/servers/http"
	httphandlers "github.com/Sanchir01/golang-avito/internal/server/servers/http/handlers"
	"github.com/fatih/color"
)

func main() {
	env, err := app.NewEnv()
	if err != nil {
		panic(err)
	}
	serve := httpserver.NewHTTPServer(env.Cfg.Servers.HTTPServer.Host, env.Cfg.Servers.HTTPServer.Port,
		env.Cfg.Servers.HTTPServer.Timeout, env.Cfg.Servers.HTTPServer.IdleTimeout)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	green := color.New(color.FgGreen).SprintFunc()
	env.Lg.Info(green("🚀 Server started successfully!"),
		slog.String("time", time.Now().Format("2006-01-02 15:04:05")),
		slog.String("port", env.Cfg.Servers.HTTPServer.Port),
	)
	defer cancel()
	go func() {
		if err := serve.Run(httphandlers.StartHTTTPHandlers(env.Handlers)); err != nil {
			if !errors.Is(err, context.Canceled) {
				env.Lg.Error("Listen server error", slog.String("error", err.Error()))
				return
			}
		}

	}()

	<-ctx.Done()
	if err := serve.Gracefull(ctx); err != nil {
		env.Lg.Error("server gracefull")
	}
	if err := env.Database.Close(); err != nil {
		env.Lg.Error("Close database", slog.String("error", err.Error()))
	}
}
