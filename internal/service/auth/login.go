package auth

import (
	"context"

	"github.com/romanfomindev/microservices-chat/clients/grpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	conn, err := grpc.Dial(
		s.chatServerCfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return "", err
	}

	client := auth.NewAuth(conn)

	refreshToken, err := client.Login(ctx, email, password)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
