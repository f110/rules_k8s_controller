workspace(name = "dev_f110_rules_k8s_controller")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("//:deps.bzl", "rules_k8s_controller_dependencies")

rules_k8s_controller_dependencies()

http_archive(
    name = "bazel_gazelle",
    sha256 = "222e49f034ca7a1d1231422cdb67066b885819885c356673cb1f72f748a3c9d4",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.17.5")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("//k8s/kind:def.bzl", "kind_binary")

kind_binary(
    name = "kind",
    version = "0.10.0",
)

load("//k8s/kustomize:def.bzl", "kustomize_binary")

kustomize_binary(
    name = "kustomize",
    version = "v4.4.0",
)

load("@bazel_skylib//lib:unittest.bzl", "register_unittest_toolchains")

register_unittest_toolchains()
