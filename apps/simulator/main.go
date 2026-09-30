package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/ginwan/ecowatch/apps/simulator/models"
)

func main() {
	sensorIDs := []int{1, 2, 3}
	for {
		value := 20 + rand.Float64()*80 // Generate a random value between 20 and 100

		var reading models.Reading
		recordedAt := time.Now().Format(time.RFC3339)
		randomIndex := rand.Intn(len(sensorIDs))
		chosenID := sensorIDs[randomIndex]

		reading.SensorID = chosenID
		reading.Value = value
		reading.RecordedAt = recordedAt

		data, err := json.Marshal(reading)
		if err != nil {
			fmt.Printf("Error marshalling reading: %v\n", err)
			return
		}
		res, err := http.Post("http://localhost:8080/api/v1/readings", "application/json", bytes.NewBuffer(data))
		if err != nil {
			fmt.Printf("Error posting reading: %v\n", err)
			continue
		}
		fmt.Println("Response status:", res.Status)
		res.Body.Close()
		sleepSeconds := 3 + rand.Intn(6)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
