<div align="center">

# SimpleHarness

*把 AI 编码代理约束在项目规则之内的 harness 模板与多语言配置 CLI*

*约定 → AGENTS → 技能 → 钩子 — Claude Code · Codex 作为主代理，Ollama · Qwen · Kimi · GLM · DeepSeek · MiniMax · OpenRouter 作为子代理*

![lang](https://img.shields.io/badge/lang-ko%20%7C%20en%20%7C%20zh-0ea5e9)
[![license](https://img.shields.io/badge/license-MIT-a855f7)](LICENSE)
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

[한국어](README.md) · [English](README.en.md) · **简体中文**

[![START](https://img.shields.io/badge/DOCS-START-8A2BE2?style=for-the-badge&logo=markdown&logoColor=white)](docs/getting-started.md)
[![ENVIRONMENT](https://img.shields.io/badge/DOCS-ENVIRONMENT-7c3aed?style=for-the-badge&logo=markdown&logoColor=white)](docs/environment.md)
[![AGENTS](https://img.shields.io/badge/DOCS-AGENTS-38bdf8?style=for-the-badge&logo=markdown&logoColor=white)](docs/agents.md)
[![HOOKS](https://img.shields.io/badge/DOCS-HOOKS-8A2BE2?style=for-the-badge&logo=markdown&logoColor=white)](docs/hooks.md)
[![EDIT TOOLS](https://img.shields.io/badge/DOCS-EDIT%20TOOLS-7c3aed?style=for-the-badge&logo=markdown&logoColor=white)](docs/edit-tools.md)
[![CODE REVIEW](https://img.shields.io/badge/DOCS-CODE%20REVIEW-38bdf8?style=for-the-badge&logo=markdown&logoColor=white)](docs/code-review.md)
[![LANGUAGES](https://img.shields.io/badge/DOCS-LANGUAGES-8A2BE2?style=for-the-badge&logo=markdown&logoColor=white)](docs/languages/README.md)

*`docs/` 下的指南以韩语编写。*

</div>

---

## SimpleHarness 是什么

AI 编码代理在会话结束后会忘记该会话中学到的规则，文档没有写到的部分则由它自行推测填补。**SimpleHarness** 是模板加 CLI：回答几个问题，它就把适合你仓库的 harness（约定、代理契约、技能、钩子、提交守卫）汇集到一个文件夹中。把这个文件夹拖入仓库根目录后，每个会话都从同一套规则开始，违反规则的工具调用与回复会在发生的那一刻被钩子拦截。

| 组成 | 角色 | 位置 |
| --- | --- | --- |
| 代理契约 | **大脑** — 规定请求 → 判断 → 执行 → 验证的顺序与禁止行为 | `AGENTS.md`（+ `CLAUDE.md`） |
| 钩子 | **控制器** — 在每次工具调用前后与每次回复结束时判定；拦截违规并注入条件 | `.harness/hooks/run_project_hook.mjs` |
| 技能 | **四肢** — 每种行为的详细流程、规格、逻辑、正例与反例 | `.claude/skills/` 或 `.agents/skills/` |
| 约定 | **方向声明** — 项目的目的、定义、已确定决定、结构与工具契约（C00–C20） | `docs/conventions.md` |
| AI 跟踪 | **跟踪** — 进度记录、已完成会话摘要、项目节点管理 | `docs/ai-tracking.md` |

## 强制方式

规则通过 Claude Code 与 Codex 的官方钩子契约执行。钩子以退出码 2 拦截工具调用或回复结束，并把理由（条款编号 + 配置键）返回给模型。

| 时机 | 钩子的动作 |
| --- | --- |
| 会话开始 | 新会话注入通读约定与 AGENTS 全文的条件；压缩后注入摘要表、跟踪、项目节点与未完成 TODO，并附原文位置与 sha256 |
| 提交指令 | 注入需求保留、工具依据、技能加载与回复格式的条件 |
| 工具调用前 | **拦截**提问工具、受保护路径的编辑、覆盖现有文件、新增注释 / 内联类型 / 样式 / 处理器、禁止命令、cmd 路径、写入临时路径以及缺少必要技能 |
| 工具调用后 | 记录已加载的技能，编辑后标记待完成的技术栈检查，检查通过时清除 |
| 回复结束 | **拦截**未加载必需技能、禁止的判断用语、引用不存在的条款或文件、缺少回复字段或依据、推测表达；仍有待完成检查时让工作继续 |
| 提交 | `.githooks/pre-commit` 以退出码 1 停止包含被拦截扩展名文件的提交 |

## 快速开始

准备：[Go](https://go.dev/dl/)（运行 CLI）、[Node.js](https://nodejs.org)（运行钩子）、[Git](https://git-scm.com/downloads)，以及 Claude Code 或 Codex。

```sh
git clone https://github.com/GarnetRapture/SimpleEngineering.git
cd SimpleEngineering
go run ./src
```

CLI 依次询问：

1. **输出语言** — 한국어 · English · 简体中文
2. **环境检查** — 识别 Windows 10/11 · macOS · Linux · WSL，逐步引导替换 shell（PowerShell 7 / Git Bash）、安装 Git、更改默认终端
3. **项目** — 名称、摘要、技能前缀、类型（Web / 通用程序）、主要语言
4. **代理** — 主代理（Claude Code / Codex）与任意数量的子代理
5. **核心规则** — 基于领域的结构与命名；禁止内联注释、类型、样式、处理器与魔法字面量（默认全部开启）
6. **生成** — 结果汇集到 `output/<项目名称>/`
7. **推荐工具** — 按项目类型与语言可选安装工具（jq、ripgrep、sqlite3、DuckDB、LLVM、GCC、工具链……）
8. **后续步骤** — 显示在屏幕上并保存到 `output/<项目名称>-SETUP.md`

然后：

1. 把 `output/<项目名称>/` 中的**所有**文件与文件夹（包括 `.claude`、`.codex`、`.agents`、`.harness`、`.githooks`）拖入仓库根目录。
2. 运行 `git config core.hooksPath .githooks`，并把 `tmp-agent/` 加入 `.gitignore`。
3. 填写 `docs/conventions.md` 中以 `（项目定义` 开头的行。
4. 重启代理并用 `/hooks` 确认登记。在 Codex 中信任该项目并批准钩子。

## 生成结果

```text
<仓库>/
├─ AGENTS.md                         代理契约（大脑）
├─ CLAUDE.md                         Claude Code 入口（@AGENTS.md）
├─ harness.config.json               harness 常量（钩子、守卫、启动器每次运行时读取）
├─ docs/
│  ├─ conventions.md                 约定 C00–C20（方向声明）
│  ├─ agent-patterns.md              文档与技能模式 P1–P10
│  ├─ ai-tracking.md                 项目节点、进度、会话摘要
│  ├─ session-audit.md               违规证据、未完成 TODO
│  └─ user-feedback.md               反馈摘要
├─ .claude/ 或 .codex/                钩子登记（settings.json / hooks.json）、MCP 登记
├─ .claude/skills/ 或 .agents/skills/
│  └─ <前缀>-workflow · fix-process · verify · structure · stack · mcp · naming · inline · subagent
├─ .harness/
│  ├─ hooks/run_project_hook.mjs     共享钩子引擎
│  └─ subagents/run_subagent.mjs     子代理启动器
├─ .githooks/pre-commit              提交守卫入口
└─ tools/guard/staged_guard.mjs      提交守卫
```

`naming` 与 `inline` 技能在对应规则开启时生成；`subagent` 技能与启动器在选择子代理时生成。

## 代理

| 主代理 | 入口 | 技能 | 钩子登记 |
| --- | --- | --- | --- |
| Claude Code | `CLAUDE.md` → `@AGENTS.md` | `.claude/skills/`（Skill 工具） | `.claude/settings.json` — 6 个事件 |
| Codex | `AGENTS.md` | `.agents/skills/`（读取 SKILL.md） | `.codex/hooks.json` — 5 个事件（信任 + `/hooks` 批准） |

子代理是主代理通过 `node .harness/subagents/run_subagent.mjs <id> <任务文件>` 运行的**只读**助手。可以任意组合另一个主 CLI（Codex / Claude Code）、本地 Ollama、Ollama Cloud、阿里云百炼（Coding Plan · Token Plan，Qwen）、Moonshot Kimi、Z.ai GLM、DeepSeek、MiniMax 与 OpenRouter。API 密钥只放在环境变量中；启动器通过每次运行创建的临时设置文件传入，运行后删除。子代理的回答是未经验证的观察，主代理只使用对照原文确认过的内容。

## 核心规则

| 规则 | 条款 | 强制方式 |
| --- | --- | --- |
| 基于领域的结构 `{领域}/{类别}/{对象}/{名称}` | C04 | 技能 |
| 基于领域的命名（禁止 `util`、`helper`、`manager`） | C05 | 技能 |
| 禁止内联注释 — 意图由名称与结构表达 | C04 | 钩子 |
| 禁止内联类型 — 只放在 `types/` 或 `*.types.ts`（TypeScript） | C04 | 钩子 |
| 禁止内联样式（TypeScript、PHP） | C04 | 钩子 |
| 禁止内联事件处理器（TypeScript、PHP） | C05 | 钩子 |
| 禁止魔法字面量 — 使用具名常量 | C04 | 技能 |

## 语言技术栈

| 语言 | 标记文件 | 检查 | 禁止 |
| --- | --- | --- | --- |
| TypeScript | `package.json` + `tsconfig.json` | `npx tsc --noEmit`、`npx eslint . --max-warnings 0` | `npm run dev`、`npm start`、`vite`、`npm run build` |
| Python | `pyproject.toml` | `ruff check .`、`ruff format --check .`、`pyright --warnings` | — |
| C++ | `xmake.lua` | `xmake build`（无警告行） | — |
| Rust | `Cargo.toml` | `cargo clippy --all-targets -- -D warnings` | `cargo build`、`cargo run` |
| Go | `go.mod` | `go vet ./...`、`gofmt -l .` | `go run`、`go build` |
| PHP | `composer.json` | `php -l <文件>`、`vendor/bin/phpstan analyse` | `php -S` |

## 仓库结构

```text
SimpleEngineering/
├─ go.mod                  module simpleharness（go 1.26）
├─ src/
│  ├─ main.go              嵌入模板并运行向导
│  ├─ internal/
│  │  ├─ wizard/           提问流程、环境引导、后续步骤
│  │  ├─ system/           操作系统、Windows 版本号、包管理器、MSVC 检测
│  │  ├─ catalog/          语言、技术栈、代理、提供商、规则、工具、回复判定
│  │  ├─ blueprint/        回答 → harness.config.json
│  │  ├─ generate/         清单生成、钩子与 MCP 登记、文档片段
│  │  ├─ render/           占位符渲染
│  │  ├─ install/          推荐工具安装（winget · brew · apt）
│  │  ├─ console/ locale/  输入界面、ko · en · zh 消息
│  │  └─ harness/          配置结构体
│  └─ template/            生成的文档、技能、钩子与守卫的原文
└─ docs/                   本 README 链接的指南（韩语）
```

## 指南（韩语）

| 指南 | 内容 |
| --- | --- |
| [入门](docs/getting-started.md) | harness 与约定的基础、快速开始、术语表 |
| [环境准备](docs/environment.md) | Windows 10/11、macOS、Linux、WSL 的 shell、Git 与终端；工具安装命令 |
| [代理](docs/agents.md) | Claude Code 与 Codex 作为主代理；子代理提供商与启动器原理 |
| [钩子](docs/hooks.md) | 各事件的判定、登记命令、在 `harness.config.json` 中修改规则、测试方法 |
| [编辑工具](docs/edit-tools.md) | Edit、Write、NotebookEdit、apply_patch 与 shell 写入的契约和行为定义 |
| [代码评审](docs/code-review.md) | 以约定的目的与定义为依据的逻辑化、科学化评审 |
| [语言](docs/languages/README.md) | C++（xmake）、Rust、Go、TypeScript、Python、PHP |

## 许可证

[MIT](LICENSE) — 生成的 harness 文件也可以自由修改和分发。
