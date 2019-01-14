package sockets

import (
	"net"
	"strings"
)

// Socket is a container for an address and listener pair
type Socket struct {
	// Address is the address for this socket
	Address *net.UnixAddr
	// Listener is the listener for this socket
	Listener *net.UnixListener
	// // SendChannel is a channel for sending bytes over the socket
	// SendChannel chan<- byte
	// // ReceiveChannel is a channel for receiving bytes over the socket
	// ReceiveChannel <-chan byte
}

// CreateSocket creates a new socket from an endpoint
func CreateSocket(endpoint string, sendChannelBufferSize int, receiveChannelBufferSize int) (*Socket, error) {
	if !strings.HasSuffix(endpoint, ".sock") {
		endpoint += ".sock"
	}

	address, err := net.ResolveUnixAddr("unix", endpoint)
	if err != nil {
		return new(Socket), err
	}

	listener, err := net.ListenUnix("unix", address)
	if err != nil {
		return new(Socket), err
	}

	if sendChannelBufferSize < 0 || receiveChannelBufferSize < 0 {
		return new(Socket), ErrChanNegativeBufSize
	}

	socket := new(Socket)
	socket.Address = address
	socket.Listener = listener
	// socket.ReceiveChannel = make(<-chan byte, receiveChannelBufferSize)
	// socket.SendChannel = make(chan<- byte, sendChannelBufferSize)

	// go socket.receive()
	// go socket.send()

	return socket, nil
}

// CloseSocket attempts to close the listener for a certain socket
func CloseSocket(socket *Socket) error {
	return socket.Listener.Close()
}

// func (socket *Socket) receive() {
// 	for {
// 		// socket.Listener.
// 	}
// }

// func (socket *Socket) send() {
// 	for {

// 	}
// }
