#!/usr/bin/env bash
set -euo pipefail

if git check-ignore --no-index cmd/usedcontract/main.go agents/go/usedcontract/agent.go agents/jvm/src/main/java/io/usedcontract/Agent.java agents/dotnet/UsedContract/Agent.cs; then
  echo 'Source files must not be excluded by .gitignore.' >&2
  exit 1
fi

go test ./...
go vet ./...
go build ./cmd/usedcontract

./usedcontract verify -contract examples/contract.json -candidate examples/candidate-compatible.json
status=0
./usedcontract verify -contract examples/contract.json -candidate examples/candidate-breaking.json || status=$?
test "$status" -eq 3

(cd agents/node && node test.mjs && node auto.test.mjs)
(cd agents/python && python3 -m unittest test_usedcontract.py test_auto_usedcontract.py)

rm -rf agents/jvm/out
mkdir -p agents/jvm/out
javac -d agents/jvm/out agents/jvm/src/main/java/io/usedcontract/Agent.java
