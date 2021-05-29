package model

import "time"

type Product struct {
	Id string `gorm:"primaryKey"`
	Name string
	Price float32
	IsActive bool
	Quantity int
	PhotoLink string
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
	ShippingDate time.Time //TODO
	Referral     int
}

type OrderProducts struct{
	OrderId string `gorm:"primaryKey"`
	ProductId string `gorm:"primaryKey"`
}

type CommonReport struct{
	Id string `gorm:"primaryKey"`
	CampaignId string
	PlacementNum int
	LinkClickNum int
	Revenue float32
}

type PostReport struct {
	CommonReport
	LikesNum int
	DislikesNum int
	CommentsNum int
}

type StoryReport struct{
	CommonReport
	ViewsNum int
}

type User struct{
	Id           string `gorm:"primaryKey"`
	FirstName    string
	LastName     string
	Email        string
	Username     string
	Role         UserRole
	BirthDate    time.Time // TODO
	ProfilePhoto string
	PhoneNumber  string
	Sex          string
	IsActive     bool
}

type UserRole string
const (
	Basic UserRole = "Basic"
	Agent           = "Agent"
)
