package main

import (
    "crypto/sha256"
    "io"
    "log"
    "os"
    "fmt"
    "encoding/hex"
    // "context"
)


// FileHashPair Type for pairs of filenames and hashes
type FileHashPair struct {
    FileName string ;
    FileHash string
}

// ListMessage Type for pairs of filenames and hashes
type ListMessage struct {
    Size int32;
    DirHashes []FileHashPair

}
func Create() (*ListMessage) {
    return new(ListMessage)
}
func hashFile(f *os.File,res chan<- FileHashPair)  {
    h := sha256.New()
    if _, err := io.Copy(h, f); err != nil {
        log.Fatal(err)
    }
    var fhp FileHashPair
    fhp.FileName = f.Name()
    fhp.FileHash = hex.EncodeToString(h.Sum(nil))

    res <- fhp
}

func hashDirectory(path string,files []os.FileInfo ) {
    hashPairs := make(chan FileHashPair,5)

    for _, file := range files {
        fd, _  := os.Open(path+file.Name())
        defer fd.Close()
        go hashFile(fd,hashPairs)
    }

    results := make([]FileHashPair, len(files))
    for i := range results {
        results[i] = <-hashPairs
    }
    for i := range results {
        fmt.Printf("Hash pair{%s,%s}\n",results[i].FileName,results[i].FileHash)
    }
}
