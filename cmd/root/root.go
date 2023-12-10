package root

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/romanfomindev/microservices-chat/internal/app"
	"github.com/romanfomindev/microservices-chat/internal/config"
	"github.com/romanfomindev/microservices-chat/internal/storage"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "chat application",
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Auth commands",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		emailStr, err := cmd.Flags().GetString("email")
		if err != nil {
			log.Fatalf("failed to get email: %s\n", err.Error())
		}
		passwordStr, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalf("failed to get password: %s\n", err.Error())
		}
		fmt.Println("data", emailStr, passwordStr)
		serviceProvicer := app.NewServiceProvider()
		refreshToken, err := serviceProvicer.AuthService().Login(ctx, emailStr, passwordStr)

		accessToken, err := serviceProvicer.AuthService().GetAccessToken(ctx, refreshToken)

		fmt.Println("access token: " + accessToken)

	},
}

var accessTokenCmd = &cobra.Command{
	Use:   "access-token",
	Short: "Get Access Token",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

var refreshTokenCmd = &cobra.Command{
	Use:   "refresh-token",
	Short: "Get Refresh Token",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat commands",
}

var createChatCmd = &cobra.Command{
	Use:   "create",
	Short: "Create chat",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken, err := cmd.Flags().GetString("token")
		if err != nil {
			log.Fatalf("failed to get email: %s\n", err.Error())
		}
		storage.SetAccessToken(accessToken)

		nameStr, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("failed to get name: %s\n", err.Error())
		}

		usernames, err := cmd.Flags().GetStringSlice("usernames")
		if err != nil {
			log.Fatalf("failed to get email: %s\n", err.Error())
		}

		ctx := context.Background()
		serviceProvicer := app.NewServiceProvider()
		chatID, err := serviceProvicer.ChatService().Create(ctx, nameStr, usernames)
		if err != nil {
			log.Fatalf("failed to create chat: %s\n", err.Error())
		}

		fmt.Println("chat id: ", chatID)
	},
}

var deleteChatCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete chat",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

var connectChatCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to chat",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken, err := cmd.Flags().GetString("token")
		if err != nil {
			log.Fatalf("failed to get email: %s\n", err.Error())
		}

		chatIDStr, err := cmd.Flags().GetString("chat")
		if err != nil {
			log.Fatalf("failed to get chatID: %s\n", err.Error())
		}

		chatID, err := strconv.ParseUint(chatIDStr, 10, 64)
		if err != nil {
			log.Fatalf("failed to parse chatID: %s\n", err.Error())
		}

		ctx := context.Background()
		serviceProvicer := app.NewServiceProvider()
		err = serviceProvicer.ChatService().Connect(ctx, accessToken, chatID)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	err := config.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	rootCmd.AddCommand(authCmd)

	authCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("email", "e", "", "Email пользователя")
	loginCmd.Flags().StringP("password", "p", "", "Password пользователя")

	createChatCmd.Flags().StringP("token", "t", "", "Access token")
	createChatCmd.Flags().StringP("name", "n", "", "Название чата")
	createChatCmd.Flags().StringSliceP("usernames", "u", []string{}, "Пользователи чата")

	connectChatCmd.Flags().StringP("token", "t", "", "Access token")
	connectChatCmd.Flags().StringP("chat", "c", "", "ID чата")

	authCmd.AddCommand(accessTokenCmd)
	authCmd.AddCommand(refreshTokenCmd)
	rootCmd.AddCommand(chatCmd)
	chatCmd.AddCommand(createChatCmd)
	chatCmd.AddCommand(deleteChatCmd)
	chatCmd.AddCommand(connectChatCmd)

}
