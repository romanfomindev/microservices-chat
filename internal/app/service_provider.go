package app

import (
	"log"

	"github.com/romanfomindev/microservices-chat/internal/config"
	"github.com/romanfomindev/microservices-chat/internal/config/env"
	"github.com/romanfomindev/microservices-chat/internal/service"
	"github.com/romanfomindev/microservices-chat/internal/service/auth"
	"github.com/romanfomindev/microservices-chat/internal/service/chat"
)

type serviceProvider struct {
	chatServerConfig config.ChatServerConfig
	authServerConfig config.AuthServerConfig
	authService      service.AuthService
	chatsService     service.ChatService
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) ChatServerConfig() config.ChatServerConfig {
	if s.chatServerConfig == nil {
		cfg, err := env.NewChatServerConfig()
		if err != nil {
			log.Fatalf("failed to get chat server config: %s", err.Error())
		}

		s.chatServerConfig = cfg
	}

	return s.chatServerConfig
}

func (s *serviceProvider) AuthServerConfig() config.AuthServerConfig {
	if s.authServerConfig == nil {
		cfg, err := env.NewAuthServerConfig()
		if err != nil {
			log.Fatalf("failed to get chat server config: %s", err.Error())
		}

		s.authServerConfig = cfg
	}

	return s.authServerConfig
}

func (s *serviceProvider) AuthService() service.AuthService {
	if s.authService == nil {
		authService := auth.NewAuthService(s.AuthServerConfig())
		s.authService = authService
	}

	return s.authService
}

func (s *serviceProvider) ChatService() service.ChatService {
	if s.chatsService == nil {
		chatService := chat.NewChatService(s.ChatServerConfig())
		s.chatsService = chatService
	}

	return s.chatsService
}
