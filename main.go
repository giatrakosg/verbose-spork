package main

import (
  "fmt"
  "net"
)

func handleConnection(conn net.Conn ) {
  adr := conn.RemoteAddr()
  fmt.Println("Handling conenction type "+adr.Network()+" from " + adr.String())
  fmt.Print("")
  conn.Close()

}

// Main function
func main() {

    fmt.Println("Hello World")

    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
    	// handle error
    }
    for {
    	conn, err := ln.Accept()
    	if err != nil {
    		// handle error
    	}
    	go handleConnection(conn)
    }

}
