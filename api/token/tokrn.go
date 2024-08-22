package token

import (
	"errors"
	"time"
	"users/config"
	"users/generated/users"
	"users/model"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId   string
	Username string
	Email    string
	jwt.StandardClaims
}

func GenerateAccessJWT(signUp *model.LoginResponse) (string, error) {
	cfg := config.Load()

	claims := Claims{
		UserId: signUp.Id,
		Username: signUp.Username,
		Email:    signUp.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(cfg.AccessToken))
}

func GenerateRefreshJWT(user *users.SignUpRequest) (string, error) {
	cfg := config.Load()
	claims := &Claims{
		Username: user.UserName,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(cfg.RefreshToken))
}

func ExtractClaims(tokenString string) (*Claims, error) {
	cfg := config.Load()
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.RefreshToken), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ExtractClaimsAccess(tokenString string) (*Claims, error) {
	cfg := config.Load()
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AccessToken), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaimsAccess(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}
