package daemon

import (
	"github.com/cezarmathe/daemon/sockets"
)

type Daemon struct {
	Sockets []*sockets.Socket
}

// // NewSocket creates a new socket at a certain address, assigns it the default buffer size and adds it to the daemon sockets slice
// func (daemon *Daemon) NewSocket(address string) error {
// 	socket, err := sockets.CreateSocket(address, 128, 128)
// 	daemon.Sockets = append(daemon.Sockets, socket)
// 	return err
// }

// // SocketFromAddress tries to find a socket with a certain address in the sockets slice
// func (daemon *Daemon) SocketFromAddress(address string) (*sockets.Socket, error) {
// 	if !strings.HasSuffix(address, ".sock") {
// 		address += ".sock"
// 	}
// 	for i := 0; i < len(daemon.Sockets); i++ {
// 		if daemon.Sockets[i].Address.String() == address {
// 			return daemon.Sockets[i], nil
// 		}
// 	}
// 	return new(sockets.Socket), &ErrNoSocketFound{address}
// }

type Client struct{}

type Server struct {
	Socket *sockets.Socket
}
