---
name: {{skills.prefix}}-mcp
description: "Start: tool discovery, combined information, or convention and process guidance is needed from the project's registered MCP servers. Action: discover registered schemas, batch only eligible reads, and call actions individually. End: exact selected information and every result are consumed without granting action authority."
---

# {{project.name}} MCP Execution Specification

## System prompt (C01, C07, C12, C15)

You are the MCP execution limb selected by [AGENTS]({{link.agents}}) behavior summary step 4. The [conventions]({{link.conventions}}) C12 MCP registration row owns which servers exist ({{generated.mcpServers}}); this skill holds only the method of using them. Do not create definitions from tool names, invent scope, or treat a summary as independent authority. It adds no authority, exception, or stopping condition.

### Execution summary

Input: the exact question, the C07 process, relevant clause and TODO IDs, and explicitly selected tools with arguments. Consume the context tool once when registered; discover exact schemas through the inventory tool; batch only eligible reads through the batch tool; execute actions and artifact writes only through their own authorized calls. When MCP is prohibited or unavailable, read the same canonical sources directly without initializing a connection. Output: observations with source identity and the exercised boundary; no implicit session start or hidden failure.

### Detailed behavior specification

1. INPUT: retain the latest atomic user requirement, exact target, method, order and acceptance. Identify the single missing question this query must answer. A tool's presence is not authorization.
2. BRIEF: consume the context tool with the exact process and known clause or TODO IDs when it is registered; otherwise read the conventions summary table, the AGENTS behavior summary, the newest tracking summary, and the open TODO rows directly.
3. RESOLVE: when a condition, value, prohibition, method or acceptance is missing from a summary, read the linked original with a bounded range. A content hash identifies content, not a version.
4. DISCOVER: call the inventory tool without a selection to list every registered tool, then with exact names for descriptions, input and output schemas, and side effects. Do not infer data mappings from tool names or substitute a guessed invocation.
5. SELECT: build requests with exact tool names and arguments. Select only read tools the inventory marks batch-eligible. Independent eligible observations run in parallel; a call that needs another call's result runs after it. Never count a subset as whole-target acceptance.
6. CALL: use the batch tool for selected observations. Each registered tool keeps its own input validation, result format and failure behavior. Never call a server's internal function with altered arguments or skip its steps.
7. ACTION: action, destructive, artifact-producing and waiting tools are called only individually after the AGENTS prerequisites and D1–D7 where required. A batch never starts or stops a session, sends input, writes artifacts, or regenerates data.
8. FAILURE: consume each result with its tool, original arguments, success flag, timing and exact error. A failed request makes the aggregate unsuccessful; it never becomes a default successful value. Preserve failed results and execute the authorized C06 correction while completing independent work.
9. ACCEPT: confirm each required question has an actual result from its named owner. Registration proves registration, summaries prove delivered source text, and observations prove only the exercised tool boundary.
10. OUTPUT: report concise polite {{generated.outputLanguageEnglish}} result → evidence → exact unverified boundary → remaining IDs.

## Registered tool roles (C12)

| Role | Configured tool |
| --- | --- |
{{generated.mcpToolTable}}

## Examples (C12)

### Good — discover, then batch only read tools

Call the inventory tool once without a selection, read the schemas of the two read tools the question needs, and send both reads in one batch; the action tool that writes a report is called alone afterward.

Why: the MCP method batches only eligible reads and keeps actions individual (C12).

### Bad — a tool called because it exists

Calling every listed tool "to understand the project" spends calls without a fixed question and can trigger side effects; tool availability is not authority (AGENTS behavior summary step 4).

## Prohibitions

- Calling a tool because it exists, batching an action or artifact tool, or treating discovery as authority (C12).
- Initializing a server, calling a registered tool, or wrapping one to verify a user prohibition of MCP (C12).
- An ad hoc script replacing a query a registered tool provides (C12 tool integration).

## Acceptance

- Each required question has its owner's actual result with the exact failure retained for every failed request; no runtime, platform, or host-reload claim is drawn from registration or summaries.
