package scraper

type scraper interface {
	configure() bool
	getName() string
	collectMetrics() float64
}