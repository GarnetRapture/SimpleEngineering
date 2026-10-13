<div align="center">

# SimpleHarness · 심플하네스

*AI 코딩 에이전트를 프로젝트 규칙에 묶는 하네스 템플릿과 다국어 설정 CLI*

*컨벤션 → AGENTS → 스킬 → 훅 — Claude Code · Codex 메인, Ollama · Qwen · Kimi · GLM · DeepSeek · MiniMax · OpenRouter 서브*

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

**한국어** · [English](README.en.md) · [简体中文](README.zh-CN.md)

[![START](https://img.shields.io/badge/DOCS-START-8A2BE2?style=for-the-badge&logo=markdown&logoColor=white)](docs/getting-started.md)
[![ENVIRONMENT](https://img.shields.io/badge/DOCS-ENVIRONMENT-7c3aed?style=for-the-badge&logo=markdown&logoColor=white)](docs/environment.md)
[![AGENTS](https://img.shields.io/badge/DOCS-AGENTS-38bdf8?style=for-the-badge&logo=markdown&logoColor=white)](docs/agents.md)
[![HOOKS](https://img.shields.io/badge/DOCS-HOOKS-8A2BE2?style=for-the-badge&logo=markdown&logoColor=white)](docs/hooks.md)
[![EDIT TOOLS](https://img.shields.io/badge/DOCS-EDIT%20TOOLS-7c3aed?style=for-the-badge&logo=markdown&logoColor=white)](docs/edit-tools.md)
[![CODE REVIEW](https://img.shields.io/badge/DOCS-CODE%20REVIEW-38bdf8?style=for-the-badge&logo=markdown&logoColor=white)](docs/code-review.md)
[![LANGUAGES](https://img.shields.io/badge/DOCS-LANGUAGES-8A2BE2?style=for-the-badge&logo=markdown&logoColor=white)](docs/languages/README.md)

</div>

---

## SimpleHarness란

AI 코딩 에이전트는 세션이 끝나면 그 세션에서 배운 규칙을 기억하지 못하고, 문서에 없는 부분은 스스로 추측해서 채워요. **SimpleHarness** 는 질문 몇 개에 답하면 내 저장소에 맞는 하네스(컨벤션·에이전트 문서·스킬·훅·커밋 가드)를 한 폴더로 만들어 주는 템플릿 + CLI예요. 결과 폴더를 저장소 루트에 끌어다 놓으면, 에이전트는 매 세션 같은 규칙으로 시작하고 규칙을 어기는 도구 호출과 답변은 훅이 그 자리에서 막아요.

| 구성 | 역할 | 위치 |
| --- | --- | --- |
| 에이전트 문서 | **뇌** — 요청 → 판단 → 실행 → 검증 순서와 금지 행동을 정한 행동 계약 | `AGENTS.md` (+ `CLAUDE.md`) |
| 훅 | **제어자** — 도구 호출 직전·직후와 답변 종료 때마다 판정하고, 위반은 차단, 조건은 주입 | `.harness/hooks/run_project_hook.mjs` |
| 스킬 | **팔다리** — 행동마다의 상세 절차·명세·논리·좋은 예시·나쁜 예시 | `.claude/skills/` 또는 `.agents/skills/` |
| 컨벤션 | **방향 선언** — 프로젝트의 목적·정의·확정 결정·구조·도구 계약 (C00~C20) | `docs/conventions.md` |
| AI 트래킹 | **추적** — 진행 기록, 완료 세션 요약, 프로젝트 노드 관리 | `docs/ai-tracking.md` |

## 강제 방식

규칙은 문서에 적히는 것으로 끝나지 않고, 공식 훅 계약으로 실행돼요. 훅은 종료 코드 2로 도구 호출·답변 종료를 차단하고, 차단 이유(조항 번호 + 설정 키)를 모델에게 돌려줘요.

| 시점 | 훅이 하는 일 |
| --- | --- |
| 세션 시작 | 새 세션은 컨벤션·AGENTS 전문 통독 조건을, 압축 후에는 요약표·트래킹·프로젝트 노드·열린 TODO를 원문 위치와 sha256과 함께 주입 |
| 지시 입력 | 요구 보존, 도구 근거, 스킬 적재, 응답 형식 조건 주입 |
| 도구 호출 직전 | 질문 도구, 보호 경로 편집, 기존 파일 덮어쓰기, 주석·인라인 타입·스타일·핸들러 추가, 금지 명령, cmd 경로, 임시 경로 쓰기, 필요한 스킬 미적재를 **차단** |
| 도구 호출 직후 | 스킬 적재 기록, 편집 후 남은 스택 검사 표시, 검사 통과 시 해제 |
| 답변 종료 | 필수 스킬 미적재, 금지 판단 표현, 없는 조항·파일 인용, 응답 필드·근거 누락, 추측 표현을 **차단**. 남은 검사가 있으면 계속 진행시킴 |
| 커밋 | `.githooks/pre-commit` 이 막힌 확장자 파일의 커밋을 종료 1로 중지 |

자세한 판정 순서와 시험 방법은 [훅 설정](docs/hooks.md)에 있어요.

## 빠른 시작

준비물: [Go](https://go.dev/dl/) (CLI 실행), [Node.js](https://nodejs.org) (훅 실행), [Git](https://git-scm.com/downloads), Claude Code 또는 Codex. 운영체제별 설치는 [환경 준비](docs/environment.md)를 보세요.

```sh
git clone https://github.com/GarnetRapture/SimpleEngineering.git
cd SimpleEngineering
go run ./src
```

CLI가 순서대로 물어봐요.

1. **출력 언어** — 한국어 · English · 简体中文
2. **환경 점검** — Windows 10/11 · macOS · Linux · WSL을 감지해 셸 교체(PowerShell 7 / Git Bash), Git 설치, 기본 터미널 변경을 단계별로 안내
3. **프로젝트** — 이름, 요약, 스킬 접두사, 종류(웹 / 일반 프로그램), 대표 언어
4. **에이전트** — 메인(Claude Code / Codex)과 서브 에이전트(여러 개)
5. **핵심 규칙** — 도메인 기반 구조·명명, 인라인 주석·타입·스타일·핸들러·매직 리터럴 금지 (기본 모두 켜짐)
6. **생성** — `output/<프로젝트 이름>/` 에 결과물 수집
7. **추천 도구** — 프로젝트 종류·언어에 맞는 도구(jq, ripgrep, sqlite3, DuckDB, LLVM, GCC, 툴체인…) 선택 설치
8. **다음 단계** — 남은 설정을 화면과 `output/<프로젝트 이름>-SETUP.md` 로 안내

그다음:

1. `output/<프로젝트 이름>/` 안의 **모든** 파일·폴더(`.claude`, `.codex`, `.agents`, `.harness`, `.githooks` 포함)를 내 저장소 루트로 끌어다 놓아요.
2. `git config core.hooksPath .githooks` 로 커밋 가드를 연결하고, `.gitignore` 에 `tmp-agent/` 를 추가해요.
3. `docs/conventions.md` 의 `(프로젝트 정의` 행을 채워요.
4. 에이전트를 다시 시작하고 `/hooks` 로 등록을 확인해요. Codex는 프로젝트를 신뢰(trust)하고 훅을 승인해요.

## 생성 결과

```text
<저장소>/
├─ AGENTS.md                         행동 계약 (뇌)
├─ CLAUDE.md                         Claude Code 진입 (@AGENTS.md)
├─ harness.config.json               하네스 상수 (훅·가드·실행기가 실행마다 읽음)
├─ docs/
│  ├─ conventions.md                 컨벤션 C00~C20 (방향 선언)
│  ├─ agent-patterns.md              문서·스킬 양식 P1~P10
│  ├─ ai-tracking.md                 프로젝트 노드·진행 기록·세션 요약
│  ├─ session-audit.md               위반 증거·미완료 TODO
│  └─ user-feedback.md               피드백 요약
├─ .claude/ 또는 .codex/              훅 등록 (settings.json / hooks.json), MCP 등록
├─ .claude/skills/ 또는 .agents/skills/
│  └─ <접두사>-workflow · fix-process · verify · structure · stack · mcp · naming · inline · subagent
├─ .harness/
│  ├─ hooks/run_project_hook.mjs     공용 훅 엔진
│  └─ subagents/run_subagent.mjs     서브 에이전트 실행기
├─ .githooks/pre-commit              커밋 가드 진입
└─ tools/guard/staged_guard.mjs      커밋 가드
```

`naming`·`inline` 스킬은 해당 규칙을 켰을 때, `subagent` 스킬과 실행기는 서브 에이전트를 골랐을 때 생성돼요.

## 에이전트

| 메인 | 진입 | 스킬 | 훅 등록 |
| --- | --- | --- | --- |
| Claude Code | `CLAUDE.md` → `@AGENTS.md` | `.claude/skills/` (Skill 도구) | `.claude/settings.json` — 6개 이벤트 |
| Codex | `AGENTS.md` | `.agents/skills/` (SKILL.md 읽기) | `.codex/hooks.json` — 5개 이벤트 (신뢰 + `/hooks` 승인) |

서브 에이전트는 메인이 `node .harness/subagents/run_subagent.mjs <id> <작업 파일>` 로 부르는 **읽기 전용** 조력자예요. 다른 메인 CLI(Codex / Claude Code), 로컬 Ollama, Ollama Cloud, Alibaba Cloud Model Studio(Coding Plan · Token Plan, Qwen), Moonshot Kimi, Z.ai GLM, DeepSeek, MiniMax, OpenRouter를 여러 개 함께 붙일 수 있어요. API 키는 환경 변수에만 두고, 실행기가 실행마다 임시 설정 파일로 넘긴 뒤 지워요. 서브의 답은 검증 전 관찰이라 메인이 원문으로 확인한 것만 써요. → [에이전트별 적용](docs/agents.md)

## 핵심 규칙

| 규칙 | 조항 | 강제 |
| --- | --- | --- |
| 도메인 기반 구조 `{도메인}/{클래스}/{대상}/{이름}` | C04 | 스킬 |
| 도메인 기반 명명 (`util`·`helper`·`manager` 금지) | C05 | 스킬 |
| 인라인 주석 금지 — 의도는 이름과 구조로 | C04 | 훅 |
| 인라인 타입 금지 — `types/`·`*.types.ts` 에만 (TypeScript) | C04 | 훅 |
| 인라인 스타일 금지 (TypeScript·PHP) | C04 | 훅 |
| 인라인 이벤트 핸들러 금지 (TypeScript·PHP) | C05 | 훅 |
| 매직 리터럴 금지 — 이름 있는 상수로 | C04 | 스킬 |

## 언어 스택

| 언어 | 표지 파일 | 검사 | 금지 |
| --- | --- | --- | --- |
| TypeScript | `package.json` + `tsconfig.json` | `npx tsc --noEmit`, `npx eslint . --max-warnings 0` | `npm run dev`, `npm start`, `vite`, `npm run build` |
| Python | `pyproject.toml` | `ruff check .`, `ruff format --check .`, `pyright --warnings` | — |
| C++ | `xmake.lua` | `xmake build` (경고 줄 0) | — |
| Rust | `Cargo.toml` | `cargo clippy --all-targets -- -D warnings` | `cargo build`, `cargo run` |
| Go | `go.mod` | `go vet ./...`, `gofmt -l .` | `go run`, `go build` |
| PHP | `composer.json` | `php -l <파일>`, `vendor/bin/phpstan analyse` | `php -S` |

## 이 저장소의 구조

```text
SimpleEngineering/
├─ go.mod                  module simpleharness (go 1.26)
├─ src/
│  ├─ main.go              템플릿 임베드 + 마법사 실행
│  ├─ internal/
│  │  ├─ wizard/           질문 흐름·환경 안내·다음 단계
│  │  ├─ system/           운영체제·Windows 빌드·패키지 매니저·MSVC 감지
│  │  ├─ catalog/          언어·스택·에이전트·제공자·규칙·도구·응답 판정 목록
│  │  ├─ blueprint/        답변 → harness.config.json 구성
│  │  ├─ generate/         매니페스트 생성, 훅·MCP 등록, 문서 조각
│  │  ├─ render/           자리표시자 치환
│  │  ├─ install/          추천 도구 설치 (winget · brew · apt)
│  │  ├─ console/ locale/  입력 UI, ko·en·zh 메시지
│  │  └─ harness/          설정 구조체
│  └─ template/            생성되는 문서·스킬·훅·가드 원문
└─ docs/                   이 README의 가이드 문서
```

## 문서

| 문서 | 내용 |
| --- | --- |
| [시작하기](docs/getting-started.md) | 하네스·컨벤션의 기초, 빠른 시작, 용어집 |
| [환경 준비](docs/environment.md) | Windows 10/11·macOS·Linux·WSL 셸·Git·터미널, 도구 설치 명령 |
| [에이전트별 적용](docs/agents.md) | Claude Code·Codex 메인, 서브 에이전트 제공자와 실행 원리 |
| [훅 설정](docs/hooks.md) | 이벤트별 판정, 등록 명령, `harness.config.json` 으로 바꾸기, 시험 방법 |
| [편집 도구 정의](docs/edit-tools.md) | Edit·Write·NotebookEdit·apply_patch·셸 쓰기의 계약과 행동 정의 |
| [코드 리뷰 방법론](docs/code-review.md) | 컨벤션의 목적·정의에 근거한 논리적·과학적 리뷰 |
| [언어별 핵심 지식](docs/languages/README.md) | C++(xmake)·Rust·Go·TypeScript·Python·PHP |

## 라이선스

[MIT](LICENSE) — 생성된 하네스 파일도 자유롭게 고치고 배포할 수 있어요.

