#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/users_srv"
exec "$CURDIR/bin/users_srv"