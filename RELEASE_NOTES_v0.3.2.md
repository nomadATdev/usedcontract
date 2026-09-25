# UsedContract v0.3.2 — GitHub Publication Fixes

This release corrects repository-publication issues found during the first live
GitHub review. Product functionality is unchanged from v0.3.1.

## Fixed

- canonical Go module path is now `github.com/nomadATdev/usedcontract`;
- all internal Go imports use the canonical repository path;
- Node package version is aligned to `0.3.2`;
- repository/version documentation is aligned to `v0.3.2`;
- GitHub CI invokes `bash scripts/verify.sh` so local and CI verification share
  the same core checks;
- GitHub security issue-template link points to the canonical repository;
- stale v0.2 architecture wording was updated;
- BSL 1.1 attribution uses the current MariaDB plc 2024 wording.

## Important limitations

UsedContract remains a research MVP. It does not claim complete whole-program
taint tracking, transparent semantic instrumentation for every runtime,
production-safe payload collection, or patent novelty.

## License

BUSL-1.1. Commercial production use generally requires a separate commercial
license. See `LICENSE` and `COMMERCIAL_USE.md`.
