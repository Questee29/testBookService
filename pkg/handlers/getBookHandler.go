package handlers

import (
	"encoding/json"
	"net/http"

	author "github.com/Questee29/testBookService/models/author"
	"github.com/Questee29/testBookService/models/book"
)

type Service interface {
	GetBook(firstName string, lastName string) ([]book.BookDBResponse, error)
	GetAuthor(book_title string) ([]author.AuthorDBResponse, error)
}

type GetBookHandler struct {
	service Service
}

func NewHandler(service Service) *GetBookHandler {
	return &GetBookHandler{
		service: service,
	}
}

func (handler *GetBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input author.AuthorInputDetails
	//reads json from user
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	booksTitles, err := handler.service.GetBook(input.FirstName, input.LastName)
	if err != nil {
		http.Error(w, "No such books", http.StatusBadRequest)
	}
	if err := json.NewEncoder(w).Encode(booksTitles); err != nil {
		http.Error(w, "error during parsing", http.StatusInternalServerError)
	}

}
