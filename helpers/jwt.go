package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// type jwtClaim struct {
// 	userId int `json:"user_id"`
// 	jwt.StandardClaims
// }

func GenerateToken(userId int) (string, error) {
	// token := jwt.New(jwt.SigningMethodHS256)
	// token.Claims["username"] = username
	// token.Claims["password"] = password
	// tokenString, _ := token.SignedString([]byte("secret"))
	// return tokenString

	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// claims := jwtClaim{userId, jwt.StandardClaims{
	// 	// ExpiresAt: jwt.TimeFunc().Add(time.Hour * 24).Unix(),
	// 	ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
	// }}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString, nil
}

func ExtractToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_id"].(int), nil
	}
	return 0, nil
}

func ExtractTokenUserId(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
			return claims["user_id"].(int), nil
		}
	}
	return 0, nil
}
