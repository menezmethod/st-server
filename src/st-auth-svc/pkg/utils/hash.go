package utils

import (
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(logger *zap.Logger, password string, cost int) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		logger.Error("Error hashing password", zap.Error(err))
		return "", err
	}
	return string(hashedBytes), nil
}

func CheckPasswordHash(logger *zap.Logger, password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		logger.Error("Error comparing password hash", zap.Error(err))
		return false, err
	}
	return true, nil
}
