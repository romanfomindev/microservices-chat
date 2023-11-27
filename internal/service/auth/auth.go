package auth

import "github.com/romanfomindev/microservices-chat/internal/config"

type AuthService struct {
	chatServerCfg config.AuthServerConfig
}

func NewAuthService(cfg config.AuthServerConfig) *AuthService {
	return &AuthService{
		chatServerCfg: cfg,
	}
}
