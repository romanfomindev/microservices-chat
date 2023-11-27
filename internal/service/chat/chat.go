package chat

import (
	"github.com/romanfomindev/microservices-chat/internal/config"
	"github.com/romanfomindev/microservices-chat/internal/service"
)

type ChatService struct {
	cfg config.ChatServerConfig
}

func NewChatService(cfg config.ChatServerConfig) service.ChatService {
	return &ChatService{
		cfg: cfg,
	}
}
