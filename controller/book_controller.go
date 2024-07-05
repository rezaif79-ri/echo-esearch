package controller

import (
	"net/http"

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
	id := c.Param("id")

	book, meta := b.bookService.Get(id)
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

func (b *BookController) Count(c echo.Context) error {
	count, meta := b.bookService.Count()
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
		echo.Map{"count": count},
	))
}

func (b *BookController) Insert(c echo.Context) error {
	var InsertBook domain.BookData
	if err := c.Bind(&InsertBook); err != nil {
		return c.JSON(http.StatusConflict, responseutil.Rest(
			http.StatusConflict,
			"Failed to bind request body",
			echo.Map{"error": err.Error()},
		))
	}

	data, meta := b.bookService.Insert(InsertBook)
	if meta.Error != nil {
		return c.JSON(meta.Status, responseutil.Rest(
			meta.Status,
			meta.Message,
			echo.Map{"error": meta.Error.Error()},
		))
	}

	return c.JSON(http.StatusCreated, responseutil.Rest(
		http.StatusCreated,
		"OK",
		data))
}

func (b *BookController) Update(c echo.Context) error {
	id := c.Param("id")

	var UpdateBook domain.BookData
	if err := c.Bind(&UpdateBook); err != nil {
		return c.JSON(http.StatusConflict, responseutil.Rest(
			http.StatusConflict,
			"Failed to bind request body",
			echo.Map{"error": err.Error()},
		))
	}
	UpdateBook.BookID = id

	data, meta := b.bookService.Update(UpdateBook)
	if meta.Error != nil {
		return c.JSON(meta.Status, responseutil.Rest(
			meta.Status,
			meta.Message,
			echo.Map{"error": meta.Error.Error()},
		))
	}

	return c.JSON(http.StatusCreated, responseutil.Rest(
		http.StatusCreated,
		"OK",
		data))
}

func (b *BookController) List(c echo.Context) error {
	data, meta := b.bookService.List(c.QueryParam("title"), c.QueryParam("sort_title"))
	if meta.Error != nil {
		return c.JSON(meta.Status, responseutil.Rest(
			meta.Status,
			meta.Message,
			echo.Map{"error": meta.Error.Error()},
		))
	}

	return c.JSON(http.StatusCreated, responseutil.Rest(
		http.StatusCreated,
		"OK",
		data))
}
