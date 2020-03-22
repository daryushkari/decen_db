package server

import (
	"testing"
)

func TestListenServer(t *testing.T) {

	go ListenServer()

	testMessages := []string{"ok"}
	serverAddress := "127.0.0.1:3333"

	for i := 0; i < len(testMessages); i++ {
		err := checkServer(serverAddress, testMessages[i])
		if err != nil {
			t.Error(err)
		}
	}

}
