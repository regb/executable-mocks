package main

import (
  "log"
  "bytes"
  "io/ioutil"
  
  "github.com/golang/protobuf/proto"
  "github.com/regb/executable-mocks/cmd/bldmock"
)

func testOutPath() {
  mc := bldmock.BuildMockConfig("example", bldmock.StrArg("-t"), bldmock.StrArg("SOME STRING"), bldmock.StrArg("-o"), bldmock.OutPath("tmp/test-output"), bldmock.InHash("4cb90a97dd7fc70b16eeefa5c3937769b7a35d7ba1a07db400b8d470dacbf030"))
  d := proto.MarshalTextString(&mc)
  t, err := ioutil.ReadFile("test/testdata/testconfig_path.textproto")
  if err != nil {
    log.Fatal(err)
  }
  
  if !bytes.Equal([]byte(d), t) {
    log.Fatal("Test OutPath: not working properly.\n")
  }
}

func testOutContent() {
  mc := bldmock.BuildMockConfig("example", bldmock.StrArg("-t"), bldmock.StrArg("SOME STRING"), bldmock.StrArg("-o"), bldmock.OutContent("I AM AN OUTPUT"), bldmock.InHash("4cb90a97dd7fc70b16eeefa5c3937769b7a35d7ba1a07db400b8d470dacbf030"))
  d := proto.MarshalTextString(&mc)
  t, err := ioutil.ReadFile("test/testdata/testconfig_content.textproto")
  if err != nil {
    log.Fatal(err)
  }
  
  if !bytes.Equal([]byte(d), t) {
    log.Fatal("Test OutContent: not working properly.\n")
  }
}

// Tests the bldmock.BuildMockConfig function.
func main() {
  testOutPath()
  testOutContent()
}
