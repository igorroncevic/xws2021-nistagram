package common

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
	"time"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey: secretKey, tokenDuration: tokenDuration}
}

func (manager *JWTManager) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.String(), "favicon.ico") {
			// Allow favicon.ico to load
			next.ServeHTTP(w, r)
		}

		authHeader := r.Header.Get("Authorization")
		splitHeader := strings.Split(authHeader, " ")
		if len(splitHeader) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jwtString := splitHeader[1]
		_, err := manager.ValidateJWT(jwtString)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (manager *JWTManager) GenerateJwt(id string, role string, email string) (string, error) {
	claims := &Claims{
		UserId: id,
		Email:  email,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(manager.secretKey))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", errors.New("unable to sign JWT token")
	}

	return tokenString, nil
}

func (manager *JWTManager) ValidateJWT(jwtString string) (*Claims, error) {
	if jwtString == "" {
		return nil, errors.New("unauthorized")
	}

	// This method will return an error if the token is invalid (if it has expired according to the expiry time
	// we set on sign in), or if the signature does not match
	token, err := jwt.ParseWithClaims(jwtString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secretKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("unauthorized")
		}
		// If string that is sent is not even a JWT
		return nil, errors.New("bad request")
	}

	if !token.Valid {
		return nil, errors.New("unauthorized")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func (manager *JWTManager) ExtractClaimsFromMetadata(ctx context.Context) (*Claims, error) {
	contextMetadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &Claims{}, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := contextMetadata["authorization"]
	if len(values) == 0 {
		return &Claims{}, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	authorizationHeader := values[0]
	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 {
		return &Claims{}, status.Errorf(codes.Unauthenticated, "authorization token is not in valid format")
	}
	accessToken := headerParts[1]

	claims, err := manager.ExtractClaims(accessToken)
	if err != nil {
		return &Claims{}, errors.New("invalid token claims")
	}

	return claims, nil
}

func (manager *JWTManager) ExtractClaims(accessToken string) (*Claims, error) {
	_, err := manager.ValidateJWT(accessToken)
	if err != nil {
		return &Claims{}, errors.New("invalid JWT")
	}

	token, _ := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secretKey), nil
	})

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return &Claims{}, errors.New("invalid token claims")
	}
	return claims, nil
}

// Create a refresh route?
func (manager *JWTManager) RefreshJWT(jwtString string) (string, time.Time, int, error) {
	claims := &Claims{}

	// This method will return an error if the token is invalid (if it has expired according to the expiry time
	// we set on sign in), or if the signature does not match
	tkn, err := jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtString, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", time.Now(), http.StatusUnauthorized, errors.New("unauthorized")
		}
		return "", time.Now(), http.StatusBadRequest, errors.New("bad request")
	}

	if !tkn.Valid {
		return "", time.Now(), http.StatusUnauthorized, errors.New("unauthorized")
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return "", time.Now(), http.StatusBadRequest, errors.New("token has not expired yet")
	}

	// Now, create a new token with previous claims but with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(manager.secretKey)
	if err != nil {
		return "", time.Now(), http.StatusInternalServerError, errors.New("unable to sign JWT token")
	}

	return tokenString, expirationTime, http.StatusOK, nil
}
