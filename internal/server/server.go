package server

import (
	"fmt"
	"net"
	"time"
)

//ListenServer is main function which listens to new connections for getting commands
func ListenServer() {
	listen, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	PanicError(err)
	defer listen.Close()
	fmt.Println("listening!")
	for {
		conn, err := listen.Accept()
		PanicError(err)
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)
	timeoutDuration := 5 * time.Second
	fmt.Println("rec")
	err := conn.SetReadDeadline(time.Now().Add(timeoutDuration))
	if err != nil {
		fmt.Println("something wrong")
		return
	}

	readConn, err := conn.Read(buf)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(readConn)

	netst := string(buf)
	fmt.Println(netst)

	conn.Write([]byte("Message received."))

}
