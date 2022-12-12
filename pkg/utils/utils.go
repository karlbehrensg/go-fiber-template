package utils

import (
	"fmt"
	"os"

	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
)

// CheckEnv validate env required
func CheckEnv() {
	envProps := []string{
		"ENV",
		"DB_HOST",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_PORT",
		"DB_SSL_MODE",
		"DB_TIME_ZONE",
		"APP_PORT",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}
