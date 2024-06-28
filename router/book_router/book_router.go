package bookrouter

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"
	"github.com/rezaif79-ri/echo-esearch/controller"
	"github.com/rezaif79-ri/echo-esearch/service"
)

func Route(e *echo.Group, es *elasticsearch.Client) {
	bookService := service.NewBookServiceES(es)
	bookController := controller.NewBookController(bookService)

	e.GET("/:id", bookController.GetBookByID)
}
