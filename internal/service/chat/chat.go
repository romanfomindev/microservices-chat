package chat

import (
	"github.com/romanfomindev/microservices-chat/internal/config"
	"github.com/romanfomindev/microservices-chat/internal/service"
	"google.golang.org/grpc"
)

type ChatService struct {
	cfg  config.ChatServerConfig
	conn *grpc.ClientConn
}

func NewChatService(cfg config.ChatServerConfig, conn *grpc.ClientConn) service.ChatService {
	return &ChatService{
		cfg:  cfg,
		conn: conn,
	}
}
