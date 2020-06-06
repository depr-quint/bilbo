package bilbo

import (
	"testing"
)

func TestSearch(t *testing.T) {
	items := Search("Peach Pit")
	if len(items) != 1 {
		t.Error("expected one item")
	}

	item := items[0]
	if item.ImageURL == "" {
		t.Error("no image url")
	}

	if item.StockStatus != NoLongerAvailable {
		t.Error("invalid stock status")
	}

	if item.ReleaseDate != "3 april 2020" {
		t.Error("invalid release date")
	}

	if item.Price != 23 {
		t.Error("invalid price")
	}

	if item.Title != "You And Your Friends (LP)" {
		t.Error("invalid title")
	}

	if item.Artist != "Peach Pit" {
		t.Error("invalid artist")
	}

	if item.Type != LP {
		t.Error("invalid type")
	}

	if item.DetailRelativeURL != "/peach-pit-you-and-your-friends-lp" {
		t.Error("invalid detail url")
	}
}
