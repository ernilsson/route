package packet

import (
	"fmt"
	"net"
)

type Context struct {
	Payload   []byte
	Responder Responder
}

type Responder interface {
	Respond([]byte) error
}

type UdpResponder struct {
	addr *net.UDPAddr
}

func (responder UdpResponder) Respond(packet []byte) error {
	fmt.Printf("Sending UDP response to host: %v\n", responder.addr.String())
	conn, err := net.Dial("udp4", responder.addr.String())
	if err != nil {
		return err
	}
	_, err = conn.Write(packet)
	return err
}

type Listener interface {
	Run(func(), func(error)) error
	Router(Router)
	Stop() error
}

func NewUdpSocketListener(port string) Listener {
	return &UdpSocketListener{
		port: port,
	}
}

type UdpSocketListener struct {
	port   string
	conn   *net.UDPConn
	router Router
}

func (listener *UdpSocketListener) Router(r Router) {
	listener.router = r
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
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			errCallback(err)
			continue
		}
		ctx := Context{
			Payload:   buffer[:n-1],
			Responder: UdpResponder{addr: addr},
		}
		go listener.handlePacket(ctx, errCallback)
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

func (listener UdpSocketListener) handlePacket(ctx Context, errCallback func(error)) {
	if err := listener.router.Consume(ctx); err != nil {
		errCallback(err)
	}
}

func (listener UdpSocketListener) Stop() error {
	err := listener.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
