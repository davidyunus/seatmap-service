package routes

import (
	"github.com/seatmap-service/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/api/seatmap", controllers.GetSeatMap)
}
