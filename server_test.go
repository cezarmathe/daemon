package daemon

import "testing"

func TestCreateSocket(t *testing.T) {

	server := NewServer()
	t.Log("server created")

	server.CreateSocket("/tmp/go_daemon_server_test")
	t.Log("socket created")

	defer server.Socket.Listener.Close()

	go func() {
		client := new(Client)
		t.Log("client created")

		err := client.OpenConnection("/tmp/go_daemon_server_test")
		t.Log("client connection opened")

		if err != nil {
			t.Fatal(err)
		}

		err = client.CloseConnection()
		t.Log("client connection closed")

		if err != nil {
			t.Fatal(err)
		}
	}()

	_, err := server.AcceptConn()
	t.Log("server connection accepted")

	if err != nil {
		t.Fatal(err)
	}

	err = server.Close()
	t.Log("server closed")

	if err != nil {
		t.Fatal(err)
	}

}
