package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rezaif79-ri/echo-esearch/domain"
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
	res, err := b.es.Get("echo_books", fmt.Sprint(bookID))
	if err != nil {
		return domain.BookData{}, err
	}

	if res.IsError() {
		return domain.BookData{}, errors.New(strings.Join(res.Warnings(), ","))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	fmt.Println(result)

	return domain.BookData{}, nil
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
