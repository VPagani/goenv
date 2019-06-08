package goenv

import (
	"errors"
	"os"
	"path"

	"github.com/joho/godotenv"
)

// Errors
var (
	ErrEnvFileNotFound = errors.New("Env file not found")
)

func fileExists(dir string, filename string) bool {
	filepath := path.Join(dir, filename)
	_, err := os.Stat(filepath)
	return err != nil
}

// LoadEnv loads env files
func LoadEnv(filename string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if filename == "" {
		filename = ".env"
	}

	dir := wd
	for !fileExists(dir, filename) {
		if dir == "/" {
			return "", ErrEnvFileNotFound
		}

		dir = path.Dir(dir)
	}

	envPath := path.Join(dir, filename)

	return envPath, godotenv.Load(envPath)
}

// Env helps fetching env vars and applying default values when needed
func Env(envVar string, defaultValue ...string) string {
	value, exists := os.LookupEnv(envVar)
	if !exists && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

// SetupEnv loads env and returns `Env` function
func SetupEnv(envFilename string) func(envVar string, defaultValue ...string) string {
	_, _ = LoadEnv(envFilename)
	return Env
}
