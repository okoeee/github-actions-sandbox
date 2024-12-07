#!/bin/sh

FILE="./.github/workflows/.release"

CURRENT_VERSION=$(cat $FILE)

NEXT_VERSION=$((CURRENT_VERSION + 1))

echo "$NEXT_VERSION" > $FILE

echo "$NEXT_VERSION"

