load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "serve.go",
    ],
    importpath = "github.wdf.sap.corp/team-pyrite/hackathon/eventgen",
    visibility = ["//visibility:private"],
    deps = [
        "//proto/commerce:commerce_go_proto_library",
        "@com_github_sirupsen_logrus//:go_default_library",
        "@com_github_urfave_cli//:go_default_library",
    ],
)

go_binary(
    name = "eventgen",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
