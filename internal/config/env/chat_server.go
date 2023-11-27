package env

import (
	"errors"
	"net"
	"os"

	"github.com/romanfomindev/microservices-chat/internal/config"
)

const (
	chatServerEnvName     = "CHAT_SERVER_HOST"
	chatServerPortEnvName = "CHAT_SERVER__PORT"
)

type serverConfig struct {
	host string
	port string
}

func NewChatServerConfig() (config.ChatServerConfig, error) {
	host := os.Getenv(chatServerEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(chatServerPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &serverConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *serverConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
