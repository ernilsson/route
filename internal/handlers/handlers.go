package handlers

import (
	"fmt"
	"github.com/ernilsson/router/internal/packet"
)

func HandleLegacyPacket(ctx packet.Context) error {
	fmt.Println("Handling legacy packet")
	return nil
}

func HandleStatePacket(ctx packet.Context) error {
	fmt.Println("Handling state packet")
	return nil
}

func HandleTelemetryPacket(ctx packet.Context) error {
	fmt.Println("Handling telemetry packet")
	err := ctx.Responder.Respond([]byte{0})
	if err != nil {
		return err
	}
	return nil
}
