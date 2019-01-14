package daemon

import (
	"testing"
)

func TestNewSocket(t *testing.T) {

	daemon := new(Daemon)

	for i := 0; i < 16; i++ {
		err := daemon.NewSocket("/tmp/go_test_TestNewSocket" + string(i))
		if err != nil {
			t.Fatal(err)
		}
		daemon.Sockets[i].Listener.Close()
	}

}

func TestSocketFromAddress(t *testing.T) {

	daemon := new(Daemon)

	for i := 0; i < 16; i++ {
		err := daemon.NewSocket("/tmp/go_test_TestNewSocket" + string(i))
		if err != nil {
			t.Fatal(err)
		}
		daemon.Sockets[i].Listener.Close()
	}
	deferClose := func(index int) {
		for i := 0; i <= index; i++ {
			daemon.Sockets[i].Listener.Close()
		}
	}

	for i := 0; i < 16; i++ {
		socket, err := daemon.SocketFromAddress("/tmp/go_test_TestNewSocket" + string(i))
		if err != nil {
			deferClose(i)
			t.Fatal(err)
		}
		socket.Listener.Close()
	}

}
