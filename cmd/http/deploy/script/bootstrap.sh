#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=alerts
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}