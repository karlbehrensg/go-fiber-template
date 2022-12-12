package requests

type UserRequest struct {
	Name                 string `form:"full_name" example:"Edwyn Rangel"`
	Email                string `form:"email" validate:"required,email" example:"edwyn.rangel.externo@zeleri.com"`
	EmailConfirmation    string `form:"email_confirmation" validate:"required,email,eqfield=Email" example:"edwyn.rangel.externo@zeleri.com"`
	Password             string `form:"password" validate:"required,min=7" example:"1234567"`
	PasswordConfirmation string `form:"password_confirmation" validate:"required,min=7,eqfield=Password" example:"1234567"`
}
