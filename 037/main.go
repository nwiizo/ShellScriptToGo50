package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte("your-secret-key")

func login(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// ユーザーの認証ロジックをここに記述
	// 例: ユーザー名とパスワードを検証
	if user.Username != "admin" || user.Password != "password" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func protected(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claims)
	username := claims.Username

	return c.JSON(http.StatusOK, map[string]string{"message": "Hello, " + username})
}

func main() {
	e := echo.New()

	e.POST("/login", login)

	r := e.Group("/api")
	r.Use(middleware.JWT(jwtSecret))
	r.GET("/protected", protected)

	e.Logger.Fatal(e.Start(":8080"))
}
