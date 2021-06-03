package model


import (
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/proto"
)


type User struct{
	UserId string
}

type Follower struct {
	UserId string
	FollowerId string
	IsMuted bool
	IsCloseFriends bool
	IsApprovedRequest bool
	IsNotificationEnabled bool
}

func (user *User) ConvertFromGrpc(u *proto.User) *User  {
	return &User {
		UserId: u.UserId,
	}
}

func (follower *Follower) ConvertFromGrpc(f *proto.Follower) *Follower {
	return &Follower{
		UserId: f.UserId,
		FollowerId: f.FollowerId,
		IsMuted: f.IsMuted,
		IsCloseFriends: f.IsCloseFriends,
		IsApprovedRequest: f.IsApprovedRequest,
		IsNotificationEnabled: f.IsNotificationEnabled,
	}
}

