load("@gazelle//:def.bzl", "gazelle_binary")
load("@rules_go//go:def.bzl", "go_library", "go_test")
load("//tools:gazelle.bzl", "gazelle_generation_test")

# Exclude all test data
# gazelle:exclude tests/

go_library(
    name = "kotlin",
    srcs = [
        "configure.go",
        "generate.go",
        "imports.go",
        "kotlin.go",
        "language.go",
        "resolver.go",
    ],
    importpath = "github.com/aspect-build/gazelle_kt/kotlin",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_aspect_build_aspect_cli//gazelle/common",
        "@com_github_aspect_build_aspect_cli//gazelle/kotlin/kotlinconfig",
        "@com_github_aspect_build_aspect_cli//gazelle/kotlin/parser",
        "@com_github_aspect_build_aspect_cli//pkg/logger",
        "@com_github_bazel_contrib_rules_jvm//java/gazelle/javaconfig",
        "@com_github_bazel_contrib_rules_jvm//java/gazelle/private/java",
        "@com_github_bazel_contrib_rules_jvm//java/gazelle/private/maven",
        "@com_github_bazel_contrib_rules_jvm//java/gazelle/private/types",
        "@com_github_emirpasic_gods//maps/treemap",
        "@com_github_emirpasic_gods//sets/treeset",
        "@com_github_emirpasic_gods//utils",
        "@com_github_rs_zerolog//:zerolog",
        "@gazelle//config",
        "@gazelle//label",
        "@gazelle//language",
        "@gazelle//repo",
        "@gazelle//resolve",
        "@gazelle//rule",
    ],
)

# Internal only for tests
gazelle_binary(
    name = "gazelle_kotlin_binary",
    languages = [":kotlin"],
    visibility = ["//visibility:private"],
)

# A separate generation test for each tests/* test case
[
    gazelle_generation_test(
        name = "%s_test" % test_workspace.replace("/WORKSPACE", "").replace("tests/", ""),
        dir = test_workspace.replace("/WORKSPACE", ""),
        gazelle_binary = "gazelle_kotlin_binary",
    )
    for test_workspace in glob(["tests/**/WORKSPACE"])
]

go_test(
    name = "kotlin_test",
    srcs = ["kotlin_test.go"],
    embed = [":kotlin"],
)
