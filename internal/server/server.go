package server

import (
	"bufio"
	"decen_db/internal/parser"
	"fmt"
	"net"
	"os"
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

	req, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// input command list for database
	cmd := strings.Split(string(buf)[:req], " ")
	//Todo: fix input later
	response := parser.CommandManager(cmd[0])

	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println(err)
		return
	}

}

func EnterCmd() {
	for {
		fmt.Println("please enter command:")
		// Todo:
		in := bufio.NewReader(os.Stdin)
		cmd, err := in.ReadString('\n')
		cmd = cmd[:len(cmd)-1]
		if err != nil {
			panic(err)
		}
		if cmd == "exit" {
			return
		}

		response := parser.CommandManager(cmd)
		fmt.Println(response)
	}
}
