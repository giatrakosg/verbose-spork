package main

import (
    "crypto/sha256"
    "io"
    "log"
    "os"
    "fmt"
)

func hashFile(f *os.File)([]byte)  {
    h := sha256.New()
    if _, err := io.Copy(h, f); err != nil {
        log.Fatal(err)
    }
    sum := h.Sum(nil)
    fmt.Printf("hash of file: %s\n %x",f.Name(), sum)

    return sum
}
