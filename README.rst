=======================
rules_k8s_controller
=======================

This repository is no longer maintain.

Tools that using code-generator are not suitable for use under the bazel sandbox.
Therefore, we give up using these tools.

rules_k8s_controller provides the rule for bazel that for code-generator and controller-tools.

Currently supported:

* Generate deepcopy
* Generate clientset
* Generate informer
* Generate lister
* Generate register
* Generate CustomResourceDefinition
* Generate the manifest of RBAC
* Support kustomize
* Create the cluster for development by kind

Setup
======

.. code::

    http_archive(
        name = "dev_f110_rules_k8s_controller",
        sha256 = "ec185cf615f01d69ce25ad0ca60c7dbb22b5f570af6f6ce199ca6c1812a86ae6",
        strip_prefix = "rules_k8s_controller-0.10.0",
        urls = [
            "https://github.com/f110/rules_k8s_controller/archive/v0.10.0.tar.gz",
        ],
    )

    load("@dev_f110_rules_k8s_controller//:deps.bzl", "rules_k8s_controller_dependencies")

    rules_k8s_controller_dependencies()

LICENSE
==========

MIT
