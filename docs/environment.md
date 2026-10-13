# 환경 준비 — 운영체제별 셸·Git·터미널

[← 시작하기](getting-started.md) · 다음: [에이전트별 적용](agents.md)

SimpleHarness CLI(`go run ./src`)는 실행할 때 운영체제를 감지해서 이 문서와 같은 안내를 단계 번호와 함께 보여 줘요. 이 문서는 그 안내를 한곳에 모은 참고서예요.

## 1. 왜 셸과 터미널부터 정하나

AI 에이전트는 명령을 셸(shell)에서 실행해요. 셸이 오래되거나 에이전트가 기대하는 문법과 다르면, 같은 명령이 실패하거나 다르게 동작해요. 그래서 하네스를 넣기 전에 다음 세 가지를 맞춰요.

| 용어 | 뜻 | 예시 |
| --- | --- | --- |
| 셸 | 명령을 해석해서 실행하는 프로그램 | cmd, Windows PowerShell 5.1, PowerShell 7(pwsh), bash, zsh |
| 터미널 | 셸을 띄우는 창(탭·글꼴·색상) | Windows Terminal, macOS Terminal, GNOME Terminal |
| 패키지 매니저 | 명령 한 줄로 프로그램을 설치하는 도구 | winget(Windows), Homebrew(macOS), apt·dnf·pacman·zypper(Linux) |

## 2. Windows

### 2.1 셸 선택: PowerShell 7 또는 Git Bash

cmd(명령 프롬프트)는 Windows의 오래된 셸이라 교체를 권장해요. Windows에 기본으로 들어 있는 **Windows PowerShell 5.1** 도 구버전이에요. 아래 둘 중 하나를 기본 셸로 쓰세요.

| 셸 | 장점 | 이런 사람에게 |
| --- | --- | --- |
| **PowerShell 7 (pwsh)** | 명령 결과를 텍스트가 아니라 **객체**(속성을 가진 데이터)로 파이프라인에 넘겨요. 그래서 필터·정렬·변환을 문자열 자르기 없이 정확하게 할 수 있어요. 5.1과 따로 설치되고 `pwsh` 로 실행해요 | Windows 중심으로 작업하는 사람 |
| **Git Bash** | Linux·macOS와 같은 bash 문법이라 AI 에이전트 문서와 예제 명령을 그대로 쓸 수 있어요. Claude Code도 Windows에서 훅 명령을 Git Bash로 실행해요 | 여러 운영체제를 오가거나 Linux 명령에 익숙한 사람 |

객체 파이프라인 예시 — 1MB보다 큰 파일을 크기순으로 보기:

```powershell
Get-ChildItem -Recurse -File | Where-Object Length -gt 1MB | Sort-Object Length -Descending | Select-Object Name, Length
```

`Length` 는 파일 객체의 속성이라, 출력 형식이 바뀌어도 필터가 깨지지 않아요. bash에서는 같은 일을 `find . -type f -size +1M` 처럼 도구별 옵션으로 해요.

### 2.2 Windows 10

Windows 10의 기본 셸은 **cmd** 예요. 아래 명령을 **cmd 창에 한 줄씩** 붙여 넣어 셸을 교체해요. Microsoft Store를 거치지 않고 winget으로 설치해요.

1. **winget 확인**

   ```bat
   winget --version
   ```

   `'winget'은(는) 내부 또는 외부 명령... 이 아닙니다` 가 나오면 등록하고 cmd 창을 새로 열어요. winget은 Windows 10 1809 이상에서 지원돼요.

   ```bat
   powershell -NoProfile -Command "Add-AppxPackage -RegisterByFamilyName -MainPackage Microsoft.DesktopAppInstaller_8wekyb3d8bbwe"
   ```

2. **PowerShell 7 설치**

   ```bat
   winget install --id Microsoft.PowerShell --source winget
   ```

3. **Windows Terminal 설치**

   ```bat
   winget install --id Microsoft.WindowsTerminal --source winget
   ```

4. **Git 설치** (git과 Git Bash가 함께 설치돼요)

   ```bat
   winget install --id Git.Git -e --source winget
   ```

5. **기본 터미널 바꾸기** — Windows Terminal → 설정 → 시작 → **기본 터미널 애플리케이션** 을 Windows Terminal로, **기본 프로필** 을 PowerShell(7)로 바꿔요. Windows 10은 22H2에 2023-05-23 업데이트(KB5026435) 이상이 설치돼야 '기본 터미널 애플리케이션' 항목이 보여요.

6. 터미널 창을 모두 닫고 새로 열어요. 새로 설치한 명령은 새 창부터 PATH에 잡혀요.

> **알려진 문제 — Windows 10 Home의 Codex 앱**: 공식 설치한 Codex 앱이 기본 셸을 cmd로 잡아 명령이 제대로 동작하지 않는 경우가 있어요. PowerShell 7을 설치하고 기본 터미널·프로필을 PowerShell 7로 바꾼 뒤 사용하세요. SimpleHarness가 만드는 Codex 훅 등록 명령은 한 줄짜리 `node -e` 라서 cmd로 실행돼도 같은 방식으로 동작해요.

### 2.3 Windows 11

Windows 11에는 Windows Terminal과 winget이 들어 있지만, 함께 들어 있는 Windows PowerShell 5.1은 구버전이에요.

1. PowerShell 7 설치: `winget install --id Microsoft.PowerShell --source winget`
2. Git 설치: `winget install --id Git.Git -e --source winget`
3. Windows Terminal → 설정 → 시작 → **기본 프로필** 을 PowerShell(7)로 바꿔요. 이후 작업은 그 창(pwsh)에서 해요.

### 2.4 Git Bash를 기본 셸로 쓰기

- **Windows Terminal**: 설정 → `+ 새 프로필 추가` → 명령줄에 `C:\Program Files\Git\bin\bash.exe -li` 를 넣어 프로필을 만들고 **기본 프로필** 로 지정해요.
- **VS Code**: 명령 팔레트(`Ctrl+Shift+P`) → `Terminal: Select Default Profile` → Git Bash 또는 PowerShell을 골라요.

### 2.5 스크립트 실행 정책

npm으로 설치한 명령(`.ps1`)이 PowerShell에서 "스크립트를 실행할 수 없습니다"로 막히면 한 번 실행해요.

```powershell
Set-ExecutionPolicy -Scope CurrentUser RemoteSigned
```

## 3. macOS

1. **Homebrew** 가 없으면 https://brew.sh 의 설치 명령을 먼저 실행해요. CLI가 도구를 자동 설치할 때 Homebrew를 써요.
2. **Git**: `xcode-select --install` 또는 `brew install git`
3. **셸**: macOS 기본 셸은 zsh예요. 바꾸려면 Terminal → 설정 → 일반 → **셸 열기 방식** 에서 '명령(전체 경로)'을 고르고 셸 경로를 입력해요.
4. 서브 에이전트 API 키 같은 환경 변수는 `~/.zshrc` 에 `export 이름=값` 으로 넣고 새 터미널을 열어요.

## 4. Linux

1. **Git 설치**

   | 패키지 매니저 | 명령 |
   | --- | --- |
   | apt (Debian·Ubuntu) | `sudo apt-get install -y git` |
   | dnf (Fedora) | `sudo dnf install -y git` |
   | pacman (Arch) | `sudo pacman -S git` |
   | zypper (openSUSE) | `sudo zypper install git` |

2. **기본 터미널 바꾸기 (Debian·Ubuntu)**: `sudo update-alternatives --config x-terminal-emulator` 를 실행하고 번호를 골라요. 다른 배포판은 데스크톱 환경의 기본 앱 설정에서 바꿔요.
3. 환경 변수는 `~/.bashrc` 에 `export 이름=값` 으로 넣고 새 터미널을 열어요.

### WSL (Windows 안의 Linux)

WSL 배포판은 Windows Terminal에 프로필이 자동으로 생겨요. Windows Terminal → 설정 → 시작 → **기본 프로필** 을 그 배포판으로 바꾸면 열자마자 Linux 셸이에요. 저장소도 WSL 파일 시스템(`~/…`) 안에 두고 그 안에서 에이전트를 실행해요.

## 5. 도구 설치 명령

CLI의 "추천 도구" 단계가 아래 명령으로 설치해요. Windows는 `winget install` 뒤에 `--accept-package-agreements --accept-source-agreements` 를 붙여 실행해요.

| 도구 | 쓰임 | winget | Homebrew | apt |
| --- | --- | --- | --- | --- |
| Git | 버전 관리·커밋 가드 | `--id Git.Git -e` | `git` | `git` |
| Node.js LTS | 훅·커밋 가드·서브 에이전트 실행기 (필수) | `--id OpenJS.NodeJS.LTS -e` | `node` | `nodejs npm` |
| PowerShell 7 | 최신 셸 (Windows) | `--id Microsoft.PowerShell` | — | — |
| Windows Terminal | 터미널 (Windows 10) | `--id Microsoft.WindowsTerminal` | — | — |
| jq | JSON 조회·변환 (인자로 쿼리) | `--id jqlang.jq -e` | `jq` | `jq` |
| ripgrep | 빠른 코드 검색 | `--id BurntSushi.ripgrep.MSVC -e` | `ripgrep` | `ripgrep` |
| SQLite | 파일 DB 조회 (인자로 SQL) | `--id SQLite.SQLite -e` | `sqlite` | `sqlite3` |
| DuckDB | CSV·JSON·Parquet 분석 SQL (인자로 SQL) | `--id DuckDB.cli -e` | `duckdb` | `curl https://install.duckdb.org \| sh` |
| Go | Go 툴체인 (SimpleHarness CLI를 소스에서 실행할 때 필요) | `--id GoLang.Go -e` | `go` | `golang-go` |
| Python 3.14 | Python 런타임 | `--id Python.Python.3.14 -e` | `python@3.14` | `python3` |
| uv | Python 패키지·환경 관리 | `--id astral-sh.uv -e` | `uv` | `curl -LsSf https://astral.sh/uv/install.sh \| sh` |
| Rustup | Rust 툴체인 | `--id Rustlang.Rustup -e` | `rustup` | `rustup` |
| LLVM | clang·clang-tidy·clangd | `--id LLVM.LLVM -e` | `llvm` | `llvm clang clang-tidy clangd` |
| GCC | GNU C/C++ 컴파일러 | `--id BrechtSanders.WinLibs.POSIX.UCRT -e` | `gcc` | `build-essential` |
| Visual Studio Build Tools | Windows C++·Rust 컴파일러·링커 | `--id Microsoft.VisualStudio.BuildTools -e --override "--wait --passive --add Microsoft.VisualStudio.Workload.VCTools --includeRecommended"` | — | — |
| xmake | C++ 빌드 도구 | `--id Xmake-io.Xmake -e` | `xmake` | `xmake` |
| PHP | PHP 런타임 | `--id PHP.PHP.8.5 -e` | `php` | `php-cli` |
| Composer | PHP 패키지 관리 | https://getcomposer.org/Composer-Setup.exe | `composer` | `composer` |
| Claude Code | 메인·서브 에이전트 | `--id Anthropic.ClaudeCode -e` | `--cask claude-code` | `curl -fsSL https://claude.ai/install.sh \| bash` |
| Codex CLI | 메인·서브 에이전트 | `--id OpenAI.Codex -e` | `--cask codex` | `npm install -g @openai/codex` |
| Ollama | 로컬·클라우드 모델 서버 | `--id Ollama.Ollama -e` | `ollama` | `curl -fsSL https://ollama.com/install.sh \| sh` |

apt로 설치하기 전에 CLI는 `sudo apt-get update` 를 먼저 실행해요.

## 6. AI를 실행하는 세 가지 방식

| 방식 | 하네스 적용 |
| --- | --- |
| 터미널에서 `claude` 또는 `codex` 실행 | 저장소 루트에서 실행하면 그 저장소의 설정·훅을 읽어요 |
| VS Code 확장 | Claude Code 확장은 프로젝트의 `.claude/settings.json` 훅과 `CLAUDE.md` 를 읽어요. Codex 확장은 CLI와 같은 설정 계층을 써요 |
| 데스크톱 앱 | 작업 폴더로 이 저장소를 열어요 |

어느 방식이든 **저장소 폴더를 작업 폴더로 열어야** 하네스가 적용돼요.
