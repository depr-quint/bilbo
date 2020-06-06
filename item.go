package bilbo

import (
	"fmt"
)

type ItemType string

const (
	CD  ItemType = "CD"
	LP  ItemType = "LP"
	DVD ItemType = "DVD"
)

type StockStatus string

const (
	InStock           StockStatus = "In Stock"
	OutOfStock        StockStatus = "Out of Stock"
	NoLongerAvailable StockStatus = "No Longer Available"
)

type Item struct {
	ShortTitle        string
	DetailRelativeURL string
	Price             float64
	Type              ItemType
	ImageURL          string
}

func (i Item) String() string {
	return fmt.Sprintf("%-27s %-3s (€ %.2f)", i.ShortTitle, i.Type, i.Price)
}

type ItemDetail struct {
	ImageURL    string
	Price       float64
	ReleaseDate string
	Type        ItemType
	StockStatus StockStatus
	ArtistURL   string
	Title       string
	Artist      string
	Body        string
	SKU         string
}

type ItemSearch struct {
	Artist, Title     string
	Price             float64
	Type              ItemType
	ReleaseDate       string
	StockStatus       StockStatus
	ImageURL          string
	DetailRelativeURL string
}
