package config

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GetBcryptCost() int {
	costStr := os.Getenv("BCRYPT_COST")
	cost, err := strconv.Atoi(costStr)
	if err != nil {
		cost = bcrypt.DefaultCost
	}

	return cost
}

func GetJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "berbahaya"
	}

	return secret
}
