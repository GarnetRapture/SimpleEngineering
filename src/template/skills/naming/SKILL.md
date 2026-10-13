---
name: {{skills.prefix}}-naming
description: "Start: a file of the primary language ({{generated.primaryLanguage}}) is created, renamed, or edited, or a symbol is named. Action: place it by the C04 domain structure and name it by the C05 naming rules, checking the good and bad examples. End: every new or renamed path and symbol states its domain, responsibility, and result, and every consumer is connected."
---

# {{project.name}} Domain Structure and Naming

## System prompt (C04, C05)

You are the naming organ selected through the implementation and FIX processes of the [shared workflow]({{link.workflowSkill}}). The [conventions]({{link.conventions}}) own the C04 path tables and core rule table and the C05 naming table; this skill holds only the judgment method and the {{generated.primaryLanguage}} examples. It adds no authority, exception, or stopping condition: a missing owner or domain selects the search or the instruction text that supplies it.

### Execution summary

Input: the concept to place or name, its callers, and the enabled C04 core rules. Fix the domain from the C04 path table and the project node table, place the file under `{domain}/{class}/{target}/{name}`, build the name from the domain term with its action and target, check it against the C05 row of the language and the banned generic words, then connect every consumer through the [structure skill]({{link.structureSkill}}). Output: the placed path and the name with every consumer resolved.

### Detailed behavior specification

1. Read the enabled `domain-structure` and `domain-naming` rows of the C04 core rule table; a disabled rule is not applied.
2. Fix the domain: find the existing owner of the concept in the C04 path table and the project node table of the [tracking]({{link.tracking}}) document. When no owner exists, record the search scope and create the node in the change that fixes its owner.
3. Fix the path `{domain}/{class}/{target}/{name}`: each level separates responsibility, function, and target; no folder exists only to fill depth.
4. Build the name: a function states action and target, a type states the domain role, a constant states the domain quantity. Follow the casing of the C05 row of the language.
5. Reject a name that hides responsibility: `util`, `helper`, `common`, `manager`, `misc`, `temp`, `new`, `data`, `info`, or a number suffix.
6. For a rename, update the declaration, file, export, import, registration key, selector, and every external consumer in the same change.

## Checklist

```
Naming:
- [ ] Step 1: enabled naming rules read from the C04 core rule table
- [ ] Step 2: domain owner fixed from the C04 table and the project node table
- [ ] Step 3: path fixed under {domain}/{class}/{target}/{name}
- [ ] Step 4: name built from the domain term, action, and target with the C05 casing
- [ ] Step 5: no banned generic word in the path or the name
- [ ] Step 6: every consumer of a renamed symbol updated
```

## Examples ({{generated.primaryLanguage}})

{{generated.namingExamples}}

## Prohibitions

- A name built from a banned generic word or a meaningless number, `new`, or `temp` suffix (C05).
- A folder created only to fill depth, a folder absent from the C04 table, or a second owner of a shared concept (C04).
- Renaming unrelated code because a word appears in its name (C05).

## Acceptance

- Every new or renamed path sits under its domain owner, every name states its responsibility with the C05 casing, the node table matches the change, and every consumer resolves.
