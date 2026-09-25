# Contributing to UsedContract

Issues, design critiques, reproducible API-breakage examples, documentation
feedback and benchmark ideas are welcome.

## Code contributions

UsedContract is source-available and is intended to support commercial
licensing. Contributor rights therefore need to be unambiguous.

**Until a formal contributor-license workflow is enabled, unsolicited source
code pull requests may be closed without merge.**

Before submitting source code, open an issue describing the proposed change.
The maintainer can confirm whether it can be accepted and under what contributor
terms.

Do not submit code copied from projects with incompatible licenses.

## Useful contributions now

- minimal schema-compatible semantic breakage examples;
- false-positive/false-negative test cases;
- runtime-instrumentation research references;
- reproducible bugs;
- documentation corrections;
- privacy/security threat-model feedback.

## Development check

```sh
bash scripts/verify.sh
```

## Design principles

1. Provider implementation language is irrelevant.
2. Consumer observations normalize into one Semantic Event IR.
3. Unobserved does not mean impossible.
4. A schema difference is not automatically a behavioral break.
5. Sensitive values should not be exported by default.
6. Unsupported instrumentation capabilities must not be described as complete.
