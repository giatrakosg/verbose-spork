package main

import (
    "net"
    "context"
    "log"
    "time"
    "fmt"
    "encoding/gob"
)
// Message type describing data send
type Message struct {
    MsgType int32;
    MsgBuffer []byte
}

var messages = make(chan string)

func sendInit(c net.Conn){
    enc := gob.NewEncoder(c)
    err := enc.Encode(Message{1, []byte("Pythagoras")})
	if err != nil {
		log.Fatal("encode error:", err)
	}
    messages <- "Send hello world message"
}
func main()  {
    var d net.Dialer

    ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    defer cancel()

    conn, err := d.DialContext(ctx, "tcp", "localhost:8080")
    if err != nil {
      log.Fatalf("Failed to dial: %v", err)
    }
    defer conn.Close()

    go sendInit(conn)
    fmt.Println(<- messages)


}
