package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type StoryGrpcController struct {
	service    *services.StoryService
	jwtManager *common.JWTManager
}

func NewStoryController(db *gorm.DB, jwtManager *common.JWTManager) (*StoryGrpcController, error) {
	service, err := services.NewStoryService(db)
	if err != nil {
		return nil, err
	}

	return &StoryGrpcController{
		service,
		jwtManager,
	}, nil
}

func (c *StoryGrpcController) GetStoriesForUser(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.StoriesArray, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoriesForUser")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == "" {
		return &protopb.StoriesArray{}, status.Errorf(codes.InvalidArgument, "no user id is provided")
	}

	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Recommendation_service_address)
	if err != nil{
		return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	followerClient := grpc_common.GetFollowersClient(conn)
	followingResponse, err := followerClient.GetFollowersConnection(ctx, &protopb.Follower{
		UserId:      claims.UserId,
		FollowerId:  userId,
	})

	privacyConn, err := grpc_common.CreateGrpcConnection(grpc_common.Users_service_address)
	if err != nil{
		return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer privacyConn.Close()

	privacyClient := grpc_common.GetPrivacyClient(conn)
	publicResponse, err := privacyClient.CheckUserProfilePublic(ctx, &protopb.PrivacyRequest{
		UserId: userId,
	})

	if !followingResponse.IsApprovedRequest || !publicResponse.Response {
		return &protopb.StoriesArray{}, nil
	}

	stories, err := c.service.GetStoriesForUser(ctx, userId, followingResponse.IsCloseFriends)
	if err != nil{
		return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responseStories := domain.ConvertMultipleStoriesToGrpc(stories)

	return &protopb.StoriesArray{
		Stories: responseStories,
	}, nil
}

func (c *StoryGrpcController) GetAllHomeStories(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.StoriesHome, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHomeStories")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == "" {
		return &protopb.StoriesHome{}, status.Errorf(codes.InvalidArgument, "no user id is provided")
	}

	userId := claims.UserId
	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Recommendation_service_address)
	if err != nil{
		return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	followerClient := grpc_common.GetFollowersClient(conn)
	followingResponse, err := followerClient.GetAllFollowingsForHomepageStories(ctx, &protopb.CreateUserRequestFollowers{
		User: &protopb.UserFollowers{ UserId: userId },
	})

	if len(followingResponse.Users) == 0 {
		return &protopb.StoriesHome{}, nil
	}

	userIds := []string{}
	for _, following := range followingResponse.Users{
		userIds = append(userIds, following.UserId)
	}

	privacyConn, err := grpc_common.CreateGrpcConnection(grpc_common.Users_service_address)
	if err != nil{
		return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer privacyConn.Close()

	privacyClient := grpc_common.GetPrivacyClient(conn)
	publicResponse, err := privacyClient.GetAllPublicUsers(ctx, &protopb.EmptyRequestPrivacy{})

	if len(publicResponse.Ids) == 0 {
		return &protopb.StoriesHome{}, nil
	}

	for _, publicUser := range publicResponse.Ids{
		userIds = append(userIds, publicUser)
	}

	if len(userIds) == 0 {
		return &protopb.StoriesHome{}, nil
	}

	stories, err := c.service.GetAllHomeStories(ctx, userIds)

	if err != nil{
		return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
	}

	// Get usernames
	usersClient := grpc_common.GetUsersClient(conn)
	for index, story := range stories.Stories {
		response, err := usersClient.GetUsedById(ctx, &protopb.RequestIdUsers{
			Id: story.UserId,
		})

		if err != nil {
			return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
		}

		stories.Stories[index].UserId = response.Username
	}

	responseStories := stories.ConvertToGrpc()

	return responseStories, nil
}

func (c *StoryGrpcController) CreateStory(ctx context.Context, in *protopb.Story) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateStory")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id is provided")
	}else if in.UserId != claims.UserId {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot create stories for other people")
	}

	var story *domain.Story
	story = story.ConvertFromGrpc(in)

	err = c.service.CreateStory(ctx, story)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create story")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *StoryGrpcController) GetStoryById(ctx context.Context, in *protopb.RequestId) (*protopb.Story, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoryById")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.Story{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == ""{
		return &protopb.Story{}, status.Errorf(codes.Unknown, err.Error())
	}else if in.Id == "" {
		return &protopb.Story{}, status.Errorf(codes.Unknown, "cannot retrieve non-existing stories")
	}

	story, err := c.service.GetStoryById(ctx, in.Id)

	if err != nil { return &protopb.Story{}, status.Errorf(codes.Unknown, err.Error()) }

	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Recommendation_service_address)
	if err != nil{
		return &protopb.Story{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()
	followerClient := grpc_common.GetFollowersClient(conn)
	followingResponse, err := followerClient.GetFollowersConnection(ctx, &protopb.Follower{
		UserId:                claims.UserId,
		FollowerId:            story.UserId,
	})

	if (!followingResponse.IsCloseFriends && story.IsCloseFriends) || !followingResponse.IsApprovedRequest {
		return &protopb.Story{}, status.Errorf(codes.PermissionDenied, "cannot retrieve this story")
	}

	grpcStory := story.ConvertToGrpc()

	return grpcStory, nil
}

func (c *StoryGrpcController) RemoveStory(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveStory")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == ""{
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove other people's posts")
	}else if in.Id == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove non-existing posts")
	}

	err = c.service.RemoveStory(ctx, in.Id, claims.UserId)

	if err != nil{
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseContent{}, nil
}