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
	config          *config.Configuration
	metrics         *metrics.Metrics
	waitForResponse bool
}

func MakeNewNewTester(config *config.Configuration, metrics *metrics.Metrics, waitForResponse bool) *NewNewTester {
	return &NewNewTester{
		config:          config,
		metrics:         metrics,
		waitForResponse: waitForResponse,
	}
}

func (t *NewNewTester) NewRun() {
	var wg sync.WaitGroup
	var startTime time.Time
	startTime = time.Now()

	for i := 0; i < t.config.Users; i++ {
		wg.Add(1)

		go t.NewRunUser(&wg, t.metrics, t.config.Time, t.config.Requests, t.config.Users, startTime, t.waitForResponse)
	}

	wg.Wait()
	time.Sleep(time.Second)
}

func getLastIndexGreaterThan(arr []int, input float64) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if float64(arr[i]) < input {
			if (i + 1) == len(arr) {
				return -1
			}
			return i
		}
	}
	return -1
}

func CalculateCurrentRPS(startTime time.Time, timePoints []int, requestPoints []int) float64 {
	if timePoints[0] != 0 {
		timePoints = append([]int{0}, timePoints...)
		requestPoints = append([]int{0}, requestPoints...)
	}

	start := time.Since(startTime).Seconds()
	if start == 0.0 {
		start = 0.001
	}
	timeIndex := getLastIndexGreaterThan(timePoints, start)
	if timeIndex != -1 {
		currentScaleBetweenPoints := (start - float64(timePoints[timeIndex])) / float64(timePoints[timeIndex+1]-timePoints[timeIndex])
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

func (t *NewNewTester) NewRunUser(wg *sync.WaitGroup, metrics *metrics.Metrics, timePoints []int, requestPoints []int, users int, startTime time.Time, waitForResponse bool) {
	if t.metrics != nil {
		metrics.IncrementUserCount()
		defer metrics.DecrementUserCount()
	}

	defer wg.Done()

	var endTime time.Time
	if t.config.Time[len(t.config.Time)-1] > 0 {
		endTime = time.Now().Add(time.Duration(t.config.Time[len(t.config.Time)-1]) * time.Second)
	}

	requestCount := 0
	for {
		if !endTime.IsZero() && time.Now().After(endTime) {
			break
		}

		RPS := CalculateCurrentRPS(startTime, timePoints, requestPoints)
		metrics.SetRPS(RPS)

		targetIndex := rand.Intn(len(t.config.Targets))
		targetURL := t.config.Targets[targetIndex].URL

		startTime := time.Now()
		if waitForResponse {
			client := network.NewHttpClient()
			response, err := client.SendRequest(targetURL)
			requestCount++
			responseTime := time.Since(startTime)

			if metrics != nil {
				success := err == nil && response.StatusCode() == 200
				metrics.RecordResponse(responseTime, success)
			}
		} else {
			fastclient := network.NewTcpClient()
			go fastclient.Send(targetURL)
			responseTime := time.Since(startTime)
			requestCount++

			if metrics != nil {
				metrics.RecordResponse(responseTime, true)
			}
		}

		RPS = RPS / float64(users)

		if RPS > 1 {
			sleepInterval := 1.0 / RPS
			time.Sleep(time.Duration(sleepInterval * float64(time.Second)))
		} else {
			time.Sleep(time.Millisecond * 250)
		}
	}
}
