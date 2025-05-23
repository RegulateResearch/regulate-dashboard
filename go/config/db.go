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
		port := getDatabasePort()
		name := os.Getenv(envDatabaseName)
		url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, name)
	}

	return url
}

func getDatabasePort() string {
	port := os.Getenv(envDatabasePort)
	if port == "" {
		port = "5432"
	}

	return port
}
