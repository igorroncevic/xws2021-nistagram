package persistence

import "time"

type Post struct{
	Id          string `gorm:"primaryKey"`
	UserId      string
	IsAd        bool
	Type        PostType
	Description string
	Location    string
	CreatedAt time.Time
}

type Story struct{
	Post
	IsCloseFriends bool
}

type Media struct{
	Id      string `gorm:"primaryKey"`
	Type    MediaType
	PostId  string
	Content string
}

type Tag struct {
	UserId string `gorm:"primaryKey"`
	MediaId string `gorm:"primaryKey"`
}

type Collection struct {
	Id string `gorm:"primaryKey"`
	Name string
	UserId string
}

type Favorites struct {
	PostId string `gorm:"primaryKey"`
	UserId string `gorm:"primaryKey"`
	CollectionId string
}

type PostLikes struct {
	PostId string `gorm:"primaryKey"`
	UserId string `gorm:"primaryKey"`
	IsLike bool
}

type PostComments struct {
	PostId string `gorm:"primaryKey"`
	UserId string `gorm:"primaryKey"`
	Content string
	CreatedAt time.Time //TODO
}

type HighLights struct {
	Id string `gorm:"primaryKey"`
	UserId string
	Name string
}

type RegistrationRequest struct {
	Id        string `gorm:"primaryKey"`
	UserId    string
	CreatedAt time.Time //TODO
	Status    RequestStatus
}

type Ad struct {
	Id string `gorm:"primaryKey"`
	Link string
	CampaignId string
	PostId string
	LinkClickNum int
}

type Campaign struct {
	Id string `gorm:"primaryKey"`
	IsOneTime bool
	StartDate time.Time //TODO
	EndDate time.Time //TODO
	PlacementNum int
	AgentId string
	IdAdCategory string
	LastUpdated time.Time //TODO
}

type CampaignInfluencerRequest struct {
	CampaignId   string `gorm:"primaryKey"`
	InfluencerId string `gorm:"primaryKey"`
	Status       RequestStatus
}

type ContentComplaint struct {
	Id       string `gorm:"primaryKey"`
	Category ComplaintCategory
	PostId   string
	Status   RequestStatus
}

type AdCategory struct {
	Id string `gorm:"primaryKey"`
	Name string
}

type UserAdCategories struct {
	UserId string `gorm:"primaryKey"`
	IdAdCategory string `gorm:"primaryKey"`
}

type CampaignChanges struct {
	CampaignId string `gorm:"primaryKey"`
	AdCategoryId string
	StartDate time.Time
	EndDate time.Time
	PlacementNum int
}



