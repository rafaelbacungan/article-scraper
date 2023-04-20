package main

import (
	"ArticleScraper/websites"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Variables used for command line parameters
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

	if strings.HasPrefix(m.Content, "/scrape") {
		keyword := grabKeyword(m.Content)
		articles := testWebScrape(keyword)

		// Iterate through the fields of the user struct and create a message embed field for each one.
		fields := make([]*discordgo.MessageEmbedField, 0)

		// TODO: refactor foreach loop
		for _, article := range articles {
			field := &discordgo.MessageEmbedField{
				Name:   "Title",
				Value:  article.title,
				Inline: true,
			}
			fields = append(fields, field)
			field = &discordgo.MessageEmbedField{
				Name:   "Author",
				Value:  article.author,
				Inline: true,
			}
			fields = append(fields, field)
			field = &discordgo.MessageEmbedField{
				Name:   "URL",
				Value:  article.url,
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

func grabKeyword(message string) string {
	keyword := strings.TrimPrefix(message, "/scrape ")
	return keyword
}

// defining a data structure to store the scraped data
type Articles struct {
	title  string
	author string
	url    string
}

func testWebScrape(keyword string) []Articles {
	// initialize articles struct
	var articles []Articles
	// setup tag for filtering purposes
	tag := "#" + keyword

	// Initialize a Colly instance
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// Loop through the list of websites from websites.go, extract their properties and start scraping the needed info
	for _, website := range websites.WebsiteList {
		pageToScrape := website.PageToScrape

		// scrape all articles from the home page
		c.OnHTML("div.crayons-story", func(e *colly.HTMLElement) {
			tags := e.ChildText("a.crayons-tag")
			if containsTag(tags, tag) {
				article := Articles{}
				article.title = e.ChildText(website.TitleHTML)
				article.author = e.ChildText(website.AuthorHTML)
				article.url = pageToScrape + e.ChildAttr(website.UrlHTML, website.UrlAttr)
				articles = append(articles, article)
			}
		})
		// begin scraping the page
		c.Visit(pageToScrape)
	}
	// convert the scraped articles into a csv file.
	return articles
}

func containsTag(tags string, tag string) bool {
	if strings.Contains(tags, tag) {
		return true
	} else {
		return false
	}
}

func csvConvert(articles []Articles) {
	// initializing csv file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("failed to create output cSV file", err)
	}

	defer file.Close()
	// initializing a file writer
	writer := csv.NewWriter(file)

	// defining the CSV headers
	headers := []string{
		"title",
		"author",
		"url",
	}

	writer.Write(headers)

	// adding each Pokemon product to the CSV output file
	for _, article := range articles {
		// converting a Pokemonproduct to an array of strings
		record := []string{
			article.title,
			article.author,
			article.url,
		}
		// writing a new CSV record
		writer.Write(record)
	}
	defer writer.Flush()
}

// Other Web Scraping Libraries for Go
/*

	Other great libraries for web scraping with Golang are:
	- ZenRows: A complete web scraping API that handles all anti-bot bypasses for you.
	It comes with headless browser capabilities, CAPTCHA bypass, rotating proxies and more.
	- GoQuery: A Go library that offers a syntax and a set of features similar to jQuery.
	You can use it to perform web scraping just like you would do in JQuery.
	- Ferret: A portable, extensible and fast web scraping system that aims to simplify
	data extraction from the web. Ferret allows users to focus on the data and is based
	on a unique declarative language.
	- Selenium: Probably the most well-known headless browser, ideal for scraping dynamic
	content. It doesn't offer official support but there's a port to use it in Go.

*/
