package chat

import (
	"context"

	chat_server "github.com/romanfomindev/microservices-chat/clients/grpc/chat-server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *ChatService) Create(ctx context.Context, accessToken, chatName string, username []string) (uint64, error) {
	conn, err := grpc.Dial(
		s.cfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return 0, err
	}
	client := chat_server.NewChatServer(conn)

	chatId, err := client.Create(ctx, accessToken, chatName, username)
	if err != nil {
		return 0, err
	}

	return chatId, nil
}
