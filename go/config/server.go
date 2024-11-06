package config

import "os"

func GetServerPort() string {
	port := os.Getenv(envServerPort)
	if port == "" {
		port = "8080"
	}

	return port
}
