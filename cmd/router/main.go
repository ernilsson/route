package main

import (
	"fmt"
	"github.com/ernilsson/router/internal/handlers"
	"github.com/ernilsson/router/internal/packet"
)

func main() {
	r := packet.NewRouter(
		packet.WithRoute(packet.HasPacketType(0), handlers.HandleLegacyPacket),
		packet.WithRoute(packet.HasPacketType(1), handlers.HandleStatePacket),
		packet.WithRoute(packet.HasPacketType(2), handlers.HandleTelemetryPacket),
	)

	l := packet.NewUdpSocketListener("localhost:1234")
	l.Router(r)
	if err := l.Run(onListenerRunning, onListenerErr); err != nil {
		panic(err)
	}
}

func onListenerRunning() {
	fmt.Println("Running udp server")
}

func onListenerErr(err error) {
	fmt.Printf("Failed to read from listener: %v\n", err)
}
