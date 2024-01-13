package tester

import (
	"http-stress-test/config"
	"http-stress-test/metrics"
	"http-stress-test/network"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Tester struct {
	config  *config.Configuration
	metrics *metrics.Metrics
}

func NewTester(config *config.Configuration, metrics *metrics.Metrics) *Tester {
	return &Tester{
		config:  config,
		metrics: metrics,
	}
}

func (t *Tester) Run() {
	var wg sync.WaitGroup

	// Calculate RPS slope between points

	for i := 0; i < t.config.Users; i++ {
		wg.Add(1)

		// Calculate interpolated RPS for the current user

		go t.runUser(&wg, t.metrics, 0.5)
	}

	wg.Wait()
	time.Sleep(time.Second)
}

func (t *Tester) runUser(wg *sync.WaitGroup, metrics *metrics.Metrics, interpolatedRPS float64) {
	if t.metrics != nil {
		// Metrics
		metrics.IncrementUserCount()
		defer metrics.DecrementUserCount()
	}

	// Send done to waiting group
	defer wg.Done()

	client := network.NewHttpClient()

	// Calculate force ending time
	var endTime time.Time
	if t.config.Time[len(t.config.Time)-1] > 0 {
		endTime = time.Now().Add(time.Duration(t.config.Time[len(t.config.Time)-1]) * time.Second)
	}

	// Initialize variables for scaling logic
	//nextPointTime := t.config.Time[0]
	nextPointRPS := t.config.Requests[0]
	// Recalculate timeInterval for the new RPS
	var timeInterval = time.Second
	if nextPointRPS > 0 {
		timeInterval = time.Second / time.Duration(nextPointRPS)
		metrics.SetRPS(float64(nextPointRPS))
	}
	timeInterval = 0
	log.Print(timeInterval)

	requestCount := 0
	for {
		// Cancel on max duration
		if !endTime.IsZero() && time.Now().After(endTime) {
			break
		}

		// Random target
		targetIndex := rand.Intn(len(t.config.Targets))
		targetURL := t.config.Targets[targetIndex].URL

		// Send and collect metric
		startTime := time.Now()
		response, err := client.SendRequest(targetURL)
		responseTime := time.Since(startTime)

		// Record metric
		if metrics != nil {
			success := err == nil && response.StatusCode() == 200
			metrics.RecordResponse(responseTime, success)
		}

		// Limiting rate if configured
		if interpolatedRPS > 0 {
			// Adjust the sleep interval to meet the desired RPS
			sleepInterval := time.Second / time.Duration(interpolatedRPS)
			time.Sleep(sleepInterval)
		}

		requestCount++

	}
}
