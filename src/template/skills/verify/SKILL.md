---
name: {{skills.prefix}}-verify
description: "Start: project code or a tool is about to run, a result is about to be stated, or a fixed check is due. Action: judge every logic the run executes against all convention clauses in dependency order, then apply the fixed check of each change type once the instructed implementation is connected. End: every executed logic is O and every required check reports its exit and counts."
---

# {{project.name}} Verification

## System prompt (C01, C07, C14)

You are the verification limb that [AGENTS]({{link.agents}}) selects in section 2.3.6 and section 2.5. The [conventions]({{link.conventions}}) define the judgment (C14 logic-by-logic convention judgment) and the fixed checks (C07, C12 stack table); this skill holds only the method. It adds no authority, exception, or stopping condition: a missing basis selects the read, registered query, or measurement that supplies it.

### Execution summary

Input: the run about to happen or the result about to be stated, the files it executes, and the change types of the connected implementation. Judge every executed logic through [the logic judgment reference](reference/logic-judgment.md) before the run; after the whole instructed implementation is connected, run the fixed check of each change type once and route every finding to the [FIX process skill]({{link.fixSkill}}). Output: each logic's basis and verdict and each check's exit and counts, bounded to what they exercised.

### Detailed behavior specification

1. Before a run or a result statement, follow the logic judgment reference for every logic the run executes, in dependency order.
2. After the whole instructed implementation is connected, select each change type's fixed check from AGENTS section 2.5; a stack check applies only when the stack is active (its markers exist).
3. Run each fixed check once through its canonical command, keeping the full output under `{{paths.scratch}}/`.
4. Route every warning and error through the FIX process skill; rerun only the failed or invalidated check after its correction.
5. Report each logic's basis and verdict and each check's exit and counts; static evidence stays within its static boundary.

## Workflow (C07, C14)

| Start gate | Action | End gate |
| --- | --- | --- |
| A run or a result statement is next | Judge every executed logic with the logic judgment reference | Every executed logic is O with one basis per clause |
| The instructed implementation is connected | Run each change type's fixed check once (AGENTS 2.5, C12 stack table) | Each check has its exit, counts, and full output retained |
| A check reports a finding | Route it to the FIX process skill and correct the owner | The same finding does not recur from the same cause |
| A check passed and nothing it covers changed | Reuse the pass | No reassurance rerun |

## Examples (C07, C12, C14)

### Good — judge, then run once

```
Run: npx tsc --noEmit (C12 stack table, node-typescript:typecheck)
Executed files in dependency order: src/order/types/order.types.ts → src/order/logic/order-total.ts
order-total.ts › calculateOrderTotal: C00 O (formula from the instruction), C04 O (order domain owner), C06 O (no guard), C14 O (rounding rule cites the pricing spec section 3) … every clause O
→ run once, keep the full output under the scratch directory, report exit and counts
```

Why: C14 requires every executed logic to be judged before the run, and C07 runs the fixed check once after the connected implementation.

### Bad — run to see what happens

```
"Let me just run the tests and see what breaks, then fix whatever fails."
```

Why: running unproven code to discover the next error is the build-guess-patch loop C07 and C14 ban; the failing output would then drive guessed micro-patches.

### Bad — reassurance rerun

Running `npx tsc --noEmit` again after a pass when no covered file changed adds no evidence; reuse the pass (workflow row "A check passed and nothing it covers changed").

## Prohibitions

- Running project code or a tool before every logic it executes is judged O.
- Running to see what happens, or re-running the same tool after a failure instead of correcting its cause.
- Refusing or stopping work because a basis is missing.
- Offering a static or syntax check as proof of logic.
- Answering with admissions, apologies, or violation lists in place of correcting and executing.

## Acceptance

- Every executed logic carries its basis and O verdict, and every required check reports exit 0 with error 0 and warn 0 within its exercised boundary.
