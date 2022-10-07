package main

import "os"

type ProjectConfig struct {
	ContentRoot        string
	ApplicationPort    string
	ApplicationAddress string
}

func CreateConfigFromEnv() ProjectConfig {
	port := os.Getenv("BACKEND_PORT")
	if len(port) == 0 {
		port = "8000"
	}

	address := os.Getenv("BACKEND_ADDRESS")
	if address == "" {
		address = "localhost"
	}
	contentRoot := os.Getenv("CONTENT_ROOT")

	return ProjectConfig{
		ContentRoot:        contentRoot,
		ApplicationPort:    port,
		ApplicationAddress: address,
	}
}
