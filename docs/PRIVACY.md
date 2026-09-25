# Privacy model

UsedContract should export constraints, not production payloads. Agents should prefer:
- JSON paths and value types;
- normalized operation kinds;
- hashes or locally classified constants where practical;
- counts/confidence instead of raw samples.

Secrets, tokens, credentials, unrestricted payload bodies and PII should not be collected by default. A production collector needs configurable path redaction, allowlists, retention controls and tenant isolation.
