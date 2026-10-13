package catalog

import "simpleharness/src/internal/locale"

const (
	ManagerWinget = "winget"
	ManagerBrew   = "brew"
	ManagerApt    = "apt"
)

type Tool struct {
	ID       string
	Label    string
	Purpose  locale.Text
	Commands map[string][]string
	Detector string
	Packages map[string][]string
	Manual   map[string]string
	Homepage string
}

func (tool Tool) PresenceCommands(operatingSystem string) []string {
	if commands, found := tool.Commands[operatingSystem]; found {
		return commands
	}
	return tool.Commands["*"]
}

func winget(id string, extra ...string) []string {
	return append([]string{"--id", id, "-e", "--source", "winget"}, extra...)
}

var Tools = map[string]Tool{
	"git": {
		ID: "git", Label: "Git", Homepage: "https://git-scm.com/downloads",
		Purpose:  locale.Text{Ko: "버전 관리·커밋 가드(.githooks)", En: "version control and the commit guard (.githooks)", Zh: "版本控制与提交守卫（.githooks）"},
		Commands: map[string][]string{"*": {"git"}},
		Packages: map[string][]string{ManagerWinget: winget("Git.Git"), ManagerBrew: {"git"}, ManagerApt: {"git"}},
	},
	"pwsh": {
		ID: "pwsh", Label: "PowerShell 7", Homepage: "https://learn.microsoft.com/powershell/scripting/install/install-powershell-on-windows",
		Purpose:  locale.Text{Ko: "cmd·Windows PowerShell 5.1을 대신할 최신 셸 (pwsh)", En: "modern shell replacing cmd and Windows PowerShell 5.1 (pwsh)", Zh: "替代 cmd 与 Windows PowerShell 5.1 的现代 shell（pwsh）"},
		Commands: map[string][]string{"*": {"pwsh"}},
		Packages: map[string][]string{ManagerWinget: {"--id", "Microsoft.PowerShell", "--source", "winget"}},
	},
	"windows-terminal": {
		ID: "windows-terminal", Label: "Windows Terminal", Homepage: "https://learn.microsoft.com/windows/terminal/install",
		Purpose:  locale.Text{Ko: "탭·한글·색상을 지원하는 터미널", En: "terminal with tabs, Unicode, and color", Zh: "支持标签页、Unicode 与颜色的终端"},
		Commands: map[string][]string{"*": {"wt"}},
		Packages: map[string][]string{ManagerWinget: {"--id", "Microsoft.WindowsTerminal", "--source", "winget"}},
	},
	"node": {
		ID: "node", Label: "Node.js LTS", Homepage: "https://nodejs.org",
		Purpose:  locale.Text{Ko: "훅·커밋 가드·서브 에이전트 실행기 실행 (필수)", En: "runs the hooks, commit guard, and subagent launcher (required)", Zh: "运行钩子、提交守卫与子代理启动器（必需）"},
		Commands: map[string][]string{"*": {"node"}},
		Packages: map[string][]string{ManagerWinget: winget("OpenJS.NodeJS.LTS"), ManagerBrew: {"node"}, ManagerApt: {"nodejs", "npm"}},
	},
	"jq": {
		ID: "jq", Label: "jq", Homepage: "https://jqlang.org",
		Purpose:  locale.Text{Ko: "JSON 조회·변환 (인자로 쿼리)", En: "JSON query and transform (query as an argument)", Zh: "JSON 查询与转换（以参数传入查询）"},
		Commands: map[string][]string{"*": {"jq"}},
		Packages: map[string][]string{ManagerWinget: winget("jqlang.jq"), ManagerBrew: {"jq"}, ManagerApt: {"jq"}},
	},
	"ripgrep": {
		ID: "ripgrep", Label: "ripgrep (rg)", Homepage: "https://github.com/BurntSushi/ripgrep",
		Purpose:  locale.Text{Ko: "빠른 코드 검색", En: "fast code search", Zh: "快速代码搜索"},
		Commands: map[string][]string{"*": {"rg"}},
		Packages: map[string][]string{ManagerWinget: winget("BurntSushi.ripgrep.MSVC"), ManagerBrew: {"ripgrep"}, ManagerApt: {"ripgrep"}},
	},
	"sqlite3": {
		ID: "sqlite3", Label: "SQLite", Homepage: "https://sqlite.org/download.html",
		Purpose:  locale.Text{Ko: "SQL 파일 DB 조회 (인자로 SQL)", En: "file database queries (SQL as an argument)", Zh: "文件数据库查询（以参数传入 SQL）"},
		Commands: map[string][]string{"*": {"sqlite3"}},
		Packages: map[string][]string{ManagerWinget: winget("SQLite.SQLite"), ManagerBrew: {"sqlite"}, ManagerApt: {"sqlite3"}},
	},
	"duckdb": {
		ID: "duckdb", Label: "DuckDB", Homepage: "https://duckdb.org/docs/installation",
		Purpose:  locale.Text{Ko: "CSV·JSON·Parquet 분석 SQL (인자로 SQL)", En: "analytical SQL over CSV, JSON, and Parquet (SQL as an argument)", Zh: "对 CSV、JSON、Parquet 进行分析型 SQL（以参数传入 SQL）"},
		Commands: map[string][]string{"*": {"duckdb"}},
		Packages: map[string][]string{ManagerWinget: winget("DuckDB.cli"), ManagerBrew: {"duckdb"}},
		Manual:   map[string]string{"linux": "curl https://install.duckdb.org | sh"},
	},
	"python": {
		ID: "python", Label: "Python 3.14", Homepage: "https://www.python.org/downloads/",
		Purpose:  locale.Text{Ko: "Python 런타임", En: "Python runtime", Zh: "Python 运行时"},
		Commands: map[string][]string{"windows": {"py"}, "*": {"python3"}},
		Packages: map[string][]string{ManagerWinget: winget("Python.Python.3.14"), ManagerBrew: {"python@3.14"}, ManagerApt: {"python3"}},
	},
	"uv": {
		ID: "uv", Label: "uv", Homepage: "https://docs.astral.sh/uv/",
		Purpose:  locale.Text{Ko: "Python 패키지·가상 환경 관리", En: "Python package and environment manager", Zh: "Python 包与环境管理器"},
		Commands: map[string][]string{"*": {"uv"}},
		Packages: map[string][]string{ManagerWinget: winget("astral-sh.uv"), ManagerBrew: {"uv"}},
		Manual:   map[string]string{"linux": "curl -LsSf https://astral.sh/uv/install.sh | sh"},
	},
	"go": {
		ID: "go", Label: "Go", Homepage: "https://go.dev/dl/",
		Purpose:  locale.Text{Ko: "Go 툴체인", En: "Go toolchain", Zh: "Go 工具链"},
		Commands: map[string][]string{"*": {"go"}},
		Packages: map[string][]string{ManagerWinget: winget("GoLang.Go"), ManagerBrew: {"go"}, ManagerApt: {"golang-go"}},
	},
	"rustup": {
		ID: "rustup", Label: "Rustup", Homepage: "https://rustup.rs",
		Purpose:  locale.Text{Ko: "Rust 툴체인 설치·관리", En: "Rust toolchain installer and manager", Zh: "Rust 工具链安装与管理"},
		Commands: map[string][]string{"*": {"rustup"}},
		Packages: map[string][]string{ManagerWinget: winget("Rustlang.Rustup"), ManagerBrew: {"rustup"}, ManagerApt: {"rustup"}},
	},
	"llvm": {
		ID: "llvm", Label: "LLVM (clang·clang-tidy·clangd)", Homepage: "https://llvm.org",
		Purpose:  locale.Text{Ko: "C/C++ 컴파일러와 정적 분석", En: "C/C++ compiler and static analysis", Zh: "C/C++ 编译器与静态分析"},
		Commands: map[string][]string{"*": {"clang"}},
		Packages: map[string][]string{ManagerWinget: winget("LLVM.LLVM"), ManagerBrew: {"llvm"}, ManagerApt: {"llvm", "clang", "clang-tidy", "clangd"}},
	},
	"gcc": {
		ID: "gcc", Label: "GCC", Homepage: "https://gcc.gnu.org",
		Purpose:  locale.Text{Ko: "GNU C/C++ 컴파일러", En: "GNU C/C++ compiler", Zh: "GNU C/C++ 编译器"},
		Commands: map[string][]string{"*": {"gcc"}},
		Packages: map[string][]string{ManagerWinget: winget("BrechtSanders.WinLibs.POSIX.UCRT"), ManagerBrew: {"gcc"}, ManagerApt: {"build-essential"}},
	},
	"msvc": {
		ID: "msvc", Label: "Visual Studio Build Tools (MSVC)", Homepage: "https://visualstudio.microsoft.com/downloads/",
		Purpose:  locale.Text{Ko: "Windows C++·Rust 링커와 컴파일러", En: "Windows C++ and Rust compiler and linker", Zh: "Windows C++ 与 Rust 的编译器和链接器"},
		Detector: "msvc",
		Packages: map[string][]string{ManagerWinget: winget("Microsoft.VisualStudio.BuildTools", "--override", "--wait --passive --add Microsoft.VisualStudio.Workload.VCTools --includeRecommended")},
	},
	"xmake": {
		ID: "xmake", Label: "xmake", Homepage: "https://xmake.io",
		Purpose:  locale.Text{Ko: "C++ 빌드 도구", En: "C++ build tool", Zh: "C++ 构建工具"},
		Commands: map[string][]string{"*": {"xmake"}},
		Packages: map[string][]string{ManagerWinget: winget("Xmake-io.Xmake"), ManagerBrew: {"xmake"}, ManagerApt: {"xmake"}},
	},
	"php": {
		ID: "php", Label: "PHP", Homepage: "https://www.php.net/downloads",
		Purpose:  locale.Text{Ko: "PHP 런타임", En: "PHP runtime", Zh: "PHP 运行时"},
		Commands: map[string][]string{"*": {"php"}},
		Packages: map[string][]string{ManagerWinget: winget("PHP.PHP.8.5"), ManagerBrew: {"php"}, ManagerApt: {"php-cli"}},
	},
	"composer": {
		ID: "composer", Label: "Composer", Homepage: "https://getcomposer.org/download/",
		Purpose:  locale.Text{Ko: "PHP 패키지 관리", En: "PHP package manager", Zh: "PHP 包管理器"},
		Commands: map[string][]string{"*": {"composer"}},
		Packages: map[string][]string{ManagerBrew: {"composer"}, ManagerApt: {"composer"}},
		Manual:   map[string]string{"windows": "https://getcomposer.org/Composer-Setup.exe"},
	},
	"claude-code": {
		ID: "claude-code", Label: "Claude Code", Homepage: "https://code.claude.com/docs/en/setup",
		Purpose:  locale.Text{Ko: "Claude Code CLI (메인·서브 에이전트 실행)", En: "Claude Code CLI (main or subagent runtime)", Zh: "Claude Code CLI（主代理或子代理运行时）"},
		Commands: map[string][]string{"*": {"claude"}},
		Packages: map[string][]string{ManagerWinget: winget("Anthropic.ClaudeCode"), ManagerBrew: {"--cask", "claude-code"}},
		Manual:   map[string]string{"linux": "curl -fsSL https://claude.ai/install.sh | bash"},
	},
	"codex": {
		ID: "codex", Label: "Codex CLI", Homepage: "https://developers.openai.com/codex",
		Purpose:  locale.Text{Ko: "Codex CLI (메인·서브 에이전트 실행)", En: "Codex CLI (main or subagent runtime)", Zh: "Codex CLI（主代理或子代理运行时）"},
		Commands: map[string][]string{"*": {"codex"}},
		Packages: map[string][]string{ManagerWinget: winget("OpenAI.Codex"), ManagerBrew: {"--cask", "codex"}},
		Manual:   map[string]string{"linux": "npm install -g @openai/codex"},
	},
	"ollama": {
		ID: "ollama", Label: "Ollama", Homepage: "https://ollama.com/download",
		Purpose:  locale.Text{Ko: "로컬·클라우드 모델 서버 (서브 에이전트)", En: "local and cloud model server (subagents)", Zh: "本地与云端模型服务（子代理）"},
		Commands: map[string][]string{"*": {"ollama"}},
		Packages: map[string][]string{ManagerWinget: winget("Ollama.Ollama"), ManagerBrew: {"ollama"}},
		Manual:   map[string]string{"linux": "curl -fsSL https://ollama.com/install.sh | sh"},
	},
}
