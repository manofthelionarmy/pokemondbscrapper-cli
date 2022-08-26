package webscraper

// Builder builds our webscraper
type Builder struct {
	webScaper *WebScraper
}

// NewBuilder makes a new builder
func NewBuilder() *Builder {
	return &Builder{webScaper: &WebScraper{}}
}

// WithURL sets the url for our webscraper
func (b *Builder) WithURL(url string) *Builder {
	b.webScaper.url = url
	return b
}

// Build builds our webscraper and returns it
func (b *Builder) Build() *WebScraper {
	return b.webScaper
}
