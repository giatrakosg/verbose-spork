package main

import (
  "fmt"
  "net"
  "io"
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
  // make a temporary bytes var to read from the connection
	tmp := make([]byte, 1024)
	// make 0 length data bytes (since we'll be appending)
	data := make([]byte, 0)
	// keep track of full length read
	length := 0

	// loop through the connection stream, appending tmp to data
	for {
		// read to the tmp var
		n, err := conn.Read(tmp)
		if err != nil {
			// log if not normal error
			if err != io.EOF {
				fmt.Printf("Read error - %s\n", err)
			}
			break
		}

		// append read data to full data
		data = append(data, tmp[:n]...)

		// update total read var
		length += n
	}

	// log bytes read
	fmt.Printf("READ  %d bytes\n", length)
  fmt.Println("message ",string(data))

  conn.Write([]byte("ok"))
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
