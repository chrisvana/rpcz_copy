// RPCZ
[
{ "proto_library": {
    "name": "rpcz_proto",
    "sources": [ "src/rpcz/proto/*.proto" ],
    "generate_cc": true,
    "licenses": [ "http://opensource.org/licenses/Apache-2.0" ]
} },
{ "cc_library": {
    "name": "rpcz",
    "cc_headers": [ "src/rpcz/*.hpp" ],
    "cc_sources": [ "src/rpcz/*.cc" ],
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
} }
]