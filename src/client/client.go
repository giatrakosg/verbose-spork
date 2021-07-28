package main

import (
    "net"
    "context"
    "log"
    "time"
    "fmt"
    "encoding/gob"
    "os/exec"
    "io/ioutil"
)
// Message type describing data send
type Message struct {
    MsgType int32;
    MsgBuffer []byte
}

var messages = make(chan string)

func sendInit(c net.Conn){

    files, err := ioutil.ReadDir("/tmp/")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }

    enc := gob.NewEncoder(c)
    err = enc.Encode(Message{1, []byte("Pythagoras")})
  	if err != nil {
  		log.Fatal("encode error:", err)
  	}
    cmd := exec.Command("ls")
    log.Printf("Running command and waiting for it to finish...")
    err = cmd.Run()
    log.Printf("Command finished with error: %v", err)
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
