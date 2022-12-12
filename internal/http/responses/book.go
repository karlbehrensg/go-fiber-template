package responses

type BookResponse struct {
	Id              uint   `json:"id" example:"1"`
	Title           string `json:"title" example:"Caballo de Troya 1"`
	PublicationYear string `json:"publication_year" example:"1984"`
	AuthorID        uint   `json:"author_id" example:"30"`
	AuthorName      string `json:"author_name" example:"J. J. Benítez"`
}
