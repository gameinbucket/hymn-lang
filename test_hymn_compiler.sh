#!/bin/bash -e
cd "$(dirname "$0")"

HYMN_PACKAGES="{\"hymn_compiler\":\"$(pwd)/hymn_compiler\"}"
export HYMN_PACKAGES

./hymn.sh test_hymn_compiler/test_tokenizer.hm -t -w out/test_hymn "$@"
