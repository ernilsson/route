package main

import (
	"fmt"
	"github.com/ernilsson/router/internal/handlers"
	"github.com/ernilsson/router/internal/listener"
	"github.com/ernilsson/router/internal/router"
)

func main() {
	r := router.NewRouter(
		router.WithRoute(router.HasPacketType(0), handlers.HandleLegacyPacket),
		router.WithRoute(router.HasPacketType(1), handlers.HandleStatePacket),
		router.WithRoute(router.HasPacketType(2), handlers.HandleTelemetryPacket),
	)

	l := listener.NewUdpSocketListener("localhost:1234", r)
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
