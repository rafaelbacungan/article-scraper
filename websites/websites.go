package websites

type Websites struct {
	PageToScrape string
	TitleHTML    string
	AuthorHTML   string
	UrlHTML      string
	UrlAttr      string
	Tag          string
	ParentDiv    string
}

var WebsiteList = []Websites{
	{"https://dev.to",
		"h2.crayons-story__title",
		"a.crayons-story__secondary",
		"a",
		"href",
		"a.crayons-tag",
		"div.crayons-story",
	},
	{
		"https://www.codecademy.com/articles",
		"h2.gamut-1flqe57-Text",
		"span.card__meta",
		"a.e14vpv2g1",
		"href",
		"span.gamut-1131uxx-Text",
		"li.gamut-1b3oiz6-Column",
	},
}
