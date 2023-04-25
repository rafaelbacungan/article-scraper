package go_gopher_bot_discord

import (
	"ArticleScraper/scraper"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func CreateSession(Token string) {
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

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "/test" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Who enters my domain?")
		if err != nil {
			fmt.Println(err)
		}
	}

	if strings.HasPrefix(m.Content, "-sauce") {
		keyword := scraper.GrabKeyword(m.Content)
		articles := scraper.TestWebScrape(keyword)

		// Iterate through the fields of the user struct and create a message embed field for each one.
		fields := make([]*discordgo.MessageEmbedField, 0)

		// TODO: refactor foreach loop
		for _, article := range articles {
			field := &discordgo.MessageEmbedField{
				Name:   "Title",
				Value:  article.Title,
				Inline: true,
			}
			fields = append(fields, field)
			field = &discordgo.MessageEmbedField{
				Name:   "Author",
				Value:  article.Author,
				Inline: true,
			}
			fields = append(fields, field)
			field = &discordgo.MessageEmbedField{
				Name:   "URL",
				Value:  article.Url,
				Inline: true,
			}
			fields = append(fields, field)
		}

		// Create a new message embed
		embed := &discordgo.MessageEmbed{
			Color:       0x00ff00, // Set the embed color to green
			Title:       "Articles",
			Description: "Articles retrieved",
			Fields:      fields,
		}

		// Send the message
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			fmt.Println("Error sending message; ", err)
			return
		}
	}
}
