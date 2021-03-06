#!/usr/bin/env bash
set -e
if [ "@@DEBUG@@" = "true" ]; then
set -x
fi

BIN=@@BIN@@
GAZELLE_PATH=@@GAZELLE@@
ARGS=@@ARGS@@
NO_GAZELLE=@@NO_GAZELLE@@

TARGET_DIRS=@@TARGET_DIRS@@
FILENAME=@@FILENAME@@
GENERATED_DIRS=@@GENERATED_DIRS@@
SRC_PACKAGE_DIRS=@@SRC_PACKAGE_DIRS@@
SRC_DIRS=@@SRC_DIRS@@
GO_ROOT=@@GO_ROOT@@
RUNFILE_DIR=$(pwd)
MODULE_NAME=@@MODULE_NAME@@
HEADER_FILE=$RUNFILE_DIR/@@HEADER_FILE@@

GEN="$RUNFILE_DIR/$BIN"
GAZELLE="$RUNFILE_DIR/$GAZELLE_PATH"

rm -rf src
mkdir -p src

if [ -n "$(ls vendor 2> /dev/null)" ]; then
    mkdir -p ${RUNFILE_DIR}/src/${MODULE_NAME}/vendor
    for f in vendor/*; do
        ln -sf $RUNFILE_DIR/$f $RUNFILE_DIR/src/${MODULE_NAME}/vendor/
    done
fi

for i in "${!SRC_PACKAGE_DIRS[@]}"; do
    mkdir -p src/$(dirname ${SRC_PACKAGE_DIRS[$i]})
    ln -sf $RUNFILE_DIR/${SRC_DIRS[$i]} src/${SRC_PACKAGE_DIRS[$i]}
done

unset GO111MODULE
export GOROOT=$RUNFILE_DIR/$GO_ROOT
export GOPATH=${RUNFILE_DIR}

cd src/${MODULE_NAME}
echo "module ${MODULE_NAME}" > go.mod
echo "go 1.16" >> go.mod
"$GEN" "--output-base=$RUNFILE_DIR" "--go-header-file=$HEADER_FILE" "${ARGS[@]}"

cd "$BUILD_WORKSPACE_DIRECTORY"

if [ -n "$FILENAME" ]; then
    for i in "${!GENERATED_DIRS[@]}"; do
        if [ -f "$RUNFILE_DIR/${GENERATED_DIRS[$i]}/$FILENAME" ]; then
            mkdir -p "${TARGET_DIRS[$i]}"
            cp "$RUNFILE_DIR/${GENERATED_DIRS[$i]}/$FILENAME" "${TARGET_DIRS[$i]}"
        else
            if [ "@@DEBUG@@" = "true" ]; then
                echo "Skip to copy $RUNFILE_DIR/${GENERATED_DIRS[$i]/$FILENAME} because it is not exists"
            fi
        fi
    done
else
    mkdir -p $(dirname "${TARGET_DIRS[0]}")
    rsync -r $RUNFILE_DIR/${GENERATED_DIRS[0]}/ ${TARGET_DIRS[0]}/

    if [ "$NO_GAZELLE" = "false" ]; then
        "$GAZELLE" update "${TARGET_DIRS[0]}"
    fi
fi

