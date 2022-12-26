#!/usr/bin/env bash

OUT=(
  "windows/amd64"
  "darwin/amd64"
  "darwin/arm64"
)

TMP="$PWD"
ROOT="$(cd "$(dirname "$0")/.." | pwd)"
DIR="dist"
NAME="gotmpl"

mkdir "$DIR" >/dev/null 2>&1
for _out in "${OUT[@]}"; do
  os="${_out%%/*}"
  arch="${_out##*/}"

  echo "Building... $_out"
  GOOS="$os" GOARCH="$arch" go build \
    -o "$DIR/$NAME-$os-$arch" \
    "github.com/kamontat/gotmpl/cli"
done

cd "$TMP" | exit 1
unset NAME DIST OUT TMP ROOT
