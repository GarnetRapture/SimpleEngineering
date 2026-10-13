# Python 핵심 지식

[← 언어별 핵심 지식](README.md)

Python은 읽기 쉬운 문법과 방대한 라이브러리를 가진 인터프리터 언어예요. 자동화, 데이터 분석, AI, 웹 서버에 써요. 타입을 실행 중에 정하므로, 타입 힌트와 정적 검사기(pyright)를 써야 실행 전에 타입 오류를 잡을 수 있어요.

## 1. 설치

| 도구 | Windows | macOS | Linux |
| --- | --- | --- | --- |
| Python 3.14 | `winget install --id Python.Python.3.14 -e --source winget` | `brew install python@3.14` | `sudo apt-get install -y python3` |
| uv (패키지·환경 관리) | `winget install --id astral-sh.uv -e --source winget` | `brew install uv` | `curl -LsSf https://astral.sh/uv/install.sh \| sh` |

검사 도구는 PATH에서 바로 실행되도록 도구로 설치해요.

```sh
uv tool install ruff
uv tool install pyright
```

## 2. 새 프로젝트

```sh
uv init order-tool
cd order-tool
uv add httpx
```

| 파일 | 역할 |
| --- | --- |
| `pyproject.toml` | 프로젝트 정보·의존성 (하네스 표지) |
| `uv.lock` | 설치된 정확한 버전 (잠금 파일) |
| `.venv/` | uv가 만드는 프로젝트 전용 환경 |

| 명령 | 하는 일 |
| --- | --- |
| `ruff check .` | 린트 |
| `ruff format --check .` | 포맷 검사 (다르면 실패) / `ruff format .` 로 적용 |
| `pyright --warnings` | 타입 검사. 경고가 있어도 종료 코드 1 (pyright 공식 문서) |
| `uv add <패키지>`, `uv remove <패키지>` | 의존성 변경 (`pyproject.toml`·`uv.lock` 을 uv가 갱신) |
| `uv run <파일>` | 프로젝트 환경에서 실행 |

## 3. 핵심 개념

| 개념 | 설명 | 실수 예 |
| --- | --- | --- |
| 들여쓰기 = 블록 | 중괄호 대신 들여쓰기로 범위를 정해요 | 탭·공백 섞기 |
| 타입 힌트 | `def total(items: list[Item]) -> int:` 실행에는 영향이 없고 pyright가 검사해요 | 힌트 없이 써서 검사기가 아무것도 못 잡음 |
| 예외 | 실패는 `raise`, 처리는 `try`/`except 구체예외` | `except:` 또는 `except Exception: pass` 로 모든 실패 삼키기 |
| 모듈·패키지 | 파일 하나가 모듈, 폴더가 패키지. `import order.logic` | 이름 충돌하는 파일 이름 (`json.py`) |
| `if __name__ == "__main__":` | 파일을 직접 실행할 때만 도는 진입점 | import 할 때 부작용 실행 |
| 가변 기본 인자 | `def add(item, bucket=[])` 의 리스트는 호출 사이에 공유돼요 | 기본값을 `None` 으로 두고 함수 안에서 생성 |

## 4. 이름 (컨벤션 C05, PEP 8)

| 대상 | 규칙 | 예 |
| --- | --- | --- |
| 함수·메서드·변수·모듈 파일 | `snake_case` | `calculate_order_total`, `order_total.py` |
| 클래스 | `PascalCase` | `OrderRecord` |
| 상수 | `UPPER_SNAKE_CASE` | `MAX_ORDER_ITEMS` |

## 5. 하네스 스택 (`python`)

| 항목 | 값 |
| --- | --- |
| 표지 | `pyproject.toml` |
| 확장자 | `.py` |
| 검사 | `ruff check .`, `ruff format --check .`, `pyright --warnings` (모두 종료 0) |
| 보호 | `uv.lock`, `poetry.lock` (패키지 매니저 전용) |

검사 명령은 위 문자열 그대로 실행해야 하네스가 그 검사로 인식해요. 그래서 `uv run ruff check .` 대신 `uv tool install` 로 설치한 `ruff check .` 를 써요.

## 6. AI가 자주 만드는 Python 실수

| 실수 | 바른 방향 |
| --- | --- |
| `except Exception: pass` | 경계 계약이 있는 구체 예외만 처리하고 나머지는 전달 |
| `# type: ignore`, `# noqa` | 원인 수정 (C06 억제 금지) |
| `utils.py`, `helpers.py` | 책임 이름의 모듈로 나누기 |
| 의존성을 `pip install` 로 전역 설치 | `uv add` 로 프로젝트에 기록 |

## 공식 자료

- Python 튜토리얼: https://docs.python.org/3/tutorial/
- PEP 8: https://peps.python.org/pep-0008/
- uv: https://docs.astral.sh/uv/
- Ruff: https://docs.astral.sh/ruff/
- Pyright 명령줄: https://github.com/microsoft/pyright/blob/main/docs/command-line.md
