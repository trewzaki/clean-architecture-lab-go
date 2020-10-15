package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	userDelivery "clean-architecture-go/feature/user/delivery"
	userRepo "clean-architecture-go/feature/user/repository/postgres"
	userUsecase "clean-architecture-go/feature/user/usecase"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB : Main database instance
var DB *gorm.DB

func init() {
	// Initial database
	var err error
	DB, err = newDB()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	e := echo.New()
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1 := e.Group("/v1")

	userDelivery.NewHandler(v1, userUsecase.NewUserUsecase(userRepo.NewUserRepository(DB)))

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

func newDB() (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		"postgres_host", "postgres_port", "postgres_user", "postgres_password", "postgres_dbname",
	)

	return gorm.Open(postgres.Open(connString), &gorm.Config{})
}
