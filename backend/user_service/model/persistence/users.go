package persistence

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type User struct {
	Id           string `gorm:"primaryKey"`
	FirstName    string
	LastName     string
	Email        string `gorm:"unique"`
	Username     string `gorm:"unique"`
	Password     string
	Role         model.UserRole
	BirthDate    time.Time
	ProfilePhoto string
	PhoneNumber  string
	Sex          string
	IsActive     bool
}

func (u User) ConvertToGrpc() *userspb.User {
	return &userspb.User{
		Id:           u.Id,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		Username:     u.Username,
		Password:     u.Password,
		Role:         u.Role.String(),
		Birthdate:    timestamppb.New(u.BirthDate),
		ProfilePhoto: u.ProfilePhoto,
		PhoneNumber:  u.PhoneNumber,
		Sex:          u.Sex,
		IsActive:     u.IsActive,
	}
}

func (u User) ConvertFromGrpc(user *userspb.User) *User {
	return &User{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Username:     user.Username,
		Password:     user.Password,
		Role:         model.GetUserRole(user.Role),
		BirthDate:    user.Birthdate.AsTime(),
		ProfilePhoto: user.ProfilePhoto,
		PhoneNumber:  user.PhoneNumber,
		Sex:          user.Sex,
		IsActive:     user.IsActive,
	}
}

type UserAdditionalInfo struct {
	Id        string `gorm:"primaryKey"`
	Biography string
	Website   string
	Category  model.UserCategory
}

func (u UserAdditionalInfo) ConvertFromGrpc(user *userspb.User) *UserAdditionalInfo {
	return &UserAdditionalInfo{
		Id:        user.Id,
		Biography: user.Biography,
		Website:   user.Website,
	}
}

type Privacy struct {
	UserId          string `gorm:"primaryKey"`
	IsProfilePublic bool
	IsDMPublic      bool
	IsTagEnabled    bool
}

type BlockedUsers struct {
	UserId        string `gorm:"primaryKey"`
	BlockedUserId string `gorm:"primaryKey"`
}

type Followers struct {
	UsedId                string `gorm:"primaryKey"`
	FollowerId            string `gorm:"primaryKey"`
	IsMuted               bool
	IsCloseFriend         bool
	IsApprovedRequest     bool
	IsNotificationEnabled bool
}

type VerificationRequest struct {
	UserId        string `gorm:"primaryKey"`
	FirstName     string
	LastName      string
	DocumentPhoto string
	IsApproved    bool
	CreatedAt     time.Time
}

type APIKeys struct {
	UserId string `gorm:"primaryKey"`
	APIKey string
}
