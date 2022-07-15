package handler

import (
	"context"
	"log"

	"github.com/Questee29/testBookService/models/author"
	"github.com/Questee29/testBookService/models/book"
	pb "github.com/Questee29/testBookService/proto/protob"
)

type BooksService interface {
	GetBook(ctx context.Context, firstName, lastName string) ([]book.BookDBResponse, error)
	GetAuthor(ctx context.Context, bookTitle string) ([]author.AuthorDBResponse, error)
}

type GrpcHandler struct {
	pb.UnimplementedBooksServiceServer
	service BooksService
}

func NewGrpcHandler(service BooksService) *GrpcHandler {
	return &GrpcHandler{
		service: service,
	}
}

func (handler *GrpcHandler) GetBook(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	uReq := author.AuthorInputDetails{
		FirstName: req.GetAuthorFirstName(),
		LastName:  req.GetAuthorLastName(),
	}

	//Get book from service
	books, err := handler.service.GetBook(ctx, uReq.FirstName, uReq.LastName)
	if err != nil {
		log.Println("error during searching books")
		return nil, err
	}
	return toPBBookModel(books), nil
}

func (handler *GrpcHandler) GetAuthor(ctx context.Context, req *pb.AuthorRequest) (*pb.AuthorResponse, error) {
	uReq := book.BookUnputDetail{
		BookTitle: req.GetBookTitle(),
	}

	//get author from service
	authors, err := handler.service.GetAuthor(ctx, uReq.BookTitle)
	if err != nil {
		log.Println("error during searching author")
		return nil, err
	}
	log.Println(authors)
	return toPBAuthorModel(authors), nil
}

//convert to protobuf model
func toPBAuthorModel(authors []author.AuthorDBResponse) *pb.AuthorResponse {
	var booksResponse []*pb.AuthorInfo
	for _, book := range authors {
		authorPB := &pb.AuthorInfo{AuthorFirstName: book.FirstName, AuthorLastName: book.LastName}
		booksResponse = append(booksResponse, authorPB)
	}
	return &pb.AuthorResponse{Author: booksResponse}
}
func toPBBookModel(books []book.BookDBResponse) *pb.BookResponse {
	var booksResponse []*pb.BookInfo
	for _, book := range books {
		bookPB := &pb.BookInfo{BookTitle: book.BookTitle, BookDescription: book.BookDescription}
		booksResponse = append(booksResponse, bookPB)
	}
	return &pb.BookResponse{Book: booksResponse}
}
