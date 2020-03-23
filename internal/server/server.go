package server

import (
	"decen_db/internal/cmdmgr"
	"fmt"
	"net"
	"strings"
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
	timeoutDuration := 7 * time.Second
	err := conn.SetReadDeadline(time.Now().Add(timeoutDuration))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// input command list for database
	cmd := strings.Split(string(buf), " ")
	response := cmdmgr.CommandManager(cmd)

	for i := 0; i < len(response); i++ {
		_, err = conn.Write([]byte(response[i]))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
