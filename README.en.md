<div align="center">

# SimpleHarness

*A harness template and multilingual setup CLI that binds AI coding agents to your project's rules*

*Conventions → AGENTS → skills → hooks — Claude Code · Codex as the main agent, Ollama · Qwen · Kimi · GLM · DeepSeek · MiniMax · OpenRouter as subagents*

![lang](https://img.shields.io/badge/lang-ko%20%7C%20en%20%7C%20zh-0ea5e9)
[![license](https://img.shields.io/badge/license-MIT-a855f7)](LICENSE)
[![release](https://img.shields.io/github/v/release/GarnetRapture/SimpleEngineering?color=22c55e)](https://github.com/GarnetRapture/SimpleEngineering/releases/latest)
![type](https://img.shields.io/badge/type-template%20%2B%20CLI-7c3aed)
![CLI](https://img.shields.io/badge/CLI-Go%201.26-00ADD8?logo=go&logoColor=white)
![hooks](https://img.shields.io/badge/hooks-Node.js-5FA04E?logo=nodedotjs&logoColor=white)
![platform](https://img.shields.io/badge/platform-Windows%20%C2%B7%20macOS%20%C2%B7%20Linux-6b7280)

![Claude Code](https://img.shields.io/badge/Claude%20Code-main%20%C2%B7%20sub-D97757?logo=claude&logoColor=white)
![Codex](https://img.shields.io/badge/Codex-main%20%C2%B7%20sub-111111)
![hook events](https://img.shields.io/badge/hook%20events-SessionStart%20%E2%86%92%20Stop-374151)
![MCP](https://img.shields.io/badge/MCP-.mcp.json%20%C2%B7%20config.toml-4b5563)

![Ollama](https://img.shields.io/badge/Ollama-local%20%C2%B7%20cloud-000000?logo=ollama&logoColor=white)
![Alibaba Model Studio](https://img.shields.io/badge/Alibaba%20Model%20Studio-Qwen-FF6A00?logo=alibabacloud&logoColor=white)
![Kimi](https://img.shields.io/badge/Kimi-Moonshot-1f2937)
![GLM](https://img.shields.io/badge/GLM-Z.ai-2563eb)
![DeepSeek](https://img.shields.io/badge/DeepSeek-Anthropic%20API-4D6BFE?logo=deepseek&logoColor=white)
![MiniMax](https://img.shields.io/badge/MiniMax-Anthropic%20API-e11d48?logo=minimax&logoColor=white)
![OpenRouter](https://img.shields.io/badge/OpenRouter-multi%20provider-6467F2?logo=openrouter&logoColor=white)

![TypeScript](https://img.shields.io/badge/TypeScript-tsc%20%C2%B7%20eslint-3178C6?logo=typescript&logoColor=white)
![Python](https://img.shields.io/badge/Python-ruff%20%C2%B7%20pyright-3776AB?logo=python&logoColor=white)
![C++](https://img.shields.io/badge/C%2B%2B-xmake-00599C?logo=cplusplus&logoColor=white)
![Rust](https://img.shields.io/badge/Rust-clippy-000000?logo=rust&logoColor=white)
![Go](https://img.shields.io/badge/Go-vet%20%C2%B7%20gofmt-00ADD8?logo=go&logoColor=white)
![PHP](https://img.shields.io/badge/PHP-phpstan-777BB4?logo=php&logoColor=white)

[한국어](README.md) · **English** · [简体中文](README.zh-CN.md)

[![DOWNLOAD](https://img.shields.io/badge/CLI-DOWNLOAD-22c55e?style=for-the-badge&logo=github&logoColor=white)](#download)

[![START](https://img.shields.io/badge/DOCS-START-8A2BE2?style=for-the-badge&logo=markdown&logoColor=white)](docs/getting-started.md)
[![ENVIRONMENT](https://img.shields.io/badge/DOCS-ENVIRONMENT-7c3aed?style=for-the-badge&logo=markdown&logoColor=white)](docs/environment.md)
[![AGENTS](https://img.shields.io/badge/DOCS-AGENTS-38bdf8?style=for-the-badge&logo=markdown&logoColor=white)](docs/agents.md)
[![HOOKS](https://img.shields.io/badge/DOCS-HOOKS-8A2BE2?style=for-the-badge&logo=markdown&logoColor=white)](docs/hooks.md)
[![EDIT TOOLS](https://img.shields.io/badge/DOCS-EDIT%20TOOLS-7c3aed?style=for-the-badge&logo=markdown&logoColor=white)](docs/edit-tools.md)
[![CODE REVIEW](https://img.shields.io/badge/DOCS-CODE%20REVIEW-38bdf8?style=for-the-badge&logo=markdown&logoColor=white)](docs/code-review.md)
[![LANGUAGES](https://img.shields.io/badge/DOCS-LANGUAGES-8A2BE2?style=for-the-badge&logo=markdown&logoColor=white)](docs/languages/README.md)

*The guides under `docs/` are written in Korean.*

</div>

---

## What SimpleHarness is

An AI coding agent forgets the rules it learned once a session ends, and it fills anything the documents leave out with its own guess. **SimpleHarness** is a template plus a CLI: answer a few questions and it assembles a harness for your repository (conventions, agent contract, skills, hooks, commit guard) in one folder. Drag that folder into your repository root and every session starts from the same rules, while the hooks block a tool call or an answer that breaks them at the moment it happens.

| Part | Role | Location |
| --- | --- | --- |
| Agent contract | **Brain** — the order of request → judgment → execution → verification and the forbidden behaviors | `AGENTS.md` (+ `CLAUDE.md`) |
| Hooks | **Controller** — judges every tool call before and after it runs and every answer before it ends; blocks violations and injects conditions | `.harness/hooks/run_project_hook.mjs` |
| Skills | **Limbs** — the detailed procedure, specification, logic, good examples, and bad examples of each behavior | `.claude/skills/` or `.agents/skills/` |
| Conventions | **Direction** — the project's purpose, definitions, settled decisions, structure, and tool contracts (C00–C20) | `docs/conventions.md` |
| AI tracking | **Tracking** — progress records, completed-session summaries, project node management | `docs/ai-tracking.md` |

## How rules are enforced

The rules run through the official hook contracts of Claude Code and Codex. A hook blocks a tool call or the end of an answer with exit code 2 and returns the reason (clause ID + config key) to the model.

| Moment | What the hook does |
| --- | --- |
| Session start | A new session gets the condition to read the conventions and AGENTS in full; after compaction it gets the summary table, tracking, project nodes, and open TODOs with source lines and sha256 |
| Prompt submit | Injects requirement retention, tool basis, skill loading, and response-format conditions |
| Before a tool call | **Blocks** question tools, protected-path edits, overwriting existing files, added comments / inline types / styles / handlers, forbidden commands, the cmd route, temp-path writes, and missing required skills |
| After a tool call | Records loaded skills, marks stack checks pending after edits, clears them when a check passes |
| End of answer | **Blocks** unloaded mandatory skills, forbidden judgments, citations of nonexistent clauses or files, missing response fields or evidence, and guesses; keeps work going while checks remain |
| Commit | `.githooks/pre-commit` stops a commit containing blocked extensions with exit 1 |

## Download

No Go installation is needed. Get the one executable that matches your computer from the [latest release](https://github.com/GarnetRapture/SimpleEngineering/releases/latest).

| OS | CPU | File |
| --- | --- | --- |
| Windows | x64 (Intel · AMD) | [simpleharness-windows-amd64.exe](https://github.com/GarnetRapture/SimpleEngineering/releases/latest/download/simpleharness-windows-amd64.exe) |
| Windows | ARM64 | [simpleharness-windows-arm64.exe](https://github.com/GarnetRapture/SimpleEngineering/releases/latest/download/simpleharness-windows-arm64.exe) |
| macOS | Apple Silicon (M1 and later) | [simpleharness-darwin-arm64](https://github.com/GarnetRapture/SimpleEngineering/releases/latest/download/simpleharness-darwin-arm64) |
| macOS | Intel | [simpleharness-darwin-amd64](https://github.com/GarnetRapture/SimpleEngineering/releases/latest/download/simpleharness-darwin-amd64) |
| Linux | x64 | [simpleharness-linux-amd64](https://github.com/GarnetRapture/SimpleEngineering/releases/latest/download/simpleharness-linux-amd64) |
| Linux | ARM64 | [simpleharness-linux-arm64](https://github.com/GarnetRapture/SimpleEngineering/releases/latest/download/simpleharness-linux-arm64) |
| Checksums | — | [SHA256SUMS.txt](https://github.com/GarnetRapture/SimpleEngineering/releases/latest/download/SHA256SUMS.txt) |

Which CPU: Windows PowerShell `$env:PROCESSOR_ARCHITECTURE` (`AMD64` → amd64, `ARM64` → arm64); macOS and Linux `uname -m` (`x86_64` → amd64, `arm64` or `aarch64` → arm64).

### Run

Put the executable in the folder where the result should go, open a terminal there, and run it. `output/<project name>/` and `output/<project name>-SETUP.md` appear in that folder.

**Windows** (PowerShell 7 or Windows Terminal)

```powershell
.\simpleharness-windows-amd64.exe
```

- Double-clicking it in Explorer also works. The window closes when it finishes, and the result and the guide stay in `output/` next to the executable.
- The executable is unsigned; if a "Windows protected your PC" window appears, choose **More info → Run anyway**.

**macOS**

```sh
chmod +x simpleharness-darwin-arm64
xattr -d com.apple.quarantine simpleharness-darwin-arm64
./simpleharness-darwin-arm64
```

A file downloaded through a browser carries the quarantine attribute, and Gatekeeper blocks it; `xattr -d` removes that attribute.

**Linux**

```sh
chmod +x simpleharness-linux-amd64
./simpleharness-linux-amd64
```

**Checksum** — the value must equal the one in `SHA256SUMS.txt`.

| OS | Command |
| --- | --- |
| Windows | `Get-FileHash .\simpleharness-windows-amd64.exe -Algorithm SHA256` |
| macOS | `shasum -a 256 simpleharness-darwin-arm64` |
| Linux | `sha256sum -c SHA256SUMS.txt --ignore-missing` |

### Run from source

With [Go](https://go.dev/dl/) 1.26 or later you can run it straight from the repository or build it yourself.

```sh
git clone https://github.com/GarnetRapture/SimpleEngineering.git
cd SimpleEngineering
go run ./src
```

Build: Windows `go build -o simpleharness.exe ./src`, macOS and Linux `go build -o simpleharness ./src`.

### Making a release (maintainers)

Pushing a tag that starts with `v` runs [release.yml](.github/workflows/release.yml): after `gofmt` and `go vet` it builds the six executables and `SHA256SUMS.txt` and publishes them as a GitHub release.

```sh
git tag v0.1.0
git push origin v0.1.0
```

The Windows executable icon is linked from `src/rsrc_windows_amd64.syso` and `src/rsrc_windows_arm64.syso`. After changing the icon (`src/assets/launcher.ico`) or the resource script (`src/assets/launcher.rc`), run `src/assets/build_windows_resources.ps1` on Windows to regenerate both files and commit them together (requires GNU windres).

## Quick start

Requirements: [Node.js](https://nodejs.org) (runs the hooks), [Git](https://git-scm.com/downloads), and Claude Code or Codex.

Run the [downloaded](#download) executable (or [run from source](#run-from-source)) in the folder where the result should go; the CLI asks, in order:

1. **Output language** — 한국어 · English · 简体中文
2. **Environment check** — detects Windows 10/11 · macOS · Linux · WSL and walks through replacing the shell (PowerShell 7 / Git Bash), installing Git, and changing the default terminal
3. **Project** — name, summary, skill prefix, kind (web / general program), primary language
4. **Agents** — the main agent (Claude Code / Codex) and any number of subagents
5. **Core rules** — domain-based structure and naming; no inline comments, types, styles, handlers, or magic literals (all on by default)
6. **Generate** — collects the result in `output/<project name>/`
7. **Recommended tools** — optional installation of tools for the project kind and language (jq, ripgrep, sqlite3, DuckDB, LLVM, GCC, toolchains…)
8. **Next steps** — printed and saved to `output/<project name>-SETUP.md`

Then:

1. Drag **every** file and folder inside `output/<project name>/` (including `.claude`, `.codex`, `.agents`, `.harness`, `.githooks`) into your repository root.
2. Run `git config core.hooksPath .githooks` and add `tmp-agent/` to `.gitignore`.
3. Fill the rows of `docs/conventions.md` that start with `(project definition`.
4. Restart the agent and check the registration with `/hooks`. In Codex, trust the project and approve the hooks.

## Generated files

```text
<repository>/
├─ AGENTS.md                         agent contract (brain)
├─ CLAUDE.md                         Claude Code entry (@AGENTS.md)
├─ harness.config.json               harness constants (read on every hook, guard, and launcher run)
├─ docs/
│  ├─ conventions.md                 conventions C00–C20 (direction)
│  ├─ agent-patterns.md              document and skill patterns P1–P10
│  ├─ ai-tracking.md                 project nodes, progress, session summaries
│  ├─ session-audit.md               violation evidence, open TODOs
│  └─ user-feedback.md               feedback summaries
├─ .claude/ or .codex/               hook registration (settings.json / hooks.json), MCP registration
├─ .claude/skills/ or .agents/skills/
│  └─ <prefix>-workflow · fix-process · verify · structure · stack · mcp · naming · inline · subagent
├─ .harness/
│  ├─ hooks/run_project_hook.mjs     shared hook engine
│  └─ subagents/run_subagent.mjs     subagent launcher
├─ .githooks/pre-commit              commit guard entry
└─ tools/guard/staged_guard.mjs      commit guard
```

The `naming` and `inline` skills are generated when their rules are on; the `subagent` skill and the launcher when subagents are selected.

## Agents

| Main | Entry | Skills | Hook registration |
| --- | --- | --- | --- |
| Claude Code | `CLAUDE.md` → `@AGENTS.md` | `.claude/skills/` (Skill tool) | `.claude/settings.json` — 6 events |
| Codex | `AGENTS.md` | `.agents/skills/` (read SKILL.md) | `.codex/hooks.json` — 5 events (trust + `/hooks` approval) |

Subagents are **read-only** helpers the main agent runs with `node .harness/subagents/run_subagent.mjs <id> <task-file>`. Any combination of the other main CLI (Codex / Claude Code), local Ollama, Ollama Cloud, Alibaba Cloud Model Studio (Coding Plan · Token Plan, Qwen), Moonshot Kimi, Z.ai GLM, DeepSeek, MiniMax, and OpenRouter can be attached. API keys stay in environment variables; the launcher passes them through a per-run temporary settings file and deletes it afterwards. A subagent answer is an unverified observation, and the main agent uses only what it confirms against the sources.

## Core rules

| Rule | Clause | Enforced by |
| --- | --- | --- |
| Domain-based structure `{domain}/{class}/{target}/{name}` | C04 | skill |
| Domain-based naming (no `util`, `helper`, `manager`) | C05 | skill |
| No inline comments — intent lives in names and structure | C04 | hook |
| No inline types — only in `types/` or `*.types.ts` (TypeScript) | C04 | hook |
| No inline styles (TypeScript, PHP) | C04 | hook |
| No inline event handlers (TypeScript, PHP) | C05 | hook |
| No magic literals — named constants | C04 | skill |

## Language stacks

| Language | Markers | Checks | Forbidden |
| --- | --- | --- | --- |
| TypeScript | `package.json` + `tsconfig.json` | `npx tsc --noEmit`, `npx eslint . --max-warnings 0` | `npm run dev`, `npm start`, `vite`, `npm run build` |
| Python | `pyproject.toml` | `ruff check .`, `ruff format --check .`, `pyright --warnings` | — |
| C++ | `xmake.lua` | `xmake build` (no warning lines) | — |
| Rust | `Cargo.toml` | `cargo clippy --all-targets -- -D warnings` | `cargo build`, `cargo run` |
| Go | `go.mod` | `go vet ./...`, `gofmt -l .` | `go run`, `go build` |
| PHP | `composer.json` | `php -l <file>`, `vendor/bin/phpstan analyse` | `php -S` |

## Repository layout

```text
SimpleEngineering/
├─ go.mod                  module simpleharness (go 1.26)
├─ .github/
│  ├─ workflows/release.yml tag push → build the six executables and publish the release
│  └─ release-notes.md     release body
├─ src/
│  ├─ main.go              embeds the template and runs the wizard
│  ├─ rsrc_windows_*.syso  Windows executable icon resources (amd64, arm64)
│  ├─ assets/              launcher.ico, launcher.rc, build_windows_resources.ps1
│  ├─ internal/
│  │  ├─ wizard/           question flow, environment guidance, next steps
│  │  ├─ system/           OS, Windows build, package manager, MSVC detection
│  │  ├─ catalog/          languages, stacks, agents, providers, rules, tools, response patterns
│  │  ├─ blueprint/        answers → harness.config.json
│  │  ├─ generate/         manifest generation, hook and MCP registration, document fragments
│  │  ├─ render/           placeholder rendering
│  │  ├─ install/          recommended tool installation (winget · brew · apt)
│  │  ├─ console/ locale/  input UI, ko · en · zh messages
│  │  └─ harness/          configuration structs
│  └─ template/            source of the generated documents, skills, hooks, and guard
└─ docs/                   the guides linked from this README (Korean)
```

## Guides (Korean)

| Guide | Contents |
| --- | --- |
| [Getting started](docs/getting-started.md) | Harness and convention basics, quick start, glossary |
| [Environment](docs/environment.md) | Shell, Git, and terminal on Windows 10/11, macOS, Linux, WSL; tool install commands |
| [Agents](docs/agents.md) | Claude Code and Codex as the main agent; subagent providers and how the launcher works |
| [Hooks](docs/hooks.md) | Per-event judgments, the registration command, changing rules in `harness.config.json`, testing |
| [Edit tools](docs/edit-tools.md) | Contracts and behavior definitions of Edit, Write, NotebookEdit, apply_patch, and shell writes |
| [Code review](docs/code-review.md) | Logical, scientific review grounded in the conventions' purpose and definitions |
| [Languages](docs/languages/README.md) | C++ (xmake), Rust, Go, TypeScript, Python, PHP |

## License

[MIT](LICENSE) — you may modify and distribute the generated harness files freely.
