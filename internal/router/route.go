package router

type Criterion func([]byte) bool

type HandlerFunc func([]byte) error

type Route struct {
	accept Criterion
	handle HandlerFunc
}

func WithRoute(criterion Criterion, handler HandlerFunc) Route {
	return Route{
		accept: criterion,
		handle: handler,
	}
}

func HasPacketType(packetType byte) Criterion {
	return func(packet []byte) bool {
		return len(packet) > 0 && packet[0] == packetType
	}
}
