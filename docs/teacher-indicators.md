> Snapshot of the claude.ai Project doc "Regulate — Teacher-Facing Dashboard: Indicators and Visualizations" (27 Nov 2026). The Project doc is the source of truth; this copy is refreshed when a PRD is approved. Indicator codes here are researcher-only and never appear as visible UI text.

# Regulate — Teacher-Facing Dashboard: Indicators and Visualizations

Working draft (Nov 2026; updated 25 Nov 2026 for the course-based teacher dashboard v.0.1.1, the v.0.1.2 additions — activity management, grade entry, activity summary, student summary — the v.0.1.3 grade components, the v.0.1.4 consistency pass and the v.0.1.5 split of Aktivitas kelas into four sub-tabs, 26 Nov 2026). Companion to `claude/Regulate SRL feature grid and trace codebook.md`, which defines every indicator variable used here.

---

## 1. Role in the conceptual model

The teacher-facing dashboard is a **monitoring tool**. It draws on the Unified SRL Trace produced by the student-facing dashboard (cycle steps 2–4) and supports cycle step 5, **Guide**.

- **Monitors** students' performance (scores, grades), progress (completion status) and process (time spent; expected vs. actual).
- **Informs two decisions:** (1) how to conduct or adjust the learning activity; (2) which students need an intervention, and what kind.
- **Action:** external monitoring and feedback (direct instruction, in Dignath & Veenman's terms), complementing the student dashboard's indirect activation.

Feasibility key: 🟢 feasible now (LMS + dashboard data) · 🟡 needs new logging, analytics rules or teacher workflow

---

## 2. Structure (course-based)

The teacher dashboard mirrors the student revamp: sidebar **Ringkasan** (all courses the lecturer teaches on Regulate) → course (IMK) → tabs. The course header shows when indicators were last computed (daily job, e.g. "diperbarui Kam 19 Nov, 06.00 WIB").

The same trace is offered in three views: **by indicator** (class-level charts in the tabs), **by activity** (Ringkasan aktivitas) and **by student** (Ringkasan mahasiswa). All three must use one shared calculation layer so numbers agree across pages.

**Indicator codes (P1–P7, G1–G5, R1–R6, X1, X3, A1–A3, I1, M1–M3, T1–T8, S1–S8) are researcher-facing only.** Lecturers see section titles, never codes. Each section carries its code as an invisible attribute (`data-indicator="P4"`), so developers and analyses can address it and the teacher-side trace can log views and interactions per indicator (§7). On the design canvas, a grey note under each teacher frame lists the codes top to bottom. Dimension names (Kinerja / Kemajuan / Proses) appear only inside detail panels such as the student summary.

| Tab / page | Content |
|---|---|
| **Nilai kelas** | P1–P5, P6, P7 (grade table for all students) |
| **Aktivitas kelas** | Four sub-tabs (segmented buttons below the course tabs; the Kelompok filter sits above them and applies to all): **Daftar aktivitas** (default) — M1 activity list (component and grade share per row, filters Semua / Dinilai / Tidak dinilai, Ringkasan / Ubah per row, Tambah aktivitas, Impor CSV, link to Kelola komponen nilai) · **Status & hambatan** — G1, G2, G5 · **Perencanaan & waktu** — G3, R1, R3, R4, link to the class process graph · **Usaha & hasil** — R2, A2, A1, A3 |
| ↳ Ubah / Tambah aktivitas | Activity form (§5.1) |
| ↳ Kelola komponen nilai | M2 grade components with their activities · Ubah komponen nilai · Input nilai per graded activity (§5.2) |
| ↳ Ringkasan aktivitas | T1–T8 (§4.5) |
| **Mahasiswa** | M3 student list · X1 · G4; names open the student summary |
| ↳ Ringkasan mahasiswa | S1–S8 (§4.6), incl. the former X2 drill-down (S5) and per-student R5 (S4) |
| **Intervensi** | X3, R6, I1; banner for unanswered questions |
| **Pertanyaan & catatan** | Inbox: "Tanya dosen" questions with reply, and reflection notes students chose to share |

Wording follows the student pages: "aktivitas" (types **Tugas**, **Kuis**, **Ujian**, **Materi**) rather than "tugas"; graded items are **aktivitas dinilai**, grouped under a **komponen nilai**, and their weight is always shown as a share of the final grade ("10% dari nilai akhir"); "penilaian" is kept only for self-assessment. Statuses Akan dikerjakan / Sedang dikerjakan / Terhambat / Selesai, plus the teacher-only derived state **Mandek** (past due, no update); all times in WIB. Student and activity names anywhere link to their summary page.

---

## 3. Feature list

### Monitoring
- **Performance** 🟢 — class distribution of current grade vs. goal; graded activities where many students fall below target; confidence vs. calibration.
- **Progress** 🟢 — completion per activity across the class; blocked-activity heatmap; students who haven't planned.
- **Process** 🟢 — planned vs. actual start per activity; estimated vs. actual time per activity and per activity type, and vs. the lecturer's expected effort.
- **SRL engagement** 🟡 — phase-coverage map from the trace (who plans, monitors, reflects).
- **Per student / per activity** 🟢 — student summary and activity summary re-slice the same indicators.

### Decision: adjust the learning activity
- 🟢 Activity-level insights and the activity summary — blocks, delays, "lebih sulit" ratings, questions, planning vs. expected effort.
- 🟢 Workload calendar — clustered deadlines within the course.
- 🟢 Preparation vs. score — time on ungraded (0%) activities of a component (e.g. Latihan Kuis) vs. the graded score in that component.
- 🟢 Activity and grade management — deadline, description, expected effort, steps, component and "Dinilai"; component weights (Sama rata / Kustom); archive. Every change is logged (`activityChanged`, `gradeComponentChanged`).

### Decision: which students need intervention, and what kind
- 🟡 Flags mapped to intervention type (§6).
- 🟢 Student summary — goals, plan vs. actual, reflections, trends, process path, contact history.

### Action: external monitoring & feedback
- 🟢 Prompts to a student or group; shown in the student dashboard (Tinjauan mingguan, Detail aktivitas).
- 🟢 Inbox — answer "Tanya dosen" questions (optionally as an announcement to all students with the same blocker); read shared notes.
- 🟢 Activity-level actions — reminders to non-planners / students past their planned start, clarification, resource, deadline change.
- 🟡 Follow-up tracking — whether the student acted within 72 h.

---

## 4. Visualizations

Reference frames: **goal** · **plan** · **peer** · **temporal** · **deadline**. The "#" column is the researcher-facing indicator code.

### 4.1 Performance Monitoring (tab Nilai kelas)

| # | Visualization & features | Goal / possible actions | Main indicators | Reference frame | Scope | Viz type |
|---|---|---|---|---|---|---|
| P1 | **Class goal overview** — % with goal set, median goal gap, % below goal | Judge whether course-wide support or goal-setting prompts are needed | `goalSet`, `goalAttainmentGap` | Goal | Course | Number cards |
| P2 | **Goal vs. current grade** — one dot per student, y = x line | Contact students behind goal; discuss over- or under-ambitious goals | `goalLevelRelative`, `goalAttainmentGap` | Goal + peer | Individuals within course | Scatter plot |
| P3 | **Goal gap by graded activity** — mean gap and % below target per graded activity (rollable to component) | Revisit topics, remedial session, check rubric clarity | `goalAttainmentGap` (A), `goalGranularityRatio` | Goal | Course | Diverging bar chart |
| P4 | **Calibration view** — expected vs. actual per graded activity; overconfident highlighted (> +15) | Self-assessment training, practice quizzes | `expectedScore`, `calibrationErrorSigned`, `calibrationErrorAbs` | Self-judgment | Individuals + course | Box plot |
| P5 | **Post-grade reflection & attribution** | Address maladaptive attributions; encourage appraisal | `gradeReflectionSubmitted`, `gradeReflectionAttribution`, `postFeedbackGoalChange`, `appraisalMode` | Goal | Course | Stacked bar + card |
| P6 | **Confidence vs. accuracy** — per confidence level 1–5 | Show examples/rubrics before high-stakes graded activities | `confidenceRating`, `calibrationErrorSigned` | Self-judgment | Course | Table |
| P7 | **Grades of all students** — one row per student: target, score per graded activity grouped by component (share of final grade in the header), current average over the graded weight, gap to own target; cells below the student's target highlighted; class averages and median gap; CSV download; names open the student summary | Spot individual low scores per activity; prepare grade entry and follow-up | scores (A), `goalAttainmentGap` (C) | Goal | Individuals within course | Table |

### 4.2 Progress and Process Monitoring (tab Aktivitas kelas, sub-tabs per §2)

| # | Visualization & features | Goal / possible actions | Main indicators | Reference frame | Scope | Viz type |
|---|---|---|---|---|---|---|
| M1 | **Activity list** — per activity: type, component and share of the final grade (or "0%, tidak dinilai" / "tanpa komponen nilai"), deadline, status, % planned, % done on time, currently blocked, questions; Ringkasan / Ubah | Entry point to activity summary and management | `planningCoverageRatio` (T), `onTimeCompletionRate` (T), `blockedEventCount`, `lecturerQuestionAsked`, `effectiveWeight` | Deadline | Course | Table |
| G1 | **Activity status matrix** — activities with deadlines around today | See who is behind on what | status events, `staleTaskCount`, `onTimeCompletionRate` | Deadline | Individuals | Matrix |
| A1 | **Summary by activity type** (Tugas, Kuis, Ujian, Materi) | State expected effort for the type that overruns most | `activityType`, `effortEstimateHours`, `timeOnTaskHours`, `onTimeCompletionRate` | Plan | Course | Table |
| G2 | **Blocked-activity hotspots** | Clarify instructions, add resources | `blockedEventCount`, `blockedDurationHours`, `blockedReasonCategory`, `unblockLatencyHours` | Across activities | Course | Stacked bar |
| G5 | **Help-seeking paths** | Signpost resources | `unblockPath`, `sceleLinkClickCount`, `lecturerQuestionAsked` | Across activities | Course | Stacked bar |
| G3 | **Planning & adherence trend** | Time planning prompts | `planningCoverageRatio`, `planAdherenceRatio` | Temporal | Course | Line chart with band |
| R1 | **Start slippage by activity** | Milestones; earlier release | `startSlippageHours`, `planningLeadTimeDays` | Plan | Course | Box plot |
| R2 | **Estimated vs. actual effort** | Recalibrate workload; communicate expected effort | `effortEstimateHours`, `timeOnTaskHours`, `effortEstimationError`, `expectedEffortHours` | Plan | Course | Dumbbell chart |
| A2 | **Felt difficulty & causes of overrun** | Examples/templates for activities rated harder | `doneDifficultyRating`, `processReflectionCategory` | Plan | Course | Table |
| R3 | **Work pattern before deadline** | Spaced work; intermediate submissions | `lastMinuteSessionRatio`, `workSessionCount` | Deadline | Course | Histogram |
| R4 | **Workload calendar** | Move clustered deadlines within the course | planned start/finish, `plannedBufferDays`, `workloadGini` | Calendar | Course | Calendar heatmap |
| A3 | **Preparation vs. quiz score** — time on Latihan Kuis (0%, komponen Kuis) vs. the graded quiz score | Remind students without logged preparation | `prepTimeInComponentHours`, graded score | Peer | Course | Table |

### 4.3 Students (tab Mahasiswa)

| # | Visualization & features | Goal / possible actions | Main indicators | Reference frame | Scope | Viz type |
|---|---|---|---|---|---|---|
| M3 | **Student list** — goal gap, on plan, blocked, reflections, flags, last active; name opens the student summary | Entry point to individual follow-up | `goalAttainmentGap`, `planAdherenceRatio`, `appraisalEngaged`, flags | Goal, peer | Individuals | Table |
| X1 | **SRL phase coverage map** | Phase-specific prompts | `phaseCoverage`, `cycleCompletionRate`, `appraisalEngaged` | Peer | Individuals | Heatmap |
| G4 | **Disengagement watchlist** | Individual outreach; advisor referral | `staleTaskCount`, `abandonedTaskCount`, `slippedTaskReplanRatio`, `statusUpdatesPerWeek` | Peer (thresholds) | Individuals | Table |
| — | **Process graphs** — class and per student (case = student + activity, graded activities included); main path Rencana disimpan → Dimulai → Waktu tercatat → Selesai → Dinilai; graded activities continue to Nilai dilihat → Nilai dievaluasi; "Kesesuaian jalur utama" = share of activities on the main path | Discuss deviations | `eventSequence` | Peer | Course / individual | Directly-follows graph |

### 4.4 Intervention (tab Intervensi)

| # | Visualization & features | Goal / possible actions | Main indicators | Reference frame | Scope | Viz type |
|---|---|---|---|---|---|---|
| X3 | **Intervention flags** | Act on who needs intervention; close cycle step 5 | rule combinations (§6) | Thresholds | Individuals | Table + action |
| R6 | **Alert effectiveness** | Adjust alert timing or wording | `alertResponded`, `alertResponseLatencyHours` | System | Course | Number cards |
| I1 | **Weekly decisions & strategies** | Remind non-openers | `digestDecision`, `strategySelected`, `progressReviewViewed` | Temporal | Course | Tables |

### 4.5 Activity summary (one activity, e.g. Tugas 2)

Header: type, component and share of the final grade ("Komponen Tugas · 15% dari nilai akhir"), release, deadline (WIB, time left), lecturer's expected effort; Ubah aktivitas; SCELE link.

| # | Section | Main indicators | Notes |
|---|---|---|---|
| T1 | **Class stages** — not planned · planned, not yet due to start · planned, start passed · in progress · done · reflected | status, `plannedStartAt`, `actualStartAt`, `processReflectionSubmitted` | Bars over all enrolled students |
| T2 | **Planning** — % planned (with change after reminders), median student estimate vs. `expectedEffortHours`, planned-finish-to-deadline gap | `planningCoverageRatio`, `effortEstimateHours`, `plannedBufferDays` | |
| T3 | **Who to nudge** — not planned; start passed; one-click reminder | status, `plannedStartAt` | Names link to the student summary |
| T4 | **Time and start** — median start slippage, time logged, comparison with a same-type activity | `startSlippageHours`, `timeOnTaskHours`, `effortEstimationError` | |
| T5 | **Blockers and questions** | `blockedEventCount`, `blockedReasonCategory`, `lecturerQuestionAsked` | Links to the inbox |
| T6 | **Felt difficulty and reflection** | `doneDifficultyRating`, `processReflectionCategory` | After students mark done |
| T7 | **Grade outcome (Hasil nilai)** — distribution, gap to target, calibration, score by start timing and by preparation time; for an ungraded activity, the score of its component | score, `goalAttainmentGap`, `calibrationErrorSigned`, `startSlippageHours`, `prepTimeInComponentHours` | After grades are published |
| T8 | **Actions** — clarification, resource, deadline change, process graph filtered to the activity | `activityChanged`, `interventionSent` | Logged as activity adjustments |

### 4.6 Student summary (one student in one course, e.g. S01)

Header: pseudonym, status, last active, open flags and questions, a short generated summary and a suggested action (reply, message).

| # | Section | Main indicators | Reference frame |
|---|---|---|---|
| S1 | **Grades** — current vs. target, average needed, expected − actual; per graded activity (with its component) target, expected, confidence, score, reflection/attribution | `goalAttainmentGap`, `expectedScore`, `confidenceRating`, `calibrationErrorSigned`, `gradeReflectionAttribution` | Own target, class median |
| S2 | **Activities** — all activities (graded ones marked with their grade share) with status, timing, estimate → actual, felt difficulty/cause; planning coverage, on-time rate, stale count | status, `effortEstimationError`, `doneDifficultyRating`, `planningCoverageRatio`, `onTimeCompletionRate`, `staleTaskCount` | Plan, deadline |
| S3 | **Time logged per week** | `timeOnTaskHours` (weekly) | Own history |
| S4 | **Estimation accuracy** — per activity in completion order, trend vs. class median (former per-student R5) | `effortEstimationError`, `estimationErrorTrend` | Own history, class |
| S5 | **SRL pattern vs. class** — 3 × 3 phase × dimension key indicators (former X2) | one key indicator per cell | Class median |
| S6 | **Reflection and weekly review** — activity and grade reflections, buffers accepted, weekly decisions and strategies, shared notes | `appraisalEngaged`, `suggestedBufferAccepted`, `digestDecision`, `strategySelected`, `reflectionNoteShared` | — |
| S7 | **Process path** — main-path fit, deviations, link to the per-student graph | `eventSequence` | Class main path |
| S8 | **Contact history** — flags, prompts, questions and the student's responses in one timeline | teacher-side trace (§7) | — |

No rankings or percentiles against classmates; comparisons use the student's own targets and the class median.

---

## 5. Management (lecturer, pilot)

### 5.1 Activity management
- **Fields:** name, type (Tugas / Kuis / Ujian / Materi), **komponen nilai** (or "tanpa komponen nilai") with a **Dinilai** switch — the form shows the resulting weight, e.g. "Sama rata: 50% dari komponen Tugas = 15% dari nilai akhir", or "Tidak dinilai (0%)" for activities that only prepare the component — release date, deadline (date + time, WIB), SCELE URL, description, **expected effort (hours)**, suggested steps, published/draft.
- **Deadline change:** the form shows how many saved student plans would end after an earlier deadline; affected students get a notice and the Persiapan form warns them. A later deadline notifies students who planned.
- **Archive, not delete,** once any student has interacted with the activity (keeps the trace intact).
- **Import:** CSV for a semester's activities. Later, SCELE-synced items become read-only (`source = scele`); Regulate-only items (readings, practice) stay editable.

### 5.2 Grade components and grade entry
- **Kelola komponen nilai** (lecturer for the pilot; admin can be added later): one card per component (name, % of the final grade, weight mode, progress, class average) listing its activities with weight in the component, share of the final grade, date, status and class average. Activities without a component are listed separately.
- **Component weights** must total 100% before saving; changes immediately affect students' average, score needed and simulation.
- **Weights inside a component:** **Sama rata** (default; graded activities share the component equally and rebalance when one is added) or **Kustom** (% of the component; graded activities must total 100%). Turning "Dinilai" off sets an activity to 0%.
- **Rules:** one activity belongs to at most one component; a component without activities (e.g. UTS) automatically gets an Ujian activity so students can plan study time; "open grade reflection on publish" is set per component.
- **Grade entry per graded activity** (manual or CSV); the header shows the component and share of the final grade. Publishing shows the grade to students and opens the grade reflection; later edits notify the affected students.

---

## 6. Draft intervention flag rules (for X3)

Thresholds are placeholders to be calibrated on pilot data (together with the student-side slip-alert grace period and on-plan tolerance).

| Flag | Rule (example) | Suggested intervention |
|---|---|---|
| No planning | `planningCoverageRatio` < 0.3 after week 3 | Planning prompt; short planning micro-lesson |
| Chronic slippage | median start slippage > 2 days across ≥ 3 activities | Time-management feedback; suggest earlier milestones |
| Behind goal | `goalAttainmentGap` < −10 while `onTimeCompletionRate` ≥ 0.8 | Content / academic support |
| Overconfidence | mean `calibrationErrorSigned` > 15 across graded activities | Self-assessment practice; calibration feedback |
| Disengagement | `staleTaskCount` ≥ 2 and no events in 7 days | Personal check-in; advisor referral |
| No appraisal | `appraisalEngaged` = false on ≥ 80% of completed activities | Appraisal prompt; explain feed-forward benefit |
| Persistent blocking | `blockedDurationHours` > 72 on any activity, or an unanswered "Tanya dosen" question > 24 h | Direct help offer; answer in the inbox |

---

## 7. Teacher-side trace (suggested)

| Variable | Description | Unit | Data type |
|---|---|---|---|
| `teacherId` | Teacher identifier (pseudonymised) | — | string |
| `indicatorViewed` | Section scrolled into view, keyed by its `data-indicator` code | view | enum (code) |
| `indicatorInteraction` | Hover, filter, drill-down or action within a section, keyed by code | interaction | enum (code) + type |
| `studentSummaryViewed` | Student summary opened (student, entry point) | view | object |
| `activitySummaryViewed` | Activity summary opened (activity, entry point) | view | object |
| `activityChanged` | Activity created, edited or archived: field (incl. component, "Dinilai"), old and new value, # students affected | activity | object |
| `gradeComponentChanged` | Component created or edited: weight, weight mode, activity weights, graded flags; replaces `assessmentChanged` | component | object |
| `gradesPublished` | Grades published or edited for a graded activity | activity | object |
| `flagRaisedAt` / `flagType` / `flagViewed` | Flag lifecycle | flag | datetime / enum / boolean |
| `interventionSent` / `interventionType` / `interventionLatencyHours` / `interventionRecipientScope` | Prompt sent in response | intervention | mixed |
| `studentActedOnIntervention` / `studentActionLatencyHours` | Student performed a related action within 72 h | intervention | boolean / float |
| `questionAnswered` / `questionReplyLatencyHours` / `replyBroadcast` | "Tanya dosen" replies | question | boolean / float / boolean |
| `sharedNoteViewed` | Teacher opened a shared reflection note | note | boolean |
| `teacherDashboardViewsPerWeek` | Teacher dashboard sessions per week | course | float |

---

## 8. Design notes

- **Privacy and transparency:** P2, P4, P7, G1, G4, M3, X1, X3, the student summary and the per-student process graph identify individual students. The student "Data yang dicatat dan alasannya" section states what lecturers can see (class summary; targets, expected scores, confidence and grades; status, plans and time logged per activity; estimation gaps and trend; activity history and process graph; flags; prompt follow-up; "Tanya dosen" questions). Reflection notes reach teachers only via the opt-in share toggle.
- **Actionability:** every view links to an action (prompt, inbox reply, activity adjustment).
- **Minimal first version:** P2, G1, G2, R1, X1, X3, the inbox, the activity list and activity form, grade components and grade entry, and the student summary.
- **Evaluation ideas:** usability and perceived usefulness; which indicators and summary pages lecturers use (`indicatorViewed`, `studentSummaryViewed`, `activitySummaryViewed`); whether flags, questions and activity changes lead to interventions and student adjustment.
- **Open item:** the per-student process graphs still use the earlier activity set (sample data).
- **Design file:** Dashboard v.0.1 canvas, frames "[v.0.1.1]" to "[v.0.1.5] Teacher — …" (see `claude/Regulate design files.md`).
