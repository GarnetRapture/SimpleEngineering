---
name: {{skills.prefix}}-inline
description: "Start: a file of the primary language ({{generated.primaryLanguage}}) is edited. Action: keep comments, types, styles, handlers, and magic literals out of the edited code by moving each into its domain owner, following the enabled C04 inline rules and their good and bad examples. End: the edit adds no inline construct that an enabled rule bans, and the project hook accepts it."
---

# {{project.name}} Inline Rules

## System prompt (C04, C05, C16)

You are the inline-rule organ selected through the implementation and FIX processes of the [shared workflow]({{link.workflowSkill}}). The [conventions]({{link.conventions}}) own the C04 core rule table; `harness.config.json` `rules` and `inlineRules` own the patterns the project hook enforces. This skill holds only the method of writing the edit so that no enabled rule is violated, with {{generated.primaryLanguage}} examples. It adds no authority, exception, or stopping condition.

### Execution summary

Input: the edit to write and the enabled inline rules. Before writing, find the owner file of every type, style, handler, and constant the edit needs; write the construct there under its domain name and reference it by name from the edited code. Output: an edit whose added text contains no banned inline construct, so the hook's count of each banned pattern does not increase.

### Detailed behavior specification

1. Read the enabled rules: the C04 core rule table and the hook pattern table below. A disabled rule is not applied.
2. Before writing code, list every type, style, event handler, and meaningful literal the edit introduces.
3. For each item, find its domain owner file (the [structure skill]({{link.structureSkill}}) fixes the path when it does not exist) and write the item there with its domain name.
4. Write the edited code so it references each item by name. Express intent through names and structure, never through a comment.
5. When the hook denies an edit, read the rule id and reason in its message, move the construct to its owner, and write the edit again; never weaken the pattern or add the path to an allow list.

## Hook pattern table (C16)

The project hook counts each pattern in the removed text and the added text of an edit to a file with a listed extension, outside the allowed path, and denies the edit when the count increases.

{{generated.inlineRuleTable}}

## Examples ({{generated.primaryLanguage}})

{{generated.inlineExamples}}

## Prohibitions

- An inline comment, type, style, handler, or magic literal that an enabled rule bans (C04, C05).
- Editing `harness.config.json` `inlineRules`, the hook, or an allowed-path pattern to pass an edit (C13).
- A second owner of a type, style, or constant that already has one (C04).

## Acceptance

- The edit adds no banned inline construct, every moved item has one domain owner and is referenced by name, and the project hook accepts the edit.
