package goroutines

import "testing"

func TestFetchAds(t *testing.T) {

	adsPerSource := FetchAdsFromMultipleSources()
	t.Logf("All ads: %+v", adsPerSource)
}
