package cmd

import (
	"fmt"
	"os"
	"sendtg/internal/bot"
	"sendtg/internal/chat"
	"sendtg/internal/config"
	"sendtg/internal/file"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sendtg",
	Short: "Send files to Telegram chats using CLI",
	Args:  cobra.MinimumNArgs(2), // Accept 2 or 3 arguments
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()

		filename, chatName := args[0], args[1]
		botName := ""
		if len(args) == 3 {
			botName = args[2]
		}

		botToken, err := bot.ResolveBotToken(botName)
		if err != nil {
			fmt.Println("Error resolving bot:", err)
			return
		}

		chatID, err := chat.ResolveChatID(chatName)
		if err != nil {
			fmt.Println("Error resolving chat:", err)
			return
		}

		toSend, cleanup, err := file.PrepareFile(filename)
		if err != nil {
			fmt.Println("File preparation failed:", err)
			return
		}
		defer cleanup()

		if err := bot.SendFile(botToken, chatID, toSend); err != nil {
			fmt.Println("Send failed:", err)
		} else {
			fmt.Printf("✅ Sent %s to %s using bot %s\n", toSend, chatName, botName)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(menuCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
