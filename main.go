package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"diikstra.fr/template/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Static("/assets", "assets")

	handlers.Init()

	// ---- Home Routes ---- //
	e.GET("/", handlers.HomeHandler)
	e.GET("/buttons", handlers.Buttons)
	e.POST("/buttons", handlers.ButtonsChange)

	// ---- Global Routes ---- //
	e.GET("/ping", handlers.GlobalPing)
	if os.Getenv("ENV") != "prod" {
		e.GET("/ws", handlers.GlobalHotReloadWS)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
