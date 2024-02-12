package scraper

type RAPLScraper struct {

}

func (r RAPLScraper) configure() bool {
	return True
}

func (r RAPLScraper) getName() string {
	return "RAPL"
}

func (r RAPLScraper) collectMetrics() float64 {
	return 1.0
}