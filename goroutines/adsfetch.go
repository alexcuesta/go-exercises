package goroutines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fetchAds(name AdSource) []Ad {
	fmt.Printf(">> Performing request to %s source \n", name)

	delay := rand.Intn(100) + 1
	requestDuration := time.Duration(100 + delay)
	time.Sleep(requestDuration * time.Millisecond)

	fmt.Printf("<< Request to %s took %d milliseconds\n", name, requestDuration)

	ads := []Ad{}
	adSize := rand.Intn(10) + 1

	for i := range adSize {
		newAd := Ad{
			id:   i,
			name: fmt.Sprintf("%s1", name),
		}
		ads = append(ads, newAd)
	}

	return ads
}

func FetchAdsFromMultipleSources() map[AdSource][]Ad {

	result := make(map[AdSource][]Ad)
	var wg sync.WaitGroup // to wait all goroutines to finish
	var mu sync.Mutex     // locks

	for _, adSource := range AllAdSources {
		wg.Add(1)
		go func(source AdSource) {
			defer wg.Done()
			ads := fetchAds(source)

			mu.Lock()
			result[source] = ads
			mu.Unlock()
		}(adSource)
	}

	wg.Wait()
	return result
}
