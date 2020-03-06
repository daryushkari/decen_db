package server

import (
    "fmt"
	"net"
	"../utilities"
)


//ListenServer is 
func ListenServer() {
    listen, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    utilities.PanicError(err)

	defer listen.Close()

	for {

		conn, err := listen.Accept()
        if err != nil {
            panic(err)
        }

		go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 32)
  // Read the incoming connection into the buffer.
  readconn, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  fmt.Println(readconn)
  
  netst := string(buf)
  fmt.Println(netst)

  // Send a response back to person contacting us.
  conn.Write([]byte("Message received."))
  // Close the connection when you're done with it.
  conn.Close()
}
