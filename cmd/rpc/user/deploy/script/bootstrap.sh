#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/user_srv"
exec "$CURDIR/bin/user_srv"