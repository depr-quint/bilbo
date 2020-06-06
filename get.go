package bilbo

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "http://bilborecords.be"

type Path string

const (
	LatestReleases  Path = "/laatste-releases"
	Vinyl           Path = "/laatste-releases-lp"
	ThreeForTwenty  Path = "/3voor20"
	Table           Path = "/tafel-cd"
	Recommendations Path = "/onze-aanraders"
	FutureReleases  Path = "/toekomstige-releases"
	VinylPromo      Path = "/vinyl-promo"
)

func Get(path Path) []Item {
	res, err := http.Get(fmt.Sprintf("%s%s", baseURL, path))
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

	var items []Item
	doc.Find("table td").Each(func(i int, s *goquery.Selection) {
		if s.First().Children().Nodes == nil {
			return
		}

		var item Item

		title := s.Find("span a")
		item.ShortTitle = title.Text()
		item.DetailRelativeURL, _ = title.Attr("href")

		price := s.Find("span.views-field-commerce-price")
		item.Price = parsePrice(price.Text())

		kind := s.Find("span.views-field-field-artikelsoort-tag")
		item.Type = ItemType(strings.TrimSpace(kind.Text()))

		image := s.Find("img")
		item.ImageURL, _ = image.Attr("src")

		items = append(items, item)
	})
	return items
}
