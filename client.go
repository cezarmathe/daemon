package daemon

import (
	"net"
	"strings"
)

type Client struct {
	Connection   *net.UnixConn
	LocalAddress *net.UnixAddr
}

// OpenConnection opens a connection at a certain address
func (client *Client) OpenConnection(address string) error {

	if !strings.HasSuffix(address, ".sock") {
		address += ".sock"
	}

	raddr, err := net.ResolveUnixAddr("unix", address)
	if err != nil {
		return err
	}

	conn, err := net.DialUnix("unix", nil, raddr)
	if err != nil {
		return err
	}
	client.Connection = conn
	client.LocalAddress = nil
	return nil
}

// CloseConnection close the client's connection to the socket
func (client *Client) CloseConnection() error {
	return client.Connection.Close()
}
