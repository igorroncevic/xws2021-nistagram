package grpc_common

import (
	"context"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckFollowInteraction(ctx context.Context, requestedUserId string, requestingUserId string) (*protopb.Follower, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckFollowInteraction")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	conn, err := CreateGrpcConnection(Recommendation_service_address)
	if err != nil{
		return nil, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	followerClient := GetFollowersClient(conn)
	followingResponse, err := followerClient.GetFollowersConnection(ctx, &protopb.Follower{
		UserId:                requestingUserId,
		FollowerId:            requestedUserId,
	})

	if err != nil {
		return nil, err
	}

	return followingResponse, err
}

func GetUsernameById(ctx context.Context, userId string) (string, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsernameById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	conn, err := GetClientConnection(Users_service_address)
	if err != nil{
		return "", status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	userClient := GetUsersClient(conn)
	response, err := userClient.GetUsernameById(ctx, &protopb.RequestIdUsers{Id: userId})
	if err != nil {
		return "", err
	}

	return response.Username, nil
}

func CheckIfPublicProfile(ctx context.Context, requestedUserId string) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIfPublicProfile")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	conn, err := GetClientConnection(Users_service_address)
	if err != nil{
		return false, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	privacyClient := GetPrivacyClient(conn)
	publicResponse, err := privacyClient.CheckUserProfilePublic(ctx, &protopb.PrivacyRequest{
		UserId: requestedUserId,
	})
	if err != nil { return false, err }

	return publicResponse.Response, nil
}

func CheckIfBlocked(ctx context.Context, requestedUserId string, requestingUserId string) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIfBlocked")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	conn, err := GetClientConnection(Users_service_address)
	if err != nil{
		return false, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	privacyClient := GetPrivacyClient(conn)
	blockedResponse, err := privacyClient.CheckIfBlocked(ctx, &protopb.CreateBlockRequest{
		Block: &protopb.Block{
			UserId:        requestingUserId,
			BlockedUserId: requestedUserId,
		},
	})
	if err != nil { return false, err }

	return blockedResponse.Response, nil
}

func GetHomepageUsers(ctx context.Context, userId string) ([]string, error){
	conn, err := CreateGrpcConnection(Recommendation_service_address)
	if err != nil{
		return []string{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	followerClient := GetFollowersClient(conn)
	followingResponse, err := followerClient.GetAllFollowingsForHomepage(ctx, &protopb.CreateUserRequestFollowers{
		User: &protopb.UserFollowers{ UserId: userId },
	})
	if err != nil{ return []string{}, status.Errorf(codes.Unknown, err.Error()) }

	userIds := []string{}
	for _, following := range followingResponse.Users{
		userIds = append(userIds, following.UserId)
	}

	privacyConn, err := CreateGrpcConnection(Users_service_address)
	if err != nil{
		return []string{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer privacyConn.Close()

	privacyClient := GetPrivacyClient(privacyConn)
	publicResponse, err := privacyClient.GetAllPublicUsers(ctx, &protopb.RequestIdPrivacy{Id: userId})
	if err != nil{ return []string{}, status.Errorf(codes.Unknown, err.Error()) }

	for _, publicUser := range publicResponse.Ids {
		found := false
		for _, userId := range userIds{
			if userId == publicUser {
				found = true
				break
			}
		}

		if !found {
			userIds = append(userIds, publicUser)
		}
	}

	return userIds, nil
}

func GetCloseFriends(ctx context.Context, userId string) ([]string, error){
	conn, err := CreateGrpcConnection(Recommendation_service_address)
	if err != nil{
		return []string{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	followerClient := GetFollowersClient(conn)
	closeFriends, err := followerClient.GetCloseFriends(ctx, &protopb.RequestIdFollowers{
		Id: userId,
	})
	if err != nil{ return []string{}, status.Errorf(codes.Unknown, err.Error()) }

	userIds := []string{}
	for _, closeFriend := range closeFriends.Users{
		userIds = append(userIds, closeFriend.UserId)
	}

	return userIds, nil
}

func GetPublicUsers(ctx context.Context) ([]string, error){
	conn, err := CreateGrpcConnection(Users_service_address)
	if err != nil{
		return []string{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	privacyClient := GetPrivacyClient(conn)
	publicUsers, err := privacyClient.GetAllPublicUsers(ctx, &protopb.RequestIdPrivacy{})
	if err != nil{ return []string{}, status.Errorf(codes.Unknown, err.Error()) }

	userIds := []string{}
	for _, userId := range publicUsers.Ids{
		userIds = append(userIds, userId)
	}

	return userIds, nil
}
