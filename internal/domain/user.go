package domain

import (
	"time"

	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
	"github.com/karlbehrensg/go-fiber-template/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"id;primary_key"`
	Name      string `gorm:"full_name"`
	Email     string `gorm:"email;not null;unique"`
	Password  string `gorm:"password;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// UserRepository port secondary
type UserRepository interface {
	SaveUser(*User) *errs.AppError
	FindUserByEmail(string) (*User, *errs.AppError)
}

// HashPassword encrypt password
func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// ComparePassword validate password
func (u *User) ComparePassword(password string) *errs.AppError {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		logger.Error(err.Error())
		return errs.NewAuthenticationError("Invalid credentials")
	}
	return nil
}

// ToNewUtilsJWTClaims convert User struct to utils.JWTClaims struct
func (u *User) ToNewUtilsJWTClaims() *utils.JWTClaims {
	return &utils.JWTClaims{
		UserID: u.ID,
		Email:  u.Email,
		Name:   u.Name,
	}
}
