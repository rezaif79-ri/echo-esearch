package domain

import responseutil "github.com/rezaif79-ri/echo-esearch/util/response_util"

type BookData struct {
	BookID string `json:"book_id"`
	Title  string `json:"title"`
	Pages  int    `json:"pages"`
	Author string `json:"author"`
}

type BookService interface {
	Get(bookID string) (BookData, responseutil.ControllerMeta)
	List(title string, sortID string) ([]BookData, responseutil.ControllerMeta)
	Insert(data BookData) (BookData, responseutil.ControllerMeta)
	Update(data BookData) (BookData, responseutil.ControllerMeta)
	Delete(bookID string) responseutil.ControllerMeta
	Count() (int, responseutil.ControllerMeta)
}
