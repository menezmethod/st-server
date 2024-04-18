package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/models"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours uint64
}

type JWTClaims struct {
	jwt.StandardClaims
	Id    uint64
	Email string
}

func (w *JwtWrapper) GenerateToken(user *models.User) (string, error) {
	claims := &JWTClaims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", err // Simply return the error
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(w.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT token has expired")
	}

	return claims, nil
}
