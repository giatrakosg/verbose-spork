package main

import (
    "net"
    "context"
    "log"
    "time"
    "fmt"
    "encoding/gob"
    // "os/exec"
    "io/ioutil"
    // "os"
)
// Message type describing data send
type Message struct {
    MsgType int32;
    MsgBuffer []byte
}

var messages = make(chan string)

func sendInit(c net.Conn){
    const searchDir = "./data/"
    files, err := ioutil.ReadDir(searchDir)
    if err != nil {
        log.Fatal("error reading dir",err)

    }
    hashDirectory(searchDir,files)
    enc := gob.NewEncoder(c)
    err = enc.Encode(Message{1, []byte("hello")})
  	if err != nil {
  		log.Fatal("encode error:", err)
  	}
    messages <- "Send init message"
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
