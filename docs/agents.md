# 에이전트별 적용 — 메인 에이전트와 서브 에이전트

[← 환경 준비](environment.md) · 다음: [훅 설정](hooks.md)

SimpleHarness는 **메인 에이전트 하나**(Claude Code 또는 Codex)를 하네스로 묶고, 필요하면 **서브 에이전트 여러 개**를 읽기 전용 조력자로 붙여요. 메인은 사용자와 대화하며 저장소를 고치고, 서브는 메인이 맡긴 조사·검토·초안을 돌려줘요.

## 1. 메인 에이전트

두 에이전트는 같은 `AGENTS.md`, 같은 컨벤션, 같은 훅 엔진(`.harness/hooks/run_project_hook.mjs`)을 써요. 다른 것은 파일 위치와 훅 등록 방식뿐이에요.

| 항목 | Claude Code | Codex |
| --- | --- | --- |
| 진입 문서 | `CLAUDE.md` (마지막 줄 `@AGENTS.md` 로 AGENTS를 함께 읽음) | `AGENTS.md` (저장소 에이전트 문서로 직접 읽음) |
| 스킬 위치 | `.claude/skills/<접두사>-<이름>/SKILL.md` | `.agents/skills/<접두사>-<이름>/SKILL.md` |
| 스킬 적재 | `Skill` 도구 호출 | 해당 `SKILL.md` 파일 읽기 |
| 훅 등록 파일 | `.claude/settings.json` | `.codex/hooks.json` |
| 등록 이벤트 | SessionStart, UserPromptSubmit, PreToolUse, PostToolUse, PostToolUseFailure, Stop | SessionStart, UserPromptSubmit, PreToolUse, PostToolUse, Stop |
| 훅이 아는 편집 도구 | `Edit`, `Write`, `NotebookEdit` | `apply_patch` |
| 훅이 아는 셸 도구 | `PowerShell`, `Bash` | `Bash` (셸 명령 도구 이름) |
| MCP 등록 | `.mcp.json` + `.claude/settings.json` `enabledMcpjsonServers` | `.codex/config.toml` `[mcp_servers.<이름>]` |
| 훅 실행 셸 (Windows) | Git Bash (없으면 PowerShell) | `cmd.exe /C` |

### 1.1 Claude Code를 메인으로

1. 설치: `winget install --id Anthropic.ClaudeCode -e --source winget` · `brew install --cask claude-code` · Linux `curl -fsSL https://claude.ai/install.sh | bash`
2. 생성물을 저장소 루트에 복사하고 저장소 루트에서 `claude` 를 실행해요(또는 VS Code 확장·데스크톱 앱에서 저장소 폴더를 열어요).
3. `/hooks` 를 입력해 6개 이벤트(SessionStart·UserPromptSubmit·PreToolUse·PostToolUse·PostToolUseFailure·Stop)에 `node -e "…" claude` 명령이 등록됐는지 확인해요.
4. 등록 파일을 바꿨다면 Claude Code를 다시 시작해요.

Windows에서는 하네스가 Claude Code의 `Bash` 도구를 차단하고 `PowerShell` 도구로 명령을 실행하게 해요(C12 셸 행). `Bash` 도구는 Git Bash를 띄우는데, 프로젝트 명령 경로는 PowerShell 하나로 정해져 있기 때문이에요.

### 1.2 Codex를 메인으로

1. 설치: `winget install --id OpenAI.Codex -e --source winget` · `brew install --cask codex` · Linux `npm install -g @openai/codex`
2. 생성물을 저장소 루트에 복사해요. 훅 등록 명령이 `git rev-parse --show-toplevel` 로 저장소 루트를 찾으므로 **git 저장소 안**이어야 해요.
3. Codex에서 이 프로젝트를 **신뢰(trust)** 해요. 신뢰하지 않은 프로젝트의 `.codex/` 설정과 훅은 읽히지 않아요.
4. `/hooks` 에서 프로젝트 훅을 검토하고 승인해요.
5. Windows 10에서 Codex 앱이 기본 셸을 cmd로 잡는 문제는 [환경 준비 2.2](environment.md#22-windows-10)를 보세요.

### 1.3 두 에이전트를 모두 쓰고 싶을 때

메인은 하나만 고르고, 다른 하나는 서브 에이전트로 붙여요. Claude Code가 메인이면 `Codex CLI` 를, Codex가 메인이면 `Claude Code CLI` 를 서브로 고를 수 있어요.

## 2. 서브 에이전트

### 2.1 실행 원리

서브 에이전트는 메인이 셸에서 실행기를 불러 돌려요.

```sh
node .harness/subagents/run_subagent.mjs <id> <작업 파일>
```

| 규칙 | 실제 동작 (`run_subagent.mjs`) |
| --- | --- |
| 작업 파일 위치 | `tmp-agent/` 안이어야 해요. 밖이면 사용법 오류로 끝나요 |
| 읽기 전용 | Claude 런타임은 `--disallowedTools Edit Write NotebookEdit`, Codex 런타임은 `--sandbox read-only` 로 실행해요 |
| 세션 분리 | Claude 런타임은 `--no-session-persistence --setting-sources user` 로 프로젝트 훅 없이 실행해요. Codex 런타임은 `-c features.hooks=false` 로 훅을 끄고 실행해요 |
| API 키 | 키는 파일·명령줄에 쓰지 않아요. 등록된 환경 변수에서 읽어 실행마다 만든 임시 설정 파일(권한 0600)로 넘기고, 실행이 끝나면 지워요 |
| 기록 | `tmp-agent/subagent/<id>/<시각>.log` 에 명령·종료 코드·출력이 남아요 |
| 결과 표시 | `[subagent <id> — unverified result; …]` 머리말과 함께 출력해요. 메인은 결과의 주장 하나하나를 원문·실측으로 확인한 뒤에만 써요(C12 서브 에이전트 행) |

### 2.2 런타임

| 런타임 | 실행 명령 | 쓰는 제공자 |
| --- | --- | --- |
| Claude 런타임 | `claude -p <작업> --output-format json …` + 제공자 엔드포인트(`ANTHROPIC_BASE_URL`)와 모델 | Claude Code CLI, Ollama(메인이 Claude일 때), Ollama Cloud, Alibaba, Kimi, Z.ai, DeepSeek, MiniMax, OpenRouter |
| Codex 런타임 | `codex exec --sandbox read-only --json -o <답변 파일> …` | Codex CLI, Ollama(메인이 Codex일 때 `--oss --local-provider ollama`) |

Anthropic 호환 엔드포인트를 제공하는 회사라면 Claude Code 프로그램을 그 회사 모델로 돌릴 수 있어요. 실행기는 그 엔드포인트·인증 변수·모델 이름을 Claude Code 설정의 `env` 로 넘겨요.

### 2.3 제공자 목록

| CLI 선택지 | 엔드포인트 | 기본 모델 | 키 환경 변수 | 모델 목록 |
| --- | --- | --- | --- | --- |
| Codex CLI | (Codex 로그인) | — | — | — |
| Claude Code CLI | (Claude 로그인) | — | — | — |
| Ollama (로컬) | `http://localhost:11434` | `qwen3.5` | 없음 (`ANTHROPIC_AUTH_TOKEN=ollama`) | https://ollama.com/search |
| Ollama Cloud | `https://ollama.com` | `glm-5.3-flash` | `OLLAMA_API_KEY` | https://ollama.com/search?c=cloud |
| Alibaba Model Studio — Coding Plan | 국제 `https://coding-intl.dashscope.aliyuncs.com/apps/anthropic` · 중국 `https://coding.dashscope.aliyuncs.com/apps/anthropic` | `qwen3.7-plus` | `DASHSCOPE_CODING_PLAN_API_KEY` (`sk-sp-` 키) | https://www.alibabacloud.com/help/en/model-studio/coding-plan |
| Alibaba Model Studio — Token Plan | `https://token-plan.ap-southeast-1.maas.aliyuncs.com/apps/anthropic` | (직접 입력) | `DASHSCOPE_TOKEN_PLAN_API_KEY` | https://www.alibabacloud.com/help/en/model-studio/claude-code |
| Moonshot Kimi | 국제 `https://api.moonshot.ai/anthropic` · 중국 `https://api.moonshot.cn/anthropic` | `kimi-k3[1m]` | `MOONSHOT_API_KEY` | https://platform.kimi.ai/docs/guide/claude-code-kimi |
| Z.ai GLM Coding Plan | `https://api.z.ai/api/anthropic` | `glm-5.3[1m]` | `ZAI_API_KEY` | https://docs.z.ai/devpack/tool/claude |
| DeepSeek | `https://api.deepseek.com/anthropic` | `deepseek-flash` | `DEEPSEEK_API_KEY` (`ANTHROPIC_API_KEY` 로 전달) | https://api-docs.deepseek.com/guides/anthropic_api |
| MiniMax | 국제 `https://api.minimax.io/anthropic` · 중국 `https://api.minimaxi.com/anthropic` | `MiniMax-M3` | `MINIMAX_API_KEY` | https://platform.minimax.io/docs/guides/text-ai-coding-tools |
| OpenRouter | `https://openrouter.ai/api` | (직접 입력) | `OPENROUTER_API_KEY` | https://openrouter.ai/models |

모델 이름은 CLI가 물어볼 때 바꿀 수 있어요. 제공자는 모델을 자주 바꾸니, 위 "모델 목록" 링크에서 현재 이름을 확인하고 입력하세요.

### 2.4 설정 순서

1. CLI의 에이전트 단계에서 서브 에이전트를 번호로 골라요(여러 개 가능). 지역·모델·키 환경 변수 이름을 물어보면 답해요.
2. API 키를 환경 변수에 넣어요. 새 터미널부터 적용돼요.

   ```powershell
   [Environment]::SetEnvironmentVariable('MOONSHOT_API_KEY', '<키>', 'User')
   ```

   ```sh
   echo 'export MOONSHOT_API_KEY=<키>' >> ~/.zshrc    # macOS
   echo 'export MOONSHOT_API_KEY=<키>' >> ~/.bashrc   # Linux
   ```

3. 로컬 Ollama는 모델을 먼저 내려받아요: `ollama pull qwen3.5`. `:cloud` 로 끝나는 모델은 Ollama 로그인이 필요해요.
4. 메인 에이전트에게 맡길 일을 `tmp-agent/` 안의 파일에 쓰게 하고, 실행기를 부르게 해요.

   ```sh
   node .harness/subagents/run_subagent.mjs kimi tmp-agent/review-auth.md
   ```

### 2.5 서브 에이전트에게 맡기기 좋은 일

| 맡기기 좋은 일 | 맡기지 않는 일 |
| --- | --- |
| 넓은 코드 조사 결과 요약, 두 번째 의견의 코드 리뷰, 문서 초안, 테스트 시나리오 목록 | 파일 편집(실행기가 막아요), 검증 없이 결과를 그대로 반영, 키를 명령줄에 쓰기 |

서브 에이전트의 답은 **검증 전 관찰**이에요. 메인은 답에 적힌 파일·줄·명령 결과를 직접 읽거나 다시 실행해서 맞을 때만 사용해요.
