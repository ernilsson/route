package handlers

import "fmt"

func HandleLegacyPacket(packet []byte) error {
	fmt.Println("Handling legacy packet")
	return nil
}

func HandleStatePacket(packet []byte) error {
	fmt.Println("Handling state packet")
	return nil
}

func HandleTelemetryPacket(packet []byte) error {
	fmt.Println("Handling telemetry packet")
	return nil
}
