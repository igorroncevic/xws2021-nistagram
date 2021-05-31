package model

import "fmt"

type UserRole string
const (
	Basic    UserRole = "Basic"
	Admin             = "Admin"
	Verified          = "Verified"
	Agent             = "Agent"
)

func (r UserRole) String() string{
	switch r {
	case Basic:
		return "Basic"
	case Admin:
		return "Admin"
	case Verified:
		return "Verified"
	case Agent:
		return "Agent"
	default:
		return fmt.Sprintf("%s", string(r))
	}
}

func GetUserRole(r string) UserRole{
	switch r {
	case "Basic":
		return Basic
	case "Admin":
		return Admin
	case "Verified":
		return Verified
	case "Agent":
		return Agent
	default:
		return ""
	}
}

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

func (c UserCategory) String() string{
	switch c {
	case Sports:
		return "Sports"
	case Influencer:
		return "Influencer"
	case News:
		return "News"
	case Business:
		return "Business"
	case Organization:
		return "Organization"
	case Government:
		return "Government"
	case NoCategory:
		return "NoCategory"
	default:
		return fmt.Sprintf("%s", string(c))
	}
}