package chat

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	chat_server "github.com/romanfomindev/microservices-chat/clients/grpc/chat-server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *ChatService) Connect(ctx context.Context, accessToken string, chatId uint64) error {
	fmt.Println("ChatService.Connect")
	conn, err := grpc.Dial(
		s.cfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	client := chat_server.NewChatServer(conn)

	stream, err := client.Connect(ctx, accessToken, chatId)
	fmt.Println("ChatService.Connect err:", err)
	if err != nil {
		return err
	}

	go func() {
		for {
			message, errRecv := stream.Recv()
			if errRecv == io.EOF {
				return // stream was closed by server
			}
			if errRecv != nil {
				log.Println("failed to receive message from stream: ", errRecv)
				return
			}

			log.Printf("[%v] - [from: %s]: %s\n",
				color.YellowString(message.GetCreatedAt().AsTime().Format(time.RFC3339)),
				color.BlueString(message.GetFrom()),
				message.GetText(),
			)
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		var lines strings.Builder

		for {
			scanner.Scan()
			line := scanner.Text()
			if len(line) == 0 {
				break
			}

			lines.WriteString(line)
		}

		err = scanner.Err()
		if err != nil {
			log.Println("failed to scan message: ", err)
		}

		text := lines.String()

		err := client.SendMessage(ctx, chatId, text)
		if err != nil {
			log.Println("failed to send message: ", err)
			return err
		}
	}
}
