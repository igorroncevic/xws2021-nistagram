package domain

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Password struct {
	OldPassword    		string
	NewPassword    		string
	RepeatedPassword    string
	Id 					string
}

func (p Password) ConvertFromGrpc(pass *userspb.Password) Password {
	return Password{
		OldPassword: pass.OldPassword,
		NewPassword: pass.NewPassword,
		RepeatedPassword: pass.NewPassword,
		Id : pass.Id,
	}
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

func (u User) ConvertToGrpc() (*userspb.UsersDTO) {
	return &userspb.UsersDTO{
		Id:           u.Id,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		Username:     u.Username,
		Role:         u.Role.String(),
		Birthdate:    timestamppb.New(u.BirthDate),
		ProfilePhoto: u.ProfilePhoto,
		PhoneNumber:  u.PhoneNumber,
		Sex:          u.Sex,
		IsActive:     u.IsActive,
		Website:      u.Website,
		Biography:    u.Biography,
		Category: 	  model.GetUserCategoriesString(u.Category),
	}
}

func (u User) ConvertFromGrpc(user *userspb.UsersDTO) User {
	return User{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Username:     user.Username,
		Role:         model.GetUserRole(user.Role),
		BirthDate:    user.Birthdate.AsTime(),
		ProfilePhoto: user.ProfilePhoto,
		PhoneNumber:  user.PhoneNumber,
		Sex:          user.Sex,
		IsActive:     user.IsActive,
		Website: user.Website,
		Biography: user.Biography,
		Category: model.GetUserCategories(user.Category),
	}
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
