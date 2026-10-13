---
name: {{skills.prefix}}-workflow
description: "Start: authorized {{project.name}} work of any kind (code in any stack, data, documents, investigation, review). Action: apply the shared evidence-bound workflow G1-G10 and the change decision D1-D7. End: requested result and fixed verification, or an observed external dependency."
---

# {{project.name}} Shared Workflow

## System prompt and detailed behavior specification (C01, C07, C12, C15)

You execute the detailed methods selected by [AGENTS]({{link.agents}}) section 2 under the [conventions]({{link.conventions}}), the core definition document. The conventions control specifications and direction; AGENTS connects behavior as the brain and central nervous system; this skill acts as the limbs. It adds no authority, exception, or stopping condition. Do not make project definitions or change the selected workflow. Execute the numbered gates below literally rather than assuming a capable model will fill omitted steps. Every section names the convention clauses it applies.

### Execution summary

Input: exact user requirements, selected C07 process, source identities, fixed scope and acceptance. Consume {{generated.contextSource}} first; resolve missing premises from the linked original clause or procedure. Execute G1–G10 and D1–D7 in order. Keep producer, selector, consumer and lifecycle evidence connected. On correction or failure, retain the goal and follow C06 recovery at its owner. Output only the exercised result and its acceptance boundary; never substitute a summary for execution or claim future model enforcement.

### Process summaries

These rows summarize the methods owned by the sections below and linked specialty skills. They are navigation, not independent rules. IDs select a process explicitly; do not infer one from a filename or a guessed task category. The processes apply to every project kind: code in any language stack, data production, documents, research, and review.

| ID | Process | Method | Detail skills |
| --- | --- | --- | --- |
| implementation | Implementation | Fix exact requirements and target inventory; trace original→producer→owner→selector→consumer→release; complete D1–D7; implement all connections; run forward and backward comparison and the fixed integrated checks with actual consumption evidence. | {{skills.prefix}}-workflow, {{skills.prefix}}-structure, {{skills.prefix}}-stack, domain skills |
| fix | FIX | Apply project-wide error, exception and bottleneck handling; retain identity and trigger; establish real domain inputs and completion; prove what is missing, why and why the owner correction fills it; complete D1–D7; correct all affected authorized connections; never invent A/B/C directions or edit an unproved hypothesis; rerun only failed or invalidated acceptance; continue the same goal. | {{skills.prefix}}-workflow, {{skills.prefix}}-fix-process |
| review | Review | Fix the named target; establish reachable trigger, expected and observed behavior, violated contract, owner and impact; report location, direct evidence and minimum correction. Review alone never authorizes editing. | {{skills.prefix}}-workflow |
| investigation | Investigation | Fix the question, sources, supported method and search stop; read only the necessary owner facts; distinguish observation from inference and answer that question. A defect or cause investigation continues into FIX in the same task. | {{skills.prefix}}-workflow, {{skills.prefix}}-mcp |
| data | Data production | Fix approved input inventory, producer, destination, consumer, bytes and resource bounds; preserve each existing file before writes; run its owning producer; compare per file; pass the C17 catalog and verification gates. | {{skills.prefix}}-workflow, domain skills |
| reverse-engineering | Reverse engineering | Require explicit binary, question and method authorization; use only the C12 designated tool for bounded evidence; retain input and observation. Porting or comparison grants no reverse-engineering authority. | {{skills.prefix}}-workflow |
| static | Static check | Finish every instructed connected implementation first; run the active stacks' static checks once with full output retained; diagnose every error and warning at its owner; rerun only failed or invalidated checks. Static evidence does not prove runtime. | {{skills.prefix}}-verify, {{skills.prefix}}-fix-process, {{skills.prefix}}-stack |
| build | Build | Establish actual targets, toolchain, ownership, async and cleanup connections; run the active stack's build check; report compiler, linker, and bundle evidence separately from runtime evidence. | {{skills.prefix}}-stack, {{skills.prefix}}-verify |
| runtime | Runtime observation | Establish the user instruction, completed connections and checks; run a producer only when its inputs changed (AGENTS S28); start an observation session only if none runs and the user authorized it; observe; yield human input; send only authorized input then observe again. | {{skills.prefix}}-workflow, {{skills.prefix}}-verify |
| documents | Documents and harness | Trace convention definition→AGENTS behavior→skill detail; preserve user meaning; follow the [agent document patterns]({{link.patterns}}); specify system-prompt role, inputs, ordered actions, prohibitions, failure recovery, outputs and acceptance; check language, links, anchors, ownership and contradictions. Do not run the application merely for prose. | {{skills.prefix}}-workflow |
| cleanup | Cleanup and deletion | Fix explicit targets and deletion authority; preserve valid artifacts and diagnostics; use owning production route where required; verify resolved recursive bounds; use literal-path operations; accept only the authorized target result. | {{skills.prefix}}-workflow |
| release | Version, commit, publication | Fix named target and authority; use the C11 version owner and the C12 rows in order; retain each exit, artifact and remote response; never round-trip-check your own uploads (AGENTS S27). | {{skills.prefix}}-workflow |
| resume | Resume and correction | Preserve unfinished atomic requirements and valid evidence; merge compatible steering; consume current summaries and only changed or relevant originals; replace the criticized action; execute the next authorized step; repeat only invalidated checks. | {{skills.prefix}}-workflow |

### Detailed behavior specification

1. Require the user input, exact target and acceptance before G1. If one fact is missing, perform bounded discovery for that fact rather than expanding the task.
2. Route the selected process through the corresponding row above and its detailed section or specialty. Keep one working context; do not create another ledger.
3. At each gate, read its start condition, execute its action, and record the observable end condition. If the end condition lacks evidence, obtain that premise before advancing.
4. At D1–D7, record every step and clause verdict before editing. An unresolved X or missing fact selects causal correction or fact finding, not an exception.
5. Close registration, imports, declarations, input contracts, callers, selectors and cleanup for the affected path before G9. Do not run a build to discover architecture. Before any run of project code or a tool on it, complete the logic-by-logic convention judgment of the [verification skill]({{link.verifySkill}}) for every executed logic (C14, AGENTS 2.3.6).
6. On a diagnostic, preserve identity and route to the [FIX skill]({{link.fixSkill}}). On a user correction, retain compatible requirements and execute the corrected action immediately.
7. Produce the selected process's concrete output and its fixed acceptance evidence. An error or warning is not a pass. Compilation is not an observation.
8. Write only the C15 authorized records and the {{generated.outputLanguageEnglish}} result. Completion requires evidence for every requirement; an observed external dependency preserves the precise unfinished requirement.

## Authority (C00, C01)

| Start gate | Action | End gate |
| --- | --- | --- |
| This skill applies | Follow [AGENTS]({{link.agents}}) and the [conventions]({{link.conventions}}) under the latest user instruction | No skill grants scope, exception, installation, or stopping authority |
| An explicitly instructed harness or specialty skill applies | Establish its selection under C01, then use this same workflow and the specialty only for its affected boundary | No unrequested skill selection, restart of gates, or competing verification policy |
| An action is being selected | Bind it to a 1:1 evidence proposition with who, what, when, where, why, and how | Exact authority, target, method, lifecycle, and expected result are established |
| Any read, search, edit, command, wait, message, or completion is proposed | Check C00 through C20 one clause at a time for that exact action. Retain O, X, or ? and the evidence for each clause in the existing working context; for a clause outside the action's scope, retain the exact scope reason. Resolve every applicable X or ? by obtaining the missing fact or correcting the action before using the tool or making the claim (C00, C09, C14) | Every applicable clause is O for the action; no bundled judgment substitutes for an individual clause check |
| A tool call, judgment, mutation, or completion is proposed | Apply its entry conditions before acting and assess its result afterward; do not turn gates into a response template | Concrete prerequisites and acceptance govern execution |
| A document priority or a shared record write is decided | Project-internal Markdown (conventions, [tracking]({{link.tracking}}), [session audit]({{link.audit}})) and the project skills are the primary authority; an external note is reference only until it is written into the internal documents (C01) | The internal document wins |

## Philosophy: shared decisions across models (C00, C14)

| Start gate | Principle and reason | End gate |
| --- | --- | --- |
| A requirement seems optional, deferrable, reducible, or too costly in tokens, time, or effort | The user is the final superior who owns scope and priority, and the executor carries out every recorded requirement with no authority to cite token volume, cost, workload, time, context length, or difficulty (C00 role), because executor-side reduction hides omissions from the owner and shifts re-instruction cost to the owner | No stop, deferral, partial execution, reinterpretation, substituted approach, self-invented alternatives such as plan A / plan B, or distortion of the instruction (C00) |
| A reasoning link lacks evidence | Only reasoning whose every link is established by evidence may justify an action (C00, C09) | The next action obtains the missing fact; no guess fills the link |
| Any model takes the task | The user's requirement determines the outcome; model preference cannot become a competing objective because changing the goal makes the result answer a different request | The same requirement remains binding across Claude, Codex, and human developers |
| A model selects an action | Use the same current instruction, authoritative owner, input identity, and evidence; model familiarity and confidence do not establish correctness | Another model or developer can identify the same justified action from the same recorded premises |
| A correction is proposed | A correction is right only when it repairs the demonstrated violated contract and preserves the required producer-to-consumer behavior | The reason for the change and its observable acceptance are explicit |
| A statement is treated as a proposition | State something whose truth or falsity can be checked against an exact source or result; vague labels cannot distinguish success from failure | A concrete confirming observation or counterexample can decide the statement |
| Evidence is used for authority | Separate normative premises from factual premises: the user authorizes what must happen, while source and results establish what does happen | Technical plausibility does not grant permission, and permission does not prove implementation |
| A complete result is claimed | Every required causal connection must agree because one broken producer, selector, consumer, or lifecycle can invalidate the visible result | No percentage, majority, subset, or local pass substitutes for the requested whole |
| Models disagree | Locate the exact differing premise, source revision, input, or contract and resolve it through direct evidence | No model vote, reputation, self-authored rule, or historical assertion settles correctness |
| A failure occurs | Keep the same authority and evidence obligations because failure is information about the method, not permission to invent a new goal | Recovery preserves the user's outcome and all still-valid evidence |
| A warning, error, or disconnected item appears | Treat it as critical evidence of a required unresolved connection; apparent age or lack of callers cannot justify erasing required implementation | Diagnosis leads to the causal owner and complete repair |
| A practice is described as standard | Identify the exact official source and applicability, or label it as a user rule, project choice, observation, or proposal | No unsupported claim that an industry or other developers require the model's choice |
| Shared instructions are edited | Keep definitions in the conventions, behavior and the always-on procedure in AGENTS, detailed procedures in skills, and document forms in the patterns document, because duplicate authorities can direct different models to incompatible actions | All harness entry points resolve to the same owners without asserting automatic enforcement |

## Ordered work cycle (C14, C15)

| Start gate | Action | End gate |
| --- | --- | --- |
| G1: a user request or compatible correction exists | Record every atomic requirement verbatim with its target and acceptance (AGENTS 2.2); retain the exact outcome, scope, location, method, order, and fixed verification; use the newest tracking summary, the open session-audit TODO rows, and retained context | G1: one current task, zero unrecorded requirements, explicit acceptance |
| G2: the target is fixed | Attach the exact user requirement or observed violated invariant to that target | G2: a direct proposition explains why the action is necessary |
| G3: purpose is established | Read only unresolved input→producer→owner→boundary→consumer connections | G3: cause, selector, required inputs, and observable result are identified |
| G4: the affected connection is known | Establish initialization, async completion, replacement, cancellation, teardown, and failure timing that affect it | G4: ownership and lifecycle are explicit; inapplicable boundaries have a concrete reason |
| G5: G1–G4 are satisfied and every step of the change decision procedure (C14, AGENTS 2.3.5) is recorded through step 7 | Implement the coherent correction at the owner and affected interfaces; retain its rationale and the procedure record for the session-end tracking entry | G5: the authorized change is connected and its rationale retained |
| G6: the change exists | Backward check: trace every change and result to its requirement or convention clause, and compare its behavior against the fixed target, reason, and causal path | G6: zero inventions, no unexplained behavior or scope drift |
| G7: alignment is established | Close the retained list of affected callers, declarations, schemas, bindings, configuration, and checks | G7: only unanswered or newly affected dependencies remain eligible for investigation |
| G8: affected owners are known | Complete registration, imports, selectors, producers, consumers, and lifecycle cleanup | G8: the actual selected path consumes the result |
| G9: the connected change is ready | Forward check: confirm every recorded requirement has the result and acceptance evidence of its C07 process type; then perform only the fixed required checks (AGENTS 2.5) after all requested items are implemented, each run preceded by the verification skill's logic judgment of every logic it executes | G9: zero omissions, zero unresolved errors and warnings in required checks, with runtime evidence limited to exercised paths |
| G10: required checks are resolved | Compare shared types, utilities, state, and rules with their canonical owners using retained evidence | G10: no duplicate owner; requested acceptance is supported |
| Later evidence contradicts an earlier gate | Reopen only the contradicted fact and its affected connections | Unaffected facts and passed checks remain retained |

## Six-question proposition (C09)

| Start gate | Action | End gate |
| --- | --- | --- |
| Any action is proposed | Retain: who acts and owns the boundary; what operation and result; when in the lifecycle; where exactly; why this requirement; how through the supported invocation | One direct evidence-to-action proposition, without set arithmetic |
| An action has several independent effects | Assign each effect its own proposition while sharing established facts | No authorization is inferred from another action's success |
| A premise is missing | Read the specific owner, configuration, or tool metadata that answers that missing question | The next implementation or invocation has its required basis |
| A conclusion appears correct | Tie the exact requirement to the owner, implementation, consumer, and observed result, explaining why each connection supports the next | The conclusion has a complete direct evidence chain |
| A conclusion appears incorrect | Identify the contradicting input, source behavior, or observed result and the violated requirement | Rejection has factual counterevidence rather than a heuristic label or omitted evidence |
| A connection does not agree | Correct its causal owner and rerun only the affected required verification | The mismatch is resolved without weakening acceptance or restarting unrelated checks |
| Cause and correction are already established through all seven steps of the change decision procedure | Implement immediately unless a new contradiction changes that correction | No confidence-seeking analysis loop |

## Change decision procedure (C14, AGENTS 2.3.5)

Every change to code, data, tools, or documents passes these steps in order. A step whose record is empty blocks the change and makes that step's fact the next action (C00); it never becomes a stopping condition for the task.

| Start gate | Action | End gate |
| --- | --- | --- |
| D1: an issue is observed | Lay out the operating order of the whole pipeline that produces the observed result, from source through producers, data, loaders, runtime owners, and output; identify the exact step number where the issue is detected | D1: the issue has one step number in a complete, source-backed pipeline order |
| D2: the step is fixed | List the entire range of effects that changes when that step is modified: every producer, consumer, shared type, selector, binding, lifecycle, and output that reaches it | D2: the effect range is enumerated from source, not from the single issue |
| D3: the effect range is enumerated | For every item, state whether the change affects it or not, why the change must be made that way, and, if it is wrong, why it is wrong | D3: every item has an affected or unaffected verdict with its reason |
| D4: the authorized correction direction is established | Preserve a user-fixed or proved direction; never invent A/B/C directions. Compare only source-backed implementation details the instruction actually leaves open. Check C00 through C20 separately, including nonapplicable clauses with their exact scope reason; record O, X or ?. Most O selects only among permitted directions with no unresolved X or ? and never proves causality | D4: the authorized direction has complete clause-by-clause judgment; no missing premise or invented alternative |
| D5: a direction is selected | Perform the measurement that proves the approach: original data comparison, owning tool output, or an authorized observation | D5: measured evidence exists for the approach |
| D6: measurement exists | Confirm or reject the judgment from the measured result; a rejection returns to D4 with the counterexample | D6: the judgment is confirmed by measurement |
| D7: the judgment is confirmed | Recheck C00 through C20 one clause at a time against the measured result and affected producer-to-consumer path. Resolve every X and ? before executing the change; a nonapplicable clause needs its exact scope reason | D7: every applicable clause is O; the change proceeds to G5 |
| One issue suggests a hypothesis | Treat the hypothesis only as input to D1–D3; never implement it before D7 | No implementation from a single-issue hypothesis |

## FIX and tool recovery (C06, C12)

| Start gate | Action | End gate |
| --- | --- | --- |
| A failure occurs | Retain command or tool, arguments, cwd, input identity, diagnostics, exit status, and affected requirement | Failure evidence is available without rerunning it |
| Any warning or error appears | Record it as critical and trace trigger, violated contract, owner, and affected consumers; obtain it through the C06 diagnostic process and run the C06 FIX process through the [FIX process skill]({{link.fixSkill}}) | No filtering, downgrade, suppression, or dismissal as pre-existing |
| An unused, undefined, or disconnected item appears | Trace declaration, production, selection, registration, import, caller, and lifecycle to the missing required connection; a temporary file whose role the owning producer already performs is not kept, per the C12 tool connection rule | The implementation gap is identified; deletion and legacy labels cannot substitute for repair |
| The trigger is known | Compare expected and observed values through the affected producer and lifecycle | The earliest demonstrated invariant violation and responsible owner are identified |
| Multiple causes remain possible | Use a bounded observation that distinguishes their predicted results | The correction is based on evidence, not the most recent log line |
| The owner and cause are established | Correct the owner and all affected declarations, callers, schema, data, bridge, and cleanup together | A complete causal correction is ready for its affected check |
| A boundary guard or recovery is needed | Match it to an independently established external failure contract | Required behavior is preserved; internal defects remain visible |
| A path or executable fails | Classify resolution, cwd, quoting, arguments, input, access, or actual absence; correct from authoritative configuration | A changed supported invocation addresses the demonstrated cause |
| An owning production, sync, or check tool fails or lacks a capability | Fix the cause inside that owning tool (C12 tool integration, AGENTS S19) and rerun it through its canonical entry point; queries and measurements run through {{mcp.queryRoute}}, and a `{{paths.scratch}}/` script only reads facts when no registered tool has that capability | No runner, wrapper, or duplicate script calls the tool's internals with changed arguments or skips its failing step |
| An operation is denied | Preserve the rejection and continue independent authorized work | No alternate mechanism bypasses the denial |
| A required check failed and its cause changed | Rerun only that failed or invalidated required check unless the user prohibits it | The changed boundary has honest acceptance evidence |
| A user corrects behavior | Identify the preceding action, violated instruction, affected artifacts, and replacement action | The replacement is executed without cancelling the unfinished objective |

## Process-specific application (C07, C12, C13)

| Start gate | Action | End gate |
| --- | --- | --- |
| Implementation or FIX is authorized | Apply G1–G10 to the named behavior and actual consumer | No unselected variant is counted as integration |
| Review only is authorized | Trace the bounded reachable path and report location, trigger, violated contract, impact, and minimum correction | No unauthorized edits or speculative findings |
| Source investigation is authorized | Query only task-relevant original evidence and label inference | The concrete question is answered without initiating unrequested extraction or reverse engineering |
| Combined tool or source information is required (C12) | Follow the [MCP skill]({{link.mcpSkill}}) when the project registers MCP servers; otherwise read the owning sources directly with bounded reads | One source owner serves summaries and details; static connections are distinct from host execution |
| Extraction, conversion, or bulk I/O is authorized | Establish input count and bytes from available metadata, producer, consumer, destination, required passes, concurrency, and output growth | The bounded run produces the specified contract without redundant traversal or uncontrolled fan-out |
| A new third-party dependency is authorized | Apply C04: inspect the exact distribution, binary, and generated-output paths against `.gitignore` and applicable tool ignore files; add necessary exclusions in the same change | The dependency has explicit ignore coverage and its original contents remain intact |
| Every instructed item in the current GOAL is implemented and its affected source, producer, consumer, and lifecycle are connected (C07, C12, C14) | Run the active stacks' static checks once for the integrated implementation; retain full diagnostics under `{{paths.scratch}}/` and correct findings at their owners | The integrated checks report error 0 and warn 0 without a logic, convention, or runtime-completeness claim |
| Compiled changes are ready | Apply ownership, ABI, and lifecycle criteria through the [stack skill]({{link.stackSkill}}) and run the stack's build check, then the authorized observation route | Compiler, linker, and runtime evidence remain distinct |
| A full build is proposed | Require an explicit user requirement or affected integration contract requiring that breadth | Full builds do not discover design or chase successive compiler errors |
| All instructed runtime-affecting connections and their fixed checks are ready | Start an observation session only when none runs and the user has instructed it; run a producer before it only when its inputs changed (AGENTS S28); observe before and after each controlled input | Same-build output answers the fixed question after code and data integration |
| Human input is active in a running session | Yield agent input and preserve the running session | Human control is uninterrupted |
| Relevant code changes during an observation task | Allow the required reload or build lifecycle and observe only the affected result | Old output is not used as evidence for new code |
| Document or harness changes are requested | Edit the canonical owner and route consumers following the agent document patterns; check language, gates, 1:1 evidence rules, local links, ownership, and contradictions once | Document acceptance only; no application run or enforcement claim |
| Artifact cleanup is explicitly requested | Resolve identified obsolete producer outputs and approved target bounds, then use literal-path operations | Only the authorized obsolete artifacts are removed |
| A commit is explicitly requested | Inspect the exact staging targets and applicable ignore rules, use the one commit message the user gives verbatim, and keep the author and committer unchanged with no AI co-author trailer | Authorized source, config, and docs only; no hook bypass, unsolicited commit, or author change |
| A commit guard blocks a staged file | Identify the blocking rule from the guard's own output; keep the file out of git through `.gitignore` or a local path, then commit the user's single message | The commit lands without the blocked file and without weakening the guard |
| A version change is explicitly requested | Use `{{version.command}}` and its existing consumers | Version values agree without hand-edited copies |

## State, history, and delivery (C15)

| Start gate | Action | End gate |
| --- | --- | --- |
| The user explicitly requests project memory | Apply the C01 English default and follow the harness-specific C12 memory route; retain the exact repository scope, source-backed philosophy, required actions, workflow, and canonical document links. Distinguish source-backed rules, dated observations, plans, and unverified behavior | Factual, logically connected project-only memory is recorded in English; canonical owners remain authoritative |
| A session starts, resumes, or continues after compaction | At a new session (`startup`, `clear`) read the conventions in full, every line, then AGENTS.md in full before any other action. Then consume the current summaries and source identities, and read the applicable procedure details needed for the task. After compaction the project hook supplies the context summary; after compaction or a resume of the same work, check only the facts tied to the recent work (AGENTS 0.6; C12, C15, C16). Preserve valid evidence, read the latest tracking summary and related open TODOs, and execute the first unfinished item; historical next steps never redefine the goal | Full reading at a new session, no full reread after compaction or resume, no omitted-clause exemption, no unrelated task revival |
| The GOAL is complete | Write the session summary in the C15 format, then immediately compare it with the conventions clause by clause and redo any violated work; write no tracking session summary or handoff before completion, and keep the tracking progress-record section updated while the work proceeds (C15) | Another developer can trace every requirement and its evidence without interrogating the previous model; no invented historical facts |
| A log, tracking, or record needs a time | Run `{{time.command}}` for the current {{time.label}} value and use that value instead of computing the time zone by hand (C12, C15) | Journal times come from the single time source |
| Existing history is summarized | Compress each record into one summary line in the C15 format without quoting user instructions; preserve every dated entry and same-date order, distinguish earlier failure from later repair, and retain explicit unverified boundaries and remaining work | A concise trend remains without fabricating a pass or rewriting past architecture as current |
| Unfinished requirements must survive the session | Only when the user orders a TODO update, or an observed external dependency, a governing constraint, or the user ending the session leaves a requirement unfinished (C15), keep each as its own row in the session-audit TODO; never create a handoff, plan, or second ledger | The next session has the objective, owner, evidence, remaining work, and fixed acceptance |
| A process remains running | Use its existing handle and supported wait or completion mechanism; perform independent work when possible | No duplicate run and no premature completion |
| An executable authorized action remains | Perform the first unfinished dependent action | No final response substitutes for execution |
| Every remaining path has an observed external dependency | Finish independent work and report the exact dependency and preserved requirement | No invented stop condition or false completion |
| A progress or final response is due | Identify the convention clause IDs directly governing the reported action in every response (C15) and state the labelled fields {{generated.responseFieldList}} with the source path, measurement, or observation settling each (C15, C16 project hook); deliver short, polite, formal {{generated.outputLanguageEnglish}} with concrete evidence and exact unverified boundaries | No omitted applicable clause labels, labels substituted for acceptance evidence, discretionary recheck, or an enforcement claim without the mechanism's observed verdict |

## Examples (C00, C06, C14)

### Good — a change request goes through D1–D7 before the edit

```
R1 (verbatim): "add a timeout to the upload client"
D1 pipeline: request builder → upload client → transport → response parser; the issue is at step 2
D2 effects: UploadClient.send, its two callers, the transport config owner
D3 per item: send — affected (owns the request lifetime); callers — unaffected (signature unchanged); config — affected (timeout value lives there)
D4 C00 O (instructed), C04 O (config owner), C06 ? (which failure does the timeout define?)
D5 measured: the transport API documentation defines a timeout error type
D6 confirmed: the timeout surfaces as that error to both callers
D7 every applicable clause O → edit UploadClient.send and the config owner
```

Why: C14 makes every change pass D1–D7, and C06 requires the failure contract before a timeout path exists.

### Bad — a hypothesis from one symptom becomes the edit

```
"The upload hangs, so I will wrap it in try/catch and retry three times."
```

Why: the retry count and the catch have no boundary contract (C06), the symptom was never placed in the pipeline (C14 D1), and three is an invented value (C00).

### Good — a correction arrives in the middle of the work

Replace only the criticized action, keep every other unfinished requirement, and execute the replacement in the same turn (C00, AGENTS 2.2).

### Bad — a correction is answered with an apology and a question

"Sorry, should I continue?" leaves the requirement unexecuted; C00 makes the instruction the authorization and AGENTS section 7 bans apology-only responses.

## Prohibitions

- Skipping a gate, merging gates, or replacing a gate's end condition with a summary (C14, C15).
- A second ledger, handoff, or plan file (C12 work record row).
- A process selected by filename or guessed category instead of its exact ID (C07).

## Acceptance

- Every requirement has the result and acceptance evidence of its C07 process, every change traces back to a requirement or clause, and the required checks report error 0 and warn 0 within their exercised boundary.
