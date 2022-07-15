package service

import (
	"github.com/Questee29/testBookService/models/author"
	"github.com/Questee29/testBookService/models/book"
)

type Repository interface {
	GetAuthor(book_title string) ([]author.AuthorDBResponse, error)
	GetBook(firstName string, lastName string) ([]book.BookDBResponse, error)
}
type service struct {
	repository Repository
}

func NewRestService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (service *service) GetBook(firstName string, lastName string) ([]book.BookDBResponse, error) {
	return service.repository.GetBook(firstName, lastName)
}
func (service *service) GetAuthor(book_title string) ([]author.AuthorDBResponse, error) {
	return service.repository.GetAuthor(book_title)
}
