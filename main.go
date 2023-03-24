package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

// Variables used for command line parameters
// Token MTA4NjQ0NTEwNjI1MDM5OTc3Ng.GrxLcW.6MZFhvGTkMeJs1OLE5-Mi-LwThswdix37V11gU
var (
	Token string
)

// Initialize the token variable which will be needed as a parameter for the Bot app.
func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

// create discord session
func main() {

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate function as a callback for MessageCreate event
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until Ctrl-C or other term signal is received.
	fmt.Println("Bot is now running. Press Ctrl-C to exit.")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	dg.Close()
}

type Gopher struct {
	Name string `json: "name"`
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "test" {
		fmt.Println("Helo")
	}

}
