package transport

import (
	"fmt"
	"net"
)

type Transport struct {
	listen      *net.UDPConn
	destination *net.UDPAddr
}

// NewConnection initiates a new UDP ipv4 connection to destinationAddr.
// It was written without any test. This is BAD.
func NewConnection(destinationAddr string) (*Transport, error) {
	listenAddr, err := net.ResolveUDPAddr("udp4", ":0")
	if err != nil {
		return nil, err
	}
	list, err := net.ListenUDP("udp4", listenAddr)
	if err != nil {
		return nil, err
	}

	addr, err := net.ResolveUDPAddr("udp4", destinationAddr)
	if err != nil {
		return nil, err
	}
	return &Transport{
		listen:      list,
		destination: addr,
	}, nil

}

// Write p to the open UDP connection (implements io.Writer)
func (t Transport) Write(p []byte) (n int, err error) {
	return t.listen.WriteTo(p, t.destination)
}

// Close must be called once we are done using this connection
func (t Transport) Close() {
	t.listen.Close()
}

// String provides a representation of this transport (listening address, emitting address)
func (t Transport) String() string {
	return fmt.Sprintf("{listens %s, sends %s}", t.listen.LocalAddr().String(), t.destination.String())
}
