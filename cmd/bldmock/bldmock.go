package bldmock

import (
  "log"
  
  pb "github.com/regb/executable-mocks/protos/mockexec"
)

type StrArg string
type InHash string
type OutPath string
type OutContent string

// Takes the name of the binary and a sequence of arguments and returns a MockCall.
// Each argument is a string of one of the following four types: StrArg, InHash, OutPath, OutContent
func BuildMockConfig(name string, args ...interface{}) pb.MockCall {
  
  mc := pb.MockCall {
    Binary: name,
    Args: []*pb.Argument{},
  }
  
  for _, arg := range args {
    switch t := arg.(type) {
      case StrArg:
        mc.Args = append(mc.Args, &pb.Argument{Arg: &pb.Argument_StrArg{string(arg.(StrArg))}})
      case InHash:
        mc.Args = append(mc.Args, &pb.Argument{Arg: &pb.Argument_InHash{string(arg.(InHash))}})
      case OutPath:
        mc.Args = append(mc.Args, &pb.Argument{Arg: &pb.Argument_OutPath{string(arg.(OutPath))}})
      case OutContent:
        mc.Args = append(mc.Args, &pb.Argument{Arg: &pb.Argument_OutContent{string(arg.(OutContent))}})
      default:
        log.Fatalf("Unexpected argument type %T", t)
    }
  }
  
  return mc
}

