package domain

type BookData struct {
	BookID int    `json:"book_id"`
	Title  string `json:"title"`
	Pages  int    `json:"pages"`
	Author string `json:"author"`
}

type BookService interface {
	Get(bookID int) (BookData, error)
	List(title string, sortID string) ([]BookData, error)
	Insert(data BookData) (BookData, error)
	Update(data BookData) (BookData, error)
	Delete(bookID int) error
}
