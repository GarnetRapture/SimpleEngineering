# Go 핵심 지식

[← 언어별 핵심 지식](README.md)

Go는 단순한 문법, 빠른 컴파일, 내장 동시성(goroutine)을 가진 언어예요. 서버, CLI, 네트워크 도구에 많이 써요. SimpleHarness CLI도 Go로 만들었어요(`go.mod`: `go 1.26`).

## 1. 설치

| 운영체제 | 명령 |
| --- | --- |
| Windows | `winget install --id GoLang.Go -e --source winget` |
| macOS | `brew install go` |
| Linux | `sudo apt-get install -y golang-go` 또는 https://go.dev/dl/ |

## 2. 새 프로젝트

```sh
mkdir order-service
cd order-service
go mod init example.com/order-service
```

`go.mod` 가 모듈 이름과 Go 버전, 의존성을 기록해요.

| 명령 | 하는 일 |
| --- | --- |
| `go vet ./...` | 의심스러운 코드 패턴 검사 (모든 패키지) |
| `gofmt -l .` | 공식 포맷과 다른 파일 목록 출력 (없으면 출력 없음) |
| `gofmt -w <파일>` | 포맷 적용 |
| `go test ./...` | 테스트 |
| `go get example.com/lib@<버전>`, `go mod tidy` | 의존성 추가·정리 (`go.mod`·`go.sum` 을 도구가 갱신) |
| `go run .`, `go build` | 실행·빌드 (제품 실행과 전체 빌드는 하네스 검증 단계가 아니에요) |

## 3. 핵심 개념

| 개념 | 설명 |
| --- | --- |
| 패키지 = 폴더 | 한 폴더의 `.go` 파일은 같은 `package` 이름을 가져요. 실행 프로그램은 `package main` + `func main()` |
| 대문자 = 공개 | 이름이 대문자로 시작하면 다른 패키지에서 쓸 수 있고(exported), 소문자면 패키지 안에서만 써요 |
| 오류는 값 | 함수가 `(결과, error)` 를 돌려주고 호출자가 `if err != nil` 로 처리해요. 예외(throw)가 없어요 |
| `defer` | 함수가 끝날 때 실행할 정리 작업 예약 (`defer file.Close()`) |
| 인터페이스는 암묵적 | 메서드를 갖추면 선언 없이 인터페이스를 만족해요 |
| 제로 값 | 선언만 한 변수는 타입별 기본값(0, `""`, `nil`)을 가져요 |
| goroutine·channel | `go f()` 로 동시 실행, channel로 값을 주고받아요. 공유 메모리는 `sync` 로 보호해요 |

```go
func loadOrder(path string) (Order, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Order{}, fmt.Errorf("load order %s: %w", path, err)
	}
	return parseOrder(data)
}
```

`%w` 로 감싸면 호출자가 `errors.Is` 로 원래 오류를 확인할 수 있어요. 오류를 버리는 `_ = err` 는 실패를 숨겨요(컨벤션 C06).

## 4. 이름 (컨벤션 C05)

| 대상 | 규칙 | 출처 |
| --- | --- | --- |
| 함수·메서드·타입·상수 | `MixedCaps` (공개는 대문자, 비공개는 소문자 시작). 상수도 `MixedCaps` | Effective Go |
| 패키지 폴더 | 소문자 한 단어, `util`·`common`·`misc` 같은 의미 없는 이름 금지 | Effective Go, Go 블로그 "Package names" |
| 파일 | `snake_case.go` | 프로젝트 선택 |

패키지 이름이 호출부에 붙으므로 중복을 피해요: `order.Order` 보다 `order.Record`, `http.HTTPServer` 보다 `http.Server`.

## 5. 하네스 스택 (`go`)

| 항목 | 값 |
| --- | --- |
| 표지 | `go.mod` |
| 확장자 | `.go` |
| 검사 | `go vet ./...` (종료 0), `gofmt -l .` (종료 0이고 `.go` 로 끝나는 줄이 없어야 통과 = 포맷이 다른 파일 0개) |
| 금지 | `go run` (제품 실행), `go build` (전체 빌드는 검증 단계 아님) |
| 보호 | `go.sum` (도구 전용) |
| 버전 원천 | git 태그 |

## 6. AI가 자주 만드는 Go 실수

| 실수 | 바른 방향 |
| --- | --- |
| `err` 를 무시하거나 `panic` 으로 처리 | 맥락을 붙여 반환 (`fmt.Errorf("…: %w", err)`) |
| `utils` 패키지 | 책임 이름의 패키지로 나누기 |
| goroutine 종료를 기다리지 않음 | `sync.WaitGroup`, `context` 로 생명주기 관리 |
| 포맷이 다른 코드 | `gofmt -w` 적용 후 `gofmt -l .` 결과 0줄 확인 |

## 공식 자료

- Effective Go: https://go.dev/doc/effective_go
- Package names: https://go.dev/blog/package-names
- Go Tour: https://go.dev/tour/
- 오류 처리: https://go.dev/blog/go1.13-errors
