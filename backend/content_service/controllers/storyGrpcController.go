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
	}else if claims.UserId == "" || in.Id == "" {
		return &protopb.StoriesArray{}, status.Errorf(codes.InvalidArgument, "no user id is provided")
	}

	followConnection, err := grpc_common.CheckFollowInteraction(ctx, in.Id, claims.UserId)
	if err != nil {
		return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
	if err != nil {
		return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	isBlocked, err := grpc_common.CheckIfBlocked(ctx, in.Id, claims.UserId)
	if err != nil {
		return &protopb.StoriesArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	// If used is blocked or his profile is private and did not approve your request
	if isBlocked || (!isPublic && !followConnection.IsApprovedRequest ) {
		return &protopb.StoriesArray{}, nil
	}

	stories, err := c.service.GetStoriesForUser(ctx, in.Id, followConnection.IsCloseFriends)
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

	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Recommendation_service_address)
	if err != nil{
		return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	userIds, err := grpc_common.GetHomepageUsers(ctx, claims.UserId)
	if err != nil {
		return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
	}

	closeFriends, err := grpc_common.GetCloseFriends(ctx, claims.UserId)
	nonCloseFriends := []string{}
	for _, userId := range userIds{
		found := false
		for _, closeFriends := range closeFriends{
			if closeFriends == userId {
				found = true
				break
			}
		}
		if !found{
			nonCloseFriends = append(nonCloseFriends, userId)
		}
	}

	closeFriendsStories, err := c.service.GetAllHomeStories(ctx, closeFriends, true)
	if err != nil{
		return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
	}
	nonCloseFriendsStories, err := c.service.GetAllHomeStories(ctx, nonCloseFriends, false)
	if err != nil{
		return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
	}

	allStories := domain.StoriesHome{}
	allStories.Stories = nonCloseFriendsStories.Stories
	for _, storyHome := range closeFriendsStories.Stories {
		allStories.Stories = append(allStories.Stories, storyHome)
	}

	// Get usernames
	for index, story := range allStories.Stories {
		username, err := grpc_common.GetUsernameById(ctx, story.UserId)
		if err != nil {
			return &protopb.StoriesHome{}, status.Errorf(codes.Unknown, err.Error())
		}

		allStories.Stories[index].UserId = username
	}

	responseStories := allStories.ConvertToGrpc()
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

	following, err := grpc_common.CheckFollowInteraction(ctx, in.Id, claims.UserId)
	if err != nil { return &protopb.Story{}, status.Errorf(codes.Unknown, err.Error()) }

	if (!following.IsCloseFriends && story.IsCloseFriends) || !following.IsApprovedRequest {
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