package config

import "os"

type AdminConfig struct {
	Email    string
	Username string
	Password string
}

func GetAdminConfig() AdminConfig {
	return AdminConfig{
		Email:    os.Getenv(envAdminEmail),
		Username: os.Getenv(envAdminUsername),
		Password: os.Getenv(envAdminPassword),
	}
}
