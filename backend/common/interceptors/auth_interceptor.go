package interceptors

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type AuthInterceptor struct {
    jwtManager  *common.JWTManager
}

func NewAuthInterceptor(jwtManager *common.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{jwtManager}
}

/* func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor{
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error){
		log.Println("---> unary interceptor: ", info.FullMethod)

		methodParts := strings.Split(info.FullMethod, "/")
		if len(methodParts) != 3 {
			return nil, errors.New("something went wrong")
		}

		ctx, err := interceptor.authorize(ctx)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
} */

// TODO Add intraservice authentication
func (interceptor *AuthInterceptor) authorize(ctx context.Context) (context.Context, error) {
	contextMetadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := contextMetadata["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	authorizationHeader := values[0]
	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not in valid format")
	}
	accessToken := headerParts[1]

	_, err := interceptor.jwtManager.ValidateJWT(accessToken)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	// RBAC Authentication goes here

	return ctx, nil
}