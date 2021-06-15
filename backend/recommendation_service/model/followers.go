package model

import (
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
)

type User struct {
	UserId string
}

type Follower struct {
	UserId                string
	FollowerId            string
	IsMuted               bool
	IsCloseFriends        bool
	IsApprovedRequest     bool
	IsNotificationEnabled bool
}

func (user *User) ConvertFromGrpc(u *protopb.UserFollowers) *User {
	return &User{
		UserId: u.UserId,
	}
}

func (user *User) ConvertToGrpc() *protopb.UserFollowers {
	return &protopb.UserFollowers{
		UserId: user.UserId,
	}
}

func (follower *Follower) ConvertToGrpc() *protopb.Follower {
	return &protopb.Follower{
		UserId:                follower.UserId,
		FollowerId:            follower.FollowerId,
		IsMuted:               follower.IsMuted,
		IsNotificationEnabled: follower.IsNotificationEnabled,
		IsApprovedRequest:     follower.IsApprovedRequest,
		IsCloseFriends:        follower.IsCloseFriends,
	}
}

func (user *User) ConvertAllToGrpc(users []User) []*protopb.UserFollowers {
	protoUsers := []*protopb.UserFollowers{}
	for _, s := range users {
		protoUsers = append(protoUsers, s.ConvertToGrpc())
	}
	return protoUsers
}

func (follower *Follower) ConvertFromGrpc(f *protopb.Follower) *Follower {
	return &Follower{
		UserId:                f.UserId,
		FollowerId:            f.FollowerId,
		IsMuted:               f.IsMuted,
		IsCloseFriends:        f.IsCloseFriends,
		IsApprovedRequest:     f.IsApprovedRequest,
		IsNotificationEnabled: f.IsNotificationEnabled,
	}
}
