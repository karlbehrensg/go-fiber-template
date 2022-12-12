package requests

type LoginRequest struct {
	Email    string `form:"email" validate:"required,email" example:"edwyn.rangel.externo@zeleri.com"`
	Password string `form:"password" validate:"required,min=7" example:"1234567"`
}
