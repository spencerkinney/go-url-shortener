package loadtest

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	endpoint       = "http://localhost:8080/shorten"
	contentType    = "application/json"
	payload        = `{"url": "https://www.example.com"}`
	requestsNumber = 10000 // number of requests to send
	concurrency    = 100  // number of concurrent requests
)

func main() {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	results := make(chan time.Duration, requestsNumber)
	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < requestsNumber/concurrency; j++ {
				start := time.Now()
				_, err := http.Post(endpoint, contentType, strings.NewReader(payload))
				if err != nil {
					fmt.Printf("HTTP request failed: %s\n", err)
					continue
				}
				results <- time.Since(start)
			}
		}()
	}

	wg.Wait()
	close(results)
	totalDuration := time.Since(startTime)

	var totalRequestTime time.Duration
	var durations []time.Duration
	for result := range results {
		totalRequestTime += result
		durations = append(durations, result)
	}

	avgRequestTime := totalRequestTime / time.Duration(requestsNumber)

	// Sort the durations to find the median
	sort.Slice(durations, func(i, j int) bool { return durations[i] < durations[j] })
	medianRequestTime := durations[requestsNumber/2]

	fmt.Printf("Total time for %d requests with %d concurrency: %s\n", requestsNumber, concurrency, totalDuration)
	fmt.Printf("Average request time: %s\n", avgRequestTime)
	fmt.Printf("Median request time: %s\n", medianRequestTime)
	fmt.Printf("Requests per second: %f\n", float64(requestsNumber)/totalDuration.Seconds())

	// Print the min and max request times
	minRequestTime := durations[0]
	maxRequestTime := durations[len(durations)-1]

	// Check if the minRequestTime is zero
	if minRequestTime == 0 {
		fmt.Println("Min request time: < 1ns")
	} else if minRequestTime < time.Microsecond {
		// Convert nanoseconds to float and format the output
		fmt.Printf("Min request time: %.2fns\n", float64(minRequestTime))
	} else {
		// Use the default String() method for durations longer than 1 microsecond
		fmt.Printf("Min request time: %s\n", minRequestTime)
	}

	fmt.Printf("Max request time: %s\n", maxRequestTime)
}