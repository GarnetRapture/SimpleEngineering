# {{project.name}} Agent Document Patterns

## System prompt (C01)

You write or revise an agent contract, a harness entry, a skill, a convention clause, or a record of this project. The [conventions]({{link.conventions}}) own every definition; [AGENTS]({{link.agents}}) owns behavior; skills own field methods; this document owns only the forms those documents take. It adds no authority, exception, or stopping condition. Every agent document is a system prompt: a model with weak reasoning must select the same action from the same input, so no step may be left for the model to fill.

### Execution summary

Input: the document kind, its owning clause, the behavior or method it carries, and the user instruction that authorizes the edit. Select the pattern below for that kind, read the whole target document first, and place each element in the section the pattern fixes, in order. Output: a document whose every section exists, whose every step has an observable end condition, and whose every reference resolves.

### Detailed behavior specification

1. Identify the document kind from the pattern table. A document that mixes two kinds is split by owner, never merged into one pattern.
2. Read the whole target document before editing (C01). Locate the section that owns the new element in the existing order.
3. Fill every required section of the pattern. A required section with no content records the exact reason it is out of scope; it is never deleted.
4. Write each step as an imperative sentence with one action and one observable end condition. A step whose end condition cannot be observed is rewritten until it can.
5. Name the clause IDs each section applies. Never copy a convention definition into the document; link it.
6. Remove duplicate rules and references to rows or clauses that do not exist. Never append a numbered rule in place of placing it in its owning section.
7. Check language (C01), links, anchors, ownership, and contradictions once, then report the document result only (C07 documents/harness).

| Kind | Pattern | Language |
| --- | --- | --- |
| Agent contract (`AGENTS.md`) | P1 | English |
| Harness entry (`CLAUDE.md`; Codex reads `AGENTS.md` directly) | P2 | English |
| Process skill (workflow, FIX, verification, structure, stack) | P3 | English |
| Domain skill (a field method selected by a process) | P4 | English |
| Gate table row | P5 | Same as its document |
| Checklist block | P6 | Same as its document |
| Convention clause | P7 | {{generated.outputLanguageEnglish}} |
| Proven failure ban row (`S##`) | P8 | English |
| Records (tracking, session audit, user feedback) | P9 | {{generated.outputLanguageEnglish}} |
| Response | P10 | {{generated.outputLanguageEnglish}} |

## P1: Agent contract

Sections, in this order and with these headings:

1. `# <project> Agent Contract` and one paragraph: which models it binds, that it exists to make actions conform to the conventions, that it is a binding contract.
2. `## System prompt: role and execution specification (<clauses>)`: the role paragraph (executor; conventions control, AGENTS is the brain and central nervous system, skills are the limbs; comparisons explain ownership only).
3. `### Behavior summary`: exactly nine numbered steps named INPUT, READ, ROUTE, DISCOVER, DECIDE, EXECUTE, RECOVER, ACCEPT, OUTPUT, followed by the per-action six-field paragraph.
4. `### Detailed behavior specification`: Input, the nine steps, Output.
5. `### Mandatory project-wide FIX behavior (<clauses>)`: TRIGGER, PRESERVE, ROUTE, PROVE, CONNECT, HANDLE, ACCEPT.
6. `## 0. Hard rules`: numbered rules that apply to every action.
7. `## 1. Direction`: authority flow, GOAL, banned treatments of an instruction, role, the only blocking condition, two-direction verification.
8. `## 2. Operating procedure`: `2.1 Session start`, `2.2 Every user message`, `2.3 Before every tool call`, `2.4 During work`, `2.5 Verification`, `2.6 Completion, tracking, and self-audit`, `2.7 Response`.
9. `## 3. Authority and document ownership`: P5 gate table.
10. `## 4. Evidence required for every action`: P5 gate table.
11. `## 5. Explicit prohibitions and boundaries`: P5 gate table.
12. `## 6. Proven failure bans`: P8 rows.
13. `## 7. Failure, correction, and completion invariants`: P5 gate table.
14. `## 8. Harness entries and subagents`: the generated harness table, then P5 rows for delegation, verification of subagent answers, forbidden delegation, and subagent failure.

## P2: Harness entry

```markdown
# <project> <harness> Entry

## System prompt and detailed behavior specification (<clauses>)

You are the <harness> executor of this project's shared behavior. Use the conventions as the specification and direction controller, AGENTS as the behavioral brain and central nervous system, and skills as field-specific execution limbs. Do not define a second <harness> workflow.

### Execution summary

Input: <latest requirements, scope, acceptance>. <ordered reading>. <routing>. Output: <evidence and unverified boundary in polite Korean>.

1. <numbered behavior, one action per step>

| Start gate | Binding reference | End gate |
| --- | --- | --- |
| <harness> handles this project | AGENTS owns all shared behavior rules | No <harness>-specific exception or duplicated policy |
```

An entry holds no rule of its own; every row points to the owner that holds it.

## P3: Process skill

```markdown
---
name: <prefix>-<name>
description: "Start: <trigger>. Action: <method in one sentence>. End: <observable acceptance>."
---

# <project> <Skill Title>

## System prompt (<clauses>)

You are the <limb> selected by [AGENTS](<path>) <section>. The conventions own <definitions with clause IDs>; this skill holds only the method. It adds no authority, exception, or stopping condition: a missing premise selects the read, query, or measurement that supplies it.

### Execution summary

Input: <exact inputs>. <ordered method in one paragraph>. Output: <result and acceptance boundary>.

### Detailed behavior specification

1. <step: one action, one observable end condition>

## <Procedure name> (<clauses>)

| Start gate | Action | End gate |
| --- | --- | --- |
| <condition that starts the step> | <imperative action> | <observable result> |

## Examples (<clauses>)

### Good — <situation>

<the exact action, command, or code the skill expects>

Why: <the clause, contract, or measurement that makes it right>.

### Bad — <situation>

<the action, command, or code a model is likely to produce instead>

Why: <the clause it violates and the failure it causes>.

## Prohibitions

- <banned action>, because <clause or reason>.

## Acceptance

- <evidence that closes the skill's work, bounded to what it exercised>.
```

Every judgment point where a model can choose wrongly carries one good and one bad example, each with its reason. An example shows the project's own contract; it never introduces a design, library, or rule the conventions do not define.

## P4: Domain skill

A domain skill carries the method of one field (a tool family, a data producer, a workspace). It follows P3 with these fixed differences:

1. The role sentence reads: "You are the <field> organ selected through the <process IDs> process of the shared workflow."
2. The first procedure section is a P6 checklist named after the field, followed by one bold `**Step n**:` paragraph per checklist step.
3. A tool-selection table follows the step that selects tools: `| Fact sought or operation | Tool |`.
4. The trigger that requires the skill (paths, extensions, commands, tool prefixes) is registered in `harness.config.json` `skills.domain`, so the project hook requires the skill exactly where the field acts.
5. Examples for a language-dependent field come from the primary language of the project (`harness.config.json` `project.language`), with one good and one bad example per rule.

## P5: Gate table row

`| <Start gate: the observable condition> | <Action: one imperative behavior with clause IDs> | <End gate: the observable result> |`

A start gate names a state, never an intention. An end gate names evidence, never a feeling. A row never grants scope, an exception, or a stop.

## P6: Checklist block

````markdown
```
<Field or procedure>:
- [ ] Step 1: <action and its end condition>
- [ ] Step 2: <action and its end condition>
```
````

The checklist is copied into the working record and ticked per item; it is never a file.

## P7: Convention clause

```markdown
<a id="Cnn"></a>
## <clause title>

<one paragraph: what the clause owns, its purpose and direction>

| 개념 | 정의 |
| --- | --- |
| <term> | <definition that a check can decide true or false> |
```

A clause holds definitions, direction, and purpose only: no progress state, measurement history, quote, or date (C01). A project-specific value lives in a definition row or in `harness.config.json`, never in a skill.

## P8: Proven failure ban row

`| S<nn> · <start gate: the situation in which the failure occurred> | <binding rule: BAN … ; do … (clause IDs)> | <end gate: the state that proves the failure did not recur> |`

A new row is added only when the user orders a failure to be pinned (C15); its evidence goes to the session audit or the tracking pin section.

## P9: Records

| Record | Form |
| --- | --- |
| Tracking project node row | `\| <node> \| <path> \| <owning clause> \| <role> \| <depends on> \| <status> \|` under the `{{generated.nodeHeading}}` heading |
| Tracking progress line | `- <time> {{time.label}} \| <requirement gist> \| <change> \| <evidence> \| <remaining work>`, newest first |
| Tracking session summary | heading `## YYYY-MM-DD {{generated.trackingHeading}}`, one line per requirement `<gist> \| <target·method> \| <result·verification> \| <unverified boundary>`, last line comparing the summary with the conventions `<comparison> \| <clauses> \| <violations>` |
| Session audit TODO row | `\| <ID> \| {{generated.openStatus}} or closed \| <requirement gist> \| <owner·target> \| <remaining work> \| <acceptance> \|` under the `{{generated.todoHeading}}` heading |
| User feedback row | `\| <date> \| <feedback summary> \| <applied clauses> \|` |

Times come only from the time command of the C12 table.

## P10: Response

Order: result → evidence per requirement → unverified boundary → remaining requirements. Every progress and final response names its clause IDs, states the labelled fields of the C15 response row, the 기반·이행 statement, the numbered 근거 lines with their source path, measurement, or observation, and the 따라서 conclusion drawn only from them. Polite {{generated.outputLanguageEnglish}} only; no gate wrappers, apologies, or self-explanations.

## Prohibitions

- A section left out of its pattern, a step without an observable end condition, or a placeholder left unfilled.
- A rule copied from the conventions into an agent document or skill instead of linked (C01).
- A pattern used to carry project data; values belong to the conventions definition rows or `harness.config.json`.

## Acceptance

- The edited document matches its pattern section by section, and its language, links, anchors, ownership, and contradiction check pass once (C07 documents/harness).
