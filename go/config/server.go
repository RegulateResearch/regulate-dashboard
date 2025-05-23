package config

import "os"

func GetServerPort() string {
	return getPort(envServerPort, "8080")
}

func GetListenerPort() string {
	return getPort(envListenerPort, "9090")
}

func getPort(envVar string, defaultPort string) string {
	port := os.Getenv(envVar)
	if port == "" {
		port = defaultPort
	}

	return port
}
