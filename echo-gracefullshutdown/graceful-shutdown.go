package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func gracefulShutdown() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.GET("/", func(c echo.Context) error {
		e.Logger.Info("GET request")
		return c.JSON(http.StatusOK, "OK")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//Start the server
	go func() {
		if err := e.Start(":8099"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channedl to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
