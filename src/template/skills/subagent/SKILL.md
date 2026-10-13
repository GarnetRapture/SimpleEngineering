---
name: {{skills.prefix}}-subagent
description: "Start: a bounded research, review, or drafting task could be delegated, or a subagent answer is about to be used. Action: write the task to a scratch file, run the registered subagent through the launcher, and verify every claim of the answer against its sources before use. End: only verified claims enter the work, and the subagent id and log path are recorded."
---

# {{project.name}} Subagent Delegation

## System prompt (C09, C12, C14)

You are the delegation organ selected through the investigation and review processes of the [shared workflow]({{link.workflowSkill}}). The [conventions]({{link.conventions}}) C12 subagent row and subagent table own which subagents exist and how they run; AGENTS section 8 owns the delegation behavior. This skill holds only the method. A subagent answers; the main agent decides, edits, and verifies.

### Execution summary

Input: a bounded question with its target files and expected answer form, and a registered subagent id. Write the task file under `{{paths.scratch}}/`, run `node {{paths.subagentLauncher}} <id> <task file>`, read the printed answer and the log path, verify each claim against the cited file, line, or measurement, and use only the verified claims. Output: the verified claims with their sources, and the subagent id and log path in the progress record.

### Detailed behavior specification

1. Delegate only a task whose answer the main agent can verify: a search, a summary of named files, a review of a named diff, a draft the main agent will rewrite.
2. Write a new task file under `{{paths.scratch}}/` with the question, the exact files or symbols to read, the expected answer form (claims with file and line), and the instruction to change nothing.
3. Run `node {{paths.subagentLauncher}} <id> <task file>`. The launcher runs the subagent read-only (Claude runtime without edit tools, Codex runtime in a read-only sandbox) and writes the full log under `{{paths.scratch}}/subagent/<id>/`.
4. Verify each claim of the answer against its cited source with a direct read or measurement. Discard a claim that has no source or that the source contradicts, and keep the counterevidence.
5. Record the subagent id, the log path, and which claims were verified in the progress record (C15).
6. On a failure, keep the exit and log, fix the cause at its owner (missing CLI, empty key variable, wrong model id in `harness.config.json` `subagents`), and run once more.

## Registered subagents (C12)

| id | Runtime | Endpoint | Model | API key |
| --- | --- | --- | --- | --- |
{{generated.subagentTable}}

## Examples (C09, C12, C14)

### Good — a bounded, verifiable task

```
{{paths.scratch}}/subagent-task-order-total.md:
  Question: which functions compute the order total, and where is rounding applied?
  Read: src/order/ only.
  Answer form: one claim per line with file:line.
  Change nothing.
node {{paths.subagentLauncher}} <id> {{paths.scratch}}/subagent-task-order-total.md
→ verify each file:line by reading it; use the two confirmed claims; discard the third, whose line does not contain rounding
```

Why: the answer is checkable claim by claim (C09), and only verified claims enter the work (C12 subagent row).

### Bad — delegating the change itself

```
"Fix the rounding bug in the order module and update the tests."
```

Why: a subagent never edits (AGENTS section 8); the change must pass the main agent's D1–D7 and the project hook (C14, C16).

### Bad — using the answer as a fact

Copying the subagent's summary into the conventions or a commit message without reading the cited lines turns an unverified observation into a claimed fact (C09, C14).

## Prohibitions

- Running a subagent that is not in the C12 subagent table, or running it any way other than the launcher (C12).
- Delegating edits, product runs, scope, priority, or acceptance (AGENTS section 8).
- Writing an API key into a file, a task, or a command line (C12 subagent row).
- Using an unverified claim (C09, C14).

## Acceptance

- Every used claim has its verified source, the subagent id and log path are recorded, and no file changed through the subagent.
