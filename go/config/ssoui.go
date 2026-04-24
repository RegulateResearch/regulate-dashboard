package config

import "os"

func GetSsoUiUrl() string {
	url := os.Getenv(envSsoUiUrl)
	return url
}
