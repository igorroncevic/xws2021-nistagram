package domain

import (
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type LoginRequest struct {
	Email    string
	Password string
}

type Order struct {
	Id          string `gorm:"primaryKey"`
	UserId      string
	ProductId   string
	Quantity    int
	DateCreated time.Time
	TotalPrice  float32
	Referral    int
	Username    string
	ProductName string
}

func (o Order) ConvertToGrpc() *protopb.Order {
	return &protopb.Order{
		Id:          o.Id,
		UserId:      o.UserId,
		ProductId:   o.ProductId,
		Quantity:    int32(o.Quantity),
		Referral:    int32(o.Referral),
		DateCreated: timestamppb.New(o.DateCreated),
		TotalPrice:  o.TotalPrice,
		Username:    o.Username,
		ProductName: o.ProductName,
	}
}

type CampaignStats struct {
	Id 			 string
	Name		 string
	IsOneTime 	 bool
	StartDate 	 time.Time
	EndDate 	 time.Time
	StartTime	 int
	EndTime		 int
	Placements   int
	Category	 string
	Type		 string
	Influencers	 []InfluencerStats
	Likes	 	 int
	Dislikes 	 int
	Comments 	 int
	Clicks   	 int
}

type AdStats struct {
	Id 		 string
	Media    []string
	Type	 string
	Hashtags []string
	Location string
	Likes	 int
	Dislikes int
	Comments int
	Clicks   int
}

type InfluencerStats struct {
	Id				string
	Username		string
	Ads			    []AdStats
	TotalLikes		int
	TotalDislikes	int
	TotalComments	int
	TotalClicks		int
}