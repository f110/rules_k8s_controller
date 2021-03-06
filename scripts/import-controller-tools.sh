#!/usr/bin/env bash
set -e

NAME="controller-tools"
VERSION="v0.5.0"

TOPDIR=$BUILD_WORKSPACE_DIRECTORY
TARGETDIR="${TOPDIR}/third_party/${NAME}-${VERSION}"

if [ -d "${TARGETDIR}" ]; then
  rm -rf "${TARGETDIR}"
fi

TMPDIR=$(mktemp -d)
echo "Create temporary directory ${TMPDIR}"
mkdir -p "${TMPDIR}"/third_party/${NAME}-${VERSION}
cp -r ${TOPDIR}/WORKSPACE ${TOPDIR}/internal/go ${TMPDIR}
cat <<EOS > ${TMPDIR}/BUILD.bazel
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(name = "gazelle")
EOS

echo "Clone: https://github.com/kubernetes-sigs/${NAME}.git into $(pwd)"
cd "${TMPDIR}/third_party"
git clone --depth 1 --branch "$VERSION" https://github.com/kubernetes-sigs/${NAME}.git "${NAME}-${VERSION}"

echo "Remove unnecessary files"
cd "${NAME}-${VERSION}"
find . -name "*_test.go" -delete
find . -name "testdata" -type d | xargs rm -rf
find . -name ".*" -maxdepth 1 | grep -v "^.$" | xargs rm -rf {} +

cat <<EOS > BUILD.bazel
load("//go:vendor.bzl", "go_vendor")

# gazelle:prefix sigs.k8s.io/${NAME}

go_vendor(name = "vendor")
EOS

bazel run //third_party/${NAME}-${VERSION}:vendor
bazel clean --expunge
mkdir -p "$(dirname "${TARGETDIR}")"
mv "${TMPDIR}/third_party/${NAME}-${VERSION}" "${TARGETDIR}"