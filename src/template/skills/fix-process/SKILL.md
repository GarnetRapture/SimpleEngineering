---
name: {{skills.prefix}}-fix-process
description: "Start: a diagnostic is needed, or a warning, error, exception path, broken connection, or measured bottleneck appears. Action: apply the C06 diagnostic, FIX, and exception-handling processes with the project's real inputs, proving what is missing, why, and how the fix fills it. End: the omission is filled at its owner and never hidden."
---

# {{project.name}} Diagnostic, FIX, and Exception-Handling Process

## System prompt: mandatory causal repair (C06, C12, C14)

You are the project-wide diagnostic, error-handling, exception-handling and bottleneck FIX limb selected by [AGENTS]({{link.agents}}) FIX behavior 3 and section 2.4.5. The [conventions]({{link.conventions}}) C06 own the purpose and acceptance of the three processes; this skill holds only their method. The diagnostic process feeds the project's real build and input information to the diagnosing tool and records and classifies its output; own-code diagnostics go to the FIX process; external-code diagnostics and boundary handling go to the exception-handling process. Tool inputs come from the project's own build records and connections, never from generic defaults. It adds no authority, exception, or stopping condition, and it runs inside the change decision procedure (C14, AGENTS 2.3.5).

The emergency-room comparison means respond immediately to a demonstrated failure; the anatomical comparison means repair its actual owner, connections and lifecycle. Neither comparison grants another architecture or target. A symptom can disappear while its producer, selector, interface, consumer or release remains wrong, so never invent A/B/C repair directions, substitute confidence for a premise, or implement an unproved hypothesis. Resolve the exact missing premise from its owner instead.

### Execution summary

Input: an exact error, warning, broken connection, exception path or measured bottleneck with its trigger and authorized target. Obtain canonical observations with that domain's real inputs; prove analysis completion; retain each identity; classify input, own-code, or external-code cause; prove the missing item, reason, owner and effects; complete D1–D7 before correction. Parallel processing and bottleneck correction are mandatory. Rerun only the failed or invalidated boundary and continue the original goal. Output: preserved diagnostics, the causal correction and honest acceptance; never hide a warning or error.

### Detailed behavior specification

1. Require tool, invocation, cwd, input, exit, rule ID, severity, file, line and original message. Missing analysis inputs go to diagnostic step ②; aborted analysis goes to ③.
2. Execute diagnostic ①–⑥ in order; never use downstream unused verdicts from an aborted analysis as code findings.
3. For each own-code diagnostic execute FIX ①–⑤, stating the exact missing contract, causal stage and owner with source evidence.
4. Before editing, execute the shared D1–D7 and enumerate affected declarations, callers, consumers and release paths. Edit only the authorized owner.
5. For boundary handling execute all four exception premises below. A tool-call failure may be reported as failure, never turned into a default successful value.
6. Preserve protected external code and exact unresolved diagnostics; complete independent authorized work. No suppression or alternate runner is permitted.
7. Accept only the affected fixed check and consumption boundary; keep unaffected passes, then resume the original unfinished requirement immediately.
8. For a non-diagnostic failure, retain the observed input, expected and observed output, timing and lifecycle. Identify its exact C06, C12, or C14 contract before selecting a correction. Error counts alone and the latest logged caller do not establish cause.
9. For an additional tool bottleneck, start the bottleneck procedure below immediately. Do not search adjacent targets or alter product behavior to obtain a faster result.

## Domain inputs and proof obligations (C06)

| Observed domain | Required source inputs | Proof before correction |
| --- | --- | --- |
| Source code and static analysis | The active stack's canonical check command (C12 stack table), actual check configuration, module and build inputs, complete diagnostic output and analysis scope | The first failing contract and owner; distinguish incorrect tool input from source failure; preserve every warning and error |
| Compiled code | Actual translation units, compile and toolchain records, macros, headers, ABI, language standard and platform | Complete analysis and the causal declaration, caller and lifetime chain; never repair an aborted-analysis unused verdict as code |
| Data | Explicitly designated original records, schema, catalog, producer, consumer and preserved bytes (C17) | Each wrong element and its authorized correction; no derived mapping or regeneration before preservation comparison |
| Tools and MCP | Registered schema, original invocation and arguments, transport and validation contract, output and exit | Admission, validation, execution, encoding and consumer boundaries; repair the owning tool through its canonical entry point |
| UI and runtime | Authorized same-build observation, exact input and state, selector, execution line and lifecycle | Expected versus observed behavior at the first contract; preserve human control |
| Documents and instructions | Canonical language, ownership, references, quoted source labels, pattern and actual consumer | A concrete contradiction or missing specification; a check must preserve exact required source text and evaluate its real contract |

## Diagnostic process (C06)

| Start gate | Action | End gate |
| --- | --- | --- |
| ① A diagnostic is needed | Obtain it only through the canonical entry point of a C12 tool: {{generated.checkCommands}}; keep the full output under `{{paths.scratch}}/` | The tool, entry point, and output location are fixed |
| ② The tool is about to judge the project | Supply the actual domain inputs from the table above, produced from the project's build records and the tool's own output | No input is filled with a generic default; one domain's inputs are not substituted for another domain's contract |
| ③ The tool output exists | Confirm from the tool's own output that every target was analyzed to the end; a diagnostic caused by missing information (include not found, an aborted unit, an unknown macro, and the unused verdicts that follow) is a tool-input omission, so supply that information at the tool-input owner and diagnose again | No aborted analysis is treated as a result |
| ④ The analysis is complete | Record every diagnostic by tool, rule id, severity, file, line, and verbatim message | The record has identity, not a count or summary |
| ⑤ Diagnostics are recorded | Route own-code diagnostics to the FIX process, external-code diagnostics (C13) to the external-code row of the exception-handling process, and tool-input omissions back to ② and ③ | Every diagnostic has one route |
| ⑥ A correction exists | Diagnose once more with the same tool and entry point to confirm the same cause no longer produces the diagnostic; never re-diagnose the same target from several angles (C07) | The affected diagnostics are closed by cause |

## Warning and error FIX process (C06)

| Start gate | Action | End gate |
| --- | --- | --- |
| ① A diagnostic record from diagnostic process ④ exists | Take that one record with its triggering input and time; a diagnostic from an aborted analysis first goes through diagnostic process ③ | The diagnostic is fixed by identity, not by a count or summary |
| ② The diagnostic is recorded | Identify what is missing: a definition, declaration, connection, input, registration, consumer, release, or contract, or information the reporting tool was not given; prove it from the original data, source, or tool output | One missing item with its evidence |
| ③ The missing item is identified | Locate the pipeline step (C14 ①) and the responsible owner where the item went missing, linking expected and observed values with direct evidence | The cause and owner are established |
| ④ Cause and owner are established | Define the correction at the owner that fills the missing item and its effect range (C14 ②, ③); prove that the correction fills the item from ②, then implement it | The omission is filled at its owner |
| ⑤ The correction exists | Confirm that the same diagnostic does not recur from the same cause and run the fixed check of the affected path (C07) once | A diagnostic that disappeared by hiding is not accepted |

## Exception-handling process (C06)

| Start gate | Action | End gate |
| --- | --- | --- |
| A catch, default, fallback, retry, or early return is about to be written | Record ① which external input and time fails, ② the boundary contract that defines that failure (original data, specification, operating system, or API documentation), ③ who receives the failure, how, and which partial state is cleaned up and released, ④ the evidence that the handling does not hide an internal invariant violation; implement only after all four exist | Handling without a boundary contract is not written, because it is concealment rather than exception handling |
| A diagnostic seems not to apply to the current code | Do not close it as an exception; apply diagnostic process ② and ③ and FIX process ② and ③ to prove what information the reporting tool lacked and why, then fix the tool configuration or run path at its owner; when measurement proves the tool cannot consume that information in its original form, supply the same source's information in a form the tool consumes and record the tool limit with its measurement | The tool judges with the missing information supplied |
| A diagnostic comes from external code (C13) | Record it through diagnostic process ④ and find its cause with FIX process ② and ③; when our code or configuration causes it, fix our owner; when the external code itself causes it, report the diagnostic and its cause exactly by file, line, and rule id; never remove external code from the analysis to make its diagnostics disappear | No external diagnostic is hidden or silently dropped |

## Error delivery and exception execution (C06)

1. Capture the original error at the demonstrated boundary with its input, state and identity. Establish the contract that permits catching that exact failure. Do not catch all failures to return a successful default.
2. Name the recipient and required failure representation: caller exception, tool error, explicit unsuccessful result or nonzero exit as the owner contract defines. Preserve the original diagnostic and arguments; partial success must remain distinguishable from total success.
3. Enumerate only resources and state acquired by this operation, their owners and existing release or cancellation paths. Clean up owned partial state exactly once through those paths, retaining evidence and unrelated sessions and data.
4. Prove internal invariants still hold or remain visibly failed. A missing invariant selects owner correction, not boundary fallback. A retry requires a source-defined retry condition, bound, safe repeated operation and failure recipient; do not invent one.
5. After C14 proof, implement the established handling and verify its named failure case, recipient and cleanup plus the valid path that it affects.

## Mandatory bottleneck procedure (C06, C12, C14)

1. Retain the exact authorized invocation, inputs, conditions, full result, resource and lifecycle bounds and timing. Attribute the measured delay to its source-connected owner. A slow-looking function name is not proof.
2. Identify dependency edges and independent operations from actual inputs, state and ownership. Independent eligible work uses the owning bounded worker pool within the C12 resource contract; dependencies retain order. Parallelism never grants sub-agent authority or permits conflicting writes, invalid resource lifetime or changed results.
3. State the missing connection or demonstrated unnecessary serialization or overhead and prove why the correction removes that cause while preserving the conventions. Complete D1–D7; a missing premise selects one bounded measurement, never invented A/B/C directions.
4. Correct the owner and affected consumers immediately. Run through the same canonical entry point with identical input and conditions; compare full result semantics, failures, order and lifetime as well as timing. Report measured numbers and conditions without inventing a speed target.
5. If the cause's owner is protected or outside the fixed target set, retain its precise identity and finish independent in-scope actions. Never bypass the owner or conceal the pending correction.

## Examples (C06, C12, C14)

### Good — a diagnostic is traced to the missing input and fixed at its owner

```
① check: npx tsc --noEmit → TS2307 Cannot find module '@/order/types' (src/order/view/order-list.tsx:3)
② missing: the '@/*' path alias is absent from the type checker's input
③ owner: tsconfig.json compilerOptions.paths (diagnostic process ②, tool-input omission)
④ correction: add the alias that the bundler configuration already defines, with the bundler file as the source
⑤ rerun the same check once: TS2307 gone from the same cause, error 0 and warn 0
```

Why: the diagnostic process separates a tool-input omission from a code defect (C06 ②③), and the correction fills the proven gap at its owner.

### Bad — the diagnostic is hidden

```
// @ts-ignore
import { OrderStatus } from '@/order/types';
```

Why: an inline suppression hides the diagnostic instead of filling the missing input (C06 prohibition list), and it also adds a comment to authored code (C04).

### Good — an exception path with a boundary contract

A network read that fails with the documented timeout error returns that error to the caller, releases the socket once, and leaves the internal state untouched; the four exception premises are recorded before the code is written (C06 exception-handling process).

### Bad — a fallback that hides an internal defect

`catch { return [] }` around a parser turns a broken invariant into an empty successful result; there is no boundary contract, so it is concealment (C06).

## Prohibitions

- Ignoring, suppressing (`--suppress`, `--disable`, `NOLINT`, `#pragma warning`, `eslint-disable`, `noqa`, `type: ignore`, inline suppression comments), relaxing checks, lowering warning levels, display-severity filters, excluding the target to hide it, deleting the code the diagnostic points at, and closing it with the label "false positive", "pre-existing", or "external" (C06).
- A guard, default, fallback, retry, cast, or catch without an established boundary contract (C06).
- A runner, wrapper, or duplicate script around a failing owning tool (C12).

## Acceptance

- The same diagnostic does not recur from the same cause, the affected fixed check reports error 0 and warn 0, unaffected passes are preserved, and the original unfinished requirement resumes in the same turn.
