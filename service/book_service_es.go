package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
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
	res, err := b.es.Count(func(r *esapi.CountRequest) {
		r.Index = append(r.Index, "echo_books")
	})

	if err != nil {
		return data, responseutil.ControllerMeta{
			Status:  http.StatusInternalServerError,
			Error:   err,
			Message: "Encountered unexpected error",
		}
	}

	type Count struct {
		Count int `json:"count"`
	}
	count, err := elasticbindutil.HandleAndDecodeResponse[Count](res.StatusCode, res.Body)
	if err != nil {
		return data, responseutil.ControllerMeta{
			Status:  res.StatusCode,
			Error:   err,
			Message: "Failed to fetch document count",
		}
	}

	bdata, err := json.Marshal(data)
	if err != nil {
		return data, responseutil.ControllerMeta{
			Status:  http.StatusConflict,
			Error:   err,
			Message: "Failed to create insert request",
		}
	}
	body := bytes.NewReader(bdata)

	res, err = b.es.Create("echo_books", fmt.Sprint(count), body)
	if err != nil {
		return data, responseutil.ControllerMeta{
			Status:  res.StatusCode,
			Error:   err,
			Message: "Failed to insert books",
		}
	}

	return data, responseutil.ControllerMeta{}
}

// List implements domain.BookService.
func (b *BookServiceES) List(title string, sortID string) ([]domain.BookData, responseutil.ControllerMeta) {
	panic("unimplemented")
}

// Update implements domain.BookService.
func (b *BookServiceES) Update(data domain.BookData) (domain.BookData, responseutil.ControllerMeta) {
	panic("unimplemented")
}

// Count implements domain.BookService.
func (b *BookServiceES) Count() (int, responseutil.ControllerMeta) {

	res, err := b.es.Count(func(r *esapi.CountRequest) {
		r.Index = append(r.Index, "echo_books")
	})

	if err != nil {
		return 0, responseutil.ControllerMeta{
			Status:  http.StatusInternalServerError,
			Error:   err,
			Message: "Encountered unexpected error",
		}
	}

	type Count struct {
		Count int `json:"count"`
	}
	count, err := elasticbindutil.HandleAndDecodeResponse[Count](res.StatusCode, res.Body)
	if err != nil {
		return 0, responseutil.ControllerMeta{
			Status:  res.StatusCode,
			Error:   err,
			Message: "Failed to fetch document count",
		}
	}
	return count.Count, responseutil.ControllerMeta{}
}

func NewBookServiceES(es *elasticsearch.Client) domain.BookService {
	return &BookServiceES{es}
}
