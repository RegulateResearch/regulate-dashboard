package config

import (
	"fmt"
	"os"
)

func GetDatabaseURL() string {
	url := os.Getenv(envDatabaseURL)
	if url == "" {
		username := os.Getenv(envDatabaseUsername)
		password := os.Getenv(envDatabasePassword)
		host := os.Getenv(envDatabaseHost)
		port := os.Getenv(envDatabasePort)
		name := os.Getenv(envDatabaseName)
		url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, name)
	}

	return url
}
