#!/bin/bash -eu

HYMN_STD="$(pwd)/hymn_std"
HYMN_BOOKS="$(pwd)/books"

NEW_PACKAGES="\"hymn\":\"${HYMN_STD}\",\"books\":\"${HYMN_BOOKS}\""

if [ ${HYMN_PACKAGES:-} ]; then
  HYMN_PACKAGES="${HYMN_PACKAGES#"{"}"
  HYMN_PACKAGES="${HYMN_PACKAGES%"}"}"
  HYMN_PACKAGES="$HYMN_PACKAGES,$NEW_PACKAGES"
else
  HYMN_PACKAGES="$NEW_PACKAGES"
fi

HYMN_PACKAGES='{'"$HYMN_PACKAGES"'}'

export HYMN_PACKAGES
