# 훅 설정 — 무엇을 언제 막고, 어떻게 바꾸고, 어떻게 시험하나

[← 에이전트별 적용](agents.md) · 다음: [편집 도구 정의](edit-tools.md)

훅(hook)은 에이전트가 정해진 순간(세션 시작, 지시 입력, 도구 호출 직전·직후, 답변 종료)에 실행하는 명령이에요. Claude Code와 Codex의 공식 훅 계약에서 훅 명령은 표준 입력으로 JSON을 받고, 다음 두 가지로 에이전트를 통제해요.

| 출력 | 효과 |
| --- | --- |
| 종료 코드 **2** + 표준 오류 메시지 | 그 도구 호출(PreToolUse) 또는 답변 종료(Stop)를 **차단**하고, 메시지를 모델에게 돌려줘요. 모델은 그 이유를 읽고 다시 행동해요 |
| 표준 출력 JSON `{"hookSpecificOutput": {"hookEventName": …, "additionalContext": "…"}}` | 그 시점에 지켜야 할 조건을 모델의 컨텍스트에 **주입**해요 |

SimpleHarness의 훅은 컨벤션 조항을 이 두 수단으로 실행해요. 차단 메시지에는 항상 근거가 되는 **조항 번호와 설정 키**가 들어 있어요(예: `C04 no-inline-comment rule (harness.config.json rules): …`).

## 1. 구조

```
.harness/hooks/run_project_hook.mjs   ← 공용 엔진 (Claude·Codex 공통)
harness.config.json                   ← 엔진이 실행마다 읽는 상수
.claude/settings.json                 ← Claude Code 등록 (메인이 Claude일 때)
.codex/hooks.json                     ← Codex 등록 (메인이 Codex일 때)
tmp-agent/hook/<에이전트>/<세션>/      ← 세션 상태
tmp-agent/hook/failure/               ← 엔진 자체 실패 기록
```

### 1.1 등록 명령

두 등록 파일 모두 같은 형태의 한 줄 명령을 모든 이벤트에 걸어요(시간 제한 30초).

```text
node -e "import(require('node:url').pathToFileURL(require('node:path').resolve(<루트>, '.harness/hooks/run_project_hook.mjs')).href).catch(error => { process.exitCode = 1; process.stderr.write(String(error.stack ?? error)); })" <에이전트>
```

| 부분 | Claude Code | Codex |
| --- | --- | --- |
| `<루트>` | `process.env.CLAUDE_PROJECT_DIR \|\| '.'` | `require('node:child_process').execFileSync('git', ['rev-parse', '--show-toplevel'], { encoding: 'utf8' }).trim()` |
| `<에이전트>` (마지막 인자) | `claude` | `codex` |

이 형태를 고른 이유:

- **개인 경로가 없어요.** 사용자 이름·드라이브·절대 경로를 쓰지 않고 실행 시점에 저장소 루트를 찾아요. 그래서 누가 클론해도 같은 등록이 동작해요.
- **셸 확장이 없어요.** `$VAR`, `%VAR%`, `~` 같은 셸 문법을 쓰지 않아서 Git Bash, PowerShell, cmd.exe(Codex가 Windows에서 훅을 실행하는 셸) 어디서든 같은 명령으로 동작해요.
- **엔진이 하나예요.** 에이전트 이름만 인자로 다르게 넘기고, 판정 로직은 `run_project_hook.mjs` 하나에 있어요.

## 2. 이벤트별 동작

### SessionStart — 세션 시작

입력의 `source` 값에 따라 주입하는 내용이 달라요.

| source | 주입 |
| --- | --- |
| `startup`, `clear` (새 세션) | 컨벤션 전문 → AGENTS 전문을 먼저 읽으라는 조건, 도구 근거·스킬 적재·응답 형식 조건 |
| `compact` (압축 후) | 컨벤션 요약표, AGENTS 행동 요약 위치, 트래킹 최신 요약, 프로젝트 노드 표, 열린 TODO를 원문 위치·sha256과 함께 (10,000자 이내). 스킬 적재 표시는 초기화돼요 |
| `resume`, `fork` (이어서 작업) | 전부 다시 읽지 말고 최근 작업과 연결된 조항·소유 파일만 확인하라는 조건 |

### UserPromptSubmit — 지시 입력

요구를 원자 단위로 보존하라는 조건(C00·C15), 도구 근거·스킬 적재·응답 형식 조건을 주입해요. 입력 기록은 세션 상태에 남아서, 나중에 "중지 지시가 있었는지" 같은 판정에 쓰여요.

### PreToolUse — 도구 호출 직전 (차단 지점)

위에서부터 차례로 검사하고, 처음 걸린 위반 하나로 차단해요(종료 2). Codex에서 엔진이 해석하는 도구는 `apply_patch`(편집)와 `Bash`(셸)이고, Claude Code에서는 `Edit`·`Write`·`NotebookEdit`(편집), `PowerShell`·`Bash`(셸), `Read`·`Grep`·`Glob`(읽기), 질문 도구, `TaskStop` 이에요.

| 검사 | 차단 조건 | 근거 |
| --- | --- | --- |
| 질문 도구 (Claude Code) | `AskUserQuestion`, `EnterPlanMode`, `ExitPlanMode` | C12 질문 행 |
| Windows의 Claude `Bash` | Windows에서 Claude Code `Bash` 도구 | C12 셸 행 |
| 작업 중지 (Claude Code) | 최신 지시에 중지 지시가 없는데 `TaskStop` | C12 백그라운드 행 |
| 보호 경로 편집 | `protectedPaths` 정규식에 맞는 경로, 저장소 밖 경로 | C13 |
| 기존 파일 덮어쓰기 | 이미 있는 파일에 `Write`, 이미 있는 파일에 `Add File`, 같은 패치에서 삭제 후 다시 추가 | C12 파일 편집 행 |
| 주석 증가 | `no-inline-comment` 규칙이 켜져 있고 편집 후 주석 수가 늘어남 | C04 |
| 인라인 규칙 | `inlineRules` 패턴의 일치 수가 편집 전보다 늘어남 (허용 경로 제외) | C04·C05 |
| 읽기 금지 경로 | `forbiddenPaths`, 에이전트 내부 임시 폴더 읽기 | C12 |
| 셸: 작업 강제 종료 | 중지 지시 없이 `Stop-Process`·`taskkill`·`kill` | C12 |
| 셸: cmd 경로 (Windows) | `cmd`, `npm.cmd`, `.bat` | C12 셸 행 |
| 셸: 전경 대기 | `Start-Sleep`, `Get-Content -Wait` (백그라운드 실행이 아닐 때) | C12 |
| 셸: 스택 금지 명령 | 켜진 스택의 `forbidden` 명령 (예: `cargo run`, `npm run build`) | C12 스택 표 |
| 셸: 거부 명령 | `deniedCommands` 정규식 | 설정한 사유 |
| 셸: 임시 경로 | 에이전트 내부 임시 폴더, `/tmp/` 쓰기 | C04·C12 스크래치 행 |
| 셸: 보호 폴더 쓰기 | `protectedPaths.shellSegments` 폴더로 쓰는 명령 | C13 |
| 스킬 미적재 | 이 호출에 필요한 스킬을 아직 적재하지 않음 | AGENTS 3장 스킬 행 |

필요한 스킬은 이렇게 정해져요.

| 상황 | 필요한 스킬 |
| --- | --- |
| 파일을 바꾸거나(편집 도구, 쓰기 셸 명령) 프로그램을 실행(`node`, `npm`, `python`, 스택 검사 명령 등) | `<접두사>-workflow`, `-fix-process`, `-verify` |
| 새 파일 생성(`tmp-agent/` 밖), 이동·이름 변경 | `-structure` |
| 켜진 스택의 확장자 편집, 스택 검사 실행 | `-stack` |
| `skills.domain` 항목의 경로·확장자·명령·도구 접두사와 일치 | 그 분야 스킬 (`-naming`, `-inline`, `-subagent` 등) |

통과한 호출에는 "이 도구와 방법을 정의한 조항·행과 확인한 근거를 먼저 밝히라"는 조건을 주입해요.

### PostToolUse — 도구 호출 직후

| 동작 | 조건 |
| --- | --- |
| 스킬 적재 표시 | Claude `Skill` 도구 호출, 또는 성공한 도구 호출의 입력에 `<스킬 위치>/<스킬>/SKILL.md` 경로가 있음 |
| 남은 검사 표시 | 켜진 스택의 확장자 파일을 편집함 → 그 스택의 모든 검사가 "남음" |
| 남은 검사 해제 | 검사 명령을 실행했고 종료 코드·필수 줄·금지 줄 조건을 모두 만족함 → 그 호출 **이전** 편집이 남긴 표시만 해제 |
| 실패 처리 (Claude `PostToolUseFailure`) | 실패한 입력·진단을 보존하고 FIX 절차를 적용하라는 조건 주입 |

### Stop — 답변 종료 (차단 지점)

차례로 검사하고 처음 걸린 것으로 차단해요.

1. 필수 스킬(`-workflow`, `-fix-process`, `-verify`)을 적재했는가
2. 금지 판단 표현(`response.bannedJudgment`, 예: "나중에", "범위 밖", "중단할게", "일부만")이 없는가
3. 사용자 전용 관문 표식(`response.gateWrapper`)을 흉내 내지 않았는가
4. 최신 지시에 보고 요청(`response.reportKeyword`)이 없는데 수긍·보고 문구(`response.acknowledgment`)로 답하지 않았는가
5. 인용한 조항 번호(`Cnn`)가 컨벤션에 실제 앵커로 있는가
6. 응답 필드(`response.fields`)가 모두 있는가, 근거 라벨(`response.basisLabels`) 중 하나가 있는가
7. 각 필드 값에 조항 번호·경로·숫자·URL 같은 증거 표지가 있는가, 첫 필드에 조항 번호가 있는가
8. 인용한 파일 경로·줄 번호가 실제로 있는가
9. 추측 표현(`response.guessedJudgment`)이 없는가
10. 근거 문장(`response.basisStatement`)과 번호 근거·결론(`response.numberedBasis`)이 있는가
11. 남은 검사가 있으면: Claude는 계속하라는 조건을 주입, Codex는 차단

판정 표현과 제목은 출력 언어(ko·en·zh)별로 `harness.config.json` `response` 에 들어 있어요.

## 3. 상태와 실패 기록

| 위치 | 내용 |
| --- | --- |
| `tmp-agent/hook/<에이전트>/<세션 id>/<시각>-<uuid>-<이벤트>.json` | 받은 입력 원문 |
| `…-denial.json`, `…-block.json` | 차단 메시지 |
| `skill_<해시>` | 그 세션에서 적재한 스킬 |
| `pending_<해시>_<uuid>` | 아직 통과하지 못한 검사 |
| `check_<해시>.json` | 검사 호출 직전의 남은 검사 목록 |
| `tmp-agent/hook/failure/` | 엔진이 처리하지 못한 입력과 오류 스택 (엔진은 종료 코드 1로 끝나요) |

서브 에이전트 입력(`agent_id` 가 있는 입력)은 판정하지 않고 상태도 바꾸지 않아요.

## 4. 바꾸는 법 — `harness.config.json`

엔진은 **실행할 때마다** `harness.config.json` 을 읽어요. 그래서 판정 규칙을 바꾸려면 이 파일만 고치면 되고, 다음 훅 실행부터 적용돼요. 등록 파일(`.claude/settings.json`, `.codex/hooks.json`)을 바꿨을 때만 에이전트를 다시 시작해요.

| 키 | 바꾸는 것 |
| --- | --- |
| `rules` | 켜진 핵심 규칙 id 목록. `no-inline-comment` 가 있으면 주석 증가를 차단 |
| `inlineRules[]` | `id`, `clause`, `extensions`, `pattern`(정규식, 줄 단위), `allowPath`(허용 경로 정규식), `reason`(차단 메시지) |
| `protectedPaths[]` | `pattern`(편집 차단 경로), `shellSegments`(셸 쓰기 차단 폴더), `reason` |
| `forbiddenPaths[]` | 읽기·명령에서 참조를 막을 경로 정규식과 사유 |
| `deniedCommands[]` | 실행을 막을 명령 정규식과 사유 |
| `stacks[]` | `markers`(모두 있을 때만 켜짐), `extensions`, `checks[]`(`command`, `acceptance.exit`, 선택 `line`·`forbidLine`), `forbidden[]` |
| `observations[]` | 편집하면 실행 관찰이 "남음"이 되는 확장자 |
| `skills.domain[]` | 분야 스킬이 필요해지는 `paths`·`extensions`·`commands`·`toolPrefixes` |
| `response` | 응답 필드, 근거 라벨, 판정 정규식, 압축 때 읽는 제목 |
| `paths` | 컨벤션·트래킹·감사 문서, 스크래치 폴더 위치 |

예: 강제 푸시를 막는 규칙 추가

```json
"deniedCommands": [
  { "pattern": "\\bgit\\s+push\\b[^;|]*\\s(--force|-f)\\b", "reason": "C13 git history row: force push rewrites the remote history" }
]
```

규칙을 바꾸면 같은 작업에서 컨벤션의 해당 표(C04 핵심 규칙 표, C12 스택 표, C13 보호 표)도 함께 고쳐서 문서와 판정이 같은 내용을 말하게 해요.

## 5. 시험하는 법

엔진에 JSON을 직접 넣어서 판정을 확인해요. 저장소 루트에서 실행해요.

PowerShell 7:

```powershell
'{"session_id":"hooktest","hook_event_name":"PreToolUse","tool_name":"AskUserQuestion","tool_input":{},"cwd":"."}' | node .harness/hooks/run_project_hook.mjs claude
"exit $LASTEXITCODE"
```

Git Bash·macOS·Linux:

```sh
echo '{"session_id":"hooktest","hook_event_name":"PreToolUse","tool_name":"AskUserQuestion","tool_input":{},"cwd":"."}' | node .harness/hooks/run_project_hook.mjs claude
echo "exit $?"
```

기대 결과: 표준 오류에 `C12 interactive question row …` 가 나오고 `exit 2`.

Codex `apply_patch` 판정 예 (주석을 추가하는 패치):

```sh
printf '%s' '{"session_id":"hooktest","hook_event_name":"PreToolUse","tool_name":"apply_patch","tool_input":{"command":"*** Begin Patch\n*** Add File: src/demo.ts\n+// note\n+export const demoValue = 1;\n*** End Patch"},"cwd":"."}' | node .harness/hooks/run_project_hook.mjs codex
echo "exit $?"
```

기대 결과: `no-inline-comment` 규칙이 켜져 있으면 `C04 no-inline-comment rule …` 와 `exit 2`.

시험이 끝나면 `tmp-agent/hook/claude/hooktest/`, `tmp-agent/hook/codex/hooktest/` 를 지워요. 실제 등록 문자열까지 시험하려면 등록 파일의 `command` 값을 그대로 복사해 같은 JSON을 넣어요(Claude는 Git Bash에서, Codex는 `cmd.exe /C "<명령>"` 으로).

## 6. 커밋 가드

훅과 별도로 git 커밋 시점의 검사가 있어요.

| 파일 | 동작 |
| --- | --- |
| `.githooks/pre-commit` | 병합 커밋 중이면 통과, 아니면 `node tools/guard/staged_guard.mjs` 실행 |
| `tools/guard/staged_guard.mjs` | 스테이징된 파일 중 `harness.config.json` `commitGuard.blockedExtensions` 확장자가 있으면 파일·크기를 출력하고 종료 1로 커밋을 멈춰요 |

연결: 저장소 루트에서 `git config core.hooksPath .githooks` 를 한 번 실행해요. 이 설정 뒤에는 에이전트와 사람의 커밋 모두에 같은 검사가 적용돼요.
