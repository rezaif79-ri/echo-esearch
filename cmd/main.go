package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaif79-ri/echo-esearch/config"
	"github.com/rezaif79-ri/echo-esearch/middleware"
	"github.com/rezaif79-ri/echo-esearch/router"
)

func main() {
	e := echo.New()
	middleware.SetEchoMiddleware(e)

	es, err := config.InitElasticClient()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.POST("/", func(c echo.Context) error {
		var breq map[string]interface{}

		if err := c.Bind(&breq); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		fmt.Println(breq)

		return c.JSON(http.StatusOK, map[string]interface{}{"data": "Hello, World!"})
	})

	router.SetupRoute(e, es)

	for _, v := range e.Routes() {
		fmt.Println(v.Method, "\t", v.Path, "-", "\t", v.Name)
	}
	e.Logger.Fatal(e.Start("127.0.0.1:1323"))

}
