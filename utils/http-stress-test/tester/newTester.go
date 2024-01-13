package tester

import (
	"http-stress-test/config"
	"http-stress-test/metrics"
	"http-stress-test/network"
	"math/rand"
	"sync"
	"time"
)

type NewNewTester struct {
	config  *config.Configuration
	metrics *metrics.Metrics
}

func MakeNewNewTester(config *config.Configuration, metrics *metrics.Metrics) *NewNewTester {
	return &NewNewTester{
		config:  config,
		metrics: metrics,
	}
}

func (t *NewNewTester) NewRun() {
	var wg sync.WaitGroup
	var startTime time.Time
	startTime = time.Now()

	for i := 0; i < t.config.Users; i++ {
		wg.Add(1)

		go t.NewRunUser(&wg, t.metrics, t.config.Time, t.config.Requests, t.config.Users, startTime)
	}

	wg.Wait()
	time.Sleep(time.Second)
}

func getLastIndexGreaterThan(arr []int, input float64) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if float64(arr[i]) < input {
			if i == len(arr) {
				return -1
			}
			return i
		}
	}

	return -1
}

func CalculateCurrentRPS(startTime time.Time, timePoints []int, requestPoints []int) float64 {
	containsZero := false
	for _, v := range timePoints {
		if v == 0 {
			containsZero = true
		}
	}

	// If 0 is not present, prepend it to the array
	if !containsZero {
		timePoints = append([]int{0}, timePoints...)
		requestPoints = append([]int{0}, requestPoints...)
	}
	//get the index of the current timevalue
	startfloat := time.Since(startTime).Seconds()
	if startfloat == 0.0 {
		startfloat = 0.001
	}
	timeIndex := getLastIndexGreaterThan(timePoints, startfloat)
	if timeIndex != -1 {
		currentScaleBetweenPoints := (startfloat - float64(timePoints[timeIndex])) / float64(timePoints[timeIndex+1]-timePoints[timeIndex])
		var currentRPS float64
		if requestPoints[timeIndex] <= requestPoints[timeIndex+1] {
			currentRPS = float64(requestPoints[timeIndex]) + (float64(requestPoints[timeIndex+1]-requestPoints[timeIndex]) * currentScaleBetweenPoints)
		} else {
			currentRPS = float64(requestPoints[timeIndex]) - (float64(requestPoints[timeIndex]-requestPoints[timeIndex+1]) * currentScaleBetweenPoints)
		}
		return currentRPS
	}
	return 0
}

func (t *NewNewTester) NewRunUser(wg *sync.WaitGroup, metrics *metrics.Metrics, timePoints []int, requestPoints []int, users int, startTime time.Time) {
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

	requestCount := 0
	for {
		// Cancel on max duration
		if !endTime.IsZero() && time.Now().After(endTime) {
			break
		}

		RPS := CalculateCurrentRPS(startTime, timePoints, requestPoints)
		metrics.SetRPS(RPS)

		// Random target
		targetIndex := rand.Intn(len(t.config.Targets))
		targetURL := t.config.Targets[targetIndex].URL

		// Send and collect metric
		startTime := time.Now()
		response, err := client.SendRequest(targetURL)
		requestCount++
		responseTime := time.Since(startTime)

		// Record metric
		if metrics != nil {
			success := err == nil && response.StatusCode() == 200
			metrics.RecordResponse(responseTime, success)
		}

		RPS = RPS / float64(users)

		// Limiting rate if configured
		if RPS > 1 {
			// Adjust the sleep interval to meet the desired RPS
			sleepInterval := 1.0 / RPS
			time.Sleep(time.Duration(sleepInterval * float64(time.Second)))
		} else {
			// If interpolatedRPS is zero, simply sleep for one fourth of a second
			time.Sleep(time.Millisecond * 250)
		}
	}
}
