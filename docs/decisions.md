> Snapshot of the claude.ai Project doc "Regulate — decisions and conventions" (27 Nov 2026). The Project doc is the source of truth; this copy is refreshed when a PRD is approved. Canvas-specific conventions (§1, §2, §5c) describe design work and do not apply to code.

# Regulate — decisions and conventions

Running log of decisions made while designing the Regulate dashboards (Claude Design canvases). Newest decisions are at the top of each section. Last updated 27 Nov 2026 (desktop scroll guideline in the layout frame; no archive backups by default; v.0.1.7 complete: landing hero reverted and SCELE wording toned down on all four landing frames; review copies deleted).

Related docs: `claude/Regulate design files.md` (where things are), `claude/Regulate SRL feature grid and trace codebook.md` (variables), `claude/Regulate teacher-facing dashboard indicators and visualizations.md` (teacher indicators), `claude/Regulate development workflow.md` (design → PRD → Claude Code).

---

## 1. Working conventions

| Convention | Detail |
|---|---|
| Review before apply | New versions go into review frames titled with the next version prefix (e.g. "[v.0.1.5] …"), below or next to the frames they replace. Existing frames are edited only on "apply" or "edit directly". |
| Apply | Review frames replace the originals in place (no backup unless asked; see Backups). Review frames are deleted only on request. Apply can be per frame (e.g. only one of four landing frames). |
| Backups (since 27 Nov 2026) | No archive backups by default, also not when applying review-first changes, to save usage and keep small changes light. Back up to Dashboard v.0.0 [archived] only when the user explicitly asks. |
| Design review skill (since 27 Nov 2026) | When the user asks for a design review, use the `regulate-design-review` skill (user skill). It reviews or builds frames against these rules: static check script + checklist, findings as a full table (Blocker / Major / Minor), insights first, fixes as review frames. Findings of a review run are tracked in a Sheet (first run: "Regulate design review — test findings", https://claude.ai/artifact/UWAgspKToh9xe4TioeM2yH; status Open / Apply / In review / Won't fix / Done). When a decision here changes, update the skill too. |
| Design comments (since 27 Nov 2026) | Manual review is done with comments on the design file. Plain comments (not sent to Claude) are collected; Claude applies them only when the user says so in chat (e.g. "apply the open comments"), then lists what it changed per comment. A comment sent to Claude without "apply" in it is answered in its thread but not applied. |
| Insights first | "Insights first" / "do not edit yet" = analysis only. |
| Ask manually | "Ask me manually" = clarifying questions as multiple-choice questions before building. |
| Canvas safety | Read the live canvas.json before publishing a layout change; keep the user's manual moves. Avoid publishing a frame the user is editing on the canvas at the same time. |
| Docs follow UI | When an indicator, field or the data model changes, update the codebook, teacher doc, design files doc and Conceptual Model too (unless "UI only"). |
| Canvas layout | Dashboard v.0.1 frames are arranged by user flow: Landing → Student → Teacher → Admin; student rows = pages, columns = phases (with gaps); teacher rows = course tabs with drill-down pages to the right. New frames go into their flow row. |
| Shared parts | One shared sidebar (`SharedSidebar`, `role` = mahasiswa / dosen / admin + `active`) and one shared course header (`SharedCourseHeader`, `role` = mahasiswa / dosen + `tab`), placed into pages with `<dc-import>` (merged 26 Nov 2026; first split per role in v.0.1.6). Styling exists once; menu labels, links, account details, meta lines and tab lists are per-role lists in the shared frame. Change them there, never per page; new pages use them. Page descriptions under the course tabs stay in each page. |
| Shared top bar, stylesheet and layout reference | `SharedTopbar` (path = "Label=Page.dc.html / … / Current page") in every page. `regulate-app.css` holds the layout shell (`.rg-shell`, `.rg-main`, `.rg-content`, `.rg-content-compact`), hover states and responsive rules; pages keep only unique rules in `<helmet>`. "Layout — Halaman Regulate" (desktop 1440 + mobile 390) documents the one page structure for all roles and uses the same classes. (26 Nov 2026, applied directly) |
| Shared table (27 Nov 2026) | `SharedTable` (shadcn "tasks" table pattern): a page passes `columns` (key, label, type = text · strong · muted · mono · num · badge · person · stacked · check · actions; optional align, sub, width, `sortable: false`, `srOnly`) and `rows`; variants `sortable` (header click: naik → turun → tanpa urutan) and `paginate` (Baris per halaman 5/10/20/50, "Halaman x dari y", first/prev/next/last), `page-size`, `size` = default · compact, `default-sort`. Search and filters stay above the table in each page. Row hover and sort-button styles in `regulate-app.css`. Applied to all admin tables (Mata kuliah; Pengguna; Peran & izin without sort/paging; Ekspor preview compact; Ekspor terbaru sorted newest first) and to Nilai kelas → Nilai semua mahasiswa (P7: grouped component headers via column `group`, `divider` lines, cell `tone` = below · danger · success · muted · strong, `link` cells, `footer` row "Rata-rata kelas"; 10 per page; filters Cari mahasiswa + Posisi terhadap target) and Aktivitas kelas → Daftar aktivitas (sorted by Tenggat, 10 per page, link buttons Ringkasan / Ubah). Header, group and footer rows are white like the body (27 Nov 2026). Other pages to follow. |
| Admin teacher mode (27 Nov 2026) | Admin sidebar has a "Gunakan fitur dosen" switch (off → teacher Ringkasan). `SharedSidebar` role = dosen with `asAdmin` shows the dosen menu with the switch on and the admin account; switching off returns to Admin — Mata kuliah. Which teacher features admins get (all courses, Ringkasan scope) is still to be designed. |
| Scroll behaviour (build guideline, 27 Nov 2026) | In the built app on desktop, the white main panel has a fixed height (viewport minus its 8px margins), the top bar is sticky, and only the content area scrolls; the sidebar stays sticky at full height. Recorded in "Layout — Halaman Regulate (desktop 1440)" only: design frames stay unclipped and show the full page height for manual review, so the rule is not in `regulate-app.css`. |
| Consistent sample data | Numbers must agree across student, teacher and admin frames; when a change forces other numbers to change, list them. |

## 2. Design files

| File | Role |
|---|---|
| Dashboard v.0.1 — https://claude.ai/artifact/E5bHnNSudx2YtgwFjXLbqX | Current design (landing, student, admin, teacher). |
| Conceptual Model — https://claude.ai/artifact/MaejXcBDgnCYT8NPWsfhqD | Model figure, caption, web version, integration, use case, data model. |
| Dashboard v.0.0 [archived] — https://claude.ai/artifact/ADP373onHn2FdhnzTbb2Yi | Earlier versions, API frames and all backup rows. |

Design system: Regulate (namespace `regulate`).

## 3. Language and wording

| Decision | Date |
|---|---|
| **Landing pages do not claim a SCELE integration.** The pilot only links activities to SCELE (manual linking by the lecturer), so the landing copy focuses on helping students take control of and regulate their learning. Hero: h1 "Take charge of how you learn." / "Kendalikan cara Anda belajar." (pre-v.0.1.7 wording; reverses #13); small label above the h1 "Self-regulated learning dashboard" / "Dasbor belajar mandiri" (was "Works with your SCELE courses" / "Terhubung dengan mata kuliah SCELE Anda"); main button "Get started" / "Mulai sekarang" (was "… with SCELE" / "Mulai dengan SCELE"); meta description without SCELE. Applied to all four landing frames (backups Bk023, Bk024). Anywhere in the product, describe SCELE as linked, never as integrated or synced. | 27 Nov 2026 |
| Attribution option "Tingkat kesulitan aktivitas" (not "…tugas"), in line with "aktivitas" for learning items. (v.0.1.7, #5) | 27 Nov 2026 |
| ~~Landing h1 names what Regulate is and for whom ("A learning dashboard for your SCELE courses." / "Dasbor belajar untuk mata kuliah SCELE Anda.")~~ (v.0.1.7, #13) — superseded the same day by the row above. | 27 Nov 2026 |
| Landing pages use the app's terms: "aktivitas" / "activities" (not "tugas" / "tasks"); sub-steps of an activity are "langkah-langkah" / "steps" (matches "Langkah-langkah" in the app); "aktivitas dinilai". Landing pages carry a meta description (EN + ID). (v.0.1.7) | 27 Nov 2026 |
| Graded items are "aktivitas dinilai", grouped under "komponen nilai"; "penilaian" only for self-assessment (also in accessible labels, e.g. the Evaluasi list is "Pilih aktivitas dinilai"). Weights shown as share of the final grade ("10% dari nilai akhir"). | 25 Nov 2026 |
| The field/column for related activities is labelled "Aktivitas" (not "Latihan" / "Aktivitas latihan"). | 25 Nov 2026 |
| Use "aktivitas" instead of "tugas" for learning activities; types Tugas, Kuis, Ujian, Materi. | Nov 2026 |
| UI in Bahasa Indonesia, formal "Anda"; all times in WIB. | Nov 2026 |

## 4. Grading model

| Decision | Date |
|---|---|
| **Option 2 adopted:** grade components (Kuis 30 · Tugas 30 · UTS 20 · Proyek akhir 20; total 100%) matching the RPS and SCELE gradebook categories. Every graded item is an activity. | 25 Nov 2026 |
| Weights inside a component: Sama rata (default) or Kustom as % of the component (graded activities must total 100%). "Dinilai" off = 0%: the activity only prepares its component (replaces `linkedAssessmentId`). | 25 Nov 2026 |
| A component without activities gets an automatic Ujian activity (e.g. UTS) so students can plan study time. | 25 Nov 2026 |
| One activity belongs to at most one component; activities may have no component ("tanpa komponen nilai"). | 25 Nov 2026 |
| What-if simulation "Simpan sebagai target" (option C): one save updates the course target to the projected final grade (one decimal) and each activity target whose simulated score differs, after a confirmation dialog listing every change (old → new). Logged as one `goalSet` per change with source = simulation. "Atur ulang" resets the simulation to the saved targets. | 27 Nov 2026 |
| Students set target, expected score and confidence (1–5) per graded activity, grouped by component. | 25 Nov 2026 |
| All 16 activities (graded and ungraded) appear on the student board by status; upcoming graded ones are plannable. | 25 Nov 2026 |

## 5. Teacher dashboard

| Decision | Date |
|---|---|
| **v.0.1.5:** Aktivitas kelas split into four sub-tabs (segmented buttons): Daftar aktivitas (default) · Status & hambatan (G1, G2, G5) · Perencanaan & waktu (G3, R1, R3, R4, process-graph link) · Usaha & hasil (R2, A2, A1, A3). Kelompok filter above the sub-tabs, applies to all. Content unchanged. All teacher links open Daftar aktivitas; the Ringkasan card "Aktivitas terhambat" opens Status & hambatan. Backup Bk019. | 26 Nov 2026 |
| Status matrix colours (now per §5a: Selesai #15803D, Terhambat #DC2626): previously Selesai #17803D, Sedang dikerjakan yellow (unchanged), Terhambat #DD2727, Mandek dark grey #44403C, Akan dikerjakan light grey; colour only, no letters. | 26 Nov 2026 |
| Nilai kelas includes a table of all students' grades (P7) grouped by component, with targets, current average and gap. | 26 Nov 2026 |
| Sidebar mirrors the student sidebar: "Ringkasan" (no sublabel), "Mata kuliah" heading, search box, course item, "Semua mata kuliah" (→ course table on Ringkasan). | 26 Nov 2026 |
| Course-based structure: Ringkasan → course tabs Nilai kelas · Aktivitas kelas · Mahasiswa · Intervensi · Pertanyaan & catatan; sub-pages Ringkasan aktivitas, Ringkasan mahasiswa, Kelola komponen nilai, Ubah aktivitas / komponen, Input nilai, process graphs. | 25 Nov 2026 |
| Indicator codes are researcher-only: never shown to lecturers; kept as `data-indicator` attributes and in grey canvas notes. | 25 Nov 2026 |
| Lecturer manages activities, components and grades for the pilot; admin role can be added later. | 25 Nov 2026 |

## 5c. v.0.1.7 design-review fixes (applied 27 Nov 2026)

First run of the `regulate-design-review` skill (16 findings, 0 Blocker, 2 Major). Numbers = rows in the findings Sheet. All went through review frames: first batch (backup `Bk021`); #5 and #13 after the user's decision and manual review (backup `Bk022`); then the landing hero was reverted and SCELE wording toned down on all four landing frames (§3; backups `Bk023`, `Bk024`). All review copies were deleted on request; three grey notes next to the design notes record the changes.

| Frames | Change |
|---|---|
| Student Nilai × 3 | (#1) Target marker on the "Posisi nilai akhir" bars: dashed amber #CA8A04 → short 2 px ink tick (same as the target bar); InfoTip text now "Garis tegak hitam pendek menandai target Anda." (#2) Simulation change rows highlighted with accent #FEF9C3 (was #FEFCE8). (#3) Unselected summary option cards: border #E7E5E4 (was yellow). (#4) aria-label "Pilih aktivitas dinilai". (#6) "Perkiraan nilai akhir" number 30 px (was 32). (#5) Attribution option "Tingkat kesulitan aktivitas" (was "…tugas"). |
| Teacher Aktivitas kelas — Perencanaan & waktu, Usaha & hasil | (#9) Legend and hover dots use radius 50% (were 5 px). (#10) KPI "41%" 30 px (was 32). |
| Shared Sidebar (all pages) | (#11) Selected menu item = accent fill #FEF9C3 only, no border (was a yellow #FACC15 border). |
| Landing × 4 (EN, ID, desktop, mobile) | (#12) "aktivitas"/"activities" and "langkah-langkah"/"steps" instead of "tugas"/"tasks"; (#14) "aktivitas dinilai"; (#15) meta description; (#16) type sizes on the scale (15/17/26/32 → 16/18/24/30 px) and radius 5 → 6 px. (#13) h1 with SCELE — replaced by the toned-down hero (§3) on all four. |
| Not changed | (#7/#8) raw stone #D6D3D1 kept as is. |

## 5b. Descriptions and info icons (applied 26 Nov 2026)

| Decision | Detail |
|---|---|
| Component | New `InfoTip` in the Regulate design system: an (i) button after a card or section title; click opens a 280px popover, a second click, Esc or a click outside closes it. Hover `Tooltip` stays for a few words of label only. |
| Student pages | How-to-read text and definitions (e.g. Rencana vs. aktual legends, Posisi nilai akhir scale, Sudah dinilai, Minggu Anda dalam angka, Kalender tenggat) moved fully behind the (i). Prompts to act keep a short visible line (~6 words), with the full text in the (i): Rencana langkah, Langkah-langkah, Simulasi, Aktivitas dinilai yang akan datang, Ringkasan minggu. |
| Teacher pages | "Lighter": a one-line subtitle stays; only the how-to-read part (legend, unit, threshold, scope) moved into the (i). Already-short one-liners unchanged. |
| Always visible | Live status and data (incl. dynamic descriptions), privacy and transparency notes, rules that prevent mistakes (e.g. timer behaviour, "Tidak ada yang wajib diisi"), provisional thresholds, research-logging notice, and "Tindakan:" lines. |
| Backups | 19 frames backed up in the archive (Bk016). The review copy KursusTinjauanID011 was not changed. |

## 5a. Colour system (applied 26 Nov 2026)

| Role | Decision |
|---|---|
| Theme | Brand yellow #FDE047 (and its ramp) stays a fill only, always with an ink (#292524) edge/text: primary action, selected state, progress, status "Sedang dikerjakan". |
| Selected menu item | Accent #FEF9C3 fill only, no border (sidebar; v.0.1.7, 27 Nov 2026). |
| Success | #15803D · light background #F0FDF4 · border #BBF7D0 · text on light #166534. Always with a word or icon. |
| Danger | #DC2626 · light background #FEF2F2 · border #FECACA · text on light #991B1B. Always with a word or icon. |
| Warning (new) | Orange-leaning amber #B45309 · light background #FFFBEB · border #FDE68A · text on light #92400E (distinct from brand yellow). |
| Status matrix | Selesai #15803D, Terhambat #DC2626 (token values replace #17803D / #DD2727), Mandek #44403C, Akan dikerjakan #F5F5F4, Sedang dikerjakan yellow + ink edge. |
| Categorical charts | Blue #2A78D6 · teal #0E9384 · violet #5B21B6 · pink #E2577F · brown #8A5A2B · grey #57534E for "Lainnya". Avoids yellow/green/red; all ≥ 3.5:1 on white; checked for protan/deutan/tritan separation (proposed violet #6E56CF and magenta #C2417A were swapped because they merged with blue / teal under deuteranopia). |
| Level charts | Yellow ramp #FEF9C3 → #FDE047 → #EAB308 → #A16207 → #713F12 (more = darker). |
| Below/above target | Blue (above) vs brown/amber (below) with a neutral midpoint; never red/green. |
| Progress bars | Soft: `primary-hover` yellow (#FACC15) fill, no ink edge, on a `border` (#E7E5E4) track; the value is always shown in text nearby. Applies to the Progress component and custom bars (Ringkasan grade bar: earned = soft yellow, lost = hatched grey, target = short ink tick instead of the dashed amber line). Target markers on any bar are short ink ticks (Nilai "Posisi nilai akhir" bars too, since v.0.1.7). Replaces the ink-edged Progress and the ink-filled grade bar (26 Nov 2026). |
| Neutrals | Stone only; replace #171717, #737373, #E5E5E5, #F5F5F5 with stone tokens. |
| Rollout | Tokens added to the Regulate design system (chart-1…5, chart-other, seq-1…5, chart-above/below, warning set, -subtle/-border/-text tints, status-stale) and copied into Dashboard v.0.1 and the Conceptual Model; frames updated directly. Backups of frames with visible changes in the archive (Bk015). Danger text on white uses #DC2626; #991B1B only on light red backgrounds. |

## 6. Student dashboard

| Decision | Date |
|---|---|
| Sidebar "Ringkasan" has no "semua mata kuliah" sublabel (same as the teacher side). | 26 Nov 2026 |
| Papan has a Selesai column (light green). An info tooltip explains that "Selesai" in Regulate does not submit anything in SCELE; cards that involve submission get a "Sudah submit di SCELE" checkbox (Materi gets a note instead). | 26 Nov 2026 |
| Card headers: title above description; plan list items: name above meta line. | 25 Nov 2026 |
| Linimasa: Aktual bars extend to "Sekarang", include blocked time, yellow dashed; data is a snapshot with "Diperbarui …" and a refresh button. | Nov 2026 |
| Tinjauan mingguan privacy text widened to match what lecturers can see. | 25 Nov 2026 |

## 7. Sample data anchors

- Aktivitas belajar "Sesuai rencana 13 dari 16" bar = 81% (was 70%, corrected 26 Nov 2026).

- Now = Kam 19 Nov 2026, 10.00 WIB. Only IMK (CSCE604123, Gasal 2026/2027, 24 students, SCELE IMK-2026-1) uses Regulate.
- Student S01 Ayu (u_0001 / p_7f3a, ayu@ui.ac.id): scores Kuis 1 82 · Tugas 1 74 · UTS 68 · Kuis 2 71 → average 72,7; target 80; needed 88,9; projection 78,8.
- Lecturer Dosen A (dosen.a@ui.ac.id); admin Admin Regulate (admin@ui.ac.id).
- Class averages: Kuis 1 79,4 · Tugas 1 76,2 · UTS 70,8 · Kuis 2 75,3; class average 74,7; 12 of 21 students with a target are below it; median gap −2,0, mean gap −3,2.
- Proyek akhir = Presentasi (30% of component → 6%) + Laporan (70% → 14%); Draf persona, Proposal and Rekrut peserta are 0%.
- Trace export: 3.427 records; class process graph: 223 cases.

## 8. Open items

- Design-review skill: an update card was offered on 27 Nov 2026 (SCELE linked, not integrated; landing h1 = student benefit; new `landing.scele-claim` check). Done once the user saves it. The skill should also follow the new backup rule (no backups unless asked).
- Next archive backup suffix (only when a backup is asked for): Bk025.
- Raw stone #D6D3D1 (underlines, histogram bars) is not a design-system token; add as a token or keep.

- Aktivitas kelas (v.0.1.5) options not yet done: "Perlu tindakan" strip, merging Titik rawan + Jalur mencari bantuan, links from Intervensi to a specific sub-tab, `teacherViewOpened` trace (which sub-tab lecturers open).

- Help use is not yet logged: a possible `helpOpened` trace (which (i), when) for the codebook — not decided.
- The Papan "Selesai" notice still uses the hover Tooltip with a CSS override; could switch to InfoTip.
- Nilai semua mahasiswa (P7) marks below-target cells with a red tint, while charts use brown/blue for below/above; to align or keep as a danger cue.

- Per-student process graphs (Graf proses mahasiswa) still use the earlier activity set.
- "Sudah submit di SCELE" is not yet in the codebook (e.g. `sceleSubmittedConfirmed`) and not yet on Detail aktivitas / Evaluasi.
- Teacher "Semua mata kuliah" has no dedicated course page.
- Project instructions for the Regulate project: draft prepared; to be pasted by the user.
