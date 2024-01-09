package scrap

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func BySelector(url string, cssSelector string) string {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	result := doc.Find(cssSelector).Text()
	return result
}
