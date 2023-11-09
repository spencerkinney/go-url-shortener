package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

const (
	endpoint       = "http://127.0.0.1:8080/shorten"
	contentType    = "application/json"
	payload        = `{"url": "https://www.example.com"}`
	concurrency    = 10   // Number of concurrent goroutines
	duration       = 15 * time.Second // Test duration
)

func main() {
	var wg sync.WaitGroup
	startTime := time.Now()
	stopTime := startTime.Add(duration)
	var totalRequests int64
	var totalErrors int64

	latencies := make(chan time.Duration, 10000) // Buffer based on expected RPS

	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxConnsPerHost: concurrency,
		},
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for time.Now().Before(stopTime) {
				start := time.Now()
				resp, err := client.Post(endpoint, contentType, bytes.NewReader([]byte(payload)))
				latency := time.Since(start)
				
				if err != nil {
					fmt.Printf("HTTP request failed: %s\n", err)
					atomic.AddInt64(&totalErrors, 1)
					continue
				}

				if resp.StatusCode < 200 || resp.StatusCode >= 300 {
					fmt.Printf("HTTP request returned non-2xx status: %d\n", resp.StatusCode)
					atomic.AddInt64(&totalErrors, 1)
				} else {
					latencies <- latency
					atomic.AddInt64(&totalRequests, 1)
				}

				resp.Body.Close()
			}
		}()
	}

	wg.Wait()
	close(latencies)

	// Collect and analyze the results
	var totalLatency time.Duration
	var maxLatency time.Duration
	durations := make([]time.Duration, 0, totalRequests)

	for latency := range latencies {
		totalLatency += latency
		if latency > maxLatency {
			maxLatency = latency
		}
		durations = append(durations, latency)
	}

	// Sort latencies to calculate median
	sort.Slice(durations, func(i, j int) bool { return durations[i] < durations[j] })

	var median time.Duration
	if len(durations) > 0 {
		median = durations[len(durations)/2]
	}
	mean := time.Duration(atomic.LoadInt64(&totalRequests) / totalRequests)

	fmt.Printf("Total time taken for test: %s\n", duration)
	fmt.Printf("Total number of requests: %d\n", totalRequests)
	fmt.Printf("Total number of errors: %d\n", totalErrors)
	fmt.Printf("Requests per second: %.2f\n", float64(totalRequests)/duration.Seconds())
	fmt.Printf("Mean request time: %s\n", mean)
	fmt.Printf("Median request time: %s\n", median)
	fmt.Printf("Max request time: %s\n", maxLatency)
}