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

type PostGrpcController struct {
	service    *services.PostService
	jwtManager *common.JWTManager
}

func NewPostController(db *gorm.DB, jwtManager *common.JWTManager) (*PostGrpcController, error) {
	service, err := services.NewPostService(db)
	if err != nil {
		return nil, err
	}

	return &PostGrpcController{
		service,
		jwtManager,
	}, nil
}

func (c *PostGrpcController) CreatePost(ctx context.Context, in *protopb.Post) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}  else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.UserId {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot create post for another user")
	}

	var post *domain.Post
	post = post.ConvertFromGrpc(in)

	err = c.service.CreatePost(ctx, post)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create post")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *PostGrpcController) GetAllPosts(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == ""{
		return &protopb.ReducedPostArray{}, status.Errorf(codes.InvalidArgument, "no user id is provided")
	}

	userId := claims.UserId
	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Recommendation_service_address)
	if err != nil{
		return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()

	followerClient := grpc_common.GetFollowersClient(conn)
	followingResponse, err := followerClient.GetAllFollowingsForHomepagePosts(ctx, &protopb.CreateUserRequestFollowers{
		User: &protopb.UserFollowers{ UserId: userId },
	})

	if len(followingResponse.Users) == 0 {
		return &protopb.ReducedPostArray{
			Posts: []*protopb.ReducedPost{},
		}, nil
	}

	userIds := []string{}
	for _, following := range followingResponse.Users{
		userIds = append(userIds, following.UserId)
	}

	privacyConn, err := grpc_common.CreateGrpcConnection(grpc_common.Users_service_address)
	if err != nil{
		return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer privacyConn.Close()

	privacyClient := grpc_common.GetPrivacyClient(conn)
	publicResponse, err := privacyClient.GetAllPublicUsers(ctx, &protopb.EmptyRequestPrivacy{})

	if len(publicResponse.Ids) == 0 {
		return &protopb.ReducedPostArray{
			Posts: []*protopb.ReducedPost{},
		}, nil
	}

	for _, publicUser := range publicResponse.Ids{
		userIds = append(userIds, publicUser)
	}

	if len(userIds) == 0 {
		return &protopb.ReducedPostArray{ Posts: []*protopb.ReducedPost{} }, nil
	}

	posts, err := c.service.GetAllPosts(ctx, userIds)

	if err != nil {
		return &protopb.ReducedPostArray{
			Posts: []*protopb.ReducedPost{},
		}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}

func (c *PostGrpcController) GetPostById(ctx context.Context, id string) (*protopb.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostById")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == ""{
		return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
	}else if id == "" {
		return &protopb.Post{}, status.Errorf(codes.Unknown, "cannot retrieve non-existing posts")
	}

	post, err := c.service.GetPostById(ctx, id)
	if err != nil {
		return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
	}

	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Recommendation_service_address)
	if err != nil{
		return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
	}
	defer conn.Close()
	followerClient := grpc_common.GetFollowersClient(conn)
	followingResponse, err := followerClient.GetFollowersConnection(ctx, &protopb.Follower{
		UserId:                claims.UserId,
		FollowerId:            post.UserId,
	})

	if !followingResponse.IsApprovedRequest {
		return &protopb.Post{}, status.Errorf(codes.PermissionDenied, "cannot retrieve this story")
	}

	grpcPost := post.ConvertToGrpc()
	return grpcPost, nil
}

func (c *PostGrpcController) RemovePost(ctx context.Context, id string) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemovePost")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == ""{
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove other people's posts")
	}else if id == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove non-existing posts")
	}

	err = c.service.RemovePost(ctx, id, claims.UserId)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *PostGrpcController) SearchContentByLocation(ctx context.Context, in *protopb.SearchLocationRequest) (*protopb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	location := in.Location

	posts, err := c.service.SearchContentByLocation(ctx, location)
	if err != nil {
		return &protopb.ReducedPostArray{
			Posts: []*protopb.ReducedPost{},
		}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}

func (c *PostGrpcController) GetPostsByHashtag(ctx context.Context, in *protopb.Hashtag) (*protopb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsByHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts, err := c.service.GetPostsByHashtag(ctx, in.Text)
	if err != nil {
		return &protopb.ReducedPostArray{
			Posts: []*protopb.ReducedPost{},
		}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}
