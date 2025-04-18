package ui

import (
	"bufio"
	"fmt"
	"os"
	"sendtg/internal/config"
	"strconv"
)

func StartMenu() {
	config.InitConfig()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n--- sendtg Menu ---")
		fmt.Println("1. Add Chat")
		fmt.Println("2. Delete Chat")
		fmt.Println("3. Add Bot")
		fmt.Println("4. Delete Bot")
		fmt.Println("5. List Config")
		fmt.Println("0. Exit")
		fmt.Print("Choose: ")
		scanner.Scan()
		switch scanner.Text() {
		case "1":
			fmt.Print("Chat name: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Print("Chat ID: ")
			scanner.Scan()
			id, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			config.UpdateChat(name, id)
		case "2":
			fmt.Print("Chat name: ")
			scanner.Scan()
			config.DeleteChat(scanner.Text())
		case "3":
			fmt.Print("Bot name: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Print("Bot token: ")
			scanner.Scan()
			token := scanner.Text()
			config.UpdateBot(name, token)
		case "4":
			fmt.Print("Bot name: ")
			scanner.Scan()
			config.DeleteBot(scanner.Text())
		case "5":
			fmt.Println("Config:", config.GetConfig())
		case "0":
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
