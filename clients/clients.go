package clients

import (
	"context"

	chatserverDesc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
)

type Auth interface {
	Login(ctx context.Context, email, password string) (string, error)
	AccessToken(ctx context.Context, refreshToken string) (string, error)
	CheckAccess(ctx context.Context, endpoint string) error
}

type ChatServer interface {
	Create(ctx context.Context, accessToken, name string, usernames []string) (uint64, error)
	Delete(ctx context.Context, accessToken string, id uint64) error
	SendMessage(ctx context.Context, accessToken string, chatId uint64, text string) error
	Connect(ctx context.Context, accessToken string, chatId uint64) (chatserverDesc.ChatApi_ConnectChatClient, error)
}
