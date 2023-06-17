package main

import (
	go_gopher_bot_discord "ArticleScraper/go-gopher-bot-discord"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Variables used for command line parameters
var (
	Token string
)

// Initialize the token variable which will be needed as a parameter for the Bot app.
func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")

	flag.StringVar(&Token, "botToken", botToken, "Bot Token")
	flag.Parse()
}

// create discord session
func main() {
	go_gopher_bot_discord.CreateSession(Token)
}
