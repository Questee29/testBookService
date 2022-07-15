package service

import (
	"context"

	"github.com/Questee29/testBookService/models/author"
	"github.com/Questee29/testBookService/models/book"
	pb "github.com/Questee29/testBookService/proto/protob"
)

type ServiceGrpc struct {
	pb.UnimplementedBooksServiceServer
	repository Repository
}

func NewGrpcSerivce(repository Repository) *ServiceGrpc {
	return &ServiceGrpc{
		repository: repository,
	}
}

func (service *ServiceGrpc) GetBook(ctx context.Context, firstName, lastName string) ([]book.BookDBResponse, error) {
	return nil, nil
}

func (service *ServiceGrpc) GetAuthor(ctx context.Context, bookTitle string) ([]author.AuthorDBResponse, error) {
	return nil, nil
}
