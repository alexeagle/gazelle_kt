load("@gazelle//:def.bzl", "gazelle")
load("@package_metadata//licenses/rules:license.bzl", "license")
load("@package_metadata//rules:package_metadata.bzl", "package_metadata")

package_metadata(
    name = "package_metadata",
    attributes = [
        ":license",
    ],
    purl = "pkg:bazel/{}@{}".format(
        module_name(),
        module_version(),
    ) if module_version() else "pkg:bazel/{}".format(module_name()),
    visibility = ["//visibility:public"],
)

license(
    name = "license",
    kind = "@package_metadata//licenses/spdx:Apache-2.0",
    text = "LICENSE",
)

# gazelle:prefix github.com/aspect-build/gazelle_kt
gazelle(
    name = "gazelle",
    gazelle = "@multitool//tools/gazelle",
)
