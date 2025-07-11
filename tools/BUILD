load("@bazel_env.bzl", "bazel_env")
load("@bazel_skylib//:bzl_library.bzl", "bzl_library")

bazel_env(
    name = "bazel_env",
    tools = {
        "go": "@rules_go//go",
    },
)

bzl_library(
    name = "gazelle",
    srcs = ["gazelle.bzl"],
    visibility = ["//visibility:public"],
    deps = ["@gazelle//:def"],
)
