load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "separatetests_lib",
    srcs = ["main.go"],
    importpath = "github.com/drone-ah/wordsonsand/post/2023/11/separatetests",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "separatetests",
    embed = [":separatetests_lib"],
    visibility = ["//visibility:public"],
)

# gazelle:build_tags unit_test,integration_test
go_test(
    name = "separatetests_test",
    srcs = [
        "main_int_test.go",
        "main_test.go",
    ],
)
