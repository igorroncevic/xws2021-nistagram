package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type Credentials struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var (
	expirationMinutes = 5
	jwtKey = []byte("some-jwt-key")
)

func GenerateJwt(email string) (string, time.Time, error){
	expirationTime := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", time.Now(), errors.New("unable to sign JWT token")
	}

	return tokenString, expirationTime, nil
}

func ValidateJWT(jwtString string) (int, error){
	// If working with cookies, first extract jwt value from it
	//c, err := r.Cookie("token")
	//	if err != nil {
	//		if err == http.ErrNoCookie {
	//			w.WriteHeader(http.StatusUnauthorized)
	//			return
	//		}
	//		w.WriteHeader(http.StatusBadRequest)
	//		return
	//	}
	//	tknStr := c.Value

	if jwtString == ""{
		return http.StatusUnauthorized, errors.New("unauthorized")
	}

	claims := &Claims{}

	// This method will return an error if the token is invalid (if it has expired according to the expiry time
	// we set on sign in), or if the signature does not match
	tkn, err := jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token)(interface{}, error){
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized, errors.New("unauthorized")
		}
		// If string that is sent is not even a JWT
		return http.StatusBadRequest, errors.New("bad request")
	}

	if !tkn.Valid{
		return http.StatusUnauthorized, errors.New("unauthorized")
	}

	return http.StatusOK, nil
}

// Create a refresh route?
func RefreshJWT(jwtString string) (string, time.Time, int, error){
	claims := &Claims{}

	// This method will return an error if the token is invalid (if it has expired according to the expiry time
	// we set on sign in), or if the signature does not match
	tkn, err := jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token)(interface{}, error){
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid{
			return "", time.Now(), http.StatusUnauthorized, errors.New("unauthorized")
		}
		return "", time.Now(), http.StatusBadRequest, errors.New("bad request")
	}

	if !tkn.Valid{
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
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Now(), http.StatusInternalServerError, errors.New("unable to sign JWT token")
	}

	return tokenString, expirationTime, http.StatusOK, nil
}