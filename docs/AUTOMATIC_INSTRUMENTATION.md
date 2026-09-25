# Automatic instrumentation status

v0.3 separates **what is proven** from the long-term design.

## Implemented
- Language-neutral semantic event IR and collector.
- Contract inference and candidate-response compatibility checks.
- Contract-guided mutation generation.
- Python experimental tracked-JSON adapter: automatic field reads, scalar comparisons, and boolean/truthiness use after an application opts into the tracked decoder.
- Node experimental Proxy adapter: transparent field reads while preserving primitive values.
- Event SDKs for JVM, .NET, Node, Python and Go.

## Not yet implemented
- Whole-program data/control-flow taint tracking.
- Transparent interception of every serializer/deserializer.
- JVM bytecode transformation.
- CLR profiler/IL rewriting.
- Node build-time AST transformation needed for strict-equality semantics.
- Go compiler/source rewriting.
- Native/LLVM instrumentation.

Why Node differs: a JavaScript Proxy cannot intercept `===` performed on an ordinary primitive while simultaneously returning that primitive unchanged. Claiming otherwise would be incorrect. The production path is build-time AST/bytecode instrumentation for semantic operations.
