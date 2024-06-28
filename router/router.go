package router

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"
	responseutil "github.com/rezaif79-ri/echo-esearch/util/response_util"
)

func SetupRoute(e *echo.Echo, es *elasticsearch.Client) {
	api := e.Group("api")
	api.GET("ok", func(c echo.Context) error {
		return c.JSON(http.StatusOK, responseutil.Rest(http.StatusOK, "OK", nil))
	})
}
