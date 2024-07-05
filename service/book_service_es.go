package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
	"github.com/rezaif79-ri/echo-esearch/domain"
	elasticbindutil "github.com/rezaif79-ri/echo-esearch/util/elastic_bind_util"
	responseutil "github.com/rezaif79-ri/echo-esearch/util/response_util"
)

type BookServiceES struct {
	es *elasticsearch.Client
}

// Delete implements domain.BookService.
func (b *BookServiceES) Delete(bookID string) responseutil.ControllerMeta {
	_, err := b.es.Delete("echo_books", bookID)
	if err != nil {
		return responseutil.ControllerMeta{
			Status:  http.StatusInternalServerError,
			Error:   err,
			Message: "Encountered unexpected error",
		}
	}
	return responseutil.ControllerMeta{}
}

// Get implements domain.BookService.
func (b *BookServiceES) Get(bookID string) (domain.BookData, responseutil.ControllerMeta) {
	res, err := b.es.GetSource("echo_books", bookID)

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
	uID, err := uuid.NewV7()
	if err != nil {
		return data, responseutil.ControllerMeta{
			Status:  http.StatusConflict,
			Error:   err,
			Message: "Failed to generate book id",
		}
	}

	data.BookID = uID.String()
	bdata, err := json.Marshal(data)
	if err != nil {
		return data, responseutil.ControllerMeta{
			Status:  http.StatusConflict,
			Error:   err,
			Message: "Failed to create insert request",
		}
	}
	body := bytes.NewReader(bdata)

	res, err := b.es.Create("echo_books", data.BookID, body)
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
func (b *BookServiceES) List(title string, sortTitle string) ([]domain.BookData, responseutil.ControllerMeta) {
	var searchQuery []string

	if title != "" {
		searchQuery = append(searchQuery, `
			"query":{
				"fuzzy":{
					"title": "`+title+`"
				}
			}`)
	}

	if sortTitle = strings.ToLower(sortTitle); sortTitle == "asc" || sortTitle == "desc" && sortTitle != "" {
		searchQuery = append(searchQuery, `
			"sort" : [
				{ "title.keyword" :"`+sortTitle+`" }
			]
		`)
	}
	res, err := b.es.Search(
		b.es.Search.WithIndex("echo_books"),
		b.es.Search.WithBody(strings.NewReader(`
		{
			`+strings.Join(searchQuery, ",")+`
		}`)),
	)

	fmt.Println(`
		{
			` + strings.Join(searchQuery, ",") + `
		}`)
	if err != nil {
		return nil, responseutil.ControllerMeta{
			Status:  res.StatusCode,
			Error:   err,
			Message: "Failed to query books",
		}
	}

	data, err := elasticbindutil.HandleAndDecodeResponse[domain.EsSearchResponse[domain.BookData]](res.StatusCode, res.Body)
	if err != nil {
		return nil, responseutil.ControllerMeta{
			Status:  res.StatusCode,
			Error:   err,
			Message: "Failed to search books",
		}
	}

	var books []domain.BookData
	for _, v := range data.Hits.Hits {
		books = append(books, v.Source)
	}

	return books, responseutil.ControllerMeta{}
}

// Update implements domain.BookService.
func (b *BookServiceES) Update(data domain.BookData) (domain.BookData, responseutil.ControllerMeta) {
	var script = map[string]interface{}{
		"doc": data,
	}
	bdata, err := json.Marshal(script)
	if err != nil {
		return data, responseutil.ControllerMeta{
			Status:  http.StatusConflict,
			Error:   err,
			Message: "Failed to create insert request",
		}
	}
	body := bytes.NewReader(bdata)

	res, err := b.es.Update("echo_books", data.BookID, body)
	if err != nil {
		return data, responseutil.ControllerMeta{
			Status:  res.StatusCode,
			Error:   err,
			Message: "Failed to insert books",
		}
	}

	return data, responseutil.ControllerMeta{}
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
