package env

import (
	"errors"
	"net"
	"os"

	"github.com/romanfomindev/microservices-chat/internal/config"
)

const (
	authServerEnvName     = "AUTH_SERVER_HOST"
	authServerPortEnvName = "AUTH_SERVER__PORT"
)

type authServerConfig struct {
	host string
	port string
}

func NewAuthServerConfig() (config.AuthServerConfig, error) {
	host := os.Getenv(authServerEnvName)
	if len(host) == 0 {
		return nil, errors.New("auth host not found")
	}

	port := os.Getenv(authServerPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("auth port not found")
	}

	return &authServerConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *authServerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
