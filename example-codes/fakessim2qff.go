// The fakessim2qff is a binary that is like ssim2qff, but faster.
//
// Currently it supports only running >>ssim2qff -t '201809010005' -o outfile -d inputfile>>
package main

import (
	"os"
	"io/ioutil"
	"crypto/sha256"
)

var inputHash = [32]byte{0x27, 0x3f, 0xec, 0x40, 0x9f, 0x32, 0xbd, 0xe, 0xcb, 0xd5, 0xd, 0x3a, 0xb9, 0xac, 0xf1, 0x40, 0x6e, 0xfb, 0x85, 0xf2, 0x4c, 0x89, 0xa2, 0x57, 0xe8, 0xa3, 0xad, 0x3e, 0xab, 0x19, 0x70, 0x46}

func verifyInput() {
	fIn, err := os.Open(os.Args[6])
	if err != nil  {
		os.Exit(1) 
	}
	content, err := ioutil.ReadFile(fIn.Name())
	if err != nil {
		os.Exit(1) 
	}
	if sha256.Sum256(content) != inputHash { 
		os.Exit(1)
	}
}

func generateOutput() {
	err := ioutil.WriteFile(os.Args[4], []byte("I'm an output file.\n"), 0666) // TODO: Write right content
        if err != nil {
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 7 || os.Args[1] != "-t" || os.Args[2] != "201809010005" || os.Args[3] != "-o" || os.Args[5] != "-d"  {
        	os.Exit(1)
	}
	verifyInput()
	generateOutput()
	os.Exit(0)
}
