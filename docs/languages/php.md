# PHP 핵심 지식

[← 언어별 핵심 지식](README.md)

PHP는 웹 서버에서 HTML을 만들어 내는 데 특화된 언어예요. WordPress 같은 CMS와 Laravel·Symfony 같은 프레임워크가 PHP로 돌아가요. 최근 PHP는 타입 선언과 정적 분석(PHPStan)으로 큰 프로젝트도 안전하게 관리할 수 있어요.

## 1. 설치

| 도구 | Windows | macOS | Linux (apt) |
| --- | --- | --- | --- |
| PHP | `winget install --id PHP.PHP.8.5 -e --source winget` | `brew install php` | `sudo apt-get install -y php-cli` |
| Composer (패키지 관리) | https://getcomposer.org/Composer-Setup.exe | `brew install composer` | `sudo apt-get install -y composer` |

## 2. 새 프로젝트

```sh
mkdir order-site
cd order-site
composer init
composer require --dev phpstan/phpstan
```

`composer.json` 에 자동 로딩 규칙(PSR-4)을 둬요. 네임스페이스 `App\` 가 `src/` 폴더에 대응해요.

```json
{
  "autoload": {
    "psr-4": { "App\\": "src/" }
  }
}
```

`composer dump-autoload` 로 반영해요. PHPStan 설정 `phpstan.neon`:

```yaml
parameters:
    level: 8
    paths:
        - src
```

| 명령 | 하는 일 |
| --- | --- |
| `php -l src/Order/OrderTotal.php` | 문법 검사 (파일 하나) |
| `vendor/bin/phpstan analyse` | 정적 분석 (`phpstan.neon` 의 경로·수준) |
| `composer require <패키지>`, `composer remove <패키지>` | 의존성 변경 (`composer.json`·`composer.lock` 을 Composer가 갱신) |
| `php -S localhost:8000` | 내장 개발 서버 (제품 실행이라 사용자가 해요) |

## 3. 핵심 개념

| 개념 | 설명 | 실수 예 |
| --- | --- | --- |
| `declare(strict_types=1);` | 파일 맨 위에 두면 스칼라 타입을 엄격하게 검사해요 | 문자열 `"3"` 이 정수 인자로 조용히 변환 |
| 타입 선언 | `function total(array $items): int` | 타입 없이 써서 PHPStan이 추론 못 함 |
| 네임스페이스·PSR-4 | `namespace App\Order;` + 폴더 구조 = 클래스 위치 | 클래스 이름과 파일 이름 불일치로 자동 로딩 실패 |
| 출력 이스케이프 | HTML에 쓸 때 `htmlspecialchars($value, ENT_QUOTES)` | 사용자 입력을 그대로 출력 → XSS |
| SQL 준비문 | PDO `prepare` + 바인딩 | 문자열 이어 붙이기 → SQL 주입 |
| 예외 | `throw new OrderNotFound(…)`, `try`/`catch (구체예외 $e)` | `@` 연산자로 오류 숨기기 |

## 4. 이름과 구조 (컨벤션 C04·C05)

| 대상 | 규칙 | 예 |
| --- | --- | --- |
| 함수·메서드 | `lowerCamelCase` | `calculateOrderTotal()` |
| 클래스 | `PascalCase`, 파일 이름 = 클래스 이름 | `OrderTotal` → `OrderTotal.php` |
| 상수 | `UPPER_SNAKE_CASE` | `MAX_ORDER_ITEMS` |

인라인 규칙을 켜면 PHP 템플릿에서도 `style="…"`, `<style>`, `onclick="…"` 를 훅이 막아요. 스타일은 CSS 파일의 클래스로, 이벤트는 스크립트 파일의 이름 있는 함수로 둬요.

## 5. 하네스 스택 (`php`)

| 항목 | 값 |
| --- | --- |
| 표지 | `composer.json` |
| 확장자 | `.php` |
| 검사 | `php -l <파일>`, `vendor/bin/phpstan analyse` (모두 종료 0) |
| 금지 | `php -S` (제품 실행) |
| 보호 | `composer.lock` (패키지 매니저 전용), `vendor/` (외부 코드) |

## 6. AI가 자주 만드는 PHP 실수

| 실수 | 바른 방향 |
| --- | --- |
| `$_GET` 값을 그대로 출력·쿼리 | 검증 + 이스케이프 + 준비문 |
| `@` 로 경고 숨기기, `// @phpstan-ignore` | 원인 수정 (C06) |
| 한 파일에 HTML·SQL·로직 혼합 | 도메인별 템플릿·저장소·로직 소유 파일로 분리 |
| `vendor/` 안 코드 수정 | Composer로 버전 변경, 외부 코드는 보호 대상 |

## 공식 자료

- PHP 매뉴얼: https://www.php.net/manual/
- Composer: https://getcomposer.org/doc/
- PSR-4: https://www.php-fig.org/psr/psr-4/
- PSR-12: https://www.php-fig.org/psr/psr-12/
- PHPStan: https://phpstan.org/user-guide/getting-started
