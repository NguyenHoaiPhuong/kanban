package auth

import (
	"reflect"
	"time"

	"github.com/NguyenHoaiPhuong/warehouse/server/models"
	"github.com/dgrijalva/jwt-go"
)

// Token struct
type Token struct {
	Access  string `json:"access_token" bson:"access_token"`
	Refresh string `json:"refresh_token" bson:"refresh_token"`
}

// GenerateTokenPair : generate access token and refresh token
func GenerateTokenPair(user *models.User, secretKey string) (*Token, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)

	v := reflect.ValueOf(user).Elem()
	t := reflect.TypeOf(user).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		claims[fieldName] = v.Field(i).Interface()
	}
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	at, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &Token{
		Access:  at,
		Refresh: rt,
	}, nil
}
