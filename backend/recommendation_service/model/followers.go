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

func (user *User) ConvertToGrpc() *proto.User {
	return &proto.User{
		UserId: user.UserId,
	}
}

func (follower *Follower) ConvertToGrpc() *proto.Follower {
	return &proto.Follower{
		UserId: follower.UserId,
		FollowerId: follower.FollowerId,
		IsMuted: follower.IsMuted,
		IsNotificationEnabled: follower.IsNotificationEnabled,
		IsApprovedRequest: follower.IsApprovedRequest,
		IsCloseFriends: follower.IsCloseFriends,
	}
}


func (user *User) ConvertAllToGrpc(users []User) []*proto.User{
	var protoUsers []*proto.User
	for _, s := range users {
		protoUsers = append(protoUsers, s.ConvertToGrpc())
	}
	return protoUsers
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

