package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "*")
			c.Response().Header().Set("Access-Control-Allow-Headers", "*")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			return next(c)
		}
	})

	e.GET("/profile", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Name  string `json:"name"`
			Level int64  `json:"level"`
			Xp    int64  `json:"xp"`
		}{
			Name:  "vlad",
			Level: 0,
			Xp:    0,
		})
	})

	e.Logger.Fatal(e.Start(":9090"))
}