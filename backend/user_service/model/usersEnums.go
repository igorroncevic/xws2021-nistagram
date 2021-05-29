package persistence

type UserRole string
const (
	Basic    UserRole = "Basic"
	Admin             = "Admin"
	Verified          = "Verified"
	Agent             = "Agent"
)

type UserCategory string
const (
	Sports       UserCategory = "Sports"
	Influencer                = "Influencer"
	News                      = "News"
	Brand                     = "Brand"
	Business                  = "Business"
	Organization              = "Organization"
	Government                = "Government"
	NoCategory                = "NoCategory"
)