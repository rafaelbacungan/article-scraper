package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
	"log"
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
		_, err := s.ChannelMessageSend(m.ChannelID, "Who enters my domain?")
		if err != nil {
			fmt.Println(err)
		}
	}

	if m.Content == "webscrape" {
		webScrape()
	}
}

// install colly into the application first

// defining a data structure to store the scraped data
type PokemonProduct struct {
	url, image, name, price string
}

// it verifies if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func webScrape() {
	// Initializing the slice of structs that will contain the scraped data
	var pokemonProducts []PokemonProduct

	// Initializing the list of pages to scrape with an empty slice
	var pagesToScrape []string

	// the first pagination URL to scrape
	pageToScrape := "https://scrapeme.live/shop/page/1/"

	// Initializing the list of pages discovered with a pageToScrape
	pagesDiscovered := []string{pageToScrape}

	// current iteration
	i := 1

	// max pages to scrape
	limit := 5

	// initializing a Colly instance
	/*
		Colly's main entity is the Collector. A Collector allows you to perform HTTP requests. Also,
		Also, it gives you access to the web scraping callbacks offered by the Colly interface.
	*/
	c := colly.NewCollector()
	// setting a valid User-Agent header
	/*

		Avoid being blocked!
		- Many websites implement anti-scraping anti-bot techniques. The most basic approach involves banning
		HTTP requests based on their headers. Specifically, they generally block HTTP requests that come with
		an invalid User-Agent header.

		Set a global User-Agent header gor all the requests performed by Colly with the UserAgent Collect field
		as follows:

	*/
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// iterating over the list of pagination links to implement the crawling logic
	c.OnHTML("a.page-numbers", func(e *colly.HTMLElement) {
		// discovering a new page
		newPaginationLink := e.Attr("href")

		// if the page discovered is new
		if !contains(pagesToScrape, newPaginationLink) {
			// if the page discovered should be scraped
			if !contains(pagesDiscovered, newPaginationLink) {
				pagesToScrape = append(pagesToScrape, newPaginationLink)
			}
			pagesDiscovered = append(pagesDiscovered, newPaginationLink)
		}
	})

	// scraping the product data
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	c.OnScraped(func(response *colly.Response) {
		// until there is still a page to scrape
		if len(pagesToScrape) != 0 && i < limit {
			// getting the current page to scrape and removing it from the list

			pageToScrape = pagesToScrape[0]
			pagesToScrape = pagesToScrape[1:]

			// incrementing the iteration counter
			i++
		}
	})

	// visiting the first page
	c.Visit(pageToScrape)

	// opening the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("failed to create output cSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// defining the CSV headers
	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}

	writer.Write(headers)

	// adding each Pokemon product to the CSV output file
	for _, pokemonProduct := range pokemonProducts {
		// converting a Pokemonproduct to an array of strings
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
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

// Conclusion
/*

	In this step-by-step Go tutorial, you saw the building blocks to get started on Golang web scraping.

	As a recap, you learned:
	- How to perform basic data scraping with Go using Colly.
	- How to achieve crawling logic to visit a whole website.
	- The reason why you may need a Go headless browser solution.
	- How to scrape a dynamic-content website with chromedp.

	Scraping can become challenging because of the anti-scraping measures implemented by
	several websites.

*/
