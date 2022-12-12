package requests

type BookRequest struct {
	Id              uint   `json:"id,omitempty"`
	Title           string `json:"title" validate:"required,min=3" example:"Caballo de Troya 1"`
	AuthorID        uint   `json:"author_id" validate:"required" example:"30"`
	PublicationYear string `json:"publication_year" validate:"required" example:"1984"`
}
