load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "autolambdatest_lib",
    srcs = ["main.go"],
    importpath = "github.com/drone-ah/wordsonsand/post/2023/11/autolambdatest",
    visibility = ["//visibility:public"],
    deps = ["@com_github_aws_aws_sdk_go_v2_service_ssm//:ssm"],
)

# gazelle:build_tags unit_test,integration_test
go_test(
    name = "autolambdatest_test",
    size = "small",
    srcs = ["main_int_test.go"],
    embed = [":autolambdatest_lib"],
    deps = [
        "@com_github_aws_aws_sdk_go_v2//aws",
        "@com_github_aws_aws_sdk_go_v2_config//:config",
        "@com_github_aws_aws_sdk_go_v2_credentials//stscreds",
        "@com_github_aws_aws_sdk_go_v2_service_ssm//:ssm",
        "@com_github_aws_aws_sdk_go_v2_service_sts//:sts",
        "@com_github_pulumi_pulumi_aws_sdk_v6//go/aws",
        "@com_github_pulumi_pulumi_aws_sdk_v6//go/aws/iam",
        "@com_github_pulumi_pulumi_aws_sdk_v6//go/aws/ssm",
        "@com_github_pulumi_pulumi_sdk_v3//go/auto",
        "@com_github_pulumi_pulumi_sdk_v3//go/auto/optup",
        "@com_github_pulumi_pulumi_sdk_v3//go/pulumi",
        "@com_github_stretchr_testify//assert",
        "@com_github_stretchr_testify//suite",
    ],
)
