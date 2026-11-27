# Regulate — development workflow (repo side)

How an approved feature goes from PRD to `main`. Design, priorities and PRD drafting happen in the claude.ai Project "Regulate" and its Claude Design canvases; building happens in the Claude Code project, on this repo. Human-in-the-loop: Claude proposes, the owner reviews, merges and deploys manually.

## Where things live

| What | Source of truth | Copy in this repo |
|---|---|---|
| UI design (frames) | Claude Design canvas "Dashboard v.0.1" | `docs/design/F-xx/` (exported HTML, per feature) |
| PRDs | claude.ai Project, `claude/prd/F-xx <name>.md` | `docs/prd/F-xx.md`, once approved |
| Decisions, codebook, teacher indicators | claude.ai Project docs | `docs/decisions.md`, `docs/codebook.md`, `docs/teacher-indicators.md` (snapshots; refreshed when a PRD is approved) |
| Code, migrations, tests | This repo | — |

If a snapshot here and an approved PRD disagree, the PRD wins; list the difference in the PR.

## Steps

| Step | Where | Output | Gate (owner) |
|---|---|---|---|
| 1. Design to MVP | Design chat | Frames with version prefix; decisions logged | "apply" |
| 2. Prioritise | claude.ai Project | Backlog row `F-xx`, priority, status | Approve order |
| 3. PRD | claude.ai Project | PRD with no `TBD` outside §10; frames exported | "PRD approved" |
| 4. Hand-off | This repo | `docs/prd/F-xx.md` + `docs/design/F-xx/*` committed (by the owner or as the first commit of the feature branch) | — |
| 5. Implement | Claude Code project thread, branch `feat/F-xx-<slug>` | Plan → code + migration + tests → draft PR with template | — |
| 6. Review | GitHub PR, comments answered in the same thread | Green checks, every AC ticked, deviations listed | Approve |
| 7. Merge + deploy | Owner, manually | `main` updated; `docker compose up -d --build` on the server | Smoke test |
| 8. Close the loop | claude.ai Project | Backlog = Deployed; Project docs updated from the PR's "Docs to update" | — |

## Rules

- **One PRD = one branch = one PR**, under ~400 changed lines. A large PRD lists tasks in §3a (`F-xx.n`); each task is its own branch and PR, merged in the listed order.
- **Design is the source of truth for UI; the PRD for behaviour.** On a conflict, stop and ask; don't guess.
- **`TBD` means stop.** An approved PRD has no `TBD` outside §10.
- **Docs follow code.** A new trace event, field or indicator goes under "Docs to update" in the PR.
- **Owner-only actions:** merge, deploy, force-push, CI changes, new dependencies, changes to migrations on `main`.

## Kick-off message (in the Claude Code project)

> Implement PRD F-xx (`docs/prd/F-xx.md`). Read CLAUDE.md first. Branch `feat/F-xx-<slug>`. Before coding, post your plan: FR → files, migrations, endpoints, trace events, tests, conflicts and TBDs. Then implement, add one test per AC, run the checks and open a draft PR with the template.

For a task of a split PRD: "Implement task F-xx.n of PRD F-xx (only the FRs and ACs listed for that task) …"

## Status values

Idea → Designed → PRD draft → PRD approved → In PR → Merged → Deployed.
