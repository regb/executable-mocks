package main

import (
  "log"

  "github.com/regb/executable-mocks/cmd/bldmock"
  "github.com/google/go-cmp/cmp"
  "google.golang.org/protobuf/testing/protocmp"
  pb "github.com/regb/executable-mocks/protos/mockexec"
)

func Test() {
  for i, tt := range []struct {
    binary string
    args []interface{}
    expected *pb.MockCall
  } {
    {
    binary: "example",
    args: []interface{} { bldmock.StrArg("-t"), bldmock.StrArg("SOME STRING"), bldmock.StrArg("-o"),
     bldmock.OutPath("tmp/test-output"), bldmock.InHash("4cb90a97dd7fc70b16eeefa5c3937769b7a35d7ba1a07db400b8d470dacbf030")},
    expected: &pb.MockCall{
    Binary: "example",
    Args: []*pb.Argument{
          {Arg: &pb.Argument_StrArg{"-t"}},
          {Arg: &pb.Argument_StrArg{"SOME STRING"}},
          {Arg: &pb.Argument_StrArg{"-o"}},
          {Arg: &pb.Argument_OutPath{"tmp/test-output"}},
          {Arg: &pb.Argument_InHash{"4cb90a97dd7fc70b16eeefa5c3937769b7a35d7ba1a07db400b8d470dacbf030"}},
        },
      },
    },

    {
    binary: "example",
    args: []interface{} { bldmock.StrArg("-t"), bldmock.StrArg("SOME STRING"), bldmock.StrArg("-o"),
     bldmock.OutContent("I AM AN OUTPUT"), bldmock.InHash("4cb90a97dd7fc70b16eeefa5c3937769b7a35d7ba1a07db400b8d470dacbf030")},
    expected: &pb.MockCall{
    Binary: "example",
    Args: []*pb.Argument{
          {Arg: &pb.Argument_StrArg{"-t"}},
          {Arg: &pb.Argument_StrArg{"SOME STRING"}},
          {Arg: &pb.Argument_StrArg{"-o"}},
          {Arg: &pb.Argument_OutContent{"I AM AN OUTPUT"}},
          {Arg: &pb.Argument_InHash{"4cb90a97dd7fc70b16eeefa5c3937769b7a35d7ba1a07db400b8d470dacbf030"}},
        },
      },
    },
  } {
    mc := bldmock.BuildMockConfig(tt.binary, tt.args...)
    if diff := cmp.Diff(tt.expected, &mc, protocmp.Transform()); diff != "" {
      log.Fatalf("Testcase %d: the diff is %s.\n", i, diff)
    }
  }
}

// Tests the bldmock.BuildMockConfig function.
func main() {
  Test()
}
