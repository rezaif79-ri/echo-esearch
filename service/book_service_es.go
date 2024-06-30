package service

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rezaif79-ri/echo-esearch/domain"
	elasticbindutil "github.com/rezaif79-ri/echo-esearch/util/elastic_bind_util"
	responseutil "github.com/rezaif79-ri/echo-esearch/util/response_util"
)

type BookServiceES struct {
	es *elasticsearch.Client
}

// Delete implements domain.BookService.
func (b *BookServiceES) Delete(bookID int) responseutil.ControllerMeta {
	panic("unimplemented")
}

// Get implements domain.BookService.
func (b *BookServiceES) Get(bookID int) (domain.BookData, responseutil.ControllerMeta) {
	res, err := b.es.GetSource("echo_books", fmt.Sprint(bookID))

	if err != nil {
		return domain.BookData{}, responseutil.ControllerMeta{
			Status:  http.StatusInternalServerError,
			Error:   err,
			Message: "Encountered unexpected error",
		}
	}

	data, err := elasticbindutil.HandleAndDecodeResponse[domain.BookData](res.StatusCode, res.Body)

	return data, responseutil.ControllerMeta{
		Status:  res.StatusCode,
		Error:   err,
		Message: "Failed to fetch requested document",
	}
}

// Insert implements domain.BookService.
func (b *BookServiceES) Insert(data domain.BookData) (domain.BookData, responseutil.ControllerMeta) {
	panic("unimplemented")
}

// List implements domain.BookService.
func (b *BookServiceES) List(title string, sortID string) ([]domain.BookData, responseutil.ControllerMeta) {
	panic("unimplemented")
}

// Update implements domain.BookService.
func (b *BookServiceES) Update(data domain.BookData) (domain.BookData, responseutil.ControllerMeta) {
	panic("unimplemented")
}

func NewBookServiceES(es *elasticsearch.Client) domain.BookService {
	return &BookServiceES{es}
}
