package persistence

type Product struct {
	Id string
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
	Id           string
	UserId       string
	PaymentType  PaymentType
	ShippingDate string //TODO
	Referral     int
}

type OrderProducts struct{
	OrderId string
	ProductId string
}

type CommonReport struct{
	Id string
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

type UserM struct{
	Id           string
	FirstName    string
	LastName     string
	Email        string
	Username     string
	Role         UserRoleM
	BirthDate    string // TODO
	ProfilePhoto string
	PhoneNumber  string
	Sex          string
	IsActive     bool
}

type UserRoleM string
const (
	BasicM UserRoleM = "Basic"
	AgentM           = "Agent"
)
