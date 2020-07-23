// The fakeflights is a binary that is like flights, but faster.
//
// Currently it supports only running >>flights --cull-time 202007210000 -o output.qff input1.qff input2.qff input3.qff [...] inputN.qff>>
package main

import (
	"fmt"
    "os"
	"io/ioutil"
)

const N int = 3  // global variable for number of input files

func verifyInput() {
	for i := 0; i < N; i++ {
		fIn, err := os.Open(os.Args[5 + i])
		if err != nil {
			os.Exit(1)
		}

		content, err := ioutil.ReadFile(fIn.Name())
		if err != nil {
			os.Exit(1)
		}

		s := fmt.Sprintf("I'm the input file #%d.\n", i + 1)
		if string(content) != s {  // TODO: Check if the content is right
			os.Exit(1)
		}
	}
	
}

func generateOutput() {
	err := ioutil.WriteFile(os.Args[4], []byte("I'm an output file.\n"), 0666) // TODO: Write right content
    if err != nil {

		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != (5 + N) || os.Args[1] != "--cull-time" || os.Args[2] != "202007210000" || os.Args[3] != "-o" {
    	os.Exit(1)
	}
	verifyInput()
	generateOutput()
	os.Exit(0)
}
