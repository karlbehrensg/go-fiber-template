package requests

type AuthorRequest struct {
	Id       uint   `json:"id,omitempty"`
	FullName string `json:"full_name" validate:"required,min=3" example:"J. J. Benítez"`
}
