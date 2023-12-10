package service

import "context"

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}

type ChatService interface {
	Create(ctx context.Context, chatName string, username []string) (uint64, error)
	Connect(ctx context.Context, accessToken string, chatId uint64) error
}
