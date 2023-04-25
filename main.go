package main

import (
	go_gopher_bot_discord "ArticleScraper/go-gopher-bot-discord"
	"flag"
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
	go_gopher_bot_discord.CreateSession(Token)
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
