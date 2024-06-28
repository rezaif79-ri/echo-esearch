package service

import "github.com/rezaif79-ri/echo-esearch/domain"

type BookServiceES struct {
}

// Delete implements domain.BookService.
func (b *BookServiceES) Delete(bookID int) error {
	panic("unimplemented")
}

// Get implements domain.BookService.
func (b *BookServiceES) Get(bookID int) (domain.BookData, error) {
	panic("unimplemented")
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

func NewBookServiceES() domain.BookService {
	return &BookServiceES{}
}
