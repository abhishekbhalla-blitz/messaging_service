package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Payload struct {
	Timestamp   string `json:"timestamp"`
	ServiceId   string `json:"service_id"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
	Queue       string `json:"queue"`
}

var payload = Payload{
	Timestamp:   "ab",
	ServiceId:   "nirman",
	Message:     strings.Repeat("some data that has to be sent again", 10),
	MessageType: "KAFKA",
	Queue:       "test-topic2",
}

var marshalledPayload, _ = json.Marshal(payload)

func (messageControllerImpl MessageControllerImpl) StartTest(context *gin.Context) {
	numRequests, _ := strconv.Atoi(context.Query("count"))

	// URL to make the POST request to
	url := "http://127.0.0.1:8081/api/message"

	// Variables to store the times
	var totalDuration time.Duration
	var maxDuration time.Duration
	var minDuration time.Duration

	// Loop to make the request X times
	for i := 0; i < numRequests; i++ {
		// Create a new POST request with JSON payload
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalledPayload))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			return
		}

		// Set appropriate headers
		req.Header.Set("Content-Type", "application/json")

		// Measure the time taken to make the request
		startTime := time.Now()
		client := &http.Client{}
		resp, err := client.Do(req)
		duration := time.Since(startTime)

		if err != nil {
			log.Error("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()

		// Read and discard the response body
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("Error reading response body: %v\n", err)
			return
		}

		// Update the total, max, and min durations
		totalDuration += duration
		if duration > maxDuration {
			maxDuration = duration
		}
		if duration < minDuration {
			minDuration = duration
		}
	}

	// Calculate the average duration
	avgDuration := totalDuration / time.Duration(numRequests)

	// Print the results
	log.Info("Total time: %v\n", totalDuration)
	log.Info("Average time: %v\n", avgDuration)
	log.Info("Max time: %v\n", maxDuration)
	log.Info("Min time: %v\n\n\n", minDuration)
}
