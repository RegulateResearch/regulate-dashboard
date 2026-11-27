<!-- Fill in every section. Write "None" rather than deleting a section. -->

## Summary

PRD: `docs/prd/F-xx.md` (task F-xx.n, if split)
<!-- One or two sentences: what this PR does. -->

## AC coverage

| AC | Test name | Status |
|---|---|---|
| AC-1 | `TestAC1_…` / `AC-1: …` | ✅ / ❌ |

## FR → files

| FR | Files |
|---|---|
| FR-1 | `…` |

## Schema, endpoints and trace events

- Schema changes (`sql/*.sql`, regenerated `sql/initdb.sql`):
- Endpoints added or changed:
- Trace events emitted:

## Checks

<!-- Paste the summary line of each check from CLAUDE.md. Say which ones could not run, and why. -->

```
frontend  npm run check:
frontend  npx vitest --run:
api       go vet ./...:
api       go test ./...:
schema    initdb.sql applied twice:
```

## Deviations

<!-- Anything that differs from the PRD or docs/design/, and why. -->

## Out of scope (noticed, not done)

## Open questions

## Docs to update

<!-- Changes to trace events, fields or indicators that the Project docs (decisions, codebook, teacher indicators) need. -->
