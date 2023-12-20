#! /usr/bin/env bash
CURDIR=$(cd "$(dirname "$0")"||exit; pwd)
echo "$CURDIR/bin/micro_user"
exec "$CURDIR/bin/micro_user"