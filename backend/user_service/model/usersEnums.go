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

func  ToString(c UserCategory) string{
	switch c {
	case Sports:
		return "Sports"
	case Influencer:
		return "Influencer"
	case News:
		return "News"
	case Business:
		return "Business"
	case Brand:
		return "Brand"
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

func  GetCategory(s string) UserCategory{
	switch s {
	case "Sports":
		return Sports
	case "Influencer":
		return Influencer
	case "News":
		return News
	case "Business":
		return Business
	case "Brand":
		return Brand
	case "Organization":
		return Organization
	case "Government":
		return Government
	case "NoCategory":
		return NoCategory
	default:
		return ""
	}
}

func GetUserCategories(categories string) UserCategory {
	return GetCategory(categories)
}

func GetUserCategoriesString(categories UserCategory) string {
	return ToString(categories)
}