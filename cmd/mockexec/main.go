package main

import (
  "io/ioutil"
  "log"
  "os"
  
  "github.com/golang/protobuf/proto"
  pb "github.com/regb/executable-mocks/protos/mockexec"
)

func main() {
  t, err := ioutil.ReadFile("configs/flights.textproto")
  if err != nil {
    log.Fatal(err)
  }
  c := &pb.CommandLine{}
  if err := proto.UnmarshalText(string(t), c); err != nil {
    log.Fatal(err)
  }
  if len(os.Args) - 1 != len(c.Args) {
    log.Fatalf("Expected %v arguments but got %v", len(c.Args), len(os.Args) - 1)
  }
  
  var outputPath string
  var outputContent string
  for i := 0; i < len(os.Args); i++ {
    if i == 0 {
      // TODO: check the name of the utility.
      continue
    }
    switch x := c.Args[i-1].Arg.(type) {
    case *pb.Argument_StrArg:
      if x.StrArg != os.Args[i] {
        log.Fatalf("String arguments don't match: want '%s'; got '%s'", x.StrArg, os.Args[i])
      }
    case *pb.Argument_InHash:
      // TODO: check the hash.
      log.Printf("Would check input file: %v", os.Args[i])
    case *pb.Argument_OutContent:
      outputPath = os.Args[i]
      outputContent = x.OutContent
    default:
      log.Fatalf("Unexpected argument type %T", x)
    }
  }
  // TODO: if we made it this far and it's necessary, write the desired content
  //       to the output file.
  log.Printf("Would produce '%s' to: %v",outputContent, outputPath)
}
