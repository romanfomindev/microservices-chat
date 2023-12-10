package interceptor

import (
	"context"

	"github.com/romanfomindev/microservices-chat/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AccessClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md := metadata.New(map[string]string{"Authorization": "Bearer " + storage.GetAccessToken()})
		ctx = metadata.NewOutgoingContext(ctx, md)
		err := invoker(ctx, method, req, reply, cc, opts...)

		return err
	}
}
