package config

import (
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GetBcryptCost() int {
	costStr := os.Getenv(envBcryptCost)
	cost, err := strconv.Atoi(costStr)
	if err != nil {
		cost = bcrypt.DefaultCost
	}

	return cost
}

func GetJwtSecret() string {
	secret := os.Getenv(envJwtSecret)
	if secret == "" {
		log.Fatalln("please provide JWT secret")
	}

	return secret
}
