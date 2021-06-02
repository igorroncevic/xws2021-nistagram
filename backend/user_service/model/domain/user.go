package domain

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"time"
)

type User struct {
	Id           string `gorm:"primaryKey"`
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

func (u *User) GenerateUserDTO(user persistence.User, userAdditionalInfo persistence.UserAdditionalInfo) {
	u.Id = user.Id
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.Username = user.Username
	u.Role = user.Role
	u.BirthDate = user.BirthDate
	u.ProfilePhoto = user.ProfilePhoto
	u.PhoneNumber = user.PhoneNumber
	u.IsActive = user.IsActive
	u.Sex = user.Sex
	u.Biography = userAdditionalInfo.Biography
	u.Category = userAdditionalInfo.Category
	u.Website = userAdditionalInfo.Website
}
