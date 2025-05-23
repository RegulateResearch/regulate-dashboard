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
	return getMandatoryAuthValue(envJwtSecret)
}

func GetJwtIssuer() string {
	return getMandatoryAuthValue(envJwtIssuer)
}

func getMandatoryAuthValue(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		log.Fatalln("please provide JWT secret")
	}

	return value
}
