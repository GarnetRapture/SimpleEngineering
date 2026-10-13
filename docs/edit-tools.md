# 편집 도구 정의 — 도구의 계약과 하네스의 행동 정의

[← 훅 설정](hooks.md) · 다음: [코드 리뷰 방법론](code-review.md)

에이전트는 정해진 편집 도구로만 파일을 바꿔요. 이 문서는 각 도구가 **무엇을 하는가(도구 계약)** 와, 하네스가 그 도구를 **언제 어떻게 쓰게 하는가(행동 정의)** 를 나눠서 정의해요. 행동 정의의 원천은 컨벤션 C12 "파일 읽기·검색·편집" 행이고, 훅이 아래 판정으로 실행해요.

## 1. 원칙 네 가지

| 원칙 | 정의 | 이유 |
| --- | --- | --- |
| 기존 파일은 부분 편집 | 이미 있는 파일은 바꿀 부분만 찾아 바꿔요 (Claude `Edit`, Codex `Update File`) | 전체 덮어쓰기는 바꾸지 않을 줄까지 다시 쓰므로, 모델이 기억으로 재작성한 부분에서 원문이 사라지거나 변형돼요 |
| 새 파일만 생성 | 생성 도구(Claude `Write`, Codex `Add File`)는 없는 파일에만 써요 | 생성 도구로 기존 파일을 쓰면 위와 같은 전체 덮어쓰기가 돼요 |
| 삭제 후 재생성 금지 | 같은 파일을 지우고 다시 만드는 방식으로 바꾸지 않아요 | 결과가 덮어쓰기와 같아요. 이름을 바꿀 때는 이동(`Move to`, `Move-Item`, `git mv`) 후 부분 편집해요 |
| 읽고 나서 편집 | 바꾸기 전에 대상 범위를 읽어 현재 내용을 확인해요 | 편집 도구는 정확한 원문 일치로 동작하고, 원문을 모르는 편집은 추측이에요 |

## 2. Claude Code 도구

| 도구 | 도구 계약 | 하네스 행동 정의 | 훅 판정 |
| --- | --- | --- | --- |
| `Read` | 파일을 줄 번호와 함께 읽어요. 기본 2000줄, `offset`·`limit` 로 범위 지정. 이미지·PDF·노트북도 읽어요 | 필요한 범위만 읽어요. 긴 파일을 통독할 때는 겹치지 않는 범위로 끝까지 읽어요 | `forbiddenPaths`, 에이전트 내부 임시 폴더 읽기 차단 |
| `Grep`, `Glob` | 내용 검색, 파일 이름 패턴 검색 (읽기 전용) | 검색 뿌리·범위·찾는 사실·종료 조건을 정한 한정 검색만 해요 | `Read` 와 같은 경로 판정 |
| `Edit` | 파일 안의 `old_string` 을 `new_string` 으로 바꿔요. `old_string` 은 정확히 일치하고 유일해야 하며(아니면 `replace_all`), 그 대화에서 먼저 `Read` 한 파일이어야 해요 | **기존 파일을 바꾸는 유일한 도구.** 고유하게 식별되는 가장 작은 원문을 골라요 | 보호 경로·저장소 밖 차단, `old_string`→`new_string` 사이 주석 수·인라인 규칙 일치 수 증가 차단 |
| `Write` | 파일 전체를 써요. 있으면 덮어써요(읽지 않은 기존 파일 덮어쓰기는 실패) | **새 파일 생성 전용** | 대상이 이미 있으면 차단(C12). 새 파일이 `tmp-agent/` 밖이면 `-structure` 스킬 필요 |
| `NotebookEdit` | `.ipynb` 의 셀 하나를 `cell_id` 로 바꾸기·삽입·삭제 | 노트북 셀 편집 전용 | 보호 경로·저장소 밖 차단, 스킬 판정 (`.ipynb` 는 주석·인라인 규칙의 확장자 목록에 없어요) |
| `PowerShell` | PowerShell 명령 실행 | Windows의 프로젝트 명령 경로. 파일 내용 변경에는 쓰지 않고 이동·검사·도구 실행에 써요 | 셸 판정 전체 (아래 4장) |
| `Bash` | Bash 명령 실행 (Windows에서는 Git Bash) | macOS·Linux의 프로젝트 명령 경로 | Windows에서는 차단(C12 셸 행), 그 외 셸 판정 전체 |

## 3. Codex `apply_patch`

Codex는 파일을 `apply_patch` 도구 하나로 바꾸고, 패치 본문은 훅 입력의 `tool_input.command` 로 전달돼요.

### 3.1 문법

```text
*** Begin Patch
*** Add File: src/order/types/orderStatus.types.ts
+export type OrderStatus = 'pending' | 'paid' | 'shipped';
*** Update File: src/order/logic/calculateOrderTotal.ts
@@ export function calculateOrderTotal
-  return items.reduce((sum, item) => sum + item.price, 0);
+  return items.reduce((sum, item) => sum + item.price * item.quantity, 0);
*** Update File: src/order/logic/orderHelper.ts
*** Move to: src/order/logic/formatOrderLabel.ts
@@
-export function helper(order: Order): string {
+export function formatOrderLabel(order: Order): string {
*** Delete File: src/order/legacy/oldOrder.ts
*** End Patch
```

| 머리줄 | 동작 | 하네스 행동 정의 |
| --- | --- | --- |
| `*** Add File: <경로>` | 새 파일 생성. 이어지는 `+` 줄이 내용 | 없는 파일에만 |
| `*** Update File: <경로>` | 기존 파일 부분 편집. `@@` 문맥, ` `(유지)·`-`(삭제)·`+`(추가) 줄 | 기존 파일을 바꾸는 유일한 방법 |
| `*** Move to: <새 경로>` | 바로 위 `Update File` 대상을 새 경로로 이동 | 이름 변경은 이동 + 부분 편집 |
| `*** Delete File: <경로>` | 파일 삭제 | 명시적으로 지시된 대상만 |

### 3.2 훅 판정

| 판정 | 조건 |
| --- | --- |
| 보호 경로 | `Update`·`Add`·`Delete` 대상 또는 `Move to` 경로가 `protectedPaths` 와 일치하거나 저장소 밖 → 차단 |
| 덮어쓰기 | `Add File` 대상이 이미 있음, 또는 같은 패치에서 `Delete File` 한 경로를 `Add File` → 차단 |
| 주석·인라인 규칙 | 각 파일의 `-` 줄 묶음과 `+` 줄 묶음을 비교해 주석 수·인라인 패턴 일치 수가 늘면 차단 |
| 스킬 | 편집이면 공통 스킬 3개, 새 파일·이동이면 `-structure`, 스택 확장자면 `-stack` |

## 4. 셸 명령으로 하는 변경

셸은 파일을 바꿀 수 있지만, **파일 내용 변경은 편집 도구로** 해요. 셸은 이동·도구 실행·검사에 써요. 훅은 셸 명령에서 아래를 인식해요.

| 인식하는 명령 | 의미 | 결과 |
| --- | --- | --- |
| `Set-Content`, `Add-Content`, `Out-File`, `New-Item`, `Copy-Item`, `Move-Item`, `Rename-Item`, `Remove-Item`, `Clear-Content` | 파일을 바꾸는 명령 | 공통 스킬 3개 필요 |
| `Move-Item`, `Rename-Item`, `git mv` | 이동·이름 변경 | `-structure` 스킬 필요 |
| 위 쓰기 명령 + 보호 폴더 이름(`shellSegments`) | 보호 폴더에 쓰기 | 차단 (C13) |
| 위 쓰기 명령 + `/tmp/`, 에이전트 내부 임시 폴더 | 프로젝트 밖 임시 쓰기 | 차단, `tmp-agent/` 사용 |
| `cmd`, `npm.cmd`, `.bat` (Windows) | cmd 경로 | 차단 |

훅은 명령 문자열에서 따옴표로 감싼 리터럴을 비우고 판정해요. 그래서 `echo "cargo run"` 처럼 문자열 안에 든 단어로는 금지 명령 판정이 걸리지 않고, `$( … )` 치환이 들어 있는 큰따옴표 문자열은 그대로 판정해요.

## 5. 의도별 도구 선택

| 하려는 일 | Claude Code | Codex |
| --- | --- | --- |
| 기존 파일의 일부 수정 | `Edit` | `Update File` |
| 같은 원문을 파일 전체에서 바꾸기 | `Edit` + `replace_all: true` | `Update File` 에 각 위치의 hunk |
| 새 파일 만들기 | `Write` | `Add File` |
| 파일 이름 변경·이동 | `PowerShell`/`Bash` 의 `Move-Item`·`git mv` 후 `Edit` | `Update File` + `Move to` |
| 지시된 파일 삭제 | `PowerShell` `Remove-Item -LiteralPath` | `Delete File` |
| 노트북 셀 수정 | `NotebookEdit` | `Update File` (JSON 편집) |

## 6. 좋은 예와 나쁜 예

**좋은 예** — 수량을 빠뜨린 합계 계산을 고칠 때, 바뀌는 한 줄만 원문으로 지정해요.

```text
Edit
  file_path:  src/order/logic/calculateOrderTotal.ts
  old_string: return items.reduce((sum, item) => sum + item.price, 0);
  new_string: return items.reduce((sum, item) => sum + item.price * item.quantity, 0);
```

이유: 바뀌는 줄만 건드리므로 나머지 내용은 원문 그대로 남고, 리뷰어는 차이(diff) 한 줄로 변경을 확인해요.

**나쁜 예** — 같은 수정을 위해 파일 전체를 `Write` 로 다시 써요.

이유: 이미 있는 파일이라 훅이 `C12 file editing row: an existing file changes in part through Edit; Write creates new files only` 로 차단해요. 차단되지 않는 환경이라도, 기억으로 다시 쓴 나머지 줄에서 원문이 바뀌는 사고가 생겨요.

**나쁜 예** — 파일을 지우고 같은 이름으로 다시 만들어 "정리"해요 (Codex `Delete File` + `Add File`).

이유: 결과가 덮어쓰기와 같아서 훅이 `Add File creates new files only, and deleting then re-adding a file overwrites it` 로 차단해요.
