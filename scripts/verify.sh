#!/usr/bin/env bash
set -euo pipefail

go test ./...
go vet ./...
go build ./cmd/usedcontract

(cd agents/node && node test.mjs && node auto.test.mjs)
(cd agents/python && python3 -m unittest test_usedcontract.py test_auto_usedcontract.py)

rm -rf agents/jvm/out
mkdir -p agents/jvm/out
javac -d agents/jvm/out agents/jvm/src/main/java/io/usedcontract/Agent.java
