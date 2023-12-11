#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=auth
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}