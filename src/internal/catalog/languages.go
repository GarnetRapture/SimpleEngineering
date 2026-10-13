package catalog

import (
	"simpleharness/src/internal/harness"
	"simpleharness/src/internal/locale"
)

type Kind struct {
	ID    string
	Label locale.Text
}

var Kinds = []Kind{
	{ID: "web", Label: locale.Text{
		Ko: "웹 서비스·웹 앱 (브라우저·서버)",
		En: "Web service or web app (browser, server)",
		Zh: "Web 服务或 Web 应用（浏览器、服务器）",
	}},
	{ID: "program", Label: locale.Text{
		Ko: "일반 프로그램 (데스크톱·CLI·게임·네이티브)",
		En: "General program (desktop, CLI, game, native)",
		Zh: "通用程序（桌面、CLI、游戏、原生）",
	}},
}

type ProtectedPath struct {
	Pattern       string
	ShellSegments []string
	Target        locale.Text
	Reason        string
	Route         locale.Text
}

type Language struct {
	ID               string
	Label            string
	Stack            string
	RecommendedKinds []string
	VersionSource    string
	VersionCommand   string
	Protected        []ProtectedPath
}

var packageManagerRoute = locale.Text{
	Ko: "패키지 매니저 명령으로만 바꾼다",
	En: "change it only through package manager commands",
	Zh: "只能通过包管理器命令修改",
}

var producerRoute = locale.Text{
	Ko: "빌드·생산 도구로만 생성한다",
	En: "produce it only through its build or producer tool",
	Zh: "只能通过构建或生成工具产生",
}

var Languages = []Language{
	{
		ID: "typescript", Label: "TypeScript (Node.js)", Stack: "node-typescript",
		RecommendedKinds: []string{"web", "program"},
		VersionSource:    "package.json", VersionCommand: "npm version <x.y.z> --no-git-tag-version",
		Protected: []ProtectedPath{
			{Pattern: `^dist(/|$)`, ShellSegments: []string{"dist"}, Target: locale.Text{Ko: "`dist/`", En: "`dist/`", Zh: "`dist/`"}, Reason: "C13: generated output changes require its owning build or producer", Route: producerRoute},
			{Pattern: `^package(-lock)?\.json$`, ShellSegments: []string{}, Target: locale.Text{Ko: "`package.json`·lockfile", En: "`package.json` and the lockfile", Zh: "`package.json` 与锁文件"}, Reason: "C13: package manager only", Route: packageManagerRoute},
		},
	},
	{
		ID: "python", Label: "Python", Stack: "python",
		RecommendedKinds: []string{"web", "program"},
		VersionSource:    "pyproject.toml",
		Protected: []ProtectedPath{
			{Pattern: `^(uv|poetry)\.lock$`, ShellSegments: []string{}, Target: locale.Text{Ko: "`uv.lock`·`poetry.lock`", En: "`uv.lock` and `poetry.lock`", Zh: "`uv.lock` 与 `poetry.lock`"}, Reason: "C13: package manager only", Route: packageManagerRoute},
		},
	},
	{
		ID: "cpp", Label: "C++ (xmake)", Stack: "cpp-xmake",
		RecommendedKinds: []string{"program"},
		VersionSource:    "xmake.lua",
		Protected: []ProtectedPath{
			{Pattern: `^(build|\.xmake)(/|$)`, ShellSegments: []string{"build", ".xmake"}, Target: locale.Text{Ko: "`build/`·`.xmake/`", En: "`build/` and `.xmake/`", Zh: "`build/` 与 `.xmake/`"}, Reason: "C13: generated output changes require its owning build or producer", Route: producerRoute},
		},
	},
	{
		ID: "rust", Label: "Rust", Stack: "rust",
		RecommendedKinds: []string{"program"},
		VersionSource:    "Cargo.toml",
		Protected: []ProtectedPath{
			{Pattern: `^target(/|$)`, ShellSegments: []string{"target"}, Target: locale.Text{Ko: "`target/`", En: "`target/`", Zh: "`target/`"}, Reason: "C13: generated output changes require its owning build or producer", Route: producerRoute},
			{Pattern: `^Cargo\.lock$`, ShellSegments: []string{}, Target: locale.Text{Ko: "`Cargo.lock`", En: "`Cargo.lock`", Zh: "`Cargo.lock`"}, Reason: "C13: package manager only", Route: packageManagerRoute},
		},
	},
	{
		ID: "go", Label: "Go", Stack: "go",
		RecommendedKinds: []string{"web", "program"},
		VersionSource:    "git tag",
		Protected: []ProtectedPath{
			{Pattern: `^go\.sum$`, ShellSegments: []string{}, Target: locale.Text{Ko: "`go.sum`", En: "`go.sum`", Zh: "`go.sum`"}, Reason: "C13: package manager only", Route: packageManagerRoute},
		},
	},
	{
		ID: "php", Label: "PHP", Stack: "php",
		RecommendedKinds: []string{"web"},
		VersionSource:    "composer.json",
		Protected: []ProtectedPath{
			{Pattern: `^composer\.lock$`, ShellSegments: []string{}, Target: locale.Text{Ko: "`composer.lock`", En: "`composer.lock`", Zh: "`composer.lock`"}, Reason: "C13: package manager only", Route: packageManagerRoute},
		},
	},
}

var CommonProtected = []ProtectedPath{
	{
		Pattern:       `(^|/)(vendor|third-party)(/|$)`,
		ShellSegments: []string{"vendor", "third-party"},
		Target:        locale.Text{Ko: "모든 `vendor/`·`third-party/`", En: "every `vendor/` and `third-party/`", Zh: "所有 `vendor/` 与 `third-party/`"},
		Reason:        "C13: external source is protected",
		Route:         locale.Text{Ko: "사용자 지시가 있을 때만 교체한다", En: "replace it only on a user instruction", Zh: "仅在用户指示时替换"},
	},
}

func LanguageByID(id string) Language {
	for _, language := range Languages {
		if language.ID == id {
			return language
		}
	}
	panic("catalog: unknown project language " + id)
}

var Stacks = []harness.Stack{
	{
		Name:       "node-typescript",
		Markers:    []string{"package.json", "tsconfig.json"},
		Extensions: []string{".ts", ".tsx", ".mts", ".cts", ".js", ".jsx", ".mjs", ".cjs"},
		Checks: []harness.Check{
			{Kind: "typecheck", Command: "npx tsc --noEmit", Acceptance: harness.Acceptance{Exit: 0}},
			{Kind: "lint", Command: "npx eslint . --max-warnings 0", Acceptance: harness.Acceptance{Exit: 0}},
		},
		Forbidden: []harness.Forbidden{
			{Command: "npm run dev", Reason: "product execution belongs to the user"},
			{Command: "npm start", Reason: "product execution belongs to the user"},
			{Command: "vite", Reason: "product execution belongs to the user"},
			{Command: "npm run build", Reason: "a full build is not a verification step of this stack"},
		},
	},
	{
		Name:       "python",
		Markers:    []string{"pyproject.toml"},
		Extensions: []string{".py"},
		Checks: []harness.Check{
			{Kind: "lint", Command: "ruff check .", Acceptance: harness.Acceptance{Exit: 0}},
			{Kind: "format", Command: "ruff format --check .", Acceptance: harness.Acceptance{Exit: 0}},
			{Kind: "typecheck", Command: "pyright --warnings", Acceptance: harness.Acceptance{Exit: 0}},
		},
		Forbidden: []harness.Forbidden{},
	},
	{
		Name:       "cpp-xmake",
		Markers:    []string{"xmake.lua"},
		Extensions: []string{".cpp", ".cc", ".cxx", ".ixx", ".hpp", ".hh", ".h"},
		Checks: []harness.Check{
			{Kind: "build", Command: "xmake build", Acceptance: harness.Acceptance{Exit: 0, ForbidLine: `\bwarning\b`}},
		},
		Forbidden: []harness.Forbidden{},
	},
	{
		Name:       "rust",
		Markers:    []string{"Cargo.toml"},
		Extensions: []string{".rs"},
		Checks: []harness.Check{
			{Kind: "clippy", Command: "cargo clippy --all-targets -- -D warnings", Acceptance: harness.Acceptance{Exit: 0}},
		},
		Forbidden: []harness.Forbidden{
			{Command: "cargo build", Reason: "a full build is not a verification step of this stack"},
			{Command: "cargo run", Reason: "product execution belongs to the user"},
		},
	},
	{
		Name:       "go",
		Markers:    []string{"go.mod"},
		Extensions: []string{".go"},
		Checks: []harness.Check{
			{Kind: "vet", Command: "go vet ./...", Acceptance: harness.Acceptance{Exit: 0}},
			{Kind: "format", Command: "gofmt -l .", Acceptance: harness.Acceptance{Exit: 0, ForbidLine: `\.go$`}},
		},
		Forbidden: []harness.Forbidden{
			{Command: "go run", Reason: "product execution belongs to the user"},
			{Command: "go build", Reason: "a full build is not a verification step of this stack"},
		},
	},
	{
		Name:       "php",
		Markers:    []string{"composer.json"},
		Extensions: []string{".php"},
		Checks: []harness.Check{
			{Kind: "syntax", Command: "php -l", Acceptance: harness.Acceptance{Exit: 0}},
			{Kind: "analyse", Command: "vendor/bin/phpstan analyse", Acceptance: harness.Acceptance{Exit: 0}},
		},
		Forbidden: []harness.Forbidden{
			{Command: "php -S", Reason: "product execution belongs to the user"},
		},
	},
}
