package main

import (
  "fmt"
  "net"
  "encoding/gob"
  "log"
  "time"
)

var allDone = make(chan string)

// Message type describing data send
type Message struct {
    MsgType int32;
    MsgBuffer []byte
}

func handleConnection(conn net.Conn) {
    dec := gob.NewDecoder(conn) // Will read from network.
    // Decode (receive) and print the values.
	var q Message
	err := dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("RECEIVED %d %s\n", q.MsgType, string(q.MsgBuffer))

	//allDone <- "done"
}

func server() {
  l, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatal(err)
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
  go server()
  time.Sleep(2 * time.Second)
  <- allDone
}
