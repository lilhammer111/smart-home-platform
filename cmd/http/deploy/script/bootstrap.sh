#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=alert
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}