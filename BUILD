// RPCZ
[
{ "config": {
  "component": "third_party/rpcz",
  "plugins": [ "//third_party/protobuf:proto_library" ]
} },
{ "plugin": {
    "name": "rpcz_proto_library",
    "command": "cd $SRC_DIR; go run rpcz_proto_library_plugin.go"
} },
{ "proto_library": {
    "name": "rpcz_proto",
    "sources": [ "src/rpcz/proto/rpcz.proto" ],
    "generate_cc": true,
    "licenses": [ "http://opensource.org/licenses/Apache-2.0" ]
  }
},
{ "cc_library": {
    "name": "rpcz",
    "cc_headers": [ "src/rpcz/*.hpp" ],
    "cc_sources": [ "src/rpcz/application.cc",
    		    "src/rpcz/clock.cc",
		    "src/rpcz/connection_manager.cc",
		    "src/rpcz/reactor.cc",
		    "src/rpcz/rpc.cc",
		    "src/rpcz/rpc_channel_impl.cc",
		    "src/rpcz/server.cc",
		    "src/rpcz/sync_event.cc",
		    "src/rpcz/zmq_utils.cc" ],
    "dependencies": [ ":rpcz_proto",
                      "//third_party/zeromq:zeromq",
                      "//third_party/boost:system",
                      "//third_party/boost:program_options",
                      "//third_party/boost:thread",
                      "//third_party/protobuf:protoc_lib"
    ],
    "cc_include_dirs": [ "src" ],
    "clang": { "cc_compile_args": [ "-Wno-error=unused-private-field" ] },
    "licenses": [ "http://opensource.org/licenses/Apache-2.0" ]
} },
{ "cc_binary": {
    "name": "protoc-gen-rpcz",
    "cc_headers": [ "src/rpcz/plugin/common/*.h", "src/rpcz/plugin/cpp/*.h" ],
    "cc_sources": [ "src/rpcz/plugin/cpp/file_generator.cc",
    		    "src/rpcz/plugin/cpp/rpcz_cpp_generator.cc",
		    "src/rpcz/plugin/cpp/rpcz_cpp_main.cc",
		    "src/rpcz/plugin/cpp/rpcz_cpp_service.cc" ],
    "dependencies": [ ":rpcz_proto",
    		      "//third_party/protobuf:protoc_lib"
		      ],
    "cc_include_dirs": [ "src" ],
    "licenses": [ "http://opensource.org/licenses/Apache-2.0" ]
} }
]