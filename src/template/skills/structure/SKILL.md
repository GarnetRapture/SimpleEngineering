---
name: {{skills.prefix}}-structure
description: "Start: an implementation adds, moves, or renames a file, folder, type, constant, function, or tool module. Action: fix the domain, class, owner, dependency direction, and name from the conventions before the file is written, then connect every consumer. End: one owner per concept and every caller connected."
---

# {{project.name}} Structure and Naming

## System prompt (C01, C04, C05, C10)

You are the structure limb selected by the implementation process of the [shared workflow]({{link.workflowSkill}}). The [conventions]({{link.conventions}}) own the structure (C04), compiled-layer paths (C10), and naming (C05); this skill holds only the method. It adds no authority, exception, or stopping condition.

### Execution summary

Input: the exact symbol or path to add, move, or rename, its concept, and its callers. Find the existing owner, fix the path from the C04 or C10 table, check every import against the allowed-dependency table, fix the name by the C05 row of its language, and connect declaration, export, import, registration, selector, consumer, and release. Output: the placed and named artifact with every consumer connected.

### Detailed behavior specification

1. Search the designated domain for an existing owner of the same type, utility, constant, or rule. Extend it when it exists; record the search scope when none exists.
2. Fix the path from the C04 table (or the C10 table for compiled code). A new domain or class folder is created only in the change that fixes its owner, and the same change updates the table.
3. Compare every import with the C04 allowed-dependency table. Correct a forbidden direction at the owner, never bypass it.
4. Fix the name by the C05 row of the language so it states the real responsibility, input, result, side effect, and lifetime. Authored code carries no comments.
5. Connect declaration, export, import, registration key, selector, consumer, and release. A rename covers every external consumer. A move keeps the file's content and uses a move operation, never delete-and-recreate.
6. When a domain, class, layer, or compiled area node is added, moved, or removed, update the project node table of the [tracking]({{link.tracking}}) document in the same change (C15).

## Checklist

```
Structure:
- [ ] Step 1: owner of the concept found, or its absence recorded with search scope
- [ ] Step 2: path fixed from the C04 or C10 table
- [ ] Step 3: dependency direction checked against the C04 table
- [ ] Step 4: name fixed by the C05 rule of its language
- [ ] Step 5: every caller, registration, and consumer connected
- [ ] Step 6: project node table updated for any added, moved, or removed node
```

## Examples (C04, C05, C15)

### Good — a new concept lands in its domain owner

```
Concept: order status values used by the order list and the checkout
Owner search: src/order/ (domain) → src/order/types/ exists → extend src/order/types/order-status.types.ts
Node table: order/types already registered; no new node
```

Why: one owner per shared concept inside its domain (C04), found by search before a file is created.

### Bad — a generic bucket file

```
src/utils/helpers.ts   // formatPrice, parseOrder, retryFetch, useModal
```

Why: `utils` and `helpers` hide responsibility (C05), mix four domains in one file (C04), and give every concept a second owner.

## Prohibitions

- A second owner of a shared concept, a folder absent from the C04 or C10 table, or a forbidden dependency direction (C04).
- A name that hides its responsibility, a case-only difference, or a meaningless `new`, `temp`, `misc`, or `manager` name (C05).
- Delete-and-recreate in place of a move or an in-place edit (C12 file editing row).

## Acceptance

- The artifact sits at its table path under one owner, its name follows C05, and every declaration, import, registration, and consumer resolves.
