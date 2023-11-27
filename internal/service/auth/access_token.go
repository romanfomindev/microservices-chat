package auth

import (
	"context"

	"github.com/romanfomindev/microservices-chat/clients/grpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *AuthService) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	conn, err := grpc.Dial(
		s.chatServerCfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return "", err
	}

	client := auth.NewAuth(conn)

	accessToken, err := client.AccessToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
