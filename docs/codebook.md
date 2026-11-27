> Snapshot of the claude.ai Project doc "Regulate — SRL Feature Grid and Unified SRL Trace Codebook" (27 Nov 2026). The Project doc is the source of truth; this copy is refreshed when a PRD is approved. Trace event names and payloads in an approved PRD (§5c) win over this copy.

# Regulate — SRL Feature Grid and Unified SRL Trace Codebook

Working draft (Nov 2026; updated 25 Nov 2026 to match the student revamp v.0.1, teacher dashboard v.0.1.1/v.0.1.2, grade components v.0.1.3 and the consistency pass v.0.1.4; 27 Nov 2026: simulation saves course and activity targets). Based on the Regulate conceptual model (student-facing dashboard as SRL practice tool, trace-data producer and insight provider; phases: preparatory / performance / appraisal; dimensions: performance / progress / process).

---

## 1. Student-facing feature grid (phase × dimension)

Key: ✅ in prototype · 🟢 feasible now · 🟡 needs new logging or rules · ★ added from the appraisal redesign (appraisal via system-generated comparisons, one-tap judgments and next-cycle decisions; free-text reflection optional)

Student pages (v.0.1): Ringkasan (cross-course) · Mata kuliah · course tabs **Nilai** (performance), **Aktivitas belajar** (progress + process, merged) and **Tinjauan mingguan**, each split into Persiapan / Pelaksanaan / Evaluasi; Detail aktivitas per activity. All times in WIB.

Grading model (v.0.1.3): every graded item is an **activity**. Activities have a type — **Tugas**, **Kuis** (kuis and latihan kuis), **Ujian** or **Materi** — and may belong to one **grade component** (komponen nilai, e.g. Kuis 30% · Tugas 30% · UTS 20% · Proyek akhir 20%; components total 100%). Inside a component, graded activities share its weight equally or by custom percentages; an activity with "Dinilai" off counts 0% and only prepares its component. A component without activities gets an automatic Ujian activity, so students can plan study time for exams. Students see graded items as **aktivitas dinilai**; "penilaian" is used only for self-assessment.

| Dimension | Preparatory | Performance | Appraisal |
|---|---|---|---|
| **Performance** (scores, grades) | ✅ Grade goal + what-if simulation (graded activities with their share of the final grade) · ✅ Target score per graded activity, grouped by component · ✅ Expected score + confidence (1–5) per graded activity · ★ Last attribution/note shown when setting next goal | ✅ Current grade vs. goal · ✅ "Score needed on remaining graded activities" · 🟡 Alert when trajectory drops below goal | ✅ Achieved vs. goal · ★ Auto expected vs. actual score · ★ Next-goal decision: keep / raise / lower · ★ Attribution chips · ★ Optional short note with "share with lecturer" toggle |
| **Progress** (completion) | ✅ One Persiapan form per activity (graded or not): objective, priority, dates and effort saved together · ✅ "Belum direncanakan" filter · ✅ Grade share or "tidak dinilai · persiapan …" shown per activity · ★ Reminder of last chosen strategy | ✅ Status board (Akan dikerjakan / Sedang dikerjakan / Terhambat / Selesai) · ✅ "Blocked" reason + "Tanya dosen" · ✅ On plan / behind plan / blocked summary · ✅ SCELE deep links | ✅ Weekly review ending in one decision · ★ One-step re-plan of unfinished activities · ★ Strategy pick (e.g., split task, ask for help sooner) |
| **Process** (time, expected vs. actual) | ✅ Planned start/finish: date required, time optional (empty start = 08.00, empty finish = 23.55 WIB, never after due) · ✅ Effort estimate (hours) · ★ Suggested estimate from past estimation error · ★ Past note shown for a similar activity | ✅ Linimasa (2 Minggu / Minggu / Hari) of plan vs. work sessions and Aktual span; snapshot with "Diperbarui" + refresh · ✅ Timer and manual sessions · ✅ Slip alert | ✅ Expected-vs-actual gap shown when an activity is done ("Perlu refleksi" filter) · ★ One-tap rating: easier / as expected / harder · ★ Cause chips (only if gap > 50%) · ★ Accept suggested buffer (extra hours) for the next similar activity · Summary table by type (estimation-trend chart removed from the student view; teachers still see it per student) |

**Rules for all appraisal cells (★):** prompt only when the gap is notable; ≤ 30 s per prompt (chips first, text optional, never required to finish an activity); return a personal pattern or tip after each response; one-line "why this matters" shown the first time only; personal learning journal collects reflections and adjustments.

Theoretical note: Zimmerman's self-reflection phase comprises self-judgment and self-reaction (incl. adaptive inferences), so behavioural adjustments are valid appraisal evidence, not only written reflections.

---

## 2. Unified SRL Trace — base event record

| Variable | Description | Data type |
|---|---|---|
| `studentId` | Student identifier (pseudonymised in exports as p_…; teacher views show S01–S24) | string |
| `courseId` | Course identifier | string |
| `activityId` | Activity ID, graded or not (null for course-level events); formerly `taskId` | string (nullable) |
| `eventType` | Interaction type, e.g. `goalSet` (one event per changed target, course or activity; payload: old and new value, `goalSetSource`), `planCreated` (= plan saved: objective, priority, dates, effort), `statusChanged` | enum |
| `phase` | preparatory / performance / appraisal | enum |
| `dimension` | performance / progress / process | enum |
| `timestamp` | When the event happened (stored UTC, shown WIB) | datetime |
| `payload` | Event details (old/new value, input) | JSON object |

---

## 3. Indicator codebook

Unit: C = per student per course · A = per student per graded activity (aktivitas dinilai) · T = per student per activity (graded or not) · E = per alert event · W→C = weekly, aggregated per course. A-level values can be rolled up per grade component (weighted by `weightInComponent`).

| Phase × Dimension | Variable | Description | Unit | Data type |
|---|---|---|---|---|
| Prep × Performance | `goalSet` | A grade goal was set | C / A | boolean |
| | `goalSetDelayDays` | Days from course opening / activity release to first goal set | C / A | float |
| | `goalLevelRelative` | Target minus prior average (or class median), in points | C / A | float |
| | `simulationRunCount` | What-if simulations run | C | integer |
| | `simulationScenariosBeforeGoal` | Simulations run before the goal was saved | C | integer |
| | `goalSetSource` | Where a target was set or changed: manual / simulation / weekly review / post-grade decision. Saving the what-if simulation ("Simpan sebagai target") replaces the course target with the projected final grade and each changed activity target with its simulated score, after a confirmation that lists every change | C / A | enum |
| | `simulationTargetsChanged` | Number of targets changed by one simulation save (course + activities) | C | integer |
| | `goalGranularityRatio` | Share of graded activities with own target (0–1) | C | float |
| | `expectedScore` | Self-estimated score before the graded activity (0–100); replaces `preAssessmentConfidence` | A | float |
| | `confidenceRating` | Confidence in that estimate (1–5) | A | integer |
| Perf × Performance | `gradeViewsPerWeek` | Grade page views per week | C | float |
| | `gradeViewDelayHours` | Hours from grade release to first view | A | float |
| | `goalGapViewRatio` | Goal-gap views while below goal ÷ all goal-gap views | C | float |
| | `scoreNeededCheckCount` | Uses of "score needed" feature | C | integer |
| | `midCourseGoalEditCount` | Goal edits after course start | C | integer |
| | `goalEditDirection` | Net direction: raised / lowered / unchanged | C | enum |
| Appraisal × Performance | `calibrationErrorAbs` | \|expected − actual score\| | A | float |
| | `calibrationErrorSigned` | Expected − actual (+ = overconfident) | A | float |
| | `goalAttainmentGap` | Achieved grade − goal | C / A | float |
| | `gradeReflectionSubmitted` | Reflection submitted after grade release | A | boolean |
| | `gradeReflectionDelayHours` | Hours from release to reflection | A | float |
| | `gradeReflectionLength` | Reflection length in words | A | integer |
| | `gradeReflectionAttribution` | effort / strategy / ability / task difficulty / external | A | enum |
| | `postFeedbackGoalChange` | Next goal: raised / lowered / unchanged | A | enum |
| | `reflectionNoteShared` | Note shared with lecturer | A / T | boolean |
| Prep × Progress | `planningCoverageRatio` | Share of activities with a saved plan before due (0–1); dates are part of every plan, so this also replaces `targetDatesSetRatio`. Per activity (T-level, across students) it is shown as "% direncanakan" | C | float |
| | `planningDelayDays` | Days from activity release (`activityReleasedAt`) to plan saved | T | float |
| | `planningLeadTimeDays` | Days from plan saved to due date | T | float |
| | `hasTaskObjective` | Personal objective entered | T | boolean |
| | `prioritySetCount` | Priority assignments | C | integer |
| | `priorityDueDateAlignment` | Rank correlation of priority vs. due-date order (−1 to 1) | C | float |
| | `activityType` | Tugas / Kuis / Ujian / Materi | T | enum |
| Perf × Progress | `statusUpdatesPerWeek` | Status changes per week | C | float |
| | `statusUpdateIntervalSd` | SD of hours between status updates | C | float |
| | `blockedEventCount` | Times marked blocked | T | integer |
| | `blockedDurationHours` | Total hours blocked | T | float |
| | `blockedReasonCategory` | content difficulty / missing resource / time / other | T | enum |
| | `unblockLatencyHours` | Hours from blocked to in progress | T | float |
| | `unblockPath` | Event types between blocked and in progress | T | array of string |
| | `lecturerQuestionAsked` | "Tanya dosen" used on a blocked activity | T | boolean |
| | `planAdherenceRatio` | Share of activities in planned status within their planned start–finish window (0–1) | C | float |
| | `sceleLinkClickCount` | Clicks from activity card to SCELE material | T | integer |
| Appraisal × Progress | `onTimeCompletionRate` | Activities done by due time ÷ activities planned (0–1) | C | float |
| | `progressReviewViewed` | Weekly review opened | W→C | boolean |
| | `progressReviewDurationSec` | Seconds spent on review | W→C | integer |
| | `replanCount` | Re-plan actions | C | integer |
| | `slippedTaskReplanRatio` | Share of slipped activities re-planned (0–1) | C | float |
| | `abandonedTaskCount` | Slipped activities neither re-planned nor completed | C | integer |
| | `staleTaskCount` | Open activities past due with no update | C | integer |
| Prep × Process | `plannedStartAt` / `plannedFinishAt` | Planned window (date + optional time; defaults 08.00 / 23.55 WIB) | T | datetime |
| | `plannedBufferDays` | Due − planned finish | T | float |
| | `hasEffortEstimate` | Effort estimate entered | T | boolean |
| | `effortEstimateHours` | Estimated hours | T | float (nullable) |
| | `estimateVsExpectedRatio` | Student estimate ÷ lecturer's `expectedEffortHours` | T | float (nullable) |
| | `workloadGini` | Unevenness of planned workload across days (0–1) | C | float |
| Perf × Process | `actualStartAt` / `actualFinishAt` | First session start; time marked done | T | datetime |
| | `startSlippageHours` | Actual − planned start in hours (+ = late); replaces `startSlippageDays` | T | float |
| | `timeOnTaskHours` | Sum of work sessions | T | float |
| | `prepTimeInComponentHours` | Time logged on ungraded (0%) activities of the same component before a graded activity's due time (e.g. Latihan Kuis 3 before Kuis 3); used by A3 and T7 | A | float |
| | `workSessionCount` | Distinct work sessions | T | integer |
| | `sessionSource` | timer / manual, per session | session | enum |
| | `lastMinuteSessionRatio` | Share of time in final 48 h before due (0–1) | T | float |
| | `alertResponded` | Started or re-planned after slip alert | E | boolean |
| | `alertResponseLatencyHours` | Hours from alert to response | E | float (nullable) |
| | `timelineViewCount` / `timelineZoomLevel` | Linimasa views and zoom used (2 Minggu / Minggu / Hari); replaces `timetableViewCount`, `timetableViewDurationSec` | C | integer / enum |
| | `timelineRefreshCount` | "Muat ulang" presses on the Linimasa | C | integer |
| Appraisal × Process | `effortEstimationError` | (actual − estimated hours) ÷ estimated hours | T | float |
| | `finishDeviationHours` | Actual finish − due time (+ = late) | T | float |
| | `estimationErrorTrend` | Slope of estimation error across activities (− = improving); teacher-side only | C | float |
| | `processReflectionSubmitted` | Done look-back submitted | T | boolean |
| | `processReflectionCategory` | underestimated difficulty / distraction / competing deadlines / other (asked only if gap > 50%) | T | enum |
| | `nextPlanAdjusted` | Next activity's estimate or buffer changed after large error | T | boolean |
| | `deadlineChangeNotified` | Student received a notice that an activity deadline changed; `planExceedsDeadline` if their saved plan now ends after it | T | boolean |
| Cross-cutting | `phaseCoverage` | Phases with ≥ 1 event | T | array of enum |
| | `phaseCoverageCount` | Phases covered (0–3) | T | integer |
| | `cycleCompleted` | Full plan → monitor → appraise sequence | T | boolean |
| | `cycleCompletionRate` | Share of activities with completed cycle (0–1) | C | float |
| | `eventSequence` | Ordered event types (for process mining; case = studentId + activityId); main path: planCreated → inProgress → timeLogged → done → doneRated; graded activities add gradeViewed → appraisalSaved | T | array of string |
| ★ Appraisal redesign | `appraisalMode` | How appraisal happened: action / chip / text / skipped | T / A | enum |
| | `appraisalEngaged` | Any appraisal action taken (not only written reflection) | T / A | boolean |
| | `strategySelected` | Strategy picked in the weekly review (e.g., start earlier, split task, ask for help sooner, shorter sessions) | W→C | enum |
| | `pastReflectionViewed` | Feed-forward note from earlier reflection was viewed | T | boolean |
| | `suggestedBufferAccepted` | Student accepted the suggested extra hours for the next similar activity | T | boolean |
| | `digestDecision` | Decision at end of weekly review (re-plan / adjust goal / no change) | W→C | enum |
| | `doneDifficultyRating` | One-tap rating at Done: easier / as expected / harder | T | enum |

Notes:
- "Waktu cadangan" means two things: on the student side, extra effort hours accepted for the next activity (`suggestedBufferAccepted`); `plannedBufferDays` is the gap between planned finish and due time. Teacher labels should name the latter "jarak rencana selesai ke tenggat".
- Variables depending on coded chips/reasons should be nullable until those features exist.
- Trace indicators are proxies (a view ≠ monitoring); pair with aptitude measures (e.g., OSLQ) for validation.
- Suggested evaluation for the redesign: compare chips + actions vs. free-text prompts on UT4/UT5 (ease of use, usefulness) and on `appraisalEngaged` rates.

---

## 4. Activity, grade component and grade metadata (lecturer-entered)

Entered in the teacher dashboard (Aktivitas kelas › Ubah / Tambah aktivitas; Kelola komponen nilai › Ubah komponen nilai; Input nilai) for the pilot; later partly synced from SCELE (gradebook categories map to grade components). Not student behaviour, but needed as reference values for the indicators above. Changes are logged on the teacher side (`activityChanged`, `gradeComponentChanged`, `gradesPublished`; see the teacher indicators doc §7).

**Activity**

| Variable | Description | Data type |
|---|---|---|
| `activityId` / `activityName` | Activity identifier and name | string |
| `activityType` | Tugas / Kuis / Ujian / Materi | enum |
| `gradeComponentId` | Component the activity belongs to (at most one); null = "tanpa komponen nilai" | string (nullable) |
| `isGraded` | "Dinilai" switch; false = 0%, the activity only prepares its component (replaces `linkedAssessmentId`) | boolean |
| `weightInComponent` | Share of the component (%); equal split under Sama rata, lecturer-set under Kustom, 0 when not graded | float |
| `effectiveWeight` | Derived share of the final grade = `componentWeight` × `weightInComponent` (e.g. Kuis 3: 30% × 33,3% = 10%) | float (derived) |
| `autoCreated` | Ujian activity created automatically for a component without activities (e.g. UTS) | boolean |
| `activityReleasedAt` | When the activity became visible to students | datetime |
| `dueAt` | Deadline or exam/quiz time (date + time, WIB) | datetime |
| `expectedEffortHours` | Lecturer's expected effort, compared with student estimates and time logged | float (nullable) |
| `suggestedSteps` | Steps students may reuse when splitting the activity | array of string |
| `sceleUrl` | Link to the SCELE item | string |
| `source` | manual / scele (SCELE-synced items become read-only) | enum |
| `activityStatus` | draft / published / archived (archived instead of deleted once students have interacted) | enum |
| `gradePublishedAt` | When grades for a graded activity were published (opens the grade reflection) | datetime (nullable) |

**Grade component**

| Variable | Description | Data type |
|---|---|---|
| `gradeComponentId` / `gradeComponentName` | Component identifier and name (follows the RPS, e.g. Kuis, Tugas, UTS, Proyek akhir) | string |
| `componentWeight` | Share of the final grade (%); all components total 100% | float |
| `weightMode` | equal (Sama rata, default) / custom (Kustom: % of component, graded activities must total 100%) | enum |
| `openReflectionOnPublish` | Open the grade reflection for students when grades in this component are published | boolean |

Rules: one activity belongs to at most one component; component weights must total 100% before saving; changes immediately affect students' average, score needed and simulation.
