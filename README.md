# UsedContract

> **Catch API breaking changes your schema cannot see.**

UsedContract is an experimental, language-neutral runtime-derived API
compatibility engine. It asks a narrower question than OpenAPI diffing:

**Does a candidate API response violate behavior that a deployed consumer has
actually demonstrated it depends on?**

![License: BUSL-1.1](https://img.shields.io/badge/license-BUSL--1.1-blue)
![Status: Research MVP](https://img.shields.io/badge/status-research%20MVP-orange)
![Version: v0.3.1](https://img.shields.io/badge/version-v0.3.1-informational)

## Why this exists

A provider can make a schema-compatible change that still breaks real consumer
behavior.

Provider response:

```json
{"status":"ACTIVE"}
```

Consumer behavior:

```text
if status == "ACTIVE":
    enablePayments()
```

Candidate response:

```json
{"status":"active"}
```

A schema diff may see `string -> string`. UsedContract is designed to learn the
consumer's observed dependency on the value and eventually report the semantic
break.

The inverse matters too: removing an unused provider field should not be called
breaking merely because the schema changed.

## Current status

**v0.3.1 is a research MVP, not a production taint-analysis system.**

Implemented:

- language-neutral Semantic Event IR;
- event collector;
- runtime-derived requirement aggregation;
- candidate-response compatibility checks;
- contract-guided response mutation;
- evidence hashing;
- adapters/SDKs for JVM, .NET, Node, Python and Go;
- Python tracked-JSON feasibility adapter;
- Node field-read feasibility adapter;
- CI verification.

Not yet claimed:

- whole-program data/control-flow taint tracking;
- transparent instrumentation of every serializer;
- automatic semantic instrumentation for every supported language;
- patent novelty;
- production-safe handling of sensitive payloads.

The provider API may be written in any language. UsedContract observes the
**consumer**.

## Architecture

```text
JVM -------\
.NET -------\
Node --------+--> Semantic Event IR --> Collector --> Inference
Python -----/                                      |
Go --------/                                       v
                                         Compatibility Engine
                                                   |
                                                   v
                                             CI / Impact Report
```

Example normalized observations:

```text
$.status READ
$.status EQ "ACTIVE"
$.status CONTROL true
```

These can aggregate into a requirement that a specific consumer reads
`$.status`, compares it with an observed constant, and uses that result in
control flow.

## Quick start

Requirements for the core:

- Go
- Node.js for Node adapter tests
- Python 3 for Python adapter tests
- JDK for the JVM adapter compile check

Run:

```sh
bash scripts/verify.sh
```

Or manually:

```sh
go test ./...
go vet ./...
go build ./cmd/usedcontract
node agents/node/test.mjs
(cd agents/python && python3 -m unittest -v)
(cd agents/jvm && javac -d out src/main/java/io/usedcontract/Agent.java)
```

The .NET adapter is built in GitHub Actions because the packaging environment
used for this release does not include the .NET SDK.

## Run the collector

```sh
go run ./cmd/usedcontract serve -addr :8080
```

POST semantic events to `POST /v1/events` and inspect inferred requirements at
`GET /v1/contracts`.

## Verify a candidate response

```sh
go run ./cmd/usedcontract verify \
  -contract examples/contract.json \
  -candidate examples/candidate-breaking.json
```

Exit code `3` means compatibility findings were produced.

## Why not JavaScript-only?

Real APIs have consumers written in Java, .NET, Python, Go, Node, Kotlin, C++,
Rust and other runtimes. The durable product boundary is therefore the
language-neutral event model and compatibility engine, not any one
instrumentation technique.

## v0.4 engineering gate: automatic JVM instrumentation

The next milestone focuses on proving the hardest differentiated mechanism:

```text
Spring/Java consumer
      |
      v
GET /account
      |
      v
Jackson deserializes JSON
      |
      v
account.status
      |
      v
status.equals("ACTIVE")
      |
      v
branch executes
      |
      v
automatic UsedContract observations
      |
      v
candidate "ACTIVE" -> "active"
      |
      v
semantic break report
```

Acceptance criteria:

1. Java agent instrumentation without handwritten UsedContract contracts.
2. Jackson response values associated with endpoint + JSON-path provenance.
3. Common comparisons observed where technically feasible.
4. Events normalized into the existing Semantic Event IR.
5. Candidate mutations replayed through the compatibility engine.
6. `"ACTIVE" -> "active"` detected when observed behavior changes.
7. Removing an unused field is not reported as breaking merely because the
   provider schema changed.
8. End-to-end demonstration runs in CI.

See [`docs/V0.4_JVM_MILESTONE.md`](docs/V0.4_JVM_MILESTONE.md).

## Security and privacy

API payloads can contain credentials, personal data and commercially sensitive
information. UsedContract should export constraints and normalized observations,
not raw production payloads, by default.

Read [`SECURITY.md`](SECURITY.md) and [`docs/PRIVACY.md`](docs/PRIVACY.md)
before experimenting with production-derived data.

## License and commercial use

UsedContract is **source-available**, not OSI open source, under the
**Business Source License 1.1 (BUSL-1.1)**.

In summary:

- development, testing and other non-production use are allowed by BUSL-1.1;
- the Additional Use Grant permits specified personal, educational and nonprofit
  non-commercial production use;
- other production use, including commercial production deployment, requires a
  separate commercial license;
- on **2030-09-25**, this version is scheduled to change to
  **Apache License 2.0**, subject to BUSL-1.1.

The [`LICENSE`](LICENSE) file is authoritative.
[`COMMERCIAL_USE.md`](COMMERCIAL_USE.md) is a plain-English guide.

## Contributing

Bug reports, design discussions, reproducible cases and research feedback are
welcome.

Because the project is intended to support commercial licensing, code
contributions require a contributor-rights process. Read
[`CONTRIBUTING.md`](CONTRIBUTING.md) before opening a pull request.

## Project maturity

UsedContract is being published early to validate both the problem and the
instrumentation approach. Do not present this repository as proof that runtime
semantic dependency inference is solved, or that the concept is legally or
patentably unique.

See [`docs/ROADMAP.md`](docs/ROADMAP.md).
