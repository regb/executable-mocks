package main

import (
    "os"
    "log"
    "crypto/sha256"
    "io"
    "fmt"
)

func main() {
    const bufferSize = 4096
    filePath := os.Args[1]
    fIn, err := os.Open(filePath)
    if err != nil {
	log.Fatal(err)
    }
    defer fIn.Close()

    buf := make([]byte, bufferSize)
    hash := sha256.New()
    _, err = io.CopyBuffer(hash, fIn, buf)
    if err != nil {
	 log.Fatal(err)
    }
    fmt.Printf("%x\n", hash.Sum(nil))
}
