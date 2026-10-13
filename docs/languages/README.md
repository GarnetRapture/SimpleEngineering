# 언어별 핵심 지식

[← 코드 리뷰 방법론](../code-review.md) · [시작하기](../getting-started.md)

AI 에이전트가 코드를 대신 써 줘도, 그 코드가 맞는지 판단하는 것은 사람이에요. 이 문서들은 각 언어로 프로젝트를 시작하고, 에이전트가 낸 결과를 읽고, 하네스의 검사 결과를 이해하는 데 필요한 핵심만 모았어요.

| 언어 | 문서 | 주 용도 |
| --- | --- | --- |
| C++ (xmake) | [cpp.md](cpp.md) | 게임·엔진·네이티브 앱·고성능 도구 |
| Rust | [rust.md](rust.md) | 메모리 안전이 중요한 네이티브 프로그램·CLI |
| Go | [go.md](go.md) | 서버·CLI·네트워크 도구 |
| TypeScript | [typescript.md](typescript.md) | 웹 프런트엔드·Node.js 서버·도구 |
| Python | [python.md](python.md) | 자동화·데이터·AI·웹 서버 |
| PHP | [php.md](php.md) | 웹 서버·CMS |

## 1. 모든 언어에 공통인 개념

| 개념 | 뜻 | 왜 알아야 하나 |
| --- | --- | --- |
| 정적 검사 | 실행하지 않고 코드를 분석해 타입 오류·위험 패턴을 찾음 (`tsc --noEmit`, `cargo clippy`, `go vet`, `pyright`, `phpstan`) | 하네스가 편집 후 "남은 검사"로 요구하는 것이 주로 이것이에요 |
| 빌드 | 소스를 실행 파일·배포물로 변환 | 빌드 성공은 "변환이 됐다"는 증거일 뿐 동작이 맞다는 증거가 아니에요 |
| 실행 | 만든 프로그램을 실제로 돌림 | 제품 실행은 사용자 몫이라 하네스가 막는 명령이 있어요 (`npm run dev`, `cargo run`, `go run`, `php -S`) |
| 패키지 매니저 | 외부 라이브러리를 받아 버전을 기록 (`npm`, `cargo`, `go`, `uv`, `composer`) | 의존성은 반드시 이 명령으로만 바꿔요 |
| 잠금 파일 | 실제 설치된 정확한 버전 목록 (`package-lock.json`, `Cargo.lock`, `go.sum`, `uv.lock`, `composer.lock`) | 손으로 고치면 설치 결과와 기록이 어긋나요. 하네스가 편집을 막아요 |
| 표지 파일 | 그 언어 프로젝트임을 알려 주는 파일 (`package.json`, `Cargo.toml`, `go.mod`, `pyproject.toml`, `composer.json`, `xmake.lua`) | 하네스 스택은 표지 파일이 루트에 모두 있을 때만 켜져요 |
| 경고 | 오류는 아니지만 결함 가능성이 있는 진단 | 하네스 검사는 경고도 실패로 다뤄요 (`--max-warnings 0`, `-D warnings`, `--warnings`, 줄에 `warning` 금지) |

## 2. 하네스 스택 대응표

CLI에서 대표 언어를 고르면 아래 스택이 `harness.config.json` `stacks` 에 들어가요.

| 언어 | 표지 파일 | 검사 명령 (수용 조건) | 금지 명령 | 보호 대상 |
| --- | --- | --- | --- | --- |
| TypeScript | `package.json` + `tsconfig.json` | `npx tsc --noEmit` (종료 0), `npx eslint . --max-warnings 0` (종료 0) | `npm run dev`, `npm start`, `vite`, `npm run build` | `dist/`, `package.json`·잠금 파일 |
| Python | `pyproject.toml` | `ruff check .`, `ruff format --check .`, `pyright --warnings` (모두 종료 0) | — | `uv.lock`, `poetry.lock` |
| C++ | `xmake.lua` | `xmake build` (종료 0, `warning` 줄 없음) | — | `build/`, `.xmake/` |
| Rust | `Cargo.toml` | `cargo clippy --all-targets -- -D warnings` (종료 0) | `cargo build`, `cargo run` | `target/`, `Cargo.lock` |
| Go | `go.mod` | `go vet ./...` (종료 0), `gofmt -l .` (종료 0, `.go` 로 끝나는 줄 없음) | `go run`, `go build` | `go.sum` |
| PHP | `composer.json` | `php -l <파일>` (종료 0), `vendor/bin/phpstan analyse` (종료 0) | `php -S` | `composer.lock` |

모든 언어 공통으로 `vendor/`·`third-party/` 는 외부 코드라 보호돼요.

## 3. 이름 규칙 요약 (컨벤션 C05)

| 언어 | 함수·메서드 | 타입·클래스 | 상수 | 파일 |
| --- | --- | --- | --- | --- |
| TypeScript | `lowerCamelCase` | `PascalCase` | `UPPER_SNAKE_CASE` | 내보내는 심볼 이름과 같게 |
| Python | `snake_case` | `PascalCase` | `UPPER_SNAKE_CASE` | `snake_case.py` |
| C++ | `lowerCamelCase` | `PascalCase` | `UPPER_SNAKE_CASE` | 선언·구현 쌍 |
| Rust | `snake_case` | `PascalCase` | `UPPER_SNAKE_CASE` | `snake_case.rs` |
| Go | `MixedCaps` (내보내면 대문자 시작) | `MixedCaps` | `MixedCaps` | `snake_case.go`, 패키지 폴더는 소문자 한 단어 |
| PHP | `lowerCamelCase` | `PascalCase` | `UPPER_SNAKE_CASE` | 클래스 이름과 같게 |

도메인 기반 명명 규칙을 켜면 `util`, `helper`, `common`, `manager`, `misc`, `temp`, `new`, `data`, `info` 처럼 책임을 숨기는 이름을 쓰지 않고 **도메인 용어 + 동작 + 대상**으로 지어요. 예: `helper()` → `formatOrderLabel()`.
