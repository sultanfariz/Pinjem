package helpers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/brianvoe/sjwt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtClaim struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userId int, role string) (string, error) {
	// token := jwt.New(jwt.SigningMethodHS256)
	// token.Claims["username"] = username
	// token.Claims["password"] = password
	// tokenString, _ := token.SignedString([]byte("secret"))
	// return tokenString

	// claims := jwt.MapClaims{}
	// claims["user_id"] = userId
	// claims["authorized"] = true
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	claims := jwtClaim{userId, role, jwt.StandardClaims{
		// ExpiresAt: jwt.TimeFunc().Add(time.Hour * 24).Unix(),
		ExpiresAt: time.Now().Local().Add(time.Hour * 2400).Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString, nil
}

func ExtractClaims(tokenStr string) (jwtClaim, bool) {
	hmacSecretString := os.Getenv("JWT_SECRET")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return jwtClaim{}, false
	}

	if claims, ok := token.Claims.(jwtClaim); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return jwtClaim{}, false
	}
}

func ExtractToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
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

func SetTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	// Http-only helps mitigate the risk of client side script accessing the protected cookie.
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

func AdminRoleValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		role, err := ExtractJWTPayloadRole(e)
		if err != nil {
			return err
		}
		if role == "admin" {
			return next(e)
		}
		return echo.ErrUnauthorized
	}
}

func UserRoleValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		role, err := ExtractJWTPayloadRole(e)
		if err != nil {
			return err
		}
		if role == "user" {
			return next(e)
		}
		return echo.ErrUnauthorized
	}
}

func ExtractJWTPayloadRole(c echo.Context) (string, error) {
	header := c.Request().Header.Clone().Get("Authorization")
	token := strings.Split(header, "Bearer ")[1]
	claims, _ := sjwt.Parse(token)
	return claims["role"].(string), nil
}

func ExtractJWTPayloadUserId(c echo.Context) (float64, error) {
	header := c.Request().Header.Clone().Get("Authorization")
	token := strings.Split(header, "Bearer ")[1]
	claims, _ := sjwt.Parse(token)
	userId := claims["user_id"].(float64)
	return userId, nil
}
