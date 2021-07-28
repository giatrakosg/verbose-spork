package main

import (
    "net"
    "context"
    "log"
    "time"
    "fmt"
    "encoding/gob"
    // "os/exec"
    // "io/ioutil"
    "os"
)
// Message type describing data send
type Message struct {
    MsgType int32;
    MsgBuffer []byte
}

var messages = make(chan string)

func sendInit(c net.Conn){

    // files, err := ioutil.ReadDir("./data/")
    // if err != nil {
    //     log.Fatal(err)
    // }
    fd, _  := os.Open("./data/file1.txt")
    go hashFile(fd)

    // for _, file := range files {
    //     fd, _  := os.Open(file.Name())
    //     go hashFile(fd)
    //
    // }

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
