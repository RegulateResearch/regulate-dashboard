---
id: F-xx
title: <Feature name>
status: draft            # draft | approved | in-pr | merged | deployed
priority: P1             # P0 | P1 | P2
version_target: v.0.x
roles: [mahasiswa]       # mahasiswa | dosen | admin
srl_phase: forethought   # forethought | performance | reflection
frames:                  # Dashboard v.0.1 file names
  - KursusNilaiPersiapanID.dc.html
design_export: docs/design/F-xx/
depends_on: []           # other F-xx IDs
branch: feat/F-xx-<slug>
updated: YYYY-MM-DD
---

# F-xx <Feature name>

## 1. Goal and SRL rationale
2–3 lines: which problem in which SRL phase this solves, and for whom.

## 2. User stories
- US-1 As a mahasiswa, I want … so that …
- US-2 As a dosen, I want … so that …

## 3. Functional requirements
| ID | Requirement (one testable statement) | Story |
|---|---|---|
| FR-1 | The system shows … when … | US-1 |
| FR-2 | … | US-1 |

## 4. Design reference
Claude Code cannot open the Design canvas. Describe everything it must match here, or put it in the exported HTML.

| Frame (file name) | State | What must match (labels, order, numbers, components) |
|---|---|---|
| KursusNilaiPersiapanID.dc.html | default | Title "…", cards in order A → B → C, InfoTip on "…", sample value 72,7 |
| KursusNilaiPersiapanID.dc.html | empty | Text "…" |
| — | loading / error | Not designed: use the shared skeleton / error banner |

## 5. Data contract

### 5a. Database changes
Migration: `NNN_<name>.up.sql` / `.down.sql`

| Table | Column | Type | Null | Default | Note |
|---|---|---|---|---|---|
| activity | is_graded | boolean | no | false | FR-2 |

### 5b. API endpoints
| Method | Path | Role | Request | Response (200) | Errors |
|---|---|---|---|---|---|
| GET | /api/courses/{id}/grades | mahasiswa | — | `{ components: [{ id, name, weight, average }] }` | 401, 403, 404 |

### 5c. Trace events (must match the codebook)
| Codebook name | Trigger (user action) | Payload fields | FR |
|---|---|---|---|
| goalSet | Student saves a target | activityId, target, confidence, ts (WIB) | FR-1 |

## 6. Rules and edge cases
| ID | Case | Expected behaviour |
|---|---|---|
| R-1 | No activities in a component | Show automatic Ujian activity (decisions doc §4) |
| R-2 | User without permission | 403, no data leaked |

Always check: empty state, permissions per role, WIB time zone, comma decimals, rounding, long text, mobile 390 px.

## 7. Acceptance criteria
| ID | Given / When / Then | Covers |
|---|---|---|
| AC-1 | Given Ayu (S01) has target 80, when she opens Nilai, then "Perlu 88,9" is shown | FR-1 |
| AC-2 | … | FR-2, R-1 |

Every FR and R must appear in at least one AC's "Covers" column.

## 8. Definition of done
- [ ] Every AC has a passing test named after its ID
- [ ] Migration up + down run cleanly
- [ ] Trace events match §5c and the codebook
- [ ] UI matches §4 and `docs/design/F-xx/`
- [ ] Project docs updated if traces, fields or indicators changed

## 9. Out of scope
- …

## 10. Open questions
| ID | Question | Owner | Answer |
|---|---|---|---|
| Q-1 | TBD: … | user | — |

An approved PRD has no `TBD` outside this section.
