package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	AuthSecret string
	HTTPPort   string
}

func NewConfig() (*Config, error) {
	var err error
	c := &Config{}

	c.DBHost = os.Getenv("DB_HOST")
	if c.DBHost == "" {
		return nil, fmt.Errorf("DB_HOST not set")
	}
	portStr := os.Getenv("DB_PORT")
	if portStr == "" {
		return nil, fmt.Errorf("DB_PORT not set")
	}
	c.DBPort, err = strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	c.DBUser = os.Getenv("DB_USER")
	c.DBPassword = os.Getenv("DB_PASSWORD")
	c.DBName = os.Getenv("DB_NAME")
	if c.DBUser == "" || c.DBPassword == "" || c.DBName == "" {
		return nil, fmt.Errorf("DB_USER/DB_PASSWORD/DB_NAME must be set")
	}

	c.AuthSecret = os.Getenv("AUTH_SECRET")
	if c.AuthSecret == "" {
		return nil, fmt.Errorf("AUTH_SECRET not set")
	}

	c.HTTPPort = os.Getenv("HTTP_PORT")
	if c.HTTPPort == "" {
		c.HTTPPort = "8000"
	}
	return c, nil
}
