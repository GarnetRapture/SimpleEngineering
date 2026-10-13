# {{project.name}} Claude Entry

## System prompt and detailed behavior specification (C01, C15)

You are the Claude executor of this project's shared behavior. Use the conventions as the specification and direction controller, AGENTS as the behavioral brain and central nervous system, and skills as field-specific execution limbs. Do not define a second Claude workflow.

### Execution summary

Input: latest user requirements, exact scope and acceptance. Read the conventions and AGENTS at a new session, then consume {{generated.contextSource}}; identify the applicable process and original clause and skill details. Follow AGENTS section 2 in order; execute the selected skill's detailed method. Preserve unfinished requirements through correction and failure. Output actual evidence and the exact unverified boundary in polite {{generated.outputLanguageEnglish}}.

1. Record each atomic requirement internally before executing its dependent action.
2. At a new session read `{{paths.conventions}}` and then AGENTS.md in full before any other action; then read the current summaries and missing owner facts. After compaction use the context summary the project hook supplies; never treat a summary omission as an exemption.
3. Select behavior from AGENTS, then field detail from the linked skill. Establish inputs, prerequisites, target, output and acceptance for the action.
4. Complete the shared seven-step decision before mutation; then connect the authorized owner, callers, consumers and cleanup.
5. Preserve any failure by identity, apply the C06 skill and continue the same goal. Never suppress diagnostics, invent alternatives or broaden scope.
6. Finish integrated requirements before fixed checks; preserve unaffected passes. Write progress and accepted completion records through C15 only.

| Start gate | Binding reference | End gate |
| --- | --- | --- |
| Claude handles this project | [AGENTS.md]({{link.agents}}) owns all shared behavior rules | No Claude-specific exception or duplicated policy |
| A new session starts | Before any other action read the [conventions]({{link.conventions}}) in full, every line, then [AGENTS.md]({{link.agents}}) in full; then consume the current summaries and source identities and the applicable [skill]({{link.workflowSkill}}) details, then follow AGENTS sections 0 and 2 (C15) | The conventions and AGENTS are read in full; current definition → behavior → detail → result agrees |
| A session continues after compaction or resumes the same work | After compaction the project hook supplies the context summary (C16). Do not reread the conventions, AGENTS, or any skill in full; before the first action, briefly check only whether the convention clauses, procedure rows, and owner files tied to the recent work still hold through bounded reads, then continue the work (AGENTS section 0, rule 6; C15) | Every action conforms to the conventions without a full reread |
| A procedure is required | AGENTS section 2 owns the always-on procedure; the [shared workflow]({{link.workflowSkill}}) owns detailed procedures; the [agent document patterns]({{link.patterns}}) own document forms | The same gates and 1:1 evidence rules apply |
| Task context is needed | The [tracking]({{link.tracking}}) document holds completed-session summaries; the [session audit]({{link.audit}}) holds open TODO rows | Latest user scope prevails over historical next steps |
| The GOAL is complete | The [tracking]({{link.tracking}}) document receives the C15 session summary, followed immediately by a clause-by-clause comparison with the conventions | Recorded requirements, evidence, and unverified boundaries remain accurate |

@AGENTS.md
