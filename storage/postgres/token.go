package postgres

import (
	"fmt"
	"time"
	pb "users/generated/users"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/form3tech-oss/jwt-go"
)

var secret_key = "key"

type Token struct {
	AccessToken  string
	RefreshToken string
}

func (repo *UserRepo) GenerateGWTToken(user *pb.SignUpRequest) (*Token, error) {

	AccessToken := jwt.New(jwt.SigningMethodHS256)

	claims := AccessToken.Claims.(jwt.MapClaims)
	claims["user_name"] = user.UserName
	claims["password"] = user.Password
	claims["email"] = user.Email
	claims["full_name"] = user.FullName
	claims["user_type"] = user.UserType
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(3 * time.Hour).Unix()

	access, err := AccessToken.SignedString([]byte(secret_key))
	if err != nil {
		fmt.Println("error  access singed my_secret key")
	}

	RefreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaim := RefreshToken.Claims.(jwt.MapClaims)
	refreshClaim["user_name"] = user.UserName
	refreshClaim["password"] = user.Password
	refreshClaim["email"] = user.Email
	refreshClaim["full_name"] = user.FullName
	refreshClaim["user_type"] = user.UserType
	refreshClaim["iat"] = time.Now().Unix()
	refreshClaim["exp"] = time.Now().Add(24 * time.Hour)

	refresh, err := RefreshToken.SignedString([]byte(secret_key))
	if err != nil {
		return nil, err
	}

	_, err = repo.DB.Exec("update users set  token=$1 where email =$2", refresh, user.Email)
	if err != nil {
		return nil, err
	}
	return &Token{
		AccessToken: access,
	}, nil

}
