package domain

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	"time"
)

type Objava struct{
	Id          string
	UserId      string
	IsAd        bool
	Type        model.PostType
	Description string
	Location    string
	Tags		[]Tag
	CreatedAt   time.Time
	Media	    []Media
}

type Post struct{
	Objava
	Comments 	[]Comment
	Likes 		[]Like
	Dislikes 	[]Like
}

type ReducedPost struct{
	Objava
	CommentsNum 	int32
	LikesNum 		int32
	DislikesNum 	int32
}

type Story struct{
	Objava
	IsCloseFriends bool
}

type Media struct{
	Id      	string
	Type    	model.MediaType
	PostId  	string
	Content 	string
	OrderNum	int32
}

type Tag struct {
	PostId 		string
	UserId 		string
	Username	string
}

type Collection struct {
	Id 			string
	Name 		string
	UserId 		string
	Posts 		[]Post
}

type Favorites struct {
	UserId 			string
	Collections 	[]Collection
	Unclassified 	[]Post
}

type Like struct {
	PostId string
	UserId string
	IsLike bool
}

type Comment struct {
	Id			string
	PostId 		string
	UserId 		string
	Username 	string
	Content 	string
	CreatedAt 	time.Time
}

type HighLight struct {
	Id 		string
	Name 	string
	Stories []Story
}

/*
type RegistrationRequest struct {
	Id        string
	UserId    string
	CreatedAt time.Time
	Status    model.RequestStatus
	Username string
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
	Status       model.RequestStatus
}

type ContentComplaint struct {
	Id       string `gorm:"primaryKey"`
	Category model.ComplaintCategory
	PostId   string
	Status   model.RequestStatus
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
*/




