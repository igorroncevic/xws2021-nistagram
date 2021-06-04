// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContentClient is the client API for Content service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentClient interface {
	//    Posts
	CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetAllPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ReducedPostArray, error)
	RemovePost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetPostById(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Post, error)
	//    Stories
	CreateStory(ctx context.Context, in *Story, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetAllStories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StoriesArray, error)
	RemoveStory(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetStoryById(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Story, error)
	//    Comments
	CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetCommentsForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*CommentsArray, error)
	// Likes & Dislikes
	CreateLike(ctx context.Context, in *Like, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetLikesForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*LikesArray, error)
	GetDislikesForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*LikesArray, error)
	// Collections & Favorites
	GetAllCollections(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*CollectionsArray, error)
	GetCollection(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Collection, error)
	CreateCollection(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*EmptyResponse, error)
	RemoveCollection(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetUserFavorites(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Favorites, error)
	CreateFavorite(ctx context.Context, in *FavoritesRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	RemoveFavorite(ctx context.Context, in *FavoritesRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type contentClient struct {
	cc grpc.ClientConnInterface
}

func NewContentClient(cc grpc.ClientConnInterface) ContentClient {
	return &contentClient{cc}
}

func (c *contentClient) CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Content/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetAllPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ReducedPostArray, error) {
	out := new(ReducedPostArray)
	err := c.cc.Invoke(ctx, "/proto.Content/GetAllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) RemovePost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Content/RemovePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetPostById(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/proto.Content/GetPostById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateStory(ctx context.Context, in *Story, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Content/CreateStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetAllStories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*StoriesArray, error) {
	out := new(StoriesArray)
	err := c.cc.Invoke(ctx, "/proto.Content/GetAllStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) RemoveStory(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Content/RemoveStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetStoryById(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Story, error) {
	out := new(Story)
	err := c.cc.Invoke(ctx, "/proto.Content/GetStoryById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Content/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetCommentsForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*CommentsArray, error) {
	out := new(CommentsArray)
	err := c.cc.Invoke(ctx, "/proto.Content/GetCommentsForPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateLike(ctx context.Context, in *Like, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Content/CreateLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetLikesForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*LikesArray, error) {
	out := new(LikesArray)
	err := c.cc.Invoke(ctx, "/proto.Content/GetLikesForPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetDislikesForPost(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*LikesArray, error) {
	out := new(LikesArray)
	err := c.cc.Invoke(ctx, "/proto.Content/GetDislikesForPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetAllCollections(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*CollectionsArray, error) {
	out := new(CollectionsArray)
	err := c.cc.Invoke(ctx, "/proto.Content/GetAllCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetCollection(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Collection, error) {
	out := new(Collection)
	err := c.cc.Invoke(ctx, "/proto.Content/GetCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateCollection(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Content/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) RemoveCollection(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Content/RemoveCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetUserFavorites(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Favorites, error) {
	out := new(Favorites)
	err := c.cc.Invoke(ctx, "/proto.Content/GetUserFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateFavorite(ctx context.Context, in *FavoritesRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Content/CreateFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) RemoveFavorite(ctx context.Context, in *FavoritesRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Content/RemoveFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServer is the server API for Content service.
// All implementations must embed UnimplementedContentServer
// for forward compatibility
type ContentServer interface {
	//    Posts
	CreatePost(context.Context, *Post) (*EmptyResponse, error)
	GetAllPosts(context.Context, *EmptyRequest) (*ReducedPostArray, error)
	RemovePost(context.Context, *RequestId) (*EmptyResponse, error)
	GetPostById(context.Context, *RequestId) (*Post, error)
	//    Stories
	CreateStory(context.Context, *Story) (*EmptyResponse, error)
	GetAllStories(context.Context, *EmptyRequest) (*StoriesArray, error)
	RemoveStory(context.Context, *RequestId) (*EmptyResponse, error)
	GetStoryById(context.Context, *RequestId) (*Story, error)
	//    Comments
	CreateComment(context.Context, *Comment) (*EmptyResponse, error)
	GetCommentsForPost(context.Context, *RequestId) (*CommentsArray, error)
	// Likes & Dislikes
	CreateLike(context.Context, *Like) (*EmptyResponse, error)
	GetLikesForPost(context.Context, *RequestId) (*LikesArray, error)
	GetDislikesForPost(context.Context, *RequestId) (*LikesArray, error)
	// Collections & Favorites
	GetAllCollections(context.Context, *RequestId) (*CollectionsArray, error)
	GetCollection(context.Context, *RequestId) (*Collection, error)
	CreateCollection(context.Context, *Collection) (*EmptyResponse, error)
	RemoveCollection(context.Context, *RequestId) (*EmptyResponse, error)
	GetUserFavorites(context.Context, *RequestId) (*Favorites, error)
	CreateFavorite(context.Context, *FavoritesRequest) (*EmptyResponse, error)
	RemoveFavorite(context.Context, *FavoritesRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedContentServer()
}

// UnimplementedContentServer must be embedded to have forward compatible implementations.
type UnimplementedContentServer struct {
}

func (UnimplementedContentServer) CreatePost(context.Context, *Post) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedContentServer) GetAllPosts(context.Context, *EmptyRequest) (*ReducedPostArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPosts not implemented")
}
func (UnimplementedContentServer) RemovePost(context.Context, *RequestId) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePost not implemented")
}
func (UnimplementedContentServer) GetPostById(context.Context, *RequestId) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}
func (UnimplementedContentServer) CreateStory(context.Context, *Story) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStory not implemented")
}
func (UnimplementedContentServer) GetAllStories(context.Context, *EmptyRequest) (*StoriesArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStories not implemented")
}
func (UnimplementedContentServer) RemoveStory(context.Context, *RequestId) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveStory not implemented")
}
func (UnimplementedContentServer) GetStoryById(context.Context, *RequestId) (*Story, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoryById not implemented")
}
func (UnimplementedContentServer) CreateComment(context.Context, *Comment) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedContentServer) GetCommentsForPost(context.Context, *RequestId) (*CommentsArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsForPost not implemented")
}
func (UnimplementedContentServer) CreateLike(context.Context, *Like) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLike not implemented")
}
func (UnimplementedContentServer) GetLikesForPost(context.Context, *RequestId) (*LikesArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikesForPost not implemented")
}
func (UnimplementedContentServer) GetDislikesForPost(context.Context, *RequestId) (*LikesArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDislikesForPost not implemented")
}
func (UnimplementedContentServer) GetAllCollections(context.Context, *RequestId) (*CollectionsArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCollections not implemented")
}
func (UnimplementedContentServer) GetCollection(context.Context, *RequestId) (*Collection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollection not implemented")
}
func (UnimplementedContentServer) CreateCollection(context.Context, *Collection) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedContentServer) RemoveCollection(context.Context, *RequestId) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCollection not implemented")
}
func (UnimplementedContentServer) GetUserFavorites(context.Context, *RequestId) (*Favorites, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavorites not implemented")
}
func (UnimplementedContentServer) CreateFavorite(context.Context, *FavoritesRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFavorite not implemented")
}
func (UnimplementedContentServer) RemoveFavorite(context.Context, *FavoritesRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFavorite not implemented")
}
func (UnimplementedContentServer) mustEmbedUnimplementedContentServer() {}

// UnsafeContentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServer will
// result in compilation errors.
type UnsafeContentServer interface {
	mustEmbedUnimplementedContentServer()
}

func RegisterContentServer(s grpc.ServiceRegistrar, srv ContentServer) {
	s.RegisterService(&Content_ServiceDesc, srv)
}

func _Content_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetAllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetAllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/GetAllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetAllPosts(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_RemovePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).RemovePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/RemovePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).RemovePost(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetPostById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetPostById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/GetPostById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetPostById(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Story)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/CreateStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateStory(ctx, req.(*Story))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetAllStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetAllStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/GetAllStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetAllStories(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_RemoveStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).RemoveStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/RemoveStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).RemoveStory(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetStoryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetStoryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/GetStoryById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetStoryById(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetCommentsForPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetCommentsForPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/GetCommentsForPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetCommentsForPost(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/CreateLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateLike(ctx, req.(*Like))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetLikesForPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetLikesForPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/GetLikesForPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetLikesForPost(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetDislikesForPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetDislikesForPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/GetDislikesForPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetDislikesForPost(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetAllCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetAllCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/GetAllCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetAllCollections(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/GetCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetCollection(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Collection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateCollection(ctx, req.(*Collection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_RemoveCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).RemoveCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/RemoveCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).RemoveCollection(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetUserFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetUserFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/GetUserFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetUserFavorites(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/CreateFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateFavorite(ctx, req.(*FavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_RemoveFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).RemoveFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Content/RemoveFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).RemoveFavorite(ctx, req.(*FavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Content_ServiceDesc is the grpc.ServiceDesc for Content service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Content_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Content",
	HandlerType: (*ContentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _Content_CreatePost_Handler,
		},
		{
			MethodName: "GetAllPosts",
			Handler:    _Content_GetAllPosts_Handler,
		},
		{
			MethodName: "RemovePost",
			Handler:    _Content_RemovePost_Handler,
		},
		{
			MethodName: "GetPostById",
			Handler:    _Content_GetPostById_Handler,
		},
		{
			MethodName: "CreateStory",
			Handler:    _Content_CreateStory_Handler,
		},
		{
			MethodName: "GetAllStories",
			Handler:    _Content_GetAllStories_Handler,
		},
		{
			MethodName: "RemoveStory",
			Handler:    _Content_RemoveStory_Handler,
		},
		{
			MethodName: "GetStoryById",
			Handler:    _Content_GetStoryById_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _Content_CreateComment_Handler,
		},
		{
			MethodName: "GetCommentsForPost",
			Handler:    _Content_GetCommentsForPost_Handler,
		},
		{
			MethodName: "CreateLike",
			Handler:    _Content_CreateLike_Handler,
		},
		{
			MethodName: "GetLikesForPost",
			Handler:    _Content_GetLikesForPost_Handler,
		},
		{
			MethodName: "GetDislikesForPost",
			Handler:    _Content_GetDislikesForPost_Handler,
		},
		{
			MethodName: "GetAllCollections",
			Handler:    _Content_GetAllCollections_Handler,
		},
		{
			MethodName: "GetCollection",
			Handler:    _Content_GetCollection_Handler,
		},
		{
			MethodName: "CreateCollection",
			Handler:    _Content_CreateCollection_Handler,
		},
		{
			MethodName: "RemoveCollection",
			Handler:    _Content_RemoveCollection_Handler,
		},
		{
			MethodName: "GetUserFavorites",
			Handler:    _Content_GetUserFavorites_Handler,
		},
		{
			MethodName: "CreateFavorite",
			Handler:    _Content_CreateFavorite_Handler,
		},
		{
			MethodName: "RemoveFavorite",
			Handler:    _Content_RemoveFavorite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content.proto",
}
