package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Questee29/testBookService/models/book"
)

type GetAuthorHandler struct {
	service Service
}

func NewGetAuthorHandler(service Service) *GetAuthorHandler {
	return &GetAuthorHandler{
		service: service,
	}
}

func (handler *GetAuthorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input book.BookUnputDetail
	//reads json from user
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authors, err := handler.service.GetAuthor(input.BookTitle)
	if err != nil {
		http.Error(w, "no such books in the database", http.StatusBadRequest)
	}

	if err := json.NewEncoder(w).Encode(authors); err != nil {
		http.Error(w, "error during parsing", http.StatusInternalServerError)
	}
}
