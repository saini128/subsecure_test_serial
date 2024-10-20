package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	// Configuration for the virtual serial port.
	config := &serial.Config{
		Name: "/dev/pts/3", // Replace with the virtual port name from socat.
		Baud: 9600,
	}

	// Open the serial port.
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatalf("Error opening serial port: %v", err)
	}
	defer port.Close()

	// Data to be sent.
	message := "W001L002W002L001W003L004W004L005"

	for {
		// Send data over the virtual serial port.
		_, err := port.Write([]byte(message + "\n"))
		if err != nil {
			log.Fatalf("Error writing to serial port: %v", err)
		}
		fmt.Printf("Sent: %s\n", message)

		// Delay to simulate periodic data transmission.
		time.Sleep(2 * time.Second)
	}
}
