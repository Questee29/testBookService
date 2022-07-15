package author

type Author struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type AuthorInputDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type AuthorDBResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
