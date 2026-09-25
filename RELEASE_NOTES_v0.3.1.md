# UsedContract v0.3.1 — First Public Research Release

UsedContract explores runtime-derived API compatibility contracts: learning what
deployed consumers actually depend on instead of treating every provider schema
difference as equally breaking.

## Highlights

- language-neutral Semantic Event IR;
- collector and compatibility engine;
- contract-guided candidate mutations;
- JVM, .NET, Node, Python and Go adapter foundations;
- Python automatic semantic-observation feasibility work;
- Node automatic field-read feasibility work;
- documented v0.4 JVM automatic-instrumentation acceptance gate.

## Important limitations

This is a research MVP. It does not claim complete whole-program taint tracking,
transparent semantic instrumentation for every runtime, production-safe payload
collection or patent novelty.

## License

BUSL-1.1. Commercial production use generally requires a separate commercial
license. See `LICENSE` and `COMMERCIAL_USE.md`.
