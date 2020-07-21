// The fakessim2qff is a binary that is like ssim2qff, but faster.
//
// Currently it supports only running >>ssim2qff -t '201809010005' -o outfile -d inputfile>>
package main

import (
        "os"
	"log"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 7 || os.Args[1] != "-t" || os.Args[2] != "201809010005" || os.Args[3] != "-o" || os.Args[5] != "-d"  {
                os.Exit(1)
        }
        fIn, err := os.Open(os.Args[6])
	if err != nil  {
		log.Fatal(err) 
	}
	content, err := ioutil.ReadFile(fIn.Name())
	if err != nil {
		log.Fatal(err)
	}
	if string(content) != "I'm an input file.\n" {  // TODO: Check if the content is right
		os.Exit(1)
	}
	err = ioutil.WriteFile(os.Args[4], []byte("I'm an output file.\n"), 0666) // TODO: Write right content
        if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}