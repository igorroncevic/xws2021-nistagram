package persistence

import (
	"github.com/david-drvar/xws2021-nistagram/agent_application/model"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"time"
)

type Product struct {
	Id       string `gorm:"primaryKey"`
	Name     string
	Price    float32
	IsActive bool
	Quantity int
	Photo    string
	AgentId  string
}

type PaymentType string

const (
	Cash      PaymentType = "Cash"
	PayPal                = "PayPal"
	Cripto                = "Cripto"
	DebitCard             = "DebitCard"
)

type Order struct {
	Id           string `gorm:"primaryKey"`
	UserId       string
	PaymentType  PaymentType
	ShippingDate time.Time
	Referral     int
}

type OrderProducts struct {
	OrderId   string `gorm:"primaryKey"`
	ProductId string `gorm:"primaryKey"`
}

type CommonReport struct {
	Id           string `gorm:"primaryKey"`
	CampaignId   string
	PlacementNum int
	LinkClickNum int
	Revenue      float32
}

type PostReport struct {
	CommonReport
	LikesNum    int
	DislikesNum int
	CommentsNum int
}

type StoryReport struct {
	CommonReport
	ViewsNum int
}

type User struct {
	Id           string `gorm:"primaryKey"`
	FirstName    string
	LastName     string
	Email        string `gorm:"unique"`
	Username     string `gorm:"unique"`
	Role         model.UserRole
	BirthDate    time.Time
	ProfilePhoto string
	PhoneNumber  string
	Sex          string
	Password     string
	IsActive     bool
}

func (u *User) ConvertFromGrpc(user *protopb.UserAgentApp) *User {
	newUser := &User{
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

	return newUser
}

func (u *Product) ConvertFromGrpc(product *protopb.Product) *Product {
	newProduct := &Product{
		Id:       product.Id,
		Name:     product.Name,
		Price:    float32(product.Price),
		IsActive: product.IsActive,
		Quantity: int(product.Quantity),
		Photo:    product.Photo,
		AgentId:  product.AgentId,
	}

	return newProduct
}
