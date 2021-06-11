package interceptors

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/interceptors/rbac"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"strings"
)

type RBACInterceptor struct {
	db  		*gorm.DB
	auth        *AuthInterceptor
	jwtManager  *common.JWTManager
}

func NewRBACInterceptor(db *gorm.DB, jwtManager *common.JWTManager) *RBACInterceptor {
	auth := NewAuthInterceptor(jwtManager)
	return &RBACInterceptor{ db, auth, jwtManager }
}

func (interceptor *RBACInterceptor) Authorize() grpc.UnaryServerInterceptor{
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error){
		log.Println("---> RBAC interceptor: ", info.FullMethod)

		methodParts := strings.Split(info.FullMethod, "/")
		if len(methodParts) != 3 {
			return nil, errors.New("something went wrong")
		}

		permissionToCheck := methodParts[2]

		log.Println("henlo 1")
		isAllowed, role, err := interceptor.checkPermission(permissionToCheck, ctx)
		log.Println(isAllowed, role, err)
		if err != nil { return nil, err }
		if !isAllowed { return nil, nil }

		// Allowing unauthorized users to hit some endpoints
		if isAllowed && role == rbac.Nonregistered { return handler(ctx, req) }

		ctx, err = interceptor.auth.authorize(ctx)
		if err != nil { return nil, err }

		return handler(ctx, req)
	}
}

func (interceptor *RBACInterceptor) checkPermission(permission string, ctx context.Context) (bool, string, error) {
	claims, err := interceptor.jwtManager.ExtractClaimsFromMetadata(ctx)
	role := ""
	if err != nil {
		role = rbac.Nonregistered
	}else {
		role = claims.Role
	}

	log.Println("henlo 1 " + role)

	permissionCheck := &rbac.RolePermission{}
	result := interceptor.db.Model(rbac.RolePermission{}).
		Joins("left join user_roles ON user_roles.id = role_permissions.role_id").
		Joins("left join permissions ON permissions.id = role_permissions.permission_id").
		Where("user_roles.name = ? AND permissions.name = ?", role, permission).
		Find(&permissionCheck)

	log.Println(permissionCheck)

	if result.Error != nil || permissionCheck == nil { return false, role, result.Error }
	if permissionCheck.PermissionId == "" || permissionCheck.RoleId == "" {
		return false, role, errors.New("you do not have a permission to access this endpoint")
	}

	return true, role, nil
}