// RPCZ
[
{ "proto_library": {
    "name": "rpcz_proto",
    "sources": [ "src/rpcz/proto/*.proto" ],
    "generate_cc": true
} },
{ "cc_library": {
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
    "include_dirs": [ "src" ],
    "clang": { "cc_compile_args": [ "-Wno-error=unused-private-field" ] }
} }
]