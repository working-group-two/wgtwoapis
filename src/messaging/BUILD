load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@build_stack_rules_proto//go:go_grpc_compile.bzl", "go_grpc_compile")

go_grpc_compile(
    name = "go_sms_compile",
    deps = ["//proto/external/messaging:sms_proto"],
)

go_library(
    name = "go_default_library",
    embed = [":sms_go_proto"],
    importpath = "github.com/omnicate/loltel/proto/external/messaging",
    visibility = ["//visibility:public"],
)

proto_library(
    name = "sms_proto",
    srcs = ["sms.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "sms_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/omnicate/loltel/proto/external/messaging",
    proto = ":sms_proto",
    visibility = ["//visibility:public"],
)
