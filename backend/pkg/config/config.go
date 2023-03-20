package config

import (
	"os"

	"cloud.google.com/go/compute/metadata"
)

type Config struct {
	Env                  string
	GoogleCloudProjectID string
}

func New() *Config {
	var projectID string
	got, err := metadata.ProjectID()
	if err == nil {
		projectID = got
	}

	return &Config{
		Env:                  os.Getenv("ENV"),
		GoogleCloudProjectID: projectID,
	}
}
