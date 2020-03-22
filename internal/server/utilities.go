package server

import (
	"errors"
	"net"
)

// PanicError if there is any kind of error panic
func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

// checks if tcp server with given address is running
func checkServer(connectionAddress string, message string) error {
	conn, err := net.Dial("tcp", connectionAddress)

	defer conn.Close()
	if err != nil {
		return errors.New("connection to server failed")
	}

	_, err = conn.Write([]byte(message))
	if err != nil {
		return errors.New("sending message to server failed")
	}

	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		return errors.New("reading message from server failed")
	}

	return nil

}
