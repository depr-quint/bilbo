package bilbo

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func Detail(path string) ItemDetail {
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

	var detail ItemDetail
	detail.ImageURL, _ = doc.Find("div.field-name-field-afbeelding img").First().Attr("src")
	detail.Price = parsePrice(doc.Find("div.field-type-commerce-price div.field-item").First().Text())
	detail.ReleaseDate = doc.Find("div.field-type-date div.field-item").First().Text()
	detail.Type = ItemType(doc.Find("div.vocabulary-artikelsoort h2 a").First().Text())
	detail.StockStatus = parseStockStatus(doc.Find("div.field-name-commerce-stock span").First().Text())
	detail.ArtistURL, _ = doc.Find("div.field-name-field-artiest-tag a").First().Attr("href")
	detail.Title = doc.Find("div.field-name-field-titel div.field-item").First().Text()
	detail.Artist = doc.Find("div.field-name-field-artiest div.field-item").First().Text()
	detail.Body = doc.Find("cite").First().Text()
	detail.SKU = strings.TrimSpace(doc.Find("div.commerce-product-sku").First().Contents().Not("div").Text())
	return detail
}
