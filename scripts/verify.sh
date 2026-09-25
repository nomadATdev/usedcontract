#!/usr/bin/env bash
set -euo pipefail
go test ./...
go vet ./...
(cd agents/node && node test.mjs && node auto.test.mjs)
(cd agents/python && python3 -m unittest test_usedcontract.py test_auto_usedcontract.py)
go build ./cmd/usedcontract
