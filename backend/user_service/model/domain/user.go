package domain

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"time"
)

type Password struct {
	OldPassword      string
	NewPassword      string
	RepeatedPassword string
	Id               string
}

type User struct {
	Id           string
	FirstName    string
	LastName     string
	Email        string
	Username     string
	Role         model.UserRole
	BirthDate    time.Time
	ProfilePhoto string
	PhoneNumber  string
	Sex          string
	IsActive     bool
	Biography    string
	Website      string
	Category     model.UserCategory
}

type LoginRequest struct{
	Email	 string
	Password string
}
