=======================
rules_k8s_controller
=======================

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

LICENSE
==========

MIT