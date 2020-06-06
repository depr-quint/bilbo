package bilbo

import (
	"strconv"
	"strings"
)

func parsePrice(price string) float64 {
	priceStr := strings.Replace(strings.TrimSuffix(strings.TrimSpace(price), " €"), ",", ".", -1)
	priceFloat, _ := strconv.ParseFloat(priceStr, 64)
	return priceFloat
}

func parseStockStatus(status string) StockStatus {
	switch {
	case strings.HasPrefix(status, "Op voorraad"):
		return InStock
	default:
		return OutOfStock
	}
}

func splitArtistAndTitle(at string) (string, string) {
	parts := strings.Split(at, " - ")
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}
func splitPriceAndType(otrd string) (float64, ItemType) {
	parts := strings.Split(otrd, " | ")
	if len(parts) != 3 {
		return 0, ""
	}
	return parsePrice(parts[0]), ItemType(parts[1])
}
