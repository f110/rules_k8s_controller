#!/usr/bin/env bash
set -e

NAME="code-generator"
VERSION="v0.23.0"

TOPDIR=$BUILD_WORKSPACE_DIRECTORY
TARGETDIR="${TOPDIR}/third_party/${NAME}-${VERSION}"

if [ -d "${TARGETDIR}" ]; then
  rm -rf "${TARGETDIR}"
fi

TMPDIR=$(mktemp -d)
echo "Create temporary directory ${TMPDIR}"
mkdir -p "${TMPDIR}"/third_party/${NAME}-${VERSION}
cp -r ${TOPDIR}/internal/go ${TMPDIR}
cat <<EOS > ${TMPDIR}/WORKSPACE
workspace(name = "dev_f110_rules_k8s_controller")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

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

go_register_toolchains(version = "1.17.2")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()
EOS
cat <<EOS > ${TMPDIR}/BUILD.bazel
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(name = "gazelle")
EOS

echo "Clone: https://github.com/kubernetes/${NAME}.git into $(pwd)"
cd "${TMPDIR}/third_party"
git clone --quiet --depth 1 --branch "$VERSION" "https://github.com/kubernetes/${NAME}.git" "${NAME}-${VERSION}"

echo "Remove unnecessary files"
cd "${NAME}-${VERSION}"
find . -name "*_test.go" -delete
find . -name "testdata" -type d | xargs rm -rf
find . -name "_examples" -type d | xargs rm -rf
find . -name ".*" -maxdepth 1 | grep -v "^.$" | xargs rm -rf {} +
rm -rf examples

cat <<EOS > BUILD.bazel
load("//go:vendor.bzl", "go_vendor")

# gazelle:prefix k8s.io/${NAME}

go_vendor(name = "vendor")
EOS

bazel run //third_party/${NAME}-${VERSION}:vendor
bazel clean --expunge
mkdir -p "$(dirname "${TARGETDIR}")"
mv "${TMPDIR}/third_party/${NAME}-${VERSION}" "${TARGETDIR}"