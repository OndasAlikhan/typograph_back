package service

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type JWTService struct {
}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (js *JWTService) GenerateToken(id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["auth_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (js *JWTService) GenerateHash(word string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(word), 14)

	return string(bytes), err
}

func (js *JWTService) IsEqual(hash, word string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(word))

	return err == nil
}
