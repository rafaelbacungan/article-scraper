package scraper

import (
	"ArticleScraper/websites"
	"github.com/gocolly/colly"
	"strings"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.

func GrabKeyword(message string) string {
	keyword := strings.TrimPrefix(message, "-sauce ")
	return keyword
}

// defining a data structure to store the scraped data
type Articles struct {
	Title  string
	Author string
	Url    string
}

func TestWebScrape(keyword string) []Articles {
	// initialize articles struct
	var articles []Articles

	// Initialize a Colly instance
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// Loop through the list of websites from websites.go, extract their properties and start scraping the needed info
	for _, website := range websites.WebsiteList {
		/*
			@pageToScrape string - a string containing the url of the webpage that needs scraping
			@article struct - initiate an Articles struct to store the scraped data that will be eventually
			pushed to the articles array
			@burst int - burst is used to retrieve a maximum of three articles per website to keep messages reasonable
			for the bot and not to exceed character limits when sending messages.
		*/
		pageToScrape := website.PageToScrape
		article := Articles{}
		burst := 0

		// scrape all articles from the home page
		c.OnHTML(website.ParentDiv, func(e *colly.HTMLElement) {
			tags := e.ChildText(website.Tag)

			if containsTag(strings.ToLower(tags), strings.ToLower(keyword)) {

				if burst%3 == 0 && burst != 0 {
					return
				} else {
					burst += 1
					article.Title = e.ChildText(website.TitleHTML)
					article.Author = e.ChildText(website.AuthorHTML)
					article.Url = pageToScrape + e.ChildAttr(website.UrlHTML, website.UrlAttr)
					articles = append(articles, article)
				}
			}
		})

		// visit the page to begin scraping
		c.Visit(pageToScrape)
	}
	// return the articles for the message to be processed.
	return articles
}

// retrieves the user input and checks for any similarities with the scraped data
func containsTag(tags string, tag string) bool {
	if strings.Contains(tags, tag) {
		return true
	} else {
		return false
	}
}
