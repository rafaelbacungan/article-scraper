package websites

type Websites struct {
	PageToScrape string
	TitleHTML    string
	AuthorHTML   string
	UrlHTML      string
	UrlAttr      string
}

var WebsiteList = []Websites{
	{"https://dev.to",
		"h2.crayons-story__title",
		"a.crayons-story__secondary",
		"a",
		"href",
	},
}
