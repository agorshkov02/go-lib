package database

import (
	"fmt"
	"os"
)

type ClientConfig interface {
	GetDriver() string
	GetDatasource() string
}

func NewClientConfig() ClientConfig {
	switch os.Getenv("DATABASE_TYPE") {
	case "postgres":
		return NewPqClientConfig()
	default:
		fmt.Fprintf(os.Stderr, "DATABASE_TYPE not specified\n")
		os.Exit(1)
	}

	return nil
}

type PqClientConfig struct {
	Host, Port, Name, User, Password string
}

func (pqcc *PqClientConfig) GetDriver() string {
	return "postgres"
}

func (pqcc *PqClientConfig) GetDatasource() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pqcc.User, pqcc.Password, pqcc.Host, pqcc.Port, pqcc.Name)
}

func NewPqClientConfig() *PqClientConfig {
	return &PqClientConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Name:     os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	}
}
