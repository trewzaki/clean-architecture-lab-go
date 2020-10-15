package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func init() {

}

func main() {
	e := echo.New()
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	serveGracefulShutdown(e)
}

func serveGracefulShutdown(e *echo.Echo) {
	go func() {
		if err := e.Start(":1323"); err != nil {
			zap.L().Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	gracefulShutdownTimeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		zap.L().Fatal(err.Error())
	}
}
