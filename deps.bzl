load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def rules_k8s_controller_dependencies():
    _maybe(
        http_archive,
        name = "io_bazel_rules_go",
        sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
            "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        ],
    )

def _maybe(repo_rule, name, **kwargs):
    if name not in native.existing_rules():
        repo_rule(name = name, **kwargs)
