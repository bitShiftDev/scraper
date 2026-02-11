package scraper

type Category struct {
	Name string
	Slug string
	URL string
}

type Product struct {
	Title string
	Price string
	Category Category
	URL string
}