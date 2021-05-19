package persistence

type User struct {
	Id           string
	FirstName    string
	LastName     string
	Email        string
	Username     string
	Role         UserRole
	BirthDate    string // TODO
	ProfilePhoto string
	PhoneNumber  string
	Sex          string
	IsActive     bool
}

type UserAdditionalInfo struct {
	Id        string
	Biography string
	Website   string
	Category  UserCategory
}

type Privacy struct {
	UserId string
	IsProfilePublic bool
	IsDMPublic bool
	IsTagEnabled bool
}

type BlockedUsers struct {
	UserId string
	BlockedUserId string
}

type Followers struct {
	UsedId string
	FollowerId string
	IsMuted bool
	IsCloseFriend bool
	IsApprovedRequest bool
	IsNotificationEnabled bool
}

type VerificationRequest struct {
	FirstName string
	LastName string
	DocumentPhoto string
	UserId string
	IsApproved bool
	DateCreated string
}

type APIKeys struct {
	UserId string
	APIKey string
}