package bilbo

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func Search(query string) []ItemSearch {
	res, err := http.Get(fmt.Sprintf("%s/search?search=%s", baseURL, strings.Replace(query, " ", "+", -1)))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var items []ItemSearch
	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		if s.First().Children().Nodes == nil {
			return
		}

		var item ItemSearch

		item.ImageURL, _ = s.Find("td.views-field-field-afbeelding img").Attr("src")
		item.DetailRelativeURL, _ = s.Find("td.views-field-field-afbeelding a").Attr("href")
		artistAndTitle := s.Find("h1 a")
		item.Artist, item.Title = splitArtistAndTitle(artistAndTitle.Text())
		priceAndType := s.Find("h3").Contents().Not("span")
		item.Price, item.Type = splitPriceAndType(priceAndType.Text())
		item.ReleaseDate = s.Find("h3 span").Text()
		item.StockStatus = parseStockStatus(s.Find("span").Text())
		if strings.Contains(s.Contents().Not("br").Text(), "Dit item is niet langer beschikbaar") {
			item.StockStatus = NoLongerAvailable
		}
		items = append(items, item)
	})
	return items
}
