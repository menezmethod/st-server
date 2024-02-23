package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"st-auth-svc/pkg/models"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours uint64
}

type JwtClaims struct {
	jwt.StandardClaims
	Id    uint64
	Email string
}

func (w *JwtWrapper) GenerateToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(time.Duration(w.ExpirationHours) * time.Hour)
	claims := &JwtClaims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JwtClaims{}, w.keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (w *JwtWrapper) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method")
	}
	return []byte(w.SecretKey), nil
}
