package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaif79-ri/echo-esearch/middleware"
)

func main() {
	e := echo.New()
	middleware.SetEchoMiddleware(e)

	e.POST("/", func(c echo.Context) error {
		var breq map[string]interface{}

		if err := c.Bind(&breq); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		fmt.Println(breq)

		return c.JSON(http.StatusOK, map[string]interface{}{"data": "Hello, World!"})
	})
	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
