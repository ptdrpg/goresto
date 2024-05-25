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
	jwtkey := os.Getenv("SECRET_KEY")
	validity := time.Now().Add(5 * time.Minute)
	claims := &AccesClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: validity.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtkey)
}

func GenerateRefreshToken(username string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file", err)
	}
	refreshkey := os.Getenv("SECRET_KEY")
	validity := time.Now().Add(24 * time.Hour)
	claims := &RefreshClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: validity.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return refreshToken.SignedString(refreshkey)
}