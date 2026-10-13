# Logic-by-logic convention judgment (C14)

## Contents

- Checklist
- Step 1: Order the executed files
- Step 2: Split each file into logic units
- Step 3: Judge one logic against every clause
- Step 4: Correct and judge again
- Step 5: Run once
- Clause questions

## Checklist

Copy into the working record and tick per logic:

```
Logic judgment:
- [ ] Step 1: executed project files listed in dependency order
- [ ] Step 2: logic units listed per file in source order
- [ ] Step 3: current logic judged C00-C20, one basis per clause
- [ ] Step 4: every X or ? corrected after obtaining its basis, then judged again
- [ ] Step 5: every executed logic O, then one run through the canonical entry point
```

## Step 1: Order the executed files

Start from the canonical entry point the run uses (a C12 command, a registered tool, or the tool's `main`) and follow its imports to the project files it executes. Order them so each file comes after every project file it imports. A file the run does not execute is outside this judgment. A run that executes no project code (a document edit, a pure query) has no logic units and records that reason.

## Step 2: Split each file into logic units

A logic unit is one constant, one function or method, or one branch whose result decides a value, a mapping, a write, a resource use, or a failure. List the units in source order.

## Step 3: Judge one logic against every clause

Answer each clause question below for the current logic with one basis: instruction text, an original record, a measurement (registered tool, owning tool output, authorized observation), or an official source. Record O, X, or ? with that basis. A clause outside the logic's scope gets its exact scope reason. Existing code is never a basis (C01).

Move to the next logic only when the current logic is O on every applicable clause.

## Step 4: Correct and judge again

- X: correct the logic at its owner through the change decision procedure (C14 D1-D7, AGENTS 2.3.5), then judge it again from Step 3.
- ?: the next action reads the owner, queries a registered tool, or measures. A missing basis is never a reason to stop or to assume.

## Step 5: Run once

When every executed logic is O, run once through the canonical entry point (C12) and keep the output under `{{paths.scratch}}/`. A failure goes to the [FIX process skill]({{link.fixSkill}}); the corrected logic is judged again before the next run.

## Clause questions

| Clause | Question for one logic |
| --- | --- |
| C00 | Does every value, mapping, name, field, and format come from instruction text, an original record, or a measurement, with no name, field, number, format, or pattern inference, no prefix or suffix composition or stripping, no fallback name, and no target the user did not instruct? |
| C01 | Does it implement a definition the conventions own without creating a second copy of a rule or a new authority? |
| C02 | Does it follow the fixed decision rows that govern its subject and the latest user instruction for that subject? |
| C03 | Does it stay in its owner's boundary and connect producer, selector, shared interface, consumer, and lifetime? |
| C04 | Is it in the designated domain and class, with one owner for each shared type, utility, constant, and rule, and does it satisfy every rule of the C04 core rule table (comments, inline types, inline styles, magic literals) that the project enabled? |
| C05 | Does its name state its real responsibility, input, result, side effect, and lifetime? |
| C06 | Does every failure surface with its identity, with no guard, default, fallback, retry, cast, or catch that lacks an established boundary contract? |
| C07 | Which acceptance evidence does its process type require, and does that evidence exist? |
| C08 | Does it respect the product and observation entry points without running a product entry as a debugging means? |
| C09 | Is a direct who, what, when, where, why, and how proposition supported for it? |
| C10 | For compiled code, does it follow the designated path, declaration and implementation pair, and build target? |
| C11 | Does it avoid any version field, number, comparison, or old/new switch and preserve originals and user changes? |
| C12 | Does it use registered tools and canonical entry points, the active stack's allowed checks, and the resource contract (CPU at most 50% of physical cores, RAM at most 80%, no hard-coded counts, shared-queue distribution)? |
| C13 | Does it change protected targets only through their allowed routes? |
| C14 | Does every formula, tolerance, and constant have an original, specification, or measurement basis, and is the change decision record D1-D7 complete? |
| C15 | Do its records and times follow the tracking format and the time command? |
| C16 | Does it avoid claiming enforcement or future behavior from a document change? |
| C17 | Does every write to designated data go through its producer and the catalog gate? |
| C18 | Does it stay inside its role group boundary? |
| C19 | Is every platform, device, or execution line support claim backed by an observation on that environment? |
| C20 | For production diagnostics, does it record through the designated log owner? |
