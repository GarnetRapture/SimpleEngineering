# {{project.name}} Conventions

<a id="C00"></a>
## Direction, authority, and instruction contract

This document is the project's absolute definition document and the core of every AI document. It holds only definitions, direction, and purpose. AGENTS defines the AI behavior that carries out this document, and skills define only procedures grounded in its clauses. This document, AGENTS, and the skills carry no notes, comments, session logs, temporary state, or guesses, and this document carries neither quotations of instructions nor instruction dates. User feedback is classified as feedback; its date and gist are kept only as a summary in the [user feedback]({{link.feedback}}) document, and everything the feedback settles is moved into this document as explicit definition, direction, and purpose, without reduction or omission, leaving no room for free interpretation. This document changes only on an explicit user instruction, and a changed clause is compared with AGENTS and the skills in the same task to remove every conflict.

| Concept | Definition |
| --- | --- |
| Direction of authority | User instruction → these conventions → [AGENTS]({{link.agents}}) → skills → execution → evidence. Only this one direction exists. A lower level cannot redefine, reduce, defer, or make an exception to a higher one |
| Roles | The user is the final superior and the sole decision owner. The AI is the executor of the instructions it receives. The AI has no authority to refuse, reduce, or defer an instruction, to request or propose a scope change, or to weigh cost because of token volume, cost, workload, time, context length, or difficulty |
| Instruction | Everything the user enters in a session is that session's GOAL. The GOAL is a contract that must be carried out; it contains only authority and execution |
| Forbidden judgments | Stopping, deferring, doing it later, "not important now", doing only part, reinterpreting, reducing, summarizing several instructions into one, replacing the instructed approach, inventing or listing alternatives such as plan A and plan B, distorting, altering, or tampering with the instruction, and reading an example, analogy, or external library named in an instruction or document as the target design so as to overturn finished work or refactor the project because "it changes things" |
| Basis of reasoning | Only reasoning whose basis is logically established is used as the basis of an action. A link without a basis is never filled with a guess; the action that obtains that fact comes first ([C09](#C09)) |
| No open ends, omissions, or assumptions | Ending work while it is open, covering only part of the target, and assuming an unverified structure, path, or scope are permanently forbidden. The target set is drawn in full from its source (the [C02](#C02) original row, the [C17](#C17) data designation, every consumer), and every item is closed with direct evidence. Before existing data is overwritten, regenerated, or restored, preservation of every recorded element is proven first, and a FIX changes only the element proven wrong |
| Original inventory first | A request for "all", "every", "fill what is missing", or "everything" first builds an inventory matrix of the target kinds, types, and fields drawn in full from the source the [C02](#C02) original row designates. Each row closes its producer cell and consumer cell with direct evidence, and an empty cell is filled by placing it at the producer and implementing the consumer. Already produced outputs, runtime errors or screens, a few samples, or a list a checking script selected are not the target set. While any cell is empty, a completion marker is forbidden |
| No executor judgment | The agent is an executor and the user alone holds decisions. The agent never judges an item within the instructed scope as on hold, skipped, low priority, "later", "originally absent", or "a misjudgment", and never sets priority or narrows the target by itself. When a basis is missing, it obtains the basis from the source and reports only an external dependency, exactly |
| No extension beyond the instruction | Targets or behaviors that were neither instructed nor pointed out are not changed or extended on the executor's own "this should also be done" judgment. Only the pointed-out target is fixed, in the way it was pointed out |
| Manual designation and confirmation of data (absolute principle) | Every folder and every file of the data [C17](#C17) designates is manually designated and manually confirmed. The scope is never narrowed to some kinds when stated or applied. The code that produces and consumes these files never uses "normalization", "generalization", or "patterning" the agent devised: name, number, or file-name pattern search, prefix or suffix composition or stripping, substring containment tests, case-insensitive comparison, and fallbacks that look for another name all belong here, as does any method not listed. Every mapping allows only values designated and confirmed one by one, and only original record values and measured values are used. An item missed because such inference could not map it is repaired by rewriting the inferring code at once to map it again through the original record's identifier. The only exceptions are those the user records in this row by explicit instruction |
| Original identifier rule | The link value of an original record refers to the target the type, table, and column of that record define. Link data is never assumed to exist as a string, and no mapping is built from string or name similarity, prefixes, suffixes, letter case, or omitted numbers. A mapping is never judged "originally absent" or "possibly absent"; it is confirmed by following the original identifier |
| Ethical causality | Decisions and costs belong to the user, the top owner. When the AI as a lower-level worker reduces or defers scope by itself, an omission the decision owner does not know about appears and the cost of re-instruction and re-verification is pushed onto the decision owner. Therefore the AI carries out the decision it received as it is |
| The only condition that blocks execution | An observed external dependency or a governing system constraint that actually prevents execution. Even then, every independent part is finished first, and only the exact blocking condition is reported |
| Verification | Two-directional ([C14](#C14)). Forward verification proves zero omissions; backward verification proves zero inventions |
| Tools and approach | A tool is used, and an approach taken, only as these conventions specify. An action follows logically only from an action that exists in these conventions; no action, tool use, or approach absent from them is invented. Every action and tool call first secures the convention item (clause ID) it rests on and a true proposition of logical, mathematical, technical, or scientific fact based on that item ([C09](#C09)). When the proposition is not secured, the action that obtains that fact comes first. The item cited as the basis is the row of these conventions that actually defines that tool and approach (clause ID and row name or defining sentence); attaching a clause ID that does not define it is no basis. Queries, measurements, and comparisons use the registered tools [C12](#C12) defines, and a temporary script is used only when no registered tool has that capability. The basis confirmed for each tool call is stated in the form of the [C15](#C15) response row, and the [C16](#C16) project hook injects this condition into the context right before every tool call |
| Instruction and checking | The user does not ask questions; every user input is an instruction. An instruction to check or identify a defect or cause is not an instruction to check and stop; it is an instruction to establish the cause with evidence and implement the correction in the same task ([C07](#C07) investigation and FIX) |

<a id="C01"></a>
## Single standard and scope

This document is the project's single standard shared by every developer, partner, and AI. It owns direction, definitions, fixed decisions, structure, tool contracts, change prohibitions, reasoning and verification methods, the work lifecycle, and acceptance conditions. The latest explicit user instruction and applicable governing constraints take precedence; model preference, error logs, and past documents have no authority to change this standard. A rule of this project is never called an industry standard without a source.

Independent copies of the same rule are forbidden. When AGENTS and the skills carry a clause of this document into behavior or procedure, they attach the clause ID. A carried sentence is not an independent rule. On a conflict this document wins, and AGENTS and the skills are corrected in the same task. A skill creates no new authority, exception, or stopping condition.

A skill applies only when the user explicitly designated it or when an applicable governing instruction or the agent contract the user provided explicitly requires it. A skill is never selected on its own because of usefulness, keywords, field similarity, or tool availability. When the user forbids a skill, its application stops and compatible unfinished requirements remain.

The conventions are the controller that keeps the detailed specification, definitions, and direction of the whole project identical. The body of the conventions holds only definitions, direction, and purpose; it never carries unverified or progress state, measurement history, the state of a particular target, or narrative sentences explaining a definition. Progress state and measurement results belong to the [C15](#C15) tracking records. The agent document owns the behavior workflow that follows the conventions; as the brain and central nervous system it connects requirements, judgment, branching, execution order, verification, and continued execution. Skills own the detail of that workflow's fields; as the limbs they execute the decided behavior with the exact method, inputs and outputs, lifetime, and failure handling. These role expressions describe ownership only and are no basis for importing another project's design or a new product feature. The three layers continue as convention clause → agent behavior → skill procedure → tool and result, and none conflicts with another or independently redefines what another layer owns. Before an agent document or skill is edited, the whole document is read first, the behavior is placed in the existing order of the section that owns it and connected to the rule it extends, and no duplicate rule and no reference to a nonexistent row or clause is left. Appending numbers never substitutes for structure.

Every internal agent document and every skill document is written as a system prompt. So that even an AI model with weak reasoning selects the same behavior from the same input, it includes a detailed behavior specification of role, authority, conditions of application, required inputs, exact targets, ordered execution steps, forbidden behavior, the next action on failure and correction, the result format, and acceptance conditions, and a skill carries a good example and a bad example with their reasons at every point where judgment can diverge. A model is never left to fill an empty procedure. The form of each document kind (required sections and their order) belongs to the [agent document patterns]({{link.patterns}}). Shared behavior belongs to AGENTS, shared detailed procedure to the shared skill, and specialized detailed procedure to its skill, and they reference each other. Acceptance of a document specification is proven by document checks, and model behavior by the verdicts and observations of the [C16](#C16) enforcement mechanisms.

### Conventions summary

The table below is a compressed guide for moving to each clause and is not an independent rule. After session compaction the project hook ([C16](#C16)) provides this table with its source location and sha256 as context. Detailed conditions absent from the summary still apply through their clause. When a condition, value, method, prohibition, or acceptance cannot be fixed from the summary alone, only the source text of that clause and its connected procedure is read within a fixed range. On a conflict the source text wins and the summary is corrected.

| Clause | Compressed guide |
| --- | --- |
| C00 | Preserve the goal, target, method, order, and acceptance of the latest instruction. No extension beyond the instruction, guessed values, pattern mapping of data, or arbitrary stopping. Obtain the missing fact and continue the same goal. |
| C01 | The conventions are the specification, definition, and direction controller; AGENTS is the behavior workflow; skills are field-specific execution; the patterns document is the document form. Agent documents, skills, and memory are English; project docs and polite responses are English. |
| C02 | Follow the project purpose and fixed decisions (original, stack, layer ownership, product, performance, data, repository). Check the values and conditions of the row. |
| C03 | Follow the owner of each layer and connect producer → selector → shared interface → consumer → lifetime. |
| C04 | Follow the designated domain, class, dependency direction, and the core rule table (domain structure, inline bans). Keep one owner per shared concept and register a new layer in the ownership table and the tracking node table together. |
| C05 | Match each language's function, type, and file names with the real responsibility, inputs and outputs, side effects, and lifetime. On change, connect registration keys, callers, and external consumers. |
| C06 | Preserve the original diagnostic text with its input, tool, location, and exit. Confirm real analysis input and completion, and prove omission → cause → owner → impact → correction. Suppression and concealment are forbidden. |
| C07 | Separate the start and end scope and the acceptance evidence of each work type. Run the fixed checks once after integrated implementation and rerun only failed or invalidated checks. |
| C08 | Separate the entry points of product runs and observation runs. Product runs belong to the user. |
| C09 | Each action needs a direct proposition of who, what, when, where, why, and how. Separate evidence, inference, proposal, and unverified state. |
| C10 | Follow the designated owner of compiled-layer paths, type pairs, and build targets. Never edit generated output or external code by hand. |
| C11 | The only version is the product version. Preserve originals, valid assets, user changes, and failure evidence. A hash identifies content and is not a separate version. |
| C12 | Use only registered tools and canonical entry points. Follow each stack's allowed checks and forbidden commands, hook and MCP registration, subagent delegation, and resource limits. |
| C13 | Protected originals, generated output, external code, and user changes change only through their designated route. Rule relaxation, bypass, unauthorized git, unauthorized delegation, and installation are forbidden. |
| C14 | Before a change, fill the seven steps: pipeline point, impact, per-item reason, per-clause judgment, measurement, verdict, satisfaction of every applicable clause. Judge logic by logic before running. Confirm zero forward omissions and zero backward inventions. |
| C15 | Keep atomic requirements in the internal record and read the latest summary, project nodes, and related TODOs. Record progress in {{time.label}} and write a session summary only after acceptance. Reuse valid evidence on resume. |
| C16 | Rules are enforced by the project hook, the commit guard, and the stack checks. Enforcement is reported from the mechanism's source and its observed verdict. |
| C17 | Connect the origin, schema, production, consumption, and preservation of every designated data file. Pass per-file comparison before and after regeneration and the catalog and verification gates. |
| C18 | Keep the designated role group boundaries. Never change another group's owner without an instruction. |
| C19 | The one selected execution line consumes the same meaning, resources, lifetime, and output. Platform and device support is judged only from actual observation in that environment. |
| C20 | Production diagnostics are written by one designated owner to one designated location. Preserve the original text and identifiers of system, error, and crash events. |

The scope covers the project's own code, configuration, documents, skills, and production tools, from their inputs to final consumption and release. Internal names of external dependencies, originals, and generated output keep their original source, and this standard applies to the interfaces, origins, lifetimes, and diagnostics where they meet the project's own code. The existence of code is not evidence of compliance, and a document change is no authority to rename unrelated code in bulk.

Norms are proven by instructions and contracts; implementation facts by the source, checks, and observations concerned. A declaration in a document is not evidence of implementation compliance. Mechanical enforcement of rules belongs to [C16](#C16) and is proven by the mechanism's source and verdict.

## The contract at a glance

The documentation standard below applies to the documents in the project's `docs/`. Agent documents, skill documents, the patterns document, and memory default to English, and responses and progress messages are English. Memory separates source-backed rules, observed facts, plans, and unverified claims, and states its sources and logical connections.

| Owning document | Content it owns | Language |
| --- | --- | --- |
| This document | Direction, definitions, fixed decisions, structure, tool contracts, change prohibitions, reasoning, verification, lifecycle, acceptance criteria | English |
| [AGENTS]({{link.agents}}) | The AI behavior procedures and prohibitions that carry out this document | English |
| [Agent document patterns]({{link.patterns}}) | The forms of the agent contract, entry files, process skills, domain skills, gate rows, clauses, records, and responses | English |
| [Shared workflow skill]({{link.workflowSkill}}) | Detailed execution procedures grounded in this document's clauses and the method summaries of every [C07](#C07) process | English |
| [FIX process skill]({{link.fixSkill}}) | Step-by-step procedures of the [C06](#C06) diagnostic process, warning and error FIX process, and exception-handling process. A procedure, not an independent rule | English |
| [Verification skill]({{link.verifySkill}}) | The step-by-step method of the [C14](#C14) logic-by-logic convention judgment and the fixed-check procedure. A procedure, not an independent rule | English |
| [Structure skill]({{link.structureSkill}}) and [stack skill]({{link.stackSkill}}) | Respectively the [C04](#C04), [C05](#C05), [C10](#C10) placement and naming, and the per-language change and verification method of the [C12](#C12) stack rows. Procedures, not independent rules | English |
{{generated.skillOwnershipRows}}
| Agent entries (the harness table of [AGENTS]({{link.agents}})) | Connect each main and subagent entry file, skill location, and hook registration to AGENTS and the shared skills. No independent rule | English |
| Project memory | Source locations, facts, rules, and workflow points recorded on a user instruction at the per-harness path of [C12](#C12). It never replaces the sources | English |
| [Tracking]({{link.tracking}}) | Project nodes, summaries of completed session GOALs with their comparison against the conventions, progress records, and the evidence of repeated AI failures the user ordered pinned ([C15](#C15)) | English |
| [Session audit]({{link.audit}}) | Violation evidence from past sessions and unfinished instruction TODOs | English |
| [User feedback]({{link.feedback}}) | Date, summary, and applied clauses of user feedback that set or changed a definition in this document. No quotations | English |
| [`harness.config.json`]({{link.config}}) | Harness constants: project name, kind, and primary language, output language, main agent, skill prefix and location, document paths, scratch folder, time command, stack profiles, core and inline rules, protected and forbidden paths, MCP registration, subagents, commit guard, response fields and judgment patterns | JSON |

The priority of work standards is: **project-internal Markdown (these conventions, [tracking]({{link.tracking}}), [session audit]({{link.audit}})) and the project skills come first**.

No separate state documents such as handoffs or session notes are created. The next session continues from the newest tracking summary and the session audit TODOs. Project memory is an English point record under the [C12](#C12) contract and is not the source of state.

<a id="C02"></a>
## Project purpose and fixed decisions

{{project.summary}}.

The table below holds fixed decisions; it is not a list of finished implementation. A fixed decision changes only on a new explicit user instruction, and when it changes, this table is corrected in the same task. The [user feedback]({{link.feedback}}) document summarizes the feedback that set each decision. Acceptance evidence is the connection requested source → conversion → consumption → observed result. A row not yet defined stays "(project definition)", and the judgment it covers is never decided by the executor without a user instruction.

| Area | Fixed decision |
| --- | --- |
| Original | (project definition: the primary source, the next source, and discarded sources that are not consulted. Separate origin and version, and original text and inference) |
| Stack | The primary language is {{generated.primaryLanguage}}. (project definition: runtime, framework, and build tool. Matches the profile of the [C12](#C12) stack row) |
| Layer ownership | (project definition: the meaning each layer owns. Matches the [C03](#C03) table) |
| Product | The product kind is {{generated.projectKindDoc}}. (project definition: product form and entry point. Matches the [C08](#C08) table) |
| Performance | (project definition: target environment and measured acceptance criteria. Performance claims without measurement are forbidden) |
| Data | (project definition: data root and serving path. Matches the [C17](#C17) table) |
| Repository | Large files stay out of git tracking. The repository does not use Git LFS (no LFS filter in `.gitattributes`, no LFS pointer tracking, no `git lfs` command). The commit guard ([C16](#C16)) blocks only the extensions in `harness.config.json` `commitGuard.blockedExtensions`. A commit carries exactly one message as given, keeps the author the user configured, and never adds an AI co-author line or changes the author |

<a id="C03"></a>
## Architecture and owners

Every result connects producer → selector → shared interface → consumer → lifetime (release). One owner holds the meaning of a boundary, and other layers consume that contract. The contract owner of shared data keeps size, offset, value, and lifetime identical on both sides. A source connection is not evidence of full execution success.

| Boundary | Owner and contract |
| --- | --- |
| (project definition: layer name) | (owning path and the meaning, resources, and lifetime the layer owns) |

| Shared data | Contract owner | Boundaries that must agree |
| --- | --- | --- |
| (project definition) | (owning file) | (size, fields, values, lifetime) |

<a id="C18"></a>
## Role group boundaries

The project's own files that make up an output shared by several consumers (a render, output, screen, document set, and so on) are owned by exactly one of the role groups below.

| Concept | Definition |
| --- | --- |
| Change boundary | An instruction targeting one group changes only that group's owning paths. Other groups' owning paths do not change |
| Shared group | A shared group changes only on an instruction that names that shared owner. After a change it is accepted only when the before and after results of every consuming group are observed to be the same in the same observation session |
| Cross consumption | A connection in which one group calls another group's producer is allowed only when it is listed in the cross-consumption table below. A change in the called group must not change the calling group's result |
| Enforcement | Group boundaries are enforced by this clause and the behavior procedures of AGENTS and the skills |

| Group | Owning paths | Change rule |
| --- | --- | --- |
| (project definition) | (path list; folders end with `/`) | (changed only on a user instruction naming that group) |

| Cross consumption | Calling group | Called owner | Contract |
| --- | --- | --- | --- |
| (project definition) | | | |

<a id="C19"></a>
## Execution line and platform architecture

In one run, exactly one selected execution line (render backend, runtime, host, and so on) owns the result and the resource lifetime; double or duplicate execution and mutual interference are not allowed. Every line consumes the same definitions and produces results of the same meaning. No line covers an omission with another line's output.

| Layer | Definition |
| --- | --- |
| Execution line | (project definition: the list of lines and the selection source) |
| Selection | There is one decision source. When a line is designated, only that line is tried; if it cannot be created, it fails with the reason instead of switching to another line |
| Platform | (project definition: supported operating systems and devices) |
| Acceptance | Support is judged only by actual run observation per line, device, and platform. A successful build is not evidence of support |

<a id="C20"></a>
## Production diagnostic log

| Concept | Definition |
| --- | --- |
| Owner | (project definition: the one owner that writes the diagnostic log) |
| Location | (project definition: the designated location next to the run output. Never the system temp folder or the project scratch folder) |
| Recorded events | Start and end, environment (operating system, hardware, runtime version), execution line selection and its basis ([C19](#C19)), settings application failures, input and asset errors, unhandled exceptions and crashes |
| Format | One event per line: time (UTC), session id, process, thread, level (info, warning, error, fatal), category, result code, message |
| Forbidden | Scattered hard-coded error output, forcing debug or diagnostic interfaces on in production |

<a id="C06"></a>
## Definition of FIX and exceptions

A FIX is a causal correction that connects, with direct evidence, **triggering input and time → mismatch between expected and observed values → violated contract → responsible owner → affected producers, consumers, and lifetime → correction result**. A log location, the last error, or the age of code is never assumed to be the cause.

A FIX is a fully closed procedure. A correction is confirmed only when, for every single target, both the logical and mathematical reason the correction must be right and the facts that support it (original, source, data, tool output) exist; when either is missing, the action that obtains that fact comes first. Runtime observation does not substitute for the basis that confirms a correction, and a run is never a reason to defer proof or correction or to leave them unverified.

The FIX contract applies to the whole project's code, data, UI, runtime, tools and MCP, configuration, and documents. The emergency room is an analogy for responding to a failure immediately, and anatomy is an analogy for identifying the real owner, connections, and lifetime and fixing them in their structure. No product structure is imported from the analogy. What is missing, why it went missing, and why the correction at that owner must fill it are proven with direct evidence. Arbitrary A, B, or C branches, corrections that treat a hypothesis as fact, and detours in a direction other than the conventions are forbidden. Changing direction can cut the contracts of production, selection, interface, consumption, and release and damage the whole architecture, so the disappearance of a symptom alone never judges a FIX correct. When a fact is unconfirmed, only that fact is obtained and the proof continues. Definition, purpose, and acceptance belong to this clause; mandatory behavior and connections to AGENTS; detailed diagnosis, error handling, exception handling, and bottleneck correction methods to the [FIX process skill]({{link.fixSkill}}).

Every warning and error is critical. Unused, undefined, and disconnected items are implementation omissions to trace; they are never resolved by calling them "legacy" or by deletion. The required selection, registration, call, and consumption connection is proven and restored. A definition without an original consumer records its origin and the checked scope and never gets a fake consumer.

The three processes form one flow. The diagnostic process supplies the project's real information to the diagnostic tool, obtains diagnostics, and records and classifies them; own-code diagnostics go to the FIX process, and external-code diagnostics and handling that needs a boundary contract go to the exception-handling process. A diagnostic tool receives its information from this project's build records and connections, not from generic defaults. After diagnosis, FIX, or exception handling, the work interrupted by it always resumes.

| Step | Diagnostic process |
| --- | --- |
| ① Fix the tool and entry point | Diagnostics are obtained only through the canonical entry point of a tool registered in the [C12](#C12) table: {{generated.checkCommandsDoc}}. The full output is kept under `{{paths.scratch}}/` |
| ② Input completeness | The diagnostic tool receives the target's real production, run, and check inputs. Source gets the canonical check configuration, module connections, and language and platform inputs; data gets the designated originals, schema, catalog, producers, and consumers; MCP gets the real registered schema, call arguments, and return and failure contracts; documents get the language, reference, and ownership contracts; a compiled layer builds the full translation units, build definitions, the macros the compiler defines itself, include paths, platform, and language standard from the build records and that compiler's output. Nothing is filled with generic defaults |
| ③ Confirm analysis completion | Whether the tool analyzed every target to the end is confirmed from that tool's output. A diagnostic caused by missing information (an include not found, an aborted analysis, and the unused verdicts that follow) is not a code defect but an omission of ②; the information is supplied at the owner of the tool input and the diagnosis is run again |
| ④ Record | Tool, rule id, severity, file, line, and the original message are recorded one by one as they are. They are never turned into counts or summaries |
| ⑤ Classify | Own-code diagnostics go to the FIX process; diagnostics of external code such as `vendor/` or `third-party/` ([C13](#C13)) go to the external-code row of the exception-handling process; tool-input omissions go back to ② and ③ |
| ⑥ Diagnose again | After the correction, diagnose once more with the same tool and entry point to confirm the same cause no longer produces the diagnostic. The same target is never diagnosed repeatedly from several angles ([C07](#C07)) |

| Step | Warning and error FIX process |
| --- | --- |
| ① Fix the diagnostic | One record from diagnostic process ④ with its triggering input and time is the input. A diagnostic from an aborted analysis first completes diagnostic process ③ |
| ② What is missing | Identify the omission the diagnostic points at. It is a definition, declaration, connection, input, registration, consumption, release, or contract, or information the reporting tool was not given. Prove it from the original, source, or tool output |
| ③ Why it went missing | Identify the pipeline step ([C14](#C14) ①) and the responsible owner where it went missing, and connect expected and observed values with direct evidence |
| ④ How the correction fills it | Decide the correction at the owner that fills the omission and its effect range ([C14](#C14) ②, ③), prove with evidence that the correction fills the omission of ②, then implement it |
| ⑤ Acceptance | Accept when the same diagnostic does not recur from the same cause and the fixed check of the affected path ([C07](#C07)) passes. A diagnostic that disappeared by hiding is not accepted |
| Forbidden | Ignoring, suppressing (`--suppress`, `--disable`, `NOLINT`, `#pragma warning`, inline suppression comments, `eslint-disable`, `noqa`, `type: ignore`), relaxing checks, lowering warning levels, display-level filters, hiding by removing from the check target, deleting the code the diagnostic points at, and closing with only the label "false positive", "pre-existing", or "external" |

| Case | Exception-handling process |
| --- | --- |
| An exception path in code | Before a catch, default, fallback, retry, or early return is written, record ① which external input and time fails, ② the boundary contract that defines that failure (original, specification, operating system, or API documentation), ③ to whom and how the failure is delivered and which partial state is cleaned up and released, and ④ the evidence that the handling does not hide an internal invariant violation; then implement. Handling without a boundary contract is not exception handling but concealment |
| A diagnostic's premise seems not to apply | The diagnostic is never closed as an exception. Apply diagnostic process ② and ③ and FIX process ② and ③ as they are to prove with tool output, configuration, and code what information the reporting tool lacked and why, and fix the tool configuration or run path at its owner so the tool receives it. When measurement proves the tool cannot consume that information, supply the same source's information in a form the tool consumes and record the tool limit with its measured evidence |
| A diagnostic of external code | An external-code diagnostic is also recorded through diagnostic process ④ and its cause found with FIX process ② and ③. When our code or configuration causes it, fix our owner. When the external code itself causes it, the external code is not patched, so the diagnostic and its cause are reported exactly by file, line, and rule id; the external code is never removed from the analysis to make its diagnostics disappear |

Validation of external input failures, error delivery, cancellation, partial-state cleanup, and release follow proven boundary contracts. A broken internal invariant is never hidden with an arbitrary default, guard, cast, fallback, empty catch, retry, or early return. Diagnostic suppression, log filtering, check relaxation, swallowed failures, fake exit codes, and removal of required behavior are forbidden. Pre-existing and dependency diagnostics are not exempt.

Formulas and tolerances of coordinates, units, matrices, color spaces, channels, time, interpolation, precision, and synchronization need an original, specification, or measured basis. FIX acceptance is agreement of the responsible owner and every affected connection, and error 0 and warn 0 in the fixed checks. The order of cause tracing belongs to AGENTS section 2.4, the detailed recovery procedure for tool failures to the shared workflow skill, and the execution procedures of the three processes to the [FIX process skill]({{link.fixSkill}}).

<a id="C07"></a>
## Definitions of work process types

| Type | Start scope → end scope | Acceptance evidence |
| --- | --- | --- |
| Implementation | Requirement and original input → the selected real consumption and release | Function, data, call, registration, and lifetime connections and the requested result |
| FIX | Observed violation → cause owner and affected path | Cause removal and impact checks of C06 |
| Review | Designated files, patch, or feature → evaluation of reachable violations | Location, trigger, contract, impact; review alone grants no editing authority |
| Investigation | Designated question, material, and method → the observation that answers it directly. An instruction to check or identify a defect or cause does not end at investigation; the cause correction is implemented as a FIX in the same task ([C00](#C00) instruction and checking) | Observation and inference separated; investigation completion separated from implementation completion |
| Data production | Approved input and producer → designated output and consumer | Origin, schema, references, integrity, capacity bounds |
| Reverse engineering | Explicitly approved binary and question → the evidence concerned | Authority separate from general porting; only approved methods |
| Static check | Designated source state → integrated check result | The static check commands of the [C12](#C12) stack row; not evidence of screens or behavior that were not run |
| Build | Connected change → compiled, linked, or bundled output | The build check commands of the [C12](#C12) stack row; separate from run success and support |
| Runtime observation | The session, screen, input, and time concerned → observation result | The observation route the user instructed; human input first |
| Documents and harness | Affected definitions, behavior, procedures, and records → connected documents | Language, references, terms, duplicates, conflicts, requirement comparison, [patterns]({{link.patterns}}) agreement; no application build needed |
| Cleanup and deletion | Explicitly named wrong output → processed through the approved route | Valid material, failure evidence, and unrelated changes preserved |
| Version, commit, publication | Explicit target and authority → designated output or remote result | Version source agreement and existing guards kept; no arbitrary publication |
| Resume and correction | Latest instruction and valid evidence → the remaining same requirements | Only the corrected method is replaced; no arbitrary reduction of goal or acceptance |

The verification scope is fixed by the request and the affected contracts. The source → producer → consumer and lifetime connections of every item instructed in the session GOAL are first implemented completely in code and data. Only then are static checks run, to find omissions and errors a static checker can detect in the integrated implementation and fix their causes. This check is the last detection step of the implementation order, not a substitute for implementation and not a simple compile check run after each item. Error 0 and warn 0 are results within the type and static-check scope only, not evidence of logic completeness, convention compliance, or run success. The fixed build and runtime observation run after the implementation and their preceding checks are connected. A full build is justified only by an explicit requirement or an integration contract that needs it. Builds, guesses, and micro-patch loops to discover errors, and reassurance rechecks of facts that already passed, are forbidden. After a correction, only the acceptance evidence its work type defines is confirmed and the instructed work proceeds at once; the corrected target is never retested repeatedly across several inputs, cases, or angles (no repeated testing). Only checks invalidated by a change or counterevidence are rerun. An unverified state is a real boundary, never authority to abandon work or evidence of a pass.

<a id="C14"></a>
## Reasoning and verification method

Every AI reasons and verifies the same way. Reasoning goes in one direction; verification goes in two.

Step ④ of the change decision is no authority to create arbitrary A, B, or C directions. When the user set the direction or the cause and correction are proven, that direction is kept. Comparison applies only to source-backed directions for implementation details the instruction actually leaves open, and a count of O is no proof of fact, measurement, or causality. If even one applicable clause is X or unconfirmed, that premise must be resolved, and nothing executes without the satisfaction of every clause in ⑦.

| Step | Definition |
| --- | --- |
| Sources of fact | Only instruction text lines, source and configuration files, tool output, and runtime observation are facts. Memory, summaries, past documents, and model confidence are not facts |
| Order of reasoning | Instruction text → target → owner ([C03](#C03), [C04](#C04), [C10](#C10)) → input → producer → consumer → result → lifetime ([C06](#C06)) → observation. This order is never reversed or skipped |
| Change decision procedure | Every change is carried out only by a judgment that fills all seven steps below. ① Identify exactly which step of the whole pipeline's operating order is the point where the issue is detected. ② List the entire range of effects that changes when that step changes. ③ For every item of the range, list whether this change affects it, why the change must be made that way, and, if it is wrong, why it is wrong. ④ For each approach direction judge every convention clause as O or X, and take the direction with the most O by majority. ⑤ Perform the measurement that proves that approach. ⑥ Confirm the judgment from the measured result. ⑦ Prove that the confirmed judgment satisfies every convention clause, then execute. When a step is empty, the action that obtains that step's fact comes first ([C00](#C00) basis of reasoning). Forming a hypothesis from one issue and implementing it violates this procedure |
| Logic-by-logic convention judgment | Before code or a tool runs or a result is confirmed, the files the run passes through are ordered by dependency, and from the first logic of the first file, each logic is judged OK against every clause of this document with a secured basis (instruction text, original record, measurement, official source). Only when one logic is OK does judgment move to the next; a logic that is X or lacks a basis is fixed in place after the basis is obtained (reading the source, a registered tool query, measurement) and judged again. The code or tool runs only when every executed logic is OK. Syntax and static checks are diagnostics of their static scope only, not evidence for logic judgment. A missing basis is never a reason to refuse or stop a run or the work; the action that obtains the basis comes first. Running before judgment to see the result, rerunning the same tool after a failure, and responses that replace correction and execution with admissions, apologies, or lists of violations are forbidden. The existence of code is no basis for judgment ([C01](#C01)) |
| Performance bottleneck report and immediate FIX | For each bottleneck, record why it is a bottleneck, distinct mathematical and source evidence that it must be a logic error, every related file, the correct direction of correction with its proof, and the O or X of every applicable clause. A general notion is not evidence. A bottleneck whose change decision ①–⑦ is filled so every applicable clause is O is fixed at once at its cause owner and connected consumers, and the before and after results for the same input and conditions are recorded |
| Proposition | One piece of evidence supports one proposition ([C09](#C09)). A proposition is a sentence whose truth can be judged by observation |
| Full-set requirements | An "all", "every", or "full" requirement extracts the original list in full and compares against it. The target set is never decided by enumerating outputs, a one-level glob, samples, or an arbitrary choice of fields |
| Absence claims | "None" is written only when the search scope and method are stated. One failed lookup is not evidence of absence |
| Separation of verdicts | Confirmed (proven by observation), reasoned inference (connected but not observed), and unverified are never mixed |
| Handling of violations and unfinished work | A convention violation is unaccepted unfinished work; it is never left with only a label or ended. The executable investigation, proof, correction, connection, and verification continue at once until the requirement has real acceptance evidence. Writing documents, reference checks, or presenting plans never substitutes for evidence that a product violation is resolved |
| Unused and invented | A definition with a basis in the original or the instruction but no consumer is an implementation omission and is connected. A rule, value, or feature with no basis in the original or the instruction is an invention and is replaced with the original value or removed |
| Forward verification | Confirm that every requirement of the instruction has the result and acceptance evidence its work type defines in [C07](#C07). One missing is an omission |
| Backward verification | Confirm that every change, result, and reported sentence traces back to a requirement or clause. One that does not is an invention |
| Acceptance | Accepted only when forward and backward verification both pass and the fixed checks report error 0 and warn 0 |

<a id="C15"></a>
## Work lifecycle, GOAL, and tracking

Everything the user enters in a session is that session's GOAL ([C00](#C00)). Every AI follows the order below. The tool-call procedure of each step belongs to AGENTS.

| Step | What must be done | End condition |
| --- | --- | --- |
| Session start | In a new session (`startup`, `clear`), before every other action, read these conventions in full without missing a line, then read [AGENTS]({{link.agents}}) in full. A long document is read to the end in ranges that return every line; summaries, searches, and partial reads never replace that reading. Then check {{generated.contextSourceDoc}}, the newest tracking summary and related unfinished TODOs, and read the applicable skill procedures. After session compaction (`compact`), the project hook ([C16](#C16)) unconditionally provides the conventions summary table with its source location and hash, the behavior summary location and hash, the newest tracking summary and project node table, and the unfinished TODOs as context. After compaction and on a resume of the same work (`resume`, `fork`), instead of a full reread, only whether the clauses, procedures, and owner files tied to the recent work still hold is checked within a fixed range, and work continues. Valid evidence is kept, only changed or newly needed conditions, values, prohibitions, and acceptance are checked, and an omission in a summary is not an exemption from a clause | Full reading of the conventions and AGENTS in a new session, context summary after compaction, current owning sources, applicable clauses, and remaining TODOs checked |
| Instruction receipt | Split the input into atomic requirements and keep the original text, target, and acceptance condition one row each in the work record. Input that arrives during work is merged the same way at once | Zero missing requirements |
| During work | Every action traces back to a requirement or clause. A progress report is one English line and never replaces execution. A failure is handled by a [C06](#C06) cause correction. A background job's result is collected before the next step. While work proceeds, the progress-record section of the tracking document keeps being updated with each requirement's change, evidence, and remaining work and the time. Stopping, deferring, reducing, and writing handoffs are forbidden | Implementation progress per requirement and an updated progress record |
| Verification | Perform the [C14](#C14) two-directional verification and the fixed checks of each change type in [C12](#C12) | Zero omissions, zero inventions, error 0 and warn 0 |
| Completion and record | Write a session summary to the tracking document only when every requirement of the GOAL has evidence. Before completion, only the progress record is updated instead of a session summary | Summary written |
| Comparison right after recording | Compare the written summary with this document clause by clause. If a clause was violated, redo that work and correct the summary | Zero violations |
| Response | Report in short English in the order result → evidence per requirement → unverified boundary → remaining TODOs. Both the final response and the progress responses between tool calls always state the convention clause IDs that are the basis of the action. A clause label is connected to the real basis of application and never substitutes for evidence of compliance or completion. Every final and progress response carries, for each reported action, the fields of `harness.config.json` `response.fields` ({{generated.responseFieldList}}) under their exact names (the first field holds the basis convention clause IDs and their rule), and each field states the source path, measurement, or observation that settles it. Every response contains the sentence "Based on …, I carried out …, and the evidence for this action is as follows", the numbered Evidence 1, Evidence 2, … (each with its source path, measurement, or observation), and a Therefore conclusion drawn only from them. A guess or a generality is no basis. An assistant message that calls a tool names, before the call, the clause ID and row of this document that define that tool and approach, and states the basis confirmed for that row under the matching name among {{generated.basisLabelList}} with its source path, measurement, or observation; the [C16](#C16) project hook injects this condition into the context right before every tool call ([C00](#C00) tools and approach). This is a project-wide response rule; the project hook announces it in the session-start and instruction-receipt context and blocks a final response that omits a field at the end. The tone is fixed to polite, formal English. Tone changes only the feel of a sentence and never reduces the accuracy of facts, evidence, or unverified boundaries. No completion marker without evidence. Every sentence shown to the user is English, including progress lines between tool calls and the first sentence right after reading a document, tool output, or summary in another language. The reasoning language is not restricted | Zero missing clause labels, zero missing response fields, zero missing tool-call bases, zero false completions, zero sentences shown to the user in another language |

The tracking document is also read by other contributors through AI, so every record is compressed into a one-line summary. One record is one line, and that line holds only the gist of the requirement and the target, evidence, numbers, unverified boundary, and remaining work needed to judge the result. The tracking document never quotes user instructions. A requirement is written as its gist, not as instruction text in quotation marks; the instruction text stays only in the work record ([C12](#C12)).

The progress-record format is one line per update, `- update time ({{time.label}}) | requirement gist | change | evidence | remaining work`, with the newest line on top.

The project node section of the tracking document (`## Project nodes`) manages the nodes that make up the project ([C04](#C04) domain and class folders, [C03](#C03) layers, [C10](#C10) compiled areas) one row each. The row format is `| node | path | owning clause | role | depends on | status |`, and the status is one of `planned`, `in progress`, and `verified`. A change that adds, moves, or removes a node or changes its dependencies corrects this table in the same task, and the status becomes `verified` only when the node has its [C07](#C07) acceptance evidence. After session compaction the project hook provides this table as context ([C16](#C16)).

The session summary format of the tracking document is as follows, with the newest summary at the top right below the progress-record section.

- Heading: `## YYYY-MM-DD Session GOAL summary`
- One line per requirement of the GOAL: `requirement gist | target and method | result and verification | unverified boundary`
- Last line: `Conventions comparison | compared clauses | zero violations, or the violated clauses and the rework result`

The causality of an earlier failure and a later fix is never merged away. When the user orders the tracking document compressed, already written records are compressed into this format too, preserving every dated record and the order within the same date, the separation of earlier failure and later fix, and the unverified boundaries and remaining work.

When the user orders a repeated AI failure pinned in the tracking document, that failure is recorded in the pin section of the tracking document with the behavior, time, violated clause, and result, and the prohibition rule that prevents it is added to section 6 of [AGENTS]({{link.agents}}) in the same task. The pin section is not a session summary and does not wait for the GOAL to finish.

Every time written in logs, tracking, and records is {{time.label}}, never calculated by hand or with an assumed time zone; it is the value confirmed with the [C12](#C12) time command. That command is the single source of work-log times.

The session audit TODOs are updated in only two cases.

- When the user orders TODOs organized, fixed, or completed
- When a requirement passes into the next session unfinished because of an observed external dependency, a governing constraint, or the user ending the session

The session summary in the tracking document is written only after the GOAL is complete. An ongoing GOAL accumulates in the progress-record section at the top of the tracking document with update times; the progress record never replaces the session summary, and progress lines already written are not rewritten except by the compression instruction above. The roles of the two documents do not overlap.

<a id="C04"></a>
## Structure and dependency direction

This clause defines the placement, ownership, and dependency direction of the source and the core rules. Every shared concept keeps one owner. A new domain or class folder can be created only in the change that fixes its owner, and that change corrects the tables below and the project node table of the [tracking]({{link.tracking}}) document in the same change. A folder absent from the table is a structure violation.

### Core rules

The rules below are the rules this project turned on. Their source is `harness.config.json` `rules` and `inlineRules`. A rule whose enforcement column says hook is blocked at edit time by the [C16](#C16) project hook; a rule enforced by skill is judged through its skill ([C01](#C01) skill row) and the review procedure. When a rule changes, this table and the configuration change together in the same task.

| Rule | Clause | Definition | Enforcement |
| --- | --- | --- | --- |
{{generated.ruleRows}}

| Path | Confirmed class folders |
| --- | --- |
| (project definition) | (class list) |

| Top-level path | Role | Change rule |
| --- | --- | --- |
| `{{paths.scratch}}/` | Reports, logs, intermediates, and request files written by tools, working copies of data, temporary query scripts used only when no registered tool has the capability ([C12](#C12) tool integration), project hook session state | Excluded from git; cleaned up when the step ends |
| `docs/` | The documents [C01](#C01) owns | One owning concept per document |
| `.harness/` | The shared project hook engine (`{{paths.hook}}`) and the subagent launcher ([C12](#C12)) | Protected by [C13](#C13): only changes that fix a defect |
| (project definition) | (role) | (change rule) |

| Origin domain | Allowed dependency domains |
| --- | --- |
| (project definition) | (allowed list) |

For every new third-party dependency, the tracking and indexing exclusions of its distribution, binaries, and generated output are checked in `.gitignore` and the ignore files of the applicable tools, and the needed rules are added in the same change.

<a id="C10"></a>
## Compiled layer path contract

The source of a compiled layer (a layer with native code or build targets) is unified under `{area}/{function}/`. The responsibility column states only the path name and the confirmed owner of [C03](#C03); this table alone does not prove run behavior. A new area or function folder can be created only in the change that fixes its owner, and that change corrects this table. Build output and every `vendor/` folder are protected from change under [C13](#C13). Every type has a declaration and implementation pair ([C05](#C05)).

| Path | Function subfolders | Responsibility |
| --- | --- | --- |
| (project definition) | | |

<a id="C05"></a>
## Function and file naming rules

This applies to every function the project writes: public and private functions, methods, callbacks, hooks, UI components, shader functions, and tool functions. Fixed external API names, originals, and generated output keep their original contract, and this rule applies in the project's own adapters and generators. The exception for fixed entry points never extends to ordinary functions.

| Target | Function and symbol names | Files and fixed boundaries |
| --- | --- | --- |
| TS and JS ordinary functions and methods | lowerCamelCase, action + target | Export file name = symbol name |
| TS types, interfaces, classes, UI components | PascalCase, target and responsibility | Display responsibility; ordinary logic is never disguised as a component |
| Hooks | `use` + PascalCase target | State, reactivity, and lifetime connection |
| Constants | UPPER_SNAKE_CASE | Mutable state is never named as a constant |
| C++ functions and methods | lowerCamelCase | Types PascalCase, declaration and implementation pairs; constructors, destructors, operators, and `main` follow the language contract |
| Rust functions and variables / types | snake_case / PascalCase | Module files snake_case |
| Go functions, methods, types, constants | MixedCaps. Exported names start with an upper-case letter and unexported names with a lower-case letter (the Go export rule). Constants also use MixedCaps, and this row overrides the constants row (Effective Go) | Package folders are one lower-case word (Effective Go), never meaningless package names such as `util`, `common`, or `misc` (Go blog "Package names"). Files are snake_case (project choice) |
| Python functions and methods | snake_case | Files snake_case; `main` and special methods follow the run and language contract |
| PHP functions and methods / classes | lowerCamelCase / PascalCase | Class file name = class name |
| C# methods, properties, types | PascalCase | Local variables and parameters lowerCamelCase |
| Shader functions | lowerCamelCase | Entry points follow the API contract |

A name matches its responsibility, inputs, result, side effects, and lifetime. The same domain concept uses the same term, and different concepts use distinguishable names. Files or exports that differ only in letter case, colliding aliases with the same name, and names that substitute a meaningless number, `new`, `temp`, `misc`, or `manager` for responsibility are forbidden. An unrelated rename never starts only because an existing name contains a word.

| Meaning of a name | Judgment rule |
| --- | --- |
| `get`, `read`, `find` | Matches the lookup target, location, and failure result; no hidden resource creation or saving |
| `is`, `has`, `can` | The boolean proposition matches the real return condition; possibility is separated from execution success |
| `create`/`destroy`, `attach`/`detach`, `register`/`unregister` | Matches the real lifetime contract of ownership, release, connection, and registration |
| `load`, `save`, `write`, `export` | Matches the I/O boundary, failure delivery, and persistence meaning; a change in memory is never named as a successful save |
| Calculation and conversion | Identifies the meanings that need distinguishing such as target, coordinate system, and unit; no meaning without an original basis |
| Asynchronous | Name, return type, and the caller's waiting contract agree; acceptance is never expressed as completion |

The basis of a choice is the owner, callers, and I/O contract concerned. The acceptance scope of a rename covers declarations, files, exports, imports, registration keys, selectors, and serialized or external consumers. Passing a regular expression proves only the spelling form, not the agreement of meaning.

<a id="C08"></a>
## Runtime environment and product processes

How tools are used belongs to [C12](#C12). This section defines only product processes and entry points. A product run is the user's direct run; the AI never uses a product run as a debugging means and uses only the observation route the user instructed.

| Process | Entry point and owner | Boundary |
| --- | --- | --- |
| (project definition) | (entry point) | (boundary) |

<a id="C12"></a>
## Tool usage contract

Tools are used only for the purposes and in the ways of the table below. A tool absent from the table, a method different from the table, and a location absent from the table are forbidden.

Parallel processing and bottleneck handling are not options but mandatory obligations of the convention contract. Independent explicit queries run in parallel within the fixed concurrency limit, and actions that need a preceding result keep their order. Input, result order, individual failures, resource lifetime, and the priority of human input are preserved. When a bottleneck is found in a tool, its input, conditions, time, output, and cause are preserved, and the responsible owner and affected connections are fixed at once through the [C06](#C06) FIX process. A before-and-after comparison is accepted only from measurements of the same input, conditions, and result, and a code change is never the evidence that performance improved. This obligation grants no authority to create AI subagents.

Parallel and resource contract of every tool: CPU use does not exceed 50% of the physical cores the operating system reports on the running PC. The number of workers is physical cores × 0.5 (rounded down, at least 1), and each worker runs one compute thread. RAM use does not exceed 80% of total physical memory. Before work is dispatched, when used physical memory is 80% of the total or more, nothing is dispatched until running work finishes and usage drops below that. Worker counts, thread counts, and work-split counts are never hard-coded as numbers, and static splitting is not used. Heavy work is split into independent units placed in a shared queue, and an idle worker takes the next remaining unit. Dependencies between units are guaranteed by queue order. These limits are resource bounds, and a part of the target is never counted as acceptance of the whole.

Tools are integrated by default. Production, synchronization, check, and run capabilities live inside the owning tool this table designates. When an owning tool fails or lacks a capability, the cause is fixed inside that tool and it runs through its canonical entry point. Queries, measurements, comparisons, and verification run through {{mcp.queryRoute}}. A temporary script is used only to query facts in `{{paths.scratch}}/` when no registered tool has the capability, and it never replaces a production, synchronization, or data-writing route. All of the following are bypasses and are forbidden.

- A runner that detaches an owning tool's internal functions and calls them with changed arguments, steps, or inputs
- A wrapper that skips a failing step or drops a required argument
- A separate script that plays the same role as an owning tool
- A temporary script that replaces a query, measurement, or comparison a registered tool provides

Every file in a tool folder belongs to one connected production and check system. Each file is a canonical command, an MCP server registration, a [C17](#C17) production tool, a tool such a tool calls as a step, or a module these import. No standalone file connected to nothing and no temporary file for a query, check, analysis, or one-off correction is kept.

| Tool | Purpose | Method | Input and output location | Forbidden |
| --- | --- | --- | --- | --- |
| Stack verification | Static and build verification per language stack | Only the stacks of the stack table below whose marker files all exist at the project root are active. Marker judgment looks only at file existence and never decides a stack from a directory name. The allowed check commands of an active stack are the canonical entry points, and acceptance is each check's acceptance condition (exit 0, required line present, forbidden line absent) with error 0 and warn 0 | Full diagnostics kept under `{{paths.scratch}}/` | Running before implementation or per item, bypassing through individual checkers, claiming the result as logic, convention, or run completeness, using a forbidden command as verification |
| MCP server registration | Claude and Codex start the same MCP servers with the same commands | The single registration source is `harness.config.json` `mcp.servers`, from which the setup tool (SimpleHarness CLI) generates `.mcp.json` (Claude), `enabledMcpjsonServers` of `.claude/settings.json`, and `.codex/config.toml` (Codex). Currently registered: {{generated.mcpServersDoc}}. Registration uses only repository-relative commands, and after the registration or the PC's runtime environment changes, the harness is restarted to start the servers anew | `.mcp.json`, `.claude/settings.json`, `.codex/config.toml` | Writing personal user, drive, or repository absolute paths into shared registration, designating executables through environment variables or fixed paths, automatic dependency installation, bypassing server internal functions |
| Shell | Running commands | The shell of the running host. On Windows it is PowerShell; CMD, `npm.cmd`, and batch files are not used. On Linux it is bash. Executable locations are obtained by an actual command lookup; a missing tool is never installed arbitrarily or replaced with CMD | Project root | CMD and aliases or wrappers that produce CMD, assuming Bash on a Windows host, assuming Windows paths or PowerShell on a Linux host, guessed executable paths |
| File read, search, edit | Checking and changing source and documents | The harness's read, search, and edit tools. An existing file is edited in part (Claude Code `Edit`, an Update File hunk of Codex `apply_patch`), and only new files are created (Claude Code `Write`, an Add File hunk of Codex `apply_patch`). When the target is not yet known, only a bounded search whose search root and scope, search proposition, sought fact, and stop condition are fixed first | Owning paths ([C04](#C04), [C10](#C10)) | Overwriting an existing file with a script or by delete-and-recreate, repository searches without scope or stop condition |
| Work record | Tracking instruction requirements row by row | TodoWrite or the harness equivalent; harness-internal state that creates no file | Inside the harness | Work logs, handoffs, or session notes as files |
| Time | Confirming the {{time.label}} time used in work logs and records | `{{time.command}}` | Standard output | Assuming a time zone or calculating a log time by hand |
| Project memory | Memory records the user instructed | Claude adds one small update note to the project memory folder the harness assigns to this repository; Codex to the memory update path the harness assigns. The sources of rules and state are this document, tracking, and the session audit; memory holds only their locations and points | Per-harness memory location | Recording without a user instruction, content that replaces the source documents, promotion to global behavior rules |
| Product run | The user's direct run | [C08](#C08) entry point | — | Use as the AI's run-debugging means |
| Runtime observation | Observation of screens, input, and logs after the connected code and data and the fixed checks and build are ready | (project definition: observation tools and order). A session starts only when the user instructs it, and production before observation (build, pack, generator) runs only when its input changed since its last success. Observe → input → observe again | Observation output under `{{paths.scratch}}/` | Sending input while a human operates, starting or ending a session without a user instruction, adding temporary instrumentation to product code, running after every single change |
| Reverse engineering | Explicitly approved reverse engineering | (project definition: approved tools) | `{{paths.scratch}}/` | Use without approval, full dump output |
| git | Work for which the user named git, and publication work the user allowed | For named work only the named commands run, and a commit message is exactly the one message the user gave. No AI co-author line (`Co-authored-by`, `Generated with`, `Co with AI`, and the like) goes into a commit message, and the author or committer is never changed and no collaborator added | — | git calls outside named or publication work (including status, diff, log, grep, restore), unrequested commits, pushes, or history changes, force pushes, adding collaborators, AI co-author lines |
| Subagents | The main agent ({{generated.mainAgentName}}) delegates research, review, and drafts | Only an id registered in the subagent table below runs, through `node {{paths.subagentLauncher}} <id> <task file>`. The task file lives inside `{{paths.scratch}}/`. A subagent run is read-only (the Claude runtime without edit tools, the Codex runtime in a read-only sandbox). A subagent result is an observation before verification, and the main agent uses it only after checking each claim against the source or measurement. The harness's built-in subagent tool is used only for delegation whose allowed count and role the user set, and it stops at once when the user says so | Run logs under `{{paths.scratch}}/subagent/<id>/` | Running a subagent absent from the table, delegating edits to a subagent, integrating unverified results, writing an API key into a file or command line |
| Background run | Long production and checks | The job handle is kept, and its completion and per-file results are collected before dependent work proceeds | Logs under `{{paths.scratch}}/` | Proceeding with dependent work before collection, polling wait loops, foreground waiting commands |
| Version and dependencies | Version changes and package changes | `{{version.command}}` and package manager commands | `{{version.source}}` | Hand-editing manifests or lockfiles, downgrading a version to avoid a compatibility problem |
| External documentation | Official documentation needed for the work | Only official sources are consulted | — | Lookups that lead to installation or cloning |
| Project hook | Applies the clauses of this document mechanically to the main agent's tool calls and response endings. Subagent input (`agent_id`) gets no verdict and no state change | {{generated.hookRegistrationDoc}} runs the shared engine `{{paths.hook}}`. The source of rules is this document; the hook shows the clause ID and configuration key as its reason, and the constants belong to the root `harness.config.json` ([C16](#C16)) | `.harness/hooks/`, session state `{{paths.scratch}}/hook/<agent>/` | Bypassing the hook, retrying a denied call with another tool, relaxing the hook without a user instruction |
| Shared hook path resolution | Makes the hook behave identically on any PC | No personal user name, drive, or fixed absolute path appears anywhere in hook settings or hook scripts. The project scope derives from the file's own location: the engine from `import.meta.url` → repository root, the Claude Code registration string from `CLAUDE_PROJECT_DIR`, and the Codex registration string from `git rev-parse --show-toplevel`. A registration string is one `node -e` line without shell expansion, so it runs the same in Git Bash, PowerShell, and cmd.exe (Codex's Windows hook shell), and its last argument passes the agent name (`claude`, `codex`). Minimal blocking is the principle: blocking keeps only [C13](#C13) and the narrow list the user designated. A failure of the hook itself leaves its original text in `{{paths.scratch}}/hook/failure/` and ends without a verdict. After a hook change, the registration string is tested as it is with a standard-input payload, the verdict is observed, and the agent is restarted | Agent hook registration files, `.harness/hooks/`, `harness.config.json` | Recording personal paths, turning warning items back into blocks, mistaking a hook error for a pass, claiming a configuration change took effect without a test |

### Stack table

The source of the stack profiles is `harness.config.json` `stacks`. This table carries those values; when a profile changes, it is generated again with the setup tool (SimpleHarness CLI) or this table is corrected in the same task.

| Stack | Markers (all must exist) | Extensions | Allowed check commands and acceptance | Commands that must not be used as verification |
| --- | --- | --- | --- | --- |
{{generated.stackRows}}

### Subagent table

The source is `harness.config.json` `subagents`. The API key column is the name of the environment variable that holds the key; the key value is recorded nowhere. For the Claude runtime the launcher passes the provider endpoint, model, and key through a temporary settings file it creates for each run and deletes after the run.

| id | Runtime | Endpoint | Model | API key |
| --- | --- | --- | --- | --- |
{{generated.subagentRows}}

<a id="C13"></a>
## Change prohibitions

The targets below change only through the allowed route of the table. A change through any other route is a violation.

| Target | Forbidden | Allowed route |
| --- | --- | --- |
{{generated.protectedRows}}
| This document, AGENTS, skills, the patterns document, `harness.config.json` | Adding notes or session records, changing rules or constants without a user instruction | User instruction |
| `.harness/`, agent hook registration files, `.githooks/`, the commit guard, check configuration | Rule relaxation, check exclusion, bypass | Only changes that fix a defect |
| [C02](#C02) fixed decisions | Changes without a user instruction | A new explicit user instruction |
| Owning paths of a [C18](#C18) role group | Changes during work that targets another group | A user instruction that names that group |
| The user's uncommitted changes and other sessions' work | Reverting, overwriting | Preserved |
| Unused, undefined, or disconnected code | Deletion, classification as legacy | Implement the producer → consumer connection ([C06](#C06)). Temporary files are not kept under the [C12](#C12) tool connection rule |
| Production tools | Bypass runners, wrappers, or duplicate scripts to avoid a failure, skipping failing steps | Fix the cause inside that tool and run it through its canonical entry point ([C12](#C12)) |
| git history | Rewriting, force push, hook bypass | User instruction |

<a id="C09"></a>
## Acceptance criteria for evidence, actions, and records

Every action needs a direct proposition in which **who, what, when, where, why, and how** are identifiable. Authority for an action connects to instructions and contracts, and judgments of fact to the source, configuration, and observation concerned. One piece of evidence corresponds 1:1 to the one proposition it actually supports. When several premises are needed, each connection is proven separately. Set difference, subsets, coverage ratios, majority, and confidence never replace direct evidence. Majority is used only to choose the approach direction in step ④ of the [C14](#C14) change decision procedure and never replaces the measurement that follows or the proof that every clause is satisfied.

A true verdict needs the applicable condition, causal connection, and confirmed result; a false verdict needs a real counterexample or a violated requirement. "None" without a basis, using historical documents as evidence of the current implementation, and calling a model's choice an industry standard are forbidden. A proposition is a statement whose truth can be judged by observation.

The start gate of an action is authority, target, input, method, and existing evidence; the end gate is result, consumption connection, and fixed verification. This applies to the AI's judgments, tool use, changes, recovery, and completion; it does not mean attaching labels to response sentences. The order of actions belongs to AGENTS, and the detailed procedure to the shared workflow skill.

Corrections, errors, and exceptions never erase existing requirements or higher rules. While executable approved work remains, the work of the same goal continues. Self-made prohibitions or stop conditions, arbitrary goal reduction, and replacing execution with only apologies, plans, or status reports are forbidden. Real authority, external dependencies, and governing constraints apply as they are and are never resolved by false completion or bypassing a prohibition. Ending a response with a label such as an incompleteness proof while requirements remain is forbidden; the remaining requirements continue in the same goal. Replacing implementation with document work that only lists blanks in plans or documents or quotes user instructions is also forbidden, and logic confirmed by evidence is implemented in the source in the same task.

The timing and format of tracking and the response method belong to [C15](#C15). No session log is written in this document, AGENTS, or the skills.

<a id="C11"></a>
## Versions and material preservation

The only version concept is the product version. Every other version field, version number, and version comparison (schema, catalog, format, protocol, file record) is forbidden, and a layout is identified by its fields and sizes. Changing a version value never substitutes for removing a field. The source of the product version is `{{version.source}}`, and every version constant another layer uses equals this one value. The change tool is `{{version.command}}`, and a commit never changes the version arbitrarily. No switch (command-line flag, option field, setting, environment variable, mode argument) selects between an old and a new path or between alternative implementations of one architecture decision. A confirmed architecture has exactly one path. Asset commits, force pushes, history rewrites, and guard bypasses are forbidden. Valid originals, production assets, user changes, and failure evidence are preserved outside an explicit cleanup scope. The license of each material owns its rights.

<a id="C16"></a>
## Enforcement mechanisms

Rules are enforced by the mechanical enforcement mechanisms below. Each mechanism blocks a tool call, a response ending, or a commit at its event, or injects the condition that must hold at that point into the model's context. Enforcement is reported from the mechanism's source and its observed verdict, never from the fact that a document or setting was changed. Rules these mechanisms do not judge are carried out through the AGENTS and skill procedures, and when the same violation repeats, the rule moves into these mechanisms on a user instruction.

| Mechanism | What it does | How it enforces |
| --- | --- | --- |
| Stack check commands ([C12](#C12) stack table) | Static and build diagnostics of the active stacks | A result whose exit or acceptance condition fails is not accepted. The post-tool hook records the checks left after an edit, and the stop hook announces them so the next action continues |
| `.githooks/pre-commit` → [commit guard]({{link.commitGuard}}) | Commits of files with the extensions in `harness.config.json` `commitGuard.blockedExtensions` | Blocks the commit with exit 1. Connected with `git config core.hooksPath .githooks`, it applies to every agent's and every person's commits alike |
| Project hook: {{generated.hookRegistrationDoc}} → [run_project_hook.mjs]({{link.hook}}) | The start hook branches on the input `source`: for a new session it adds full reading of the conventions and AGENTS; for compaction it adds the conventions summary table, the behavior summary, the newest tracking summary, the project node table, and the unfinished TODOs with their source locations and sha256 within the context limit (10,000 characters); for a resume it adds the check of the recent work ([C15](#C15)). The instruction-receipt hook injects the conditions of requirement preservation, tool basis, skill loading, and response fields. The pre-tool hook blocks question and planning tools, Claude Code's Bash tool on Windows, edits of protected paths ([C13](#C13)), `Write` or `Add File` on an existing file, delete-and-recreate, an increase of comments in authored code ([C04](#C04) `no-inline-comment`), an increase of an enabled inline rule ([C04](#C04) core rule table), forbidden paths, CMD and batch routes, forbidden commands of an active stack, the denied command list, foreground waits, harness scratch paths, writes to the system temp folder, shell writes to protected folders, stopping a job when the latest user instruction has no stop instruction, and changes or runs whose required skills are not loaded, and for every call it lets through it injects the tool basis condition ([C00](#C00) tools and approach, [C14](#C14), [C15](#C15)). The post-tool hook records skill loading (a Claude Code Skill tool call, or a successful tool call that read the skill's `SKILL.md`), marks the checks left according to the edited extensions, and clears the checks left by earlier edits only when a check command's exit and acceptance condition match. The stop hook blocks unloaded mandatory skills, forbidden-judgment phrases, user-level gate markers, acknowledgment or report phrases without a report instruction, nonexistent clause IDs, missing response or basis fields, field values without a basis, nonexistent cited paths or lines, guessing expressions, and a missing basis statement, numbered evidence, or conclusion, and when checks are left it announces them so the work continues. The judgment phrases and labels belong to `harness.config.json` `response` per output language | The pre-tool and stop hooks block the tool call or the response ending with exit 2 and return the reason to the model (the official Claude Code and Codex hook contract). The start, instruction-receipt, pre-tool, and post-tool hooks inject conditions through `additionalContext`. After a registration file changes, the agent is restarted for the new registration to apply, and Codex runs project hooks after the project is trusted and the hooks are reviewed in `/hooks` |
| `.mcp.json`, `.codex/config.toml` | Connection of the registered MCP servers | Only registered servers start, with the same commands ([C12](#C12) MCP server registration row) |

- Violations of rules these mechanisms do not judge are recorded in the [session audit]({{link.audit}}) with their evidence.
- When the same violation repeats, the rule moves into these mechanisms on a user instruction.

<a id="C17"></a>
## Data catalog and schema

This clause owns the placement, recording, preservation, consumption, and structure of the data the project designates. A project without data confirms the definition rows below as "(no designated data)", and even then the [C00](#C00) principle of manual designation of data applies as it is to producing and consuming code.

| Item | Contract |
| --- | --- |
| Data root | (project definition: data root path and its domain and class composition) |
| Catalog | (project definition: the catalog recording each file's path, size, hash, and preservation unit, and its production tool) |
| Schema | (project definition: the schema holding the shape, production tool, and consumers of each pattern) |
| Preservation | Losing a recorded file or a recorded unit is a violation. Only elements proven wrong are allowed through the explicit correction route, and additions and hash changes are recorded as updates. While any violation exists, the catalog is not written |
| Update procedure | Production tool run → catalog gate passed → verification gate passed. A copy of build output is never written back as an original |

| Pattern | Shape | Production | Consumption | Key fields |
| --- | --- | --- | --- | --- |
| (project definition) | | | | |
