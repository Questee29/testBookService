package repository

import (
	"database/sql"
	"log"

	"github.com/Questee29/testBookService/models/author"
	"github.com/Questee29/testBookService/models/book"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (repo *Repository) GetBook(firstName string, lastName string) ([]book.BookDBResponse, error) {
	log.Printf("Finding book with author %s %s", firstName, lastName)
	var result []book.BookDBResponse

	//create prepared state
	query := `SELECT book_title,book_description FROM view_authorsbooks
	WHERE first_name=? AND last_name=?`

	//execute our query
	rows, err := repo.db.Query(query, firstName, lastName)
	if err != nil {
		return nil, err
	}

	//check every row
	for rows.Next() {
		book := new(book.BookDBResponse)
		if err := rows.Scan(&book.BookTitle, &book.BookDescription); err != nil {
			return nil, err
		}
		result = append(result, *book)
	}

	//if found any error ->return error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *Repository) GetAuthor(book_title string) ([]author.AuthorDBResponse, error) {
	var result []author.AuthorDBResponse
	query := `SELECT first_name,last_name FROM view_authorsbooks
	WHERE book_title=?`

	rows, err := repo.db.Query(query, book_title)
	if err != nil {
		return nil, nil
	}

	defer rows.Close()

	for rows.Next() {
		author := new(author.AuthorDBResponse)
		if err := rows.Scan(&author.FirstName, &author.LastName); err != nil {
			return nil, err
		}
		result = append(result, *author)
	}

	//if found any error ->return error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
