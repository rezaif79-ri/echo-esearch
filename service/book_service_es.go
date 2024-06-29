package service

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rezaif79-ri/echo-esearch/domain"
	elasticbindutil "github.com/rezaif79-ri/echo-esearch/util/elastic_bind_util"
)

type BookServiceES struct {
	es *elasticsearch.Client
}

// Delete implements domain.BookService.
func (b *BookServiceES) Delete(bookID int) error {
	panic("unimplemented")
}

// Get implements domain.BookService.
func (b *BookServiceES) Get(bookID int) (domain.BookData, error) {
	res, err := b.es.GetSource("echo_books", fmt.Sprint(bookID))

	if err != nil {
		return domain.BookData{}, err
	}

	return elasticbindutil.HandleAndDecodeResponse[domain.BookData](res.StatusCode, res.Body)
}

// Insert implements domain.BookService.
func (b *BookServiceES) Insert(data domain.BookData) (domain.BookData, error) {
	panic("unimplemented")
}

// List implements domain.BookService.
func (b *BookServiceES) List(title string, sortID string) ([]domain.BookData, error) {
	panic("unimplemented")
}

// Update implements domain.BookService.
func (b *BookServiceES) Update(data domain.BookData) (domain.BookData, error) {
	panic("unimplemented")
}

func NewBookServiceES(es *elasticsearch.Client) domain.BookService {
	return &BookServiceES{es}
}
