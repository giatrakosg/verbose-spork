package main

import (
  "fmt"
  "net"
  "encoding/gob"
  "log"
)

var allDone = make(chan string)

// Message type describing data send
type Message struct {
    MsgType int32;
    MsgBuffer []byte
}

func handleConnection(conn net.Conn) {
    fmt.Printf("Accepted connection from %s\n",conn.RemoteAddr())

    // Instantiate new gob decoder for decoding type Message
    // from bytes

    dec := gob.NewDecoder(conn)
	var q Message

    err := dec.Decode(&q)
	if err != nil {
		log.Fatal("Error decoding object in client", err)
	}

    fmt.Printf("RECEIVED %d %s\n", q.MsgType, string(q.MsgBuffer))

	//allDone <- "done"
}

func server() {
  l, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatal("Server: Error in listening   ",err)
  }
  defer l.Close()
  for {
    // Wait for a connection.
    conn, err := l.Accept()
    if err != nil {
      log.Fatal(err)
    }
    // Handle the connection in a new goroutine.
    // The loop then returns to accepting, so that
    // multiple connections may be served concurrently.
    go handleConnection(conn)
  }
  allDone <- "done"

}
// Main function
func main() {
    fmt.Println("Starting server....")
    go server()
    <- allDone
}
