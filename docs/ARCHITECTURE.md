# Architecture

UsedContract is consumer-centric. The provider can be implemented in any language; only the consuming runtime needs an observation adapter.

```
consumer agents -> semantic event IR -> collector -> inference -> compatibility engine -> CI report
```

## Capability levels

1. Protocol observation: endpoint, method, response shape. Language independent.
2. Field consumption: which response fields the consumer reads. Runtime adapter required.
3. Semantic dependency: comparisons, null assumptions, arithmetic, propagation and control influence. Runtime/compiler instrumentation required.

v0.2 implements the IR, collector, basic inference and compatibility engine. The supplied agents are transport SDKs for emitting semantic events. They are NOT yet transparent bytecode/AST/CLR taint agents. That work is intentionally separate so the core is not coupled to one language.

## OpenTelemetry

The production design should integrate with OTLP/OpenTelemetry for identity, trace correlation and transport. UsedContract events remain a separate semantic schema because normal tracing does not encode field-level data-flow semantics.
