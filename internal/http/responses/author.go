package responses

type AuthorResponse struct {
	Id       uint   `json:"id" example:"30"`
	FullName string `json:"full_name" example:"J. J. Benítez"`
}
