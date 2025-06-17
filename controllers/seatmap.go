package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func GetSeatMap(c echo.Context) error {
	file, err := os.Open("SeatMapResponse.json")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to open file"})
	}
	defer file.Close()

	var seatMapData map[string]interface{}
	if err := json.NewDecoder(file).Decode(&seatMapData); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse JSON"})
	}

	return c.JSON(http.StatusOK, seatMapData)
}
