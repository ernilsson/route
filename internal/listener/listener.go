package listener

import (
	"github.com/ernilsson/router/internal/router"
	"net"
)

type Listener interface {
	Run(func(), func(error)) error
	Stop() error
}

func NewUdpSocketListener(port string, router router.Router) Listener {
	return UdpSocketListener{
		port:   port,
		router: router,
	}
}

type UdpSocketListener struct {
	port   string
	conn   *net.UDPConn
	router router.Router
}

func (listener UdpSocketListener) Run(runningCallback func(), errCallback func(error)) error {
	conn, err := listener.createConn()
	if err != nil {
		return err
	}
	listener.conn = conn
	buffer := make([]byte, 4096)
	runningCallback()
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			errCallback(err)
			continue
		}
		if err := listener.router.Consume(buffer[:n-1]); err != nil {
			errCallback(err)
		}
	}
}

func (listener UdpSocketListener) createConn() (*net.UDPConn, error) {
	s, err := net.ResolveUDPAddr("udp4", listener.port)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (listener UdpSocketListener) Stop() error {
	err := listener.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
