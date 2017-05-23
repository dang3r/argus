package main

import (
	"github.com/dang3r/argus/sensors"
	"log"
)

func main() {
	sensors, err := sensors.Sensors()
	if err != nil {
		log.Fatalf("Error retrieving sensor data : %v\n", err)
	}
	log.Println(sensors)
}
