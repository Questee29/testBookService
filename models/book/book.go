package book

type Book struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BookUnputDetail struct {
	BookTitle string `json:"book_title"`
}
type BookDBResponse struct {
	BookTitle       string `json:"book_title"`
	BookDescription string `json:"book_description"`
}
