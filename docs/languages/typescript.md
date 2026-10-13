# TypeScript 핵심 지식

[← 언어별 핵심 지식](README.md)

TypeScript는 JavaScript에 **정적 타입**을 더한 언어예요. 브라우저 프런트엔드(React 등), Node.js 서버, 도구 스크립트에 써요. 타입은 컴파일 때만 검사되고 실행 코드에서는 지워져요. 그래서 타입 검사 통과는 "타입이 서로 맞다"는 증거이지, 실행 결과가 맞다는 증거는 아니에요.

## 1. 설치

| 운영체제 | Node.js LTS |
| --- | --- |
| Windows | `winget install --id OpenJS.NodeJS.LTS -e --source winget` |
| macOS | `brew install node` |
| Linux | `sudo apt-get install -y nodejs npm` |

## 2. 새 프로젝트

```sh
mkdir order-web
cd order-web
npm init -y
npm install --save-dev typescript
npx tsc --init
npm init @eslint/config@latest
```

| 파일 | 역할 |
| --- | --- |
| `package.json` | 패키지 정보·스크립트·의존성 (하네스 표지) |
| `package-lock.json` | 설치된 정확한 버전 (잠금 파일) |
| `tsconfig.json` | 타입 검사 설정 (하네스 표지). `"strict": true` 를 켜요 |
| `eslint.config.*` | 린트 규칙 |

| 명령 | 하는 일 |
| --- | --- |
| `npx tsc --noEmit` | 타입 검사만 (파일을 만들지 않음) |
| `npx eslint . --max-warnings 0` | 린트, 경고 1개라도 있으면 실패 |
| `npm install <패키지>` / `npm uninstall <패키지>` | 의존성 변경 (`package.json`·잠금 파일을 npm이 갱신) |
| `npm pkg set scripts.check="tsc --noEmit"` | `package.json` 스크립트를 명령으로 수정 |
| `npm run dev`, `npm start`, `vite`, `npm run build` | 개발 서버·실행·빌드 (사용자가 해요) |

## 3. 핵심 개념

| 개념 | 설명 | 실수 예 |
| --- | --- | --- |
| `strict` 모드 | `null`·`undefined` 검사, 암묵적 `any` 금지 등 엄격한 검사 묶음 | 끄고 시작해서 나중에 대량 오류 |
| `unknown` vs `any` | `any` 는 검사를 끄고, `unknown` 은 확인(좁히기) 전에는 못 쓰게 해요 | 외부 JSON을 `any` 로 받아 오타 필드도 통과 |
| 유니온과 좁히기 | `type Status = 'pending' \| 'paid'`, `if (status === 'paid')` 안에서 타입이 좁혀져요 | 문자열 아무거나 허용하는 `string` |
| 타입 단언 `as` | "내가 더 잘 안다"고 검사를 덮어써요 | `data as Order` 로 실제 검증 없이 통과 |
| 모듈 | `export`/`import` 로 파일 간 연결 | 순환 import |
| `async`/`await` | 비동기 결과를 기다림. `Promise` 를 돌려줘요 | `await` 없이 호출해 실패가 사라짐 |

## 4. 구조와 이름 (컨벤션 C04·C05)

- 함수 `lowerCamelCase`, 타입·컴포넌트 `PascalCase`, 상수 `UPPER_SNAKE_CASE`, 훅 `use` + `PascalCase`.
- 내보내는 파일 이름은 심볼 이름과 같게: `calculateOrderTotal.ts` → `export function calculateOrderTotal`.
- **인라인 타입 금지** 규칙을 켜면 `type`·`interface`·`enum` 선언은 도메인의 `types/` 폴더나 `*.types.ts`(또는 `.d.ts`)에만 둬요.

```text
src/order/
  types/orderStatus.types.ts      ← OrderStatus 타입의 단일 소유자
  logic/calculateOrderTotal.ts
  ui/OrderSummary.tsx
  ui/orderSummary.module.css      ← 스타일 소유 파일
```

## 5. 하네스 스택 (`node-typescript`)

| 항목 | 값 |
| --- | --- |
| 표지 | `package.json` + `tsconfig.json` (둘 다 있어야 켜짐) |
| 확장자 | `.ts`, `.tsx`, `.mts`, `.cts`, `.js`, `.jsx`, `.mjs`, `.cjs` |
| 검사 | `npx tsc --noEmit`, `npx eslint . --max-warnings 0` (모두 종료 0) |
| 금지 | `npm run dev`, `npm start`, `vite` (제품 실행), `npm run build` (전체 빌드는 검증 단계 아님) |
| 보호 | `dist/` (생성물), `package.json`·`package-lock.json` (패키지 매니저 전용) |

켜진 인라인 규칙(훅 차단):

| 규칙 | 차단하는 코드 | 바른 방향 |
| --- | --- | --- |
| `no-inline-type` | 컴포넌트·로직 파일 안 `type Props = {…}` | `types/` 또는 `*.types.ts` 로 이동 |
| `no-inline-style` | `<div style={{ padding: 8 }}>` | 스타일 소유 파일의 이름 있는 클래스 |
| `no-inline-handler` | `<button onClick={() => save()}>` | `function handleOrderSave()` 정의 후 `onClick={handleOrderSave}` |
| `no-inline-comment` | `// 합계 계산` | 이름이 의도를 말하게: `calculateOrderTotal` |

## 6. AI가 자주 만드는 TypeScript 실수

| 실수 | 바른 방향 |
| --- | --- |
| `any`, `as` 로 타입 오류 덮기 | 실제 타입 정의, `unknown` + 검증 |
| `// @ts-ignore`, `eslint-disable` | 원인 수정 (C06 억제 금지) |
| `?? []`, `catch { return null }` 같은 조용한 폴백 | 실패를 호출자에게 전달 |
| `package.json` 직접 편집 | `npm install`, `npm pkg set` |

## 공식 자료

- TypeScript Handbook: https://www.typescriptlang.org/docs/handbook/intro.html
- `strict`: https://www.typescriptlang.org/tsconfig/#strict
- ESLint 시작하기: https://eslint.org/docs/latest/use/getting-started
- npm pkg: https://docs.npmjs.com/cli/commands/npm-pkg
