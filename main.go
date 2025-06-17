package main

import (
	"github.com/seatmap-service/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
