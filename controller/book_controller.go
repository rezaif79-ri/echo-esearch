package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rezaif79-ri/echo-esearch/domain"
	responseutil "github.com/rezaif79-ri/echo-esearch/util/response_util"
)

type BookController struct {
	bookService domain.BookService
}

func NewBookController(bookService domain.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

func (b *BookController) GetBookByID(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, responseutil.Rest(
			http.StatusConflict,
			err.Error(),
			nil,
		))
	}

	book, meta := b.bookService.Get(bookID)
	if meta.Error != nil {
		return c.JSON(meta.Status, responseutil.Rest(
			meta.Status,
			meta.Message,
			echo.Map{"error": meta.Error.Error()},
		))
	}

	return c.JSON(http.StatusOK, responseutil.Rest(
		http.StatusOK,
		"OK",
		book,
	))
}
