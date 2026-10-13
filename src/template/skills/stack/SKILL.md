---
name: {{skills.prefix}}-stack
description: "Start: a file of an active language stack changes or a stack check runs. Action: fix the owner, path, and shared contract, judge every executed logic, verify only through the stack's allowed checks, and observe the affected path from the same build when runtime behavior changed. End: the stack's checks pass with error 0 and warn 0 and no forbidden command was used as verification."
---

# {{project.name}} Language Stack

## System prompt (C01, C03, C10, C12, C19)

You are the stack limb selected by the implementation, FIX, static, or build process of the [shared workflow]({{link.workflowSkill}}). The [conventions]({{link.conventions}}) own the layer ownership (C03), compiled-layer paths (C10), the stack table with each stack's markers, allowed checks, and forbidden commands (C12), and execution-line support (C19); this skill holds only the method. It adds no authority, exception, or stopping condition.

### Execution summary

Input: the changed files, their stack, and the requirement. Identify the active stacks from their marker files, fix the owner and path, keep every cross-layer contract identical on both sides, judge every executed logic, run only the stack's allowed checks once the connected change set is ready, and observe the affected path from that same build through the authorized route when runtime behavior changed. Output: check exits and counts, and observation evidence bounded to what was exercised.

### Detailed behavior specification

1. Identify the active stacks: a stack is active only when every marker file of its C12 row exists at the project root. Never decide a stack from a directory name.
2. Fix the owner and path from C03, C04, and C10; for compiled code keep one declaration and implementation pair per type.
3. When two layers share data (an interface, a buffer, a schema, a message), keep the C03 contract owner identical on both sides in fields, sizes, offsets, and lifetime.
4. Apply the [verification skill]({{link.verifySkill}}) logic judgment to every changed logic the check will execute.
5. Run each allowed check of the active stack once, through its exact command, only when its inputs changed (AGENTS S28). Every diagnostic goes through the [FIX process skill]({{link.fixSkill}}).
6. When runtime behavior changed, observe the affected path from that same build through the C12 observation row, started only on the user's instruction. Platform, device, and execution-line support is stated only from an observation on that environment (C19).

## Active stack table

{{generated.stackSkillTable}}

## Checklist

```
Stack change:
- [ ] Step 1: active stacks identified from their markers
- [ ] Step 2: owner, path, and declaration/implementation pair fixed from C03, C04, and C10
- [ ] Step 3: shared cross-layer contracts and lifetime fixed
- [ ] Step 4: every executed logic judged through the verification skill
- [ ] Step 5: allowed checks run once, every diagnostic resolved
- [ ] Step 6: affected path observed from the same build when runtime behavior changed
```

## Examples (C07, C12)

### Good — the allowed check once, after the connected change set

```
Changed: src/order/logic/order_total.rs, src/order/logic/mod.rs
Active stack: rust (Cargo.toml exists)
Check: cargo clippy --all-targets -- -D warnings → exit 0, error 0, warn 0
```

Why: the stack table names the allowed check and its acceptance (C12), and C07 runs it once after the instructed implementation is connected.

### Bad — a forbidden command used as verification

```
cargo build && cargo run
```

Why: the stack table forbids `cargo build` as verification and `cargo run` as a product run (C08, C12); the project hook denies both.

### Bad — a check after every single edit

Running the type checker after each of five related edits uses the check to discover what to implement next; C07 runs it once on the integrated change.

## Prohibitions

- A forbidden command of the stack table used as verification, or a product run used to debug (C08, C12).
- A check run per edit, a build to discover architecture, or a passed check rerun for reassurance (C07).
- A support claim for a platform, device, or execution line without an observation there (C19).

## Acceptance

- Every allowed check of each active stack touched by the change reports exit 0 and meets its acceptance with error 0 and warn 0; runtime claims rest on same-build observation.
