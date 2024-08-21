package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const logFilePath = "cpu_temp_log.txt" // Change as needed

func main() {
	// Open log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// Main loop of the daemon
	for {
		temp, err := getCPUTemperature()
		if err != nil {
			log.Println("Error fetching CPU temperature:", err)
		} else {
			log.Printf("Current CPU Temperature: %.2f°C\n", temp)
		}
		time.Sleep(10 * time.Second) // Fetch every 10 seconds
	}
}

// getCPUTemperature retrieves the CPU temperature from the system
func getCPUTemperature() (float64, error) {
	// Read the temperature from the appropriate file
	data, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}

	var temp float64
	// Convert the temperature to Celsius (the value is in millidegrees Celsius)
	fmt.Sscanf(string(data), "%f", &temp)
	temp /= 1000 // Convert from millidegrees to degrees

	return temp, nil
}
