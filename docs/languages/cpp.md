# C++ 핵심 지식과 세팅 (xmake)

[← 언어별 핵심 지식](README.md)

C++은 하드웨어에 가깝게 동작하는 컴파일 언어예요. 게임·엔진·브라우저·데이터베이스처럼 성능과 메모리 배치를 직접 다뤄야 하는 프로그램에 써요. 대신 메모리를 잘못 다루면 프로그램이 조용히 잘못된 값을 내거나 멈춰요. 그래서 C++ 프로젝트에서는 **컴파일러 경고를 0으로 유지하는 것**이 가장 중요한 습관이에요.

## 1. 준비물

| 구성 | 역할 | Windows | macOS | Linux (apt) |
| --- | --- | --- | --- | --- |
| 컴파일러 | 소스를 기계어로 변환 | MSVC (Visual Studio Build Tools), LLVM clang, GCC(WinLibs) | `xcode-select --install` (Apple clang), `brew install llvm`, `brew install gcc` | `sudo apt-get install -y build-essential` (GCC), `llvm clang clang-tidy clangd` |
| 빌드 도구 | 어떤 파일을 어떤 옵션으로 컴파일·링크할지 관리 | `winget install --id Xmake-io.Xmake -e --source winget` | `brew install xmake` | `sudo apt-get install -y xmake` |
| 정적 분석 | 컴파일러가 못 잡는 위험 패턴 | LLVM의 `clang-tidy`, `clangd` | `brew install llvm` | `clang-tidy clangd` |

Windows에서 MSVC는 다음 명령으로 C++ 빌드 도구 워크로드만 설치해요(SimpleHarness CLI의 추천 도구와 같은 명령).

```powershell
winget install --id Microsoft.VisualStudio.BuildTools -e --source winget --override "--wait --passive --add Microsoft.VisualStudio.Workload.VCTools --includeRecommended"
```

## 2. 새 프로젝트 만들기

```sh
xmake create -l c++ -t console order-tool
cd order-tool
```

생성되는 파일(실측: xmake 3.1.1): `xmake.lua`, `src/main.cpp`, `.gitignore`.

`xmake.lua` 를 다음처럼 정리해요. 언어 표준과 경고 수준을 프로젝트 전체에 한 번 선언해요.

```lua
add_rules("mode.debug", "mode.release")
set_languages("c++26")
set_warnings("all", "error")

target("order-tool")
    set_kind("binary")
    add_files("src/*.cpp")
```

| 줄 | 뜻 |
| --- | --- |
| `add_rules("mode.debug", "mode.release")` | `xmake f -m debug` / `-m release` 로 빌드 모드 전환 |
| `set_languages("c++26")` | C++26 표준. MSVC에서는 xmake가 `-std:c++latest` 로 전달해요 (실측: xmake 3.1.1 + MSVC, 빌드 성공) |
| `set_warnings("all", "error")` | 경고를 모두 켜고 경고를 오류로 다룸 |
| `set_kind("binary")` | 실행 파일. 라이브러리는 `static`·`shared` |

| 명령 | 하는 일 |
| --- | --- |
| `xmake f -m debug` | 설정(구성) — 모드·툴체인 선택. 예: `xmake f --toolchain=clang` |
| `xmake` 또는 `xmake build` | 컴파일·링크 |
| `xmake run` | 실행 (제품 실행이라 사용자가 해요) |
| `xmake project -k compile_commands` | `compile_commands.json` 생성 → clangd·clang-tidy가 실제 빌드 옵션으로 분석 |

## 3. 핵심 개념

| 개념 | 설명 | 실수 예 |
| --- | --- | --- |
| 선언과 정의 | 헤더(`.hpp`)에 "이런 함수가 있다"(선언), 소스(`.cpp`)에 "이렇게 동작한다"(정의) | 헤더에 정의를 넣어 여러 파일에서 중복 정의 링크 오류 |
| 번역 단위 | 컴파일러는 `.cpp` 파일 하나(와 그것이 포함한 헤더)를 한 단위로 컴파일하고, 링커가 단위들을 합쳐요 | 선언만 있고 정의가 링크되지 않아 `unresolved external symbol` |
| RAII | 자원(메모리·파일·락)을 객체 생성 때 얻고 소멸 때 놓는 방식. 범위를 벗어나면 자동 해제 | `new` 후 `delete` 를 빠뜨려 누수. `std::unique_ptr`·`std::vector` 를 쓰면 해결 |
| 소유권 | 누가 자원을 해제할 책임이 있는가. `std::unique_ptr` = 단독 소유, `std::shared_ptr` = 공동 소유, 참조·포인터 = 빌려 씀 | 이미 해제된 객체를 가리키는 참조(dangling) 사용 |
| 값·참조 전달 | 큰 객체는 `const T&` 로 넘겨 복사를 피하고, 바꿀 때만 `T&` | 매번 복사해서 느려지거나, 바꾸려 했는데 복사본만 바뀜 |
| 정의되지 않은 동작(UB) | 배열 범위 밖 접근, 해제 후 사용, 부호 있는 정수 오버플로 등은 결과가 정해져 있지 않아요 | "내 PC에서는 됐는데" 다른 컴파일러·최적화에서 깨짐 |
| `const` 정확성 | 바꾸지 않는 값·메서드에 `const` 를 붙여 의도를 타입으로 표현 | 읽기 전용이어야 할 상태를 실수로 변경 |

## 4. 이름과 구조 (컨벤션 C04·C05·C10)

- 함수·메서드 `lowerCamelCase`, 타입 `PascalCase`, 선언(`.hpp`)과 구현(`.cpp`)은 쌍으로 둬요.
- 소스는 `{영역}/{기능}/` 으로 둬요(C10). 예: `src/render/mesh/MeshBuffer.hpp`, `src/render/mesh/MeshBuffer.cpp`.
- 생성물 `build/`, `.xmake/` 는 하네스가 편집을 막아요(빌드 도구로만 생성).

## 5. 하네스 스택 (`cpp-xmake`)

| 항목 | 값 |
| --- | --- |
| 표지 | `xmake.lua` |
| 확장자 | `.cpp`, `.cc`, `.cxx`, `.ixx`, `.hpp`, `.hh`, `.h` |
| 검사 | `xmake build` — 종료 0이고 출력에 `warning` 이 들어간 줄이 없어야 통과 |
| 보호 | `build/`, `.xmake/` |

에이전트가 C++ 파일을 편집하면 `xmake build` 가 "남은 검사"로 표시되고, 통과해야 해제돼요.

## 6. AI가 자주 만드는 C++ 실수

| 실수 | 바른 방향 |
| --- | --- |
| 원시 `new`/`delete` | `std::make_unique`, 컨테이너 |
| 경고를 `#pragma warning(disable…)` 로 끄기 | 경고 원인 수정 (컨벤션 C06이 억제를 금지) |
| 헤더에 `using namespace std;` | 이름 충돌이 모든 포함 파일로 퍼져요. 필요한 곳에서 `std::` 를 써요 |
| 매직 숫자 (`if (state == 3)`) | `enum class` 와 이름 있는 상수 |
| 오류를 무시하는 `catch (...) {}` | 경계 계약이 있는 실패만 처리하고 나머지는 전달 |

## 공식 자료

- xmake: https://xmake.io
- C++ Core Guidelines: https://isocpp.github.io/CppCoreGuidelines/CppCoreGuidelines
- cppreference: https://en.cppreference.com
- MSVC 설치: https://learn.microsoft.com/cpp/build/vscpp-step-0-installation
