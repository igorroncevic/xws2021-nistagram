package persistence

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"time"
)

type User struct {
	Id           string `gorm:"primaryKey"`
	FirstName    string
	LastName     string
	Email        string
	Username     string `gorm:"unique"`
	Password     string
	Role         model.UserRole
	BirthDate    time.Time
	ProfilePhoto string
	PhoneNumber  string
	Sex          string
	IsActive     bool
}

type UserAdditionalInfo struct {
	Id        string `gorm:"primaryKey"`
	Biography string
	Website   string
	Category  model.UserCategory
}

type Privacy struct {
	UserId string `gorm:"primaryKey"`
	IsProfilePublic bool
	IsDMPublic bool
	IsTagEnabled bool
}

type BlockedUsers struct {
	UserId string `gorm:"primaryKey"`
	BlockedUserId string `gorm:"primaryKey"`
}

type Followers struct {
	UsedId string `gorm:"primaryKey"`
	FollowerId string `gorm:"primaryKey"`
	IsMuted bool
	IsCloseFriend bool
	IsApprovedRequest bool
	IsNotificationEnabled bool
}

type VerificationRequest struct {
	UserId string `gorm:"primaryKey"`
	FirstName string
	LastName string
	DocumentPhoto string
	IsApproved bool
	CreatedAt time.Time
}

type APIKeys struct {
	UserId string `gorm:"primaryKey"`
	APIKey string
}