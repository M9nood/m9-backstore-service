package main

import (
	"context"
	"m9-backstore-service/database"
	"m9-backstore-service/pkg"
	"os"
	"os/signal"
	"time"

	router "m9-backstore-service/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Validator = pkg.NewValidationUtil()

	db := database.Connect()
	defer db.Close()

	e = router.RouterSetup(e, db)

	port := os.Getenv("PORT")
	go func() {
		if err := e.Start(":" + port); err != nil {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
