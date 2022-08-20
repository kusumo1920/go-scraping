package scraping

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Scraping(url string) {
	var res, resErr = http.Get(url)

	if resErr != nil {
		log.Fatal(resErr)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	var doc, docErr = goquery.NewDocumentFromReader(res.Body)
	if docErr != nil {
		log.Fatal(docErr)
	}

	var rows = make([]Article, 0)

	doc.Find(".et_pb_ajax_pagination_container").Children().Each(func(i int, s *goquery.Selection) {
		var row = new(Article)
		row.Title = s.Find(".entry-title").Text()
		row.Url, _ = s.Find(".entry-featured-image-url").Attr("href")
		row.Category = s.Find(".post-meta").Find("a").Text()
		
		rows = append(rows, *row)
	})

	var jsonRes, jsonResErr = json.MarshalIndent(rows, "", "  ")
	
	if jsonResErr != nil {
		log.Fatal(jsonResErr)
	}

	fmt.Println(string(jsonRes))
}