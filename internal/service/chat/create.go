package chat

import (
	"context"

	chat_server "github.com/romanfomindev/microservices-chat/clients/grpc/chat-server"
)

func (s *ChatService) Create(ctx context.Context, chatName string, username []string) (uint64, error) {
	client := chat_server.NewChatServer(s.conn)

	chatId, err := client.Create(ctx, chatName, username)
	if err != nil {
		return 0, err
	}

	return chatId, nil
}
