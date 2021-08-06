package model

import (
	protopb "github.com/igorroncevic/xws2021-nistagram/common/proto"
)

type User struct {
	UserId string
}

type Follower struct {
	UserId                       string
	FollowerId                   string
	IsMuted                      bool
	IsCloseFriends               bool
	IsApprovedRequest            bool
	RequestIsPending             bool
	IsMessageNotificationEnabled bool
	IsPostNotificationEnabled    bool
	IsStoryNotificationEnabled   bool
	IsCommentNotificationEnabled bool
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
		UserId:                       follower.UserId,
		FollowerId:                   follower.FollowerId,
		IsMuted:                      follower.IsMuted,
		IsMessageNotificationEnabled: follower.IsMessageNotificationEnabled,
		IsPostNotificationEnabled:    follower.IsPostNotificationEnabled,
		IsStoryNotificationEnabled:   follower.IsStoryNotificationEnabled,
		IsCommentNotificationEnabled: follower.IsCommentNotificationEnabled,
		IsApprovedRequest:            follower.IsApprovedRequest,
		RequestIsPending:             follower.RequestIsPending,
		IsCloseFriends:               follower.IsCloseFriends,
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
		UserId:                       f.UserId,
		FollowerId:                   f.FollowerId,
		IsMuted:                      f.IsMuted,
		IsCloseFriends:               f.IsCloseFriends,
		IsApprovedRequest:            f.IsApprovedRequest,
		IsMessageNotificationEnabled: f.IsMessageNotificationEnabled,
		IsPostNotificationEnabled:    f.IsPostNotificationEnabled,
		IsStoryNotificationEnabled:   f.IsStoryNotificationEnabled,
		IsCommentNotificationEnabled: f.IsCommentNotificationEnabled,
		RequestIsPending:             f.RequestIsPending,
	}
}
