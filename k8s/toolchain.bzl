load("//k8s/kind:def.bzl", _kind_binary = "kind_binary")
load("//k8s/kustomize:def.bzl", _kustomize_binary = "kustomize_binary")

def register_rules_k8s_controller_toolchain(**kwargs):
    if "kind" in kwargs:
        _kind_binary(name = "kind", version = kwargs["kind"])

    if "kustomize" in kwargs:
        _kustomize_binary(name = "kustomize", version = kwargs["kustomize"])
