package bilbo

import "testing"

func TestGet(t *testing.T) {
	for _, path := range []Path{
		LatestReleases,
		Vinyl,
		ThreeForTwenty,
		Table,
		Recommendations,
		FutureReleases,
		VinylPromo,
	} {
		items := Get(path)
		if len(items) == 0 {
			t.Error("no items found")
		}
	}
}
