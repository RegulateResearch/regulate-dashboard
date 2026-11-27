# Regulate dashboard

Self-regulated learning (SRL) learning-analytics dashboard for a PhD project. Roles: mahasiswa (student), dosen (lecturer), admin. Pilot: one course (IMK) linked to SCELE (the university LMS) by manual links set by the lecturer. Describe SCELE as linked, never as integrated or synced.

## Stack and checks

Active development happens on `dev`; `main` is behind and is updated from `dev`. Feature branches start from `dev` unless the owner says otherwise.

| Part | Folder | Checks (run before every PR) |
|---|---|---|
| Frontend | `sveltekit/` — SvelteKit 2, Svelte 5, TypeScript, Tailwind 4, Storybook | `cd sveltekit && npm ci && npm run check && npx vitest --run` (e2e: `npx playwright test --trace off`) |
| API | `go/` — Go 1.25 (module `frascati`), Gin, pgx | `cd go && go vet ./... && go test ./...` |
| Database schema | `sql/` — one idempotent SQL file per table, merged into the generated `sql/initdb.sql` by `sqlmerger/` | `make initdb-sync`, then apply `sql/initdb.sql` twice to an empty test database with `psql -v ON_ERROR_STOP=1` (see below) |
| Runner | `docker-compose.yaml` + `Makefile` (Postgres 16, pgAdmin, API, frontend, nginx) | Local and production only; not available in cloud sessions |

Other folders: `nginx/` (reverse proxy config), `protobuf/` and `grpc-client-coba/` (gRPC experiment, separate Go modules), `sqlmerger/` (the schema merge tool, separate Go module).

**Frontend tests.** `vite.config.ts` defines three Vitest projects: `server` (`src/**/*.{test,spec}.ts`, Node), `client` (`src/**/*.svelte.{test,spec}.ts`, Chromium via Playwright) and `storybook` (story tests, Chromium). `npm run test` runs Vitest plus the Playwright e2e suite in `e2e/` (it builds and previews the app on port 4173). `npm run lint` (Prettier + ESLint) exists but the codebase is not Prettier-clean yet; don't reformat files you aren't changing.

**Database schema (no up/down migrations).** The schema is not a numbered migration set. Each table has a file in `sql/` with `CREATE TABLE IF NOT EXISTS` and later `ALTER TABLE … ADD COLUMN IF NOT EXISTS` statements. `sqlmerger/main.go` lists the files in order and writes `sql/initdb.sql` (`make initdb-sync`); Docker runs `initdb.sql` on a fresh volume, and `make apply-migrate` re-applies it to a running container. So:
- A schema change = append idempotent statements to the table's file (or add a new file and list it in `sqlmerger/main.go`), then regenerate `initdb.sql` and commit both.
- Never edit a statement that is already on `main`; append a new one.
- Check: apply `initdb.sql` to an empty database, then apply it again; both runs must succeed.

```bash
psql "$DATABASE_URL" -v ON_ERROR_STOP=1 -f sql/initdb.sql   # first run: creates everything
psql "$DATABASE_URL" -v ON_ERROR_STOP=1 -f sql/initdb.sql   # second run: must be a no-op
```

- Tests that need Postgres use `DATABASE_URL` from the environment (the API reads `DATABASE_URL`, or falls back to `DATABASE_USERNAME/PASSWORD/HOST/PORT/NAME`). If it is not set or Postgres is not reachable, say so; don't skip the tests silently.
- Secrets live in `.env` files (ignored): root `.env` for Docker Compose, `go/.env` and `sveltekit/.env` for local runs. `make init` copies them from `.env.empty`. `.env.example` holds placeholder values; update `.env.example` and `.env.empty` in the same folder when you add a variable.

## Docs

| Path | What | Use |
|---|---|---|
| `docs/prd/F-xx.md` | Approved PRD per feature (template: `docs/prd/_template.md`) | The spec. Section numbers and IDs (US, FR, R, AC, Q) are stable. |
| `docs/design/F-xx/` | Exported design frames (HTML) for that PRD | Visual reference: match labels, order, states and numbers. |
| `docs/decisions.md` | Product decisions, wording, colours, sample data | Conventions behind the PRDs. |
| `docs/codebook.md` | Trace events and indicator variables | Names and payloads for trace events. |
| `docs/teacher-indicators.md` | Lecturer views and indicators | Context for teacher pages. |
| `docs/workflow.md` | How a feature goes from PRD to merge | Process rules. |

The design canvases themselves are not reachable from here. If a PRD and `docs/design/` disagree, stop and ask.

## Conventions

**Language and formats**
- UI text in Bahasa Indonesia, formal "Anda" (never "kamu"). Sentence case, no exclamation marks; buttons are verb + object ("Simpan target").
- Times shown in WIB (Asia/Jakarta), stored in UTC. Format: "Kam 19 Nov, 06.00 WIB".
- Decimals with a comma (72,7); thousands with a dot (2.488).

**Wording**
- Learning items are "aktivitas" (types Tugas, Kuis, Ujian, Materi). "Tugas" only as a type, component name or activity name.
- Graded items are "aktivitas dinilai", grouped under "komponen nilai"; weights as share of the final grade ("10% dari nilai akhir"). "penilaian" only for self-assessment.
- The field for related activities is labelled "Aktivitas".

**Privacy and research**
- Indicator codes (G1, P7, R2 …) are researcher-only: `data-indicator` attributes, never visible text for any role.
- Every learner action listed in `docs/codebook.md` emits its trace event with the payload the PRD gives in §5c.
- Research exports are pseudonymised (`p_…`). No real student or trace data in code, tests, fixtures or logs.

**UI**
- Use the Regulate design tokens (CSS variables), never raw hex unless it is a token value. Neutrals are stone (`#292524`, `#6F6863`, `#E7E5E4`, `#F5F5F4`).
- Brand yellow `#FDE047` is a fill only, with ink `#292524` text: one primary action per view, selected state. Selected menu item: `#FEF9C3` fill, no border.
- States: success `#15803D`, danger `#DC2626`, warning `#B45309`, always with a word or icon. Below/above target: brown `#A16207` / blue `#2A78D6`, never red/green.
- Charts: categorical blue `#2A78D6`, teal `#0E9384`, violet `#5B21B6`, pink `#E2577F`, brown `#8A5A2B`, grey `#57534E`; levels use the yellow ramp `#FEF9C3 → #FDE047 → #EAB308 → #A16207 → #713F12`.
- One `AppShell`, sidebar, top bar, course header and `DataTable` component, shared by all roles. Pages never copy their markup.
- How-to-read text goes behind an InfoTip (i); live status, privacy notes and "Tindakan:" lines stay visible.
- Accessibility: named controls, `aria-current` for current items, `aria-sort` on sorted headers, 2px ink focus ring, contrast ≥ 4.5:1.

**Seed data** (must match `docs/decisions.md` §7)
- Now = Kam 19 Nov 2026, 10.00 WIB. Course IMK (CSCE604123, Gasal 2026/2027, 24 students).
- S01 Ayu (u_0001 / p_7f3a, ayu@ui.ac.id): Kuis 1 82 · Tugas 1 74 · UTS 68 · Kuis 2 71 → average 72,7; target 80; needed 88,9.
- Dosen A (dosen.a@ui.ac.id); Admin Regulate (admin@ui.ac.id).

## Working rules

1. One PRD (or one task F-xx.n of a split PRD) per branch and PR: `feat/F-xx-<slug>`, under ~400 changed lines.
2. First commit of a feature: `docs/prd/F-xx.md` and `docs/design/F-xx/*` if they are not in the repo yet.
3. Plan first: FR → files, migrations, endpoints, trace events, tests, conflicts, TBDs.
4. A `TBD` means ask. Never invent a value, label, colour, endpoint or wording.
5. One test per acceptance criterion, named after its ID (`TestAC3_…` in Go, `AC-3: …` in Vitest/Playwright).
6. Run all checks; open a draft PR with the template filled in; list deviations and "Docs to update".
7. Don't merge, deploy, force-push or change CI.
