// RPCZ
[
{ "config": {
    "component": "rpcz",
    "component_root": "src/rpcz"
} },
{ "proto_library": {
    "name": "rpcz_proto",
    "sources": [ "src/rpcz/proto/*.proto" ],
    "dependencies" : [ "../protobuf:proto" ]
} },
{ "cc_binary": {
    "name": "rpcz",
    "cc_headers": [ "src/rpcz/*.hpp" ],
    "cc_sources": [ "src/rpcz/*.cc" ],
    "dependencies": [ ":rpcz_proto",
                      "../zeromq:zeromq",
                      "../boost:system",
                      "../boost:program_options",
                      "../boost:thread",
                      "../protobuf:protoc_lib"
    ],
    "cc_compile_args": [ "-Wno-error=unused-private-field" ]
} }
]