# Roadmap

## v0.3 (implemented)
- Canonical event IR + collector
- Contract inference + compatibility engine
- Contract-guided response mutation
- Python automatic tracked-JSON feasibility adapter
- Node automatic field-read feasibility adapter
- JVM/.NET/Node/Python/Go event SDKs
- CI and end-to-end fixtures

## v0.4 research gates
1. JVM: Jackson boundary + Byte Buddy/agent experiment for comparisons/branches.
2. .NET: System.Text.Json boundary + profiler/IL-rewrite experiment.
3. Node: AST transform for equality/null/branch expressions fed by tracked response paths.
4. Python: bytecode/AST transform beyond wrapper operators.
5. Go: source/build transform around decoded values.
6. Compare predicted breakages against replayed consumer tests to measure precision/recall.

Do not market whole-program semantic taint tracking until those gates are demonstrated.

## v0.4 — JVM automatic semantic dependency proof

### Objective

Prove, on one production-relevant runtime, that UsedContract can automatically infer a consumer dependency from real execution rather than from a manually authored contract.

### Implementation plan

- Java Agent using the standard Instrumentation API.
- Byte Buddy-based bytecode transformation where appropriate.
- Jackson deserialization-boundary integration.
- Provenance association: consumer service/version, HTTP method + route, JSON path, originating response value.
- Observation of common semantic operations: `String.equals`, `Objects.equals`, null/non-null checks where instrumentable, enum comparisons, and numeric comparisons where instrumentable.
- Emission into the existing language-neutral Semantic Event IR.
- Contract inference from those events.
- Candidate-response mutation generation and replay.
- Spring Boot zero-manual-contract demonstration.
- GitHub Actions end-to-end verification.

### Required acceptance test

Baseline provider response:

```json
{"status":"ACTIVE","unused_internal_id":123}
```

Consumer behavior:

```java
if (account.getStatus().equals("ACTIVE")) {
    enablePayments();
}
```

Candidate A:

```json
{"status":"active","unused_internal_id":123}
```

Expected: semantic break attributed to `$.status == "ACTIVE"`.

Candidate B:

```json
{"status":"ACTIVE"}
```

Expected: no observed break because `unused_internal_id` had no observed consumer dependency.

### Explicit non-goals for v0.4

- Claiming complete JVM taint analysis.
- Claiming all Java frameworks/serializers are supported.
- Claiming all branches can be causally attributed.
- Claiming automatic instrumentation for .NET, Python, Node, Go, Rust, or native code in the same milestone.
- Treating an unobserved field as proof that it can never matter; evidence is scoped to observed executions.

### Exit decision

If the JVM experiment cannot reliably distinguish the semantic-breaking case from the unused-field case without invasive application changes, pause multi-language expansion and reassess the core product thesis.

If it succeeds, v0.5 repeats the same IR contract on .NET (`System.Text.Json` + CLR instrumentation) before broader language expansion.
