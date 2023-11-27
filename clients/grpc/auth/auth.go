package auth

import (
	"context"

	accessDesc "github.com/romanfomindev/microservices-auth/pkg/access_v1"
	authDesc "github.com/romanfomindev/microservices-auth/pkg/auth_v1"
	"google.golang.org/grpc"
)

type Auth struct {
	connection *grpc.ClientConn
}

func NewAuth(connection *grpc.ClientConn) *Auth {
	return &Auth{
		connection: connection,
	}
}

func (a *Auth) Login(ctx context.Context, email, password string) (string, error) {
	cl := authDesc.NewAuthServiceClient(a.connection)

	response, err := cl.Login(ctx, &authDesc.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	return response.GetRefreshToken(), nil
}

func (a *Auth) AccessToken(ctx context.Context, refreshToken string) (string, error) {
	cl := authDesc.NewAuthServiceClient(a.connection)
	response, err := cl.GetAccessToken(ctx, &authDesc.GetAccessTokenRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return "", err
	}

	return response.GetAccessToken(), nil
}

func (a *Auth) CheckAccess(ctx context.Context, endpoint string) error {
	cl := accessDesc.NewAccessServiceClient(a.connection)

	_, err := cl.Check(ctx, &accessDesc.CheckRequest{
		EndpointAddress: endpoint,
	})

	return err
}
