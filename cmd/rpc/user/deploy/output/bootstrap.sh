#! /usr/bin/env bash
CURDIR=$(cd "$(dirname "$0")"||exit; pwd)
echo "$CURDIR/bin/user_micro"
exec "$CURDIR/bin/user_micro"