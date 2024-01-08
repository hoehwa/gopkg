package goquery

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func queryBySelector(url string, cssSelector string) string {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	result := doc.Find(cssSelector).Text()
	return result
}
