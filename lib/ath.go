package lib

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// var jwtkey = []byte("myfirstrestogo")
// var refreshkey = []byte("refreshkeyrestogo")

type AccesClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file", err)
	}
	jwtkey := []byte(os.Getenv("SECRET_KEY"))
	validity := time.Now().Add(5 * time.Minute)
	claims := &AccesClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: validity.Unix(),
			Issuer: "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtkey)
}

func GenerateRefreshToken(username string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file", err)
	}
	refreshkey := []byte(os.Getenv("SECRET_KEY"))
	validity := time.Now().Add(24 * time.Hour)
	claims := &AccesClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: validity.Unix(),
			Issuer: "test",
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refreshToken.SignedString(refreshkey)
}