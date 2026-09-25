# Security policy

## Project maturity

UsedContract v0.3.1 is a research MVP. It is not yet intended to receive raw,
unfiltered production API payloads.

## Reporting a vulnerability

Please **do not open a public GitHub issue** for a suspected security
vulnerability.

Preferred reporting path:

1. Use GitHub Private Vulnerability Reporting for the repository, if enabled.
2. If private reporting is not enabled, contact the owner of the canonical
   UsedContract repository privately.

Include the affected version/commit, reproduction steps, expected/actual
behavior, impact assessment and the minimum proof-of-concept needed to reproduce
the issue.

Do not include real credentials, production tokens, customer data or unrelated
personal data.

## Sensitive-data model

The intended production design is:

- metadata/constraint-first collection;
- local normalization and redaction;
- no raw response-body collection by default;
- configurable path allowlists/denylists;
- bounded retention;
- tenant isolation;
- secret/token suppression;
- explicit opt-in for operand capture that could contain sensitive data.

The current research adapters may emit explicitly supplied operands. Do not use
that functionality with sensitive production values without appropriate
filtering, governance and authorization.

## Supported versions

Only the latest tagged release is expected to receive security fixes while the
project remains pre-1.0.

| Version | Supported |
| --- | --- |
| latest v0.x | Yes |
| older v0.x | Best effort |
