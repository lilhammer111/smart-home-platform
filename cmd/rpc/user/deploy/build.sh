#!/usr/bin/env bash
RUN_NAME="micro_user"
mkdir -p output/bin output/conf/file
cp script/* output/
cp -r ../conf/file/* output/conf/file
cp ../.env output
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME} ../main.go