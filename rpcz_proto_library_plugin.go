// Copyright 2013
// Author: Christopher Van Arsdale
// Modified by Mark Vandevorde

package main

import (
  "encoding/json"
  "log"
  "io/ioutil"
  "os"
)

func main() {
  bytes, err := ioutil.ReadAll(os.Stdin)
  if err != nil {
    log.Fatal("Could not read input: ", err)
  }

  raw_input := make(map[string]map[string]interface{})
  err = json.Unmarshal(bytes, &raw_input)
  if err != nil {
    log.Fatal("Could not parse json: ", err)
  }
  node := raw_input["rpcz_proto_library"]
  if node["name"] == nil {
    log.Fatal("Require component Name.")
  }
  if node["generate_java"] ! = nil {
    log.Fatal("rpcz does not support java.")
  }
  if node["generate_go"] ! = nil {
    log.Fatal("rpcz does not support go.")
  }
  if node["generate_py"] ! = nil {
    log.Fatal("rpcz supports python, but rpcz_proto_library_plugin does not.")
  }

  node["translator"] = "protoc"

  // Add "cc": {} section
  cc_section := make(map[string]interface{})
  cc_section["support_library"] = "//third_party/rpcz:rpcz"
  cc_section["translator_args"] = [3]string{
    "--cpp_out=$TRANSLATOR_OUTPUT",
    "--rpcz_out=$TRANSLATOR_OUTPUT",
    "--plugin=protoc-gen-rpcz" }
  cc_source_suffixes := [2]string{ ".pb.cc", ".rpcz.cc" }
  cc_header_suffixes := [2]string{ ".pb.h", ".rpcz.h" }
  cc_section["source_suffixes"] = cc_source_suffixes
  cc_section["header_suffixes"] = cc_header_suffixes
  node["cc"] = cc_section
  node["dependencies"] = [1]string{ "//third_party/rpcz:protoc-gen-rpcz" }

  // Output
  raw_output := make(map[string]map[string]interface{})
  raw_output["translate_and_compile"] = node
  enc := json.NewEncoder(os.Stdout)
  if err := enc.Encode(&raw_output); err != nil {
    log.Fatal("Json encoding error: ", err)
  }
}
