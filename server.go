package daemon

import (
	"net"

	"github.com/cezarmathe/daemon/sockets"
)

type bgConnMsg struct {
	Conn *net.UnixConn
	Err  error
}

type Server struct {
	Socket      *sockets.Socket
	Connections []*net.UnixConn
}

// NewServer returns a new server
func NewServer() *Server {
	return new(Server)
}

// CreateSocket creates a new socket
func (s *Server) CreateSocket(endpoint string) error {
	var err error
	s.Socket, err = sockets.CreateSocket(endpoint, 1024, 1024)
	return err
}

// Close closes all the connections and then the socket
func (s *Server) Close() error {
	for i := range s.Connections {
		s.Connections[i].Close()
	}
	return s.Socket.Listener.Close()
}

// AcceptConn waits until a connection is received and accepts it
func (s *Server) AcceptConn() (*net.UnixConn, error) {
	conn, err := s.Socket.Listener.AcceptUnix()
	if err == nil {
		s.Connections = append(s.Connections, conn)
	}
	return conn, err
}
