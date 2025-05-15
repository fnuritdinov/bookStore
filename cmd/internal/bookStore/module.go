package bookStore

type Service interface {
	AddBook(book Book) (Book, error)
}

type service struct {
	books  []Book
	nextID int
}

func New() Service {
	return &service{
		books:  make([]Book, 0),
		nextID: 1,
	}
}
