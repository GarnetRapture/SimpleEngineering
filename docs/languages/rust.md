# Rust 핵심 지식

[← 언어별 핵심 지식](README.md)

Rust는 C++만큼 빠르면서, 메모리 안전 규칙을 **컴파일러가 강제**하는 언어예요. 해제 후 사용, 데이터 경쟁 같은 결함을 컴파일 단계에서 막는 대신, 그 규칙(소유권·빌림)을 이해해야 컴파일이 통과해요.

## 1. 설치

| 운영체제 | 명령 |
| --- | --- |
| Windows | `winget install --id Rustlang.Rustup -e --source winget` |
| macOS | `brew install rustup` |
| Linux | `curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs \| sh` 또는 apt `rustup` |

Windows의 기본 툴체인(MSVC ABI)은 링커와 Windows SDK가 필요해요. Visual Studio가 없으면 `rustup-init` 이 필요한 구성 요소 설치를 제안해요(rustup 공식 문서). C++ 빌드 도구를 직접 설치하려면 [cpp.md](cpp.md#1-준비물)의 MSVC 명령을 써요.

## 2. 새 프로젝트

```sh
cargo new order-tool
cd order-tool
```

생성물: `Cargo.toml`(패키지 정보·의존성), `src/main.rs`.

| 명령 | 하는 일 |
| --- | --- |
| `cargo check` | 타입·빌림 검사 (실행 파일은 만들지 않음, 빠름) |
| `cargo clippy --all-targets -- -D warnings` | 린트. 경고를 오류로 처리 |
| `cargo fmt` | 공식 포맷 적용 |
| `cargo test` | 테스트 실행 |
| `cargo add serde` | 의존성 추가 (`Cargo.toml`·`Cargo.lock` 을 패키지 매니저가 갱신) |
| `cargo run` | 실행 (제품 실행이라 사용자가 해요) |

## 3. 핵심 개념

### 소유권 (The Rust Programming Language 4.1장)

1. 모든 값은 **소유자** 하나를 가져요.
2. 한 시점에 소유자는 **하나**뿐이에요.
3. 소유자가 범위를 벗어나면 값은 **해제(drop)** 돼요.

```rust
let order = Order::new();
let moved = order;          // 소유권이 moved로 이동
// println!("{:?}", order); // 컴파일 오류: order는 더 이상 소유자가 아님
```

### 빌림 (참조)

| 규칙 | 의미 |
| --- | --- |
| 불변 참조 `&T` 는 여러 개 동시에 가능 | 여럿이 읽기만 할 수 있어요 |
| 가변 참조 `&mut T` 는 한 번에 하나만, 그동안 불변 참조 불가 | 쓰는 쪽이 하나일 때만 바꿀 수 있어요 → 데이터 경쟁이 컴파일 단계에서 사라져요 |
| 참조는 대상보다 오래 살 수 없음 | 해제된 값을 가리키는 참조가 생기지 않아요 |

### 오류 처리

- `null` 이 없어요. 값이 없을 수 있으면 `Option<T>`(`Some`/`None`), 실패할 수 있으면 `Result<T, E>`(`Ok`/`Err`).
- `?` 연산자는 `Err` 를 호출자에게 그대로 돌려줘요.
- `unwrap()`·`expect()` 는 실패 시 프로그램을 멈춰요(panic). 경계 계약이 "여기서 실패는 불가능"을 보장할 때만 써요.

### 트레이트와 모듈

- 트레이트(trait)는 타입이 가져야 할 동작의 계약이에요(`impl Display for Order`).
- 모듈은 파일·폴더 구조와 같아요(`mod order;` → `src/order.rs` 또는 `src/order/mod.rs`). 기본은 비공개, `pub` 로 공개해요.

## 4. 이름 (컨벤션 C05)

| 대상 | 규칙 | 예 |
| --- | --- | --- |
| 함수·변수·모듈·파일 | `snake_case` | `calculate_order_total`, `order_total.rs` |
| 타입·트레이트·열거형 | `PascalCase` | `OrderStatus` |
| 상수·정적 변수 | `UPPER_SNAKE_CASE` | `MAX_ORDER_ITEMS` |

## 5. 하네스 스택 (`rust`)

| 항목 | 값 |
| --- | --- |
| 표지 | `Cargo.toml` |
| 확장자 | `.rs` |
| 검사 | `cargo clippy --all-targets -- -D warnings` (종료 0) |
| 금지 | `cargo build` (전체 빌드는 이 스택의 검증 단계가 아님), `cargo run` (제품 실행은 사용자 몫) |
| 보호 | `target/` (생성물), `Cargo.lock` (패키지 매니저 전용) |

## 6. AI가 자주 만드는 Rust 실수

| 실수 | 바른 방향 |
| --- | --- |
| 빌림 오류를 `.clone()` 남발로 피하기 | 소유권 흐름을 다시 설계하거나 참조로 넘기기 |
| `unwrap()` 으로 오류 무시 | `?` 로 전달하고 경계에서 처리 |
| `#[allow(...)]` 로 경고 끄기 | 경고 원인 수정 (C06) |
| `unsafe` 로 컴파일러 우회 | 안전한 API로 표현. 꼭 필요하면 불변식을 증명할 수 있는 최소 범위만 |

## 공식 자료

- The Rust Programming Language: https://doc.rust-lang.org/book/
- 소유권: https://doc.rust-lang.org/book/ch04-01-what-is-ownership.html
- rustup Windows MSVC: https://rust-lang.github.io/rustup/installation/windows-msvc.html
- 이름 규칙(API Guidelines): https://rust-lang.github.io/api-guidelines/naming.html
- Clippy: https://doc.rust-lang.org/clippy/
