package bilbo

import (
	"testing"
)

func TestDetail(t *testing.T) {
	item := Detail("/peach-pit-you-and-your-friends-lp")
	if item.ImageURL == "" {
		t.Error("no image url")
	}

	if item.StockStatus != OutOfStock {
		t.Error("invalid stock status")
	}

	if item.ReleaseDate != "3 april 2020" {
		t.Error("invalid release date")
	}

	if item.Price != 23 {
		t.Error("invalid price")
	}

	if item.Title != "You And Your Friends" {
		t.Error("invalid title")
	}

	if item.Artist != "Peach Pit" {
		t.Error("invalid artist")
	}

	if item.Type != LP {
		t.Error("invalid type")
	}

	if item.SKU != "0194397202014" {
		t.Error("invalid sku")
	}
}
