package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/tbtec/tremligeiro/internal/env"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver/controller"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver/middleware"
)

var (
	ValidVideoExtension = []string{".mp4", ".avi", ".mov", ".mkv", ".wmv", ".flv", ".webm"}
)

type HTTPServer struct {
	Server *fiber.App
	Config env.Config
}

func New(container *container.Container, config env.Config) *HTTPServer {
	slog.InfoContext(context.Background(), "Creating HTTP Server...")

	app := fiber.New(fiber.Config{ReadBufferSize: 8192})

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go gracefullyShutdown(app, trap, *container)

	app.Get("/live", adaptRoute(controller.NewLivenessController()))

	baseRouter := app.Group("/api/v1")

	baseRouter.Post("/upload", adaptRoute(controller.NewUploadRestController(container)))

	app.Use(middleware.NewNotFound())

	return &HTTPServer{
		Server: app,
		Config: config,
	}

}

func (server *HTTPServer) Listen() {
	slog.InfoContext(context.Background(), fmt.Sprintf("Starting HTTP Server on port:%v", server.Config.Port))
	err := server.Server.Listen(fmt.Sprintf(":%v", server.Config.Port))
	if err != nil {
		log.Panic(err)
	}
}

func gracefullyShutdown(app *fiber.App, trap chan os.Signal, container container.Container) {
	<-trap
	slog.InfoContext(context.Background(), "Gracefully closing resources...")
	errContainer := container.Stop()
	if errContainer != nil {
		slog.ErrorContext(context.Background(), "Error on closing resources: "+errContainer.Error())
		return
	}
	err := app.Shutdown()
	if err != nil {
		slog.InfoContext(context.Background(), "Error on shutdown Fiber app: "+err.Error())
		return
	}
	slog.InfoContext(context.Background(), "Successfully closing resources...")
}
