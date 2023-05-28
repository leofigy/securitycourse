#!/usr/bin/env bash
# compile on windows only
go build -buildmode=c-shared -ldflags="-w -s -H=windowsgui" -o event.dll