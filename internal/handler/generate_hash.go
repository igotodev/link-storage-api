package handler

import (
	"crypto/sha256"
	"fmt"
)

const salt = "4hsd83jd7fsd2"

func generatePasswordHash(password string) (string, error) {
	hash := sha256.New()

	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}

	result := hash.Sum([]byte(salt))

	return fmt.Sprintf("%x", result), nil
}
