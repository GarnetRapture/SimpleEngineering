package locale

const LanguageQuestion = "출력 언어를 고르세요 / Choose the output language / 选择输出语言"

var (
	ChooseHint = Text{
		Ko: "번호를 입력하세요 (Enter = %d)",
		En: "Enter a number (Enter = %d)",
		Zh: "请输入编号（回车 = %d）",
	}
	ChooseRequiredHint = Text{
		Ko: "번호를 입력하세요",
		En: "Enter a number",
		Zh: "请输入编号",
	}
	InvalidChoice = Text{
		Ko: "%q 은(는) 목록에 없는 번호예요. 1부터 %d 사이로 입력하세요.",
		En: "%q is not in the list. Enter a number from 1 to %d.",
		Zh: "%q 不在列表中。请输入 1 到 %d 之间的编号。",
	}
	ManyHint = Text{
		Ko: "켜고 끌 번호를 입력하세요 (예: 2 4). Enter = 이대로 확정",
		En: "Enter numbers to toggle (for example: 2 4). Enter = keep as shown",
		Zh: "输入要切换的编号（例如：2 4）。回车 = 按当前状态确认",
	}
	TextRequiredHint = Text{
		Ko: "(필수)",
		En: "(required)",
		Zh: "（必填）",
	}
	ConfirmYesHint = Text{Ko: "(Y/n)", En: "(Y/n)", Zh: "(Y/n)"}
	ConfirmNoHint  = Text{Ko: "(y/N)", En: "(y/N)", Zh: "(y/N)"}
	InvalidConfirm = Text{
		Ko: "y 또는 n으로 답하세요.",
		En: "Answer y or n.",
		Zh: "请回答 y 或 n。",
	}
	InputEnded = Text{
		Ko: "입력이 끝나서 필수 항목에 답하지 못했어요. 아무것도 만들지 않았어요.",
		En: "Input ended before a required answer. Nothing was written.",
		Zh: "输入在必填项之前结束。未写入任何内容。",
	}

	Title = Text{
		Ko: "SimpleHarness 설정 도우미",
		En: "SimpleHarness setup wizard",
		Zh: "SimpleHarness 设置向导",
	}
	Intro = Text{
		Ko: "AI 코딩 에이전트(Claude Code·Codex)가 프로젝트 규칙을 지키도록 묶는 하네스 파일을 만들어요.\n질문에 답하면 output 폴더에 결과가 모이고, 그 폴더 안의 내용을 저장소 루트로 끌어다 놓으면 끝나요.\n모르는 항목은 Enter를 눌러 추천값을 쓰면 돼요.",
		En: "This wizard creates the harness files that bind AI coding agents (Claude Code, Codex) to your project rules.\nAnswer the questions, and the result is collected in the output folder; drag its contents into your repository root and you are done.\nPress Enter on anything you are unsure about to accept the recommended value.",
		Zh: "本向导会生成约束 AI 编程代理（Claude Code、Codex）遵守项目规则的 harness 文件。\n回答问题后，结果会汇集到 output 文件夹，把其中的内容拖到仓库根目录即可。\n不确定的项目直接按回车使用推荐值。",
	}

	StepLanguage = Text{Ko: "1/5 출력 언어", En: "1/5 Output language", Zh: "1/5 输出语言"}
	StepProject  = Text{Ko: "2/5 프로젝트", En: "2/5 Project", Zh: "2/5 项目"}
	StepAgents   = Text{Ko: "3/5 에이전트", En: "3/5 Agents", Zh: "3/5 代理"}
	StepRules    = Text{Ko: "4/5 핵심 규칙", En: "4/5 Core rules", Zh: "4/5 核心规则"}
	StepGenerate = Text{Ko: "5/5 생성", En: "5/5 Generate", Zh: "5/5 生成"}
	StepTools    = Text{Ko: "추천 도구 (선택)", En: "Recommended tools (optional)", Zh: "推荐工具（可选）"}
	StepNext     = Text{Ko: "다음 단계", En: "Next steps", Zh: "后续步骤"}

	EnvHeading   = Text{Ko: "사용 환경", En: "Your environment", Zh: "运行环境"}
	EnvWindows11 = Text{
		Ko: "운영체제: Windows 11 (빌드 %d)",
		En: "Operating system: Windows 11 (build %d)",
		Zh: "操作系统：Windows 11（版本号 %d）",
	}
	EnvWindows10 = Text{
		Ko: "운영체제: Windows 10 (빌드 %d)",
		En: "Operating system: Windows 10 (build %d)",
		Zh: "操作系统：Windows 10（版本号 %d）",
	}
	EnvLinux = Text{
		Ko: "운영체제: Linux — %s",
		En: "Operating system: Linux — %s",
		Zh: "操作系统：Linux — %s",
	}
	EnvLinuxWSL = Text{
		Ko: "운영체제: Linux — %s (WSL: Windows 안에서 실행되는 리눅스)",
		En: "Operating system: Linux — %s (WSL: Linux running inside Windows)",
		Zh: "操作系统：Linux — %s（WSL：运行在 Windows 中的 Linux）",
	}
	EnvMac = Text{
		Ko: "운영체제: macOS %s",
		En: "Operating system: macOS %s",
		Zh: "操作系统：macOS %s",
	}
	EnvOther = Text{
		Ko: "운영체제: %s (자동 설치 미지원)",
		En: "Operating system: %s (automatic installation unsupported)",
		Zh: "操作系统：%s（不支持自动安装）",
	}
	EnvPackageManager = Text{
		Ko: "패키지 관리자: %s",
		En: "Package manager: %s",
		Zh: "包管理器：%s",
	}
	EnvNoPackageManager = Text{
		Ko: "패키지 관리자: 찾지 못했어요",
		En: "Package manager: not found",
		Zh: "包管理器：未找到",
	}
	Installed = Text{Ko: "설치됨", En: "installed", Zh: "已安装"}
	Missing   = Text{Ko: "없음", En: "missing", Zh: "未安装"}

	GuideWindows10Cmd = Text{
		Ko: "Windows 10의 기본 셸은 cmd(명령 프롬프트)예요. 아래 명령을 cmd 창에 한 줄씩 붙여 넣어 셸을 교체하세요.",
		En: "The default shell on Windows 10 is cmd (Command Prompt). Paste the commands below into the cmd window one line at a time to replace it.",
		Zh: "Windows 10 的默认 shell 是 cmd（命令提示符）。请将下面的命令逐行粘贴到 cmd 窗口中，以替换它。",
	}
	GuideShellChoice = Text{
		Ko: "cmd는 Windows의 오래된 셸이라 교체를 권장해요. 둘 중 하나를 기본 셸로 쓰세요:\n   · PowerShell 7 (pwsh, 최신 버전 네이티브 설치): 명령 결과를 텍스트가 아니라 객체(속성을 가진 데이터)로 파이프라인에 넘겨서 필터·정렬·변환을 정확하게 할 수 있어요.\n   · Git Bash: Linux·macOS와 같은 bash 문법이라 AI 에이전트 문서와 예제 명령을 그대로 쓸 수 있어요. Claude Code도 Windows에서 훅을 Git Bash로 실행해요.",
		En: "cmd is the legacy Windows shell, so replacing it is recommended. Use one of these as the default shell:\n   · PowerShell 7 (pwsh, latest version installed natively): passes command results down the pipeline as objects (data with properties) rather than text, so filtering, sorting, and transforming are exact.\n   · Git Bash: the same bash syntax as Linux and macOS, so agent documentation and example commands work as written. Claude Code also runs hooks through Git Bash on Windows.",
		Zh: "cmd 是 Windows 的旧式 shell，建议替换。请选择以下其中之一作为默认 shell：\n   · PowerShell 7（pwsh，原生安装最新版本）：管道传递的是对象（带属性的数据）而不是文本，因此筛选、排序和转换都很精确。\n   · Git Bash：与 Linux、macOS 相同的 bash 语法，代理文档和示例命令可以直接使用。Claude Code 在 Windows 上也通过 Git Bash 运行钩子。",
	}
	GuideWindows10Winget = Text{
		Ko: "① winget 확인: winget --version\n   '내부 또는 외부 명령이 아닙니다'가 나오면 등록 후 cmd 창을 새로 여세요 (winget은 Windows 10 1809 이상 지원):\n   powershell -NoProfile -Command \"Add-AppxPackage -RegisterByFamilyName -MainPackage Microsoft.DesktopAppInstaller_8wekyb3d8bbwe\"",
		En: "① Check winget: winget --version\n   If cmd says it is not recognized, register it and open a new cmd window (winget supports Windows 10 1809 or later):\n   powershell -NoProfile -Command \"Add-AppxPackage -RegisterByFamilyName -MainPackage Microsoft.DesktopAppInstaller_8wekyb3d8bbwe\"",
		Zh: "① 检查 winget：winget --version\n   如果提示不是内部或外部命令，请先注册再打开新的 cmd 窗口（winget 支持 Windows 10 1809 及以上）：\n   powershell -NoProfile -Command \"Add-AppxPackage -RegisterByFamilyName -MainPackage Microsoft.DesktopAppInstaller_8wekyb3d8bbwe\"",
	}
	GuideWindowsPowerShell7 = Text{
		Ko: "%s PowerShell 7 설치: winget install --id Microsoft.PowerShell --source winget\n   Windows에 들어 있는 Windows PowerShell 5.1은 구버전이에요. PowerShell 7은 5.1과 따로 설치되고 pwsh 명령으로 실행해요.",
		En: "%s Install PowerShell 7: winget install --id Microsoft.PowerShell --source winget\n   The Windows PowerShell 5.1 that ships with Windows is the legacy version. PowerShell 7 installs side by side and runs with the pwsh command.",
		Zh: "%s 安装 PowerShell 7：winget install --id Microsoft.PowerShell --source winget\n   Windows 自带的 Windows PowerShell 5.1 是旧版本。PowerShell 7 与 5.1 并行安装，使用 pwsh 命令运行。",
	}
	GuideWindows10Terminal = Text{
		Ko: "③ Windows Terminal 설치: winget install --id Microsoft.WindowsTerminal --source winget",
		En: "③ Install Windows Terminal: winget install --id Microsoft.WindowsTerminal --source winget",
		Zh: "③ 安装 Windows Terminal：winget install --id Microsoft.WindowsTerminal --source winget",
	}
	GuideWindowsGit = Text{
		Ko: "%s Git 설치: winget install --id Git.Git -e --source winget\n   git과 Git Bash가 함께 설치돼요. Claude Code는 훅 명령을 Git Bash로 실행해요.",
		En: "%s Install Git: winget install --id Git.Git -e --source winget\n   This installs git together with Git Bash, which Claude Code uses to run hook commands.",
		Zh: "%s 安装 Git：winget install --id Git.Git -e --source winget\n   会同时安装 git 和 Git Bash，Claude Code 使用 Git Bash 运行钩子命令。",
	}
	GuideWindows10DefaultTerminal = Text{
		Ko: "%s 기본 터미널 바꾸기: Windows Terminal → 설정 → 시작 → '기본 터미널 애플리케이션'을 Windows Terminal로, '기본 프로필'을 PowerShell(7)로 바꾸세요.\n   Windows 10은 22H2에 2023-05-23 업데이트(KB5026435) 이상이 설치돼야 '기본 터미널 애플리케이션' 항목이 보여요.",
		En: "%s Change the default terminal: Windows Terminal → Settings → Startup → set 'Default terminal application' to Windows Terminal and 'Default profile' to PowerShell (7).\n   On Windows 10 the 'Default terminal application' setting appears on 22H2 with the May 23, 2023 update (KB5026435) or later.",
		Zh: "%s 更改默认终端：Windows Terminal → 设置 → 启动 → 将“默认终端应用程序”设为 Windows Terminal，将“默认配置文件”设为 PowerShell（7）。\n   Windows 10 需为 22H2 并安装 2023-05-23 更新（KB5026435）或更高版本，才会显示“默认终端应用程序”设置。",
	}
	GuideWindows11Shell = Text{
		Ko: "Windows 11에는 Windows Terminal이 들어 있지만, 함께 들어 있는 Windows PowerShell 5.1은 구버전이에요.",
		En: "Windows 11 includes Windows Terminal, but the Windows PowerShell 5.1 it includes is the legacy version.",
		Zh: "Windows 11 自带 Windows Terminal，但其自带的 Windows PowerShell 5.1 是旧版本。",
	}
	GuideWindows11DefaultProfile = Text{
		Ko: "%s 기본 프로필 바꾸기: Windows Terminal → 설정 → 시작 → '기본 프로필'을 PowerShell(7)로 바꾸세요. 이후 작업은 그 창(pwsh)에서 하세요.",
		En: "%s Change the default profile: Windows Terminal → Settings → Startup → set 'Default profile' to PowerShell (7). Work in that window (pwsh) from then on.",
		Zh: "%s 更改默认配置文件：Windows Terminal → 设置 → 启动 → 将“默认配置文件”设为 PowerShell（7）。之后请在该窗口（pwsh）中工作。",
	}
	GuideWindowsGitBash = Text{
		Ko: "Git Bash를 기본 셸로 쓰려면: Windows Terminal → 설정 → '+ 새 프로필 추가' → 명령줄 C:\\Program Files\\Git\\bin\\bash.exe -li 로 프로필을 만들고 '기본 프로필'로 지정하세요. VS Code는 명령 팔레트의 'Terminal: Select Default Profile'에서 Git Bash 또는 PowerShell을 고르세요.",
		En: "To use Git Bash as the default shell: Windows Terminal → Settings → '+ Add a new profile' → command line C:\\Program Files\\Git\\bin\\bash.exe -li, then make it the 'Default profile'. In VS Code, pick Git Bash or PowerShell with the Command Palette command 'Terminal: Select Default Profile'.",
		Zh: "若要使用 Git Bash 作为默认 shell：Windows Terminal → 设置 → “+ 添加新配置文件” → 命令行填写 C:\\Program Files\\Git\\bin\\bash.exe -li，并设为“默认配置文件”。在 VS Code 中，通过命令面板的“Terminal: Select Default Profile”选择 Git Bash 或 PowerShell。",
	}
	GuideCodexWindows10 = Text{
		Ko: "알려진 문제: Windows 10(Home)에서 공식 설치한 Codex 앱이 기본 셸을 cmd로 잡아 명령이 제대로 동작하지 않는 경우가 있어요. PowerShell 7을 설치하고 기본 터미널·프로필을 PowerShell 7로 바꾼 뒤 사용하세요. 이 하네스가 만드는 Codex 훅 명령은 cmd로 실행돼도 동작하는 형태예요.",
		En: "Known issue: on Windows 10 (Home), the officially installed Codex app can pick cmd as its default shell, which breaks commands. Install PowerShell 7 and switch the default terminal and profile to PowerShell 7 before using it. The Codex hook commands this harness generates work even when run through cmd.",
		Zh: "已知问题：在 Windows 10（家庭版）上，官方安装的 Codex 应用可能把默认 shell 设为 cmd，导致命令无法正常运行。请先安装 PowerShell 7，并将默认终端和配置文件切换为 PowerShell 7。本 harness 生成的 Codex 钩子命令即使通过 cmd 运行也能工作。",
	}
	GuideRunModes = Text{
		Ko: "AI 실행 방식은 세 가지예요: ① 터미널에서 claude 또는 codex 실행 ② VS Code 확장 — Claude Code 확장은 프로젝트의 .claude/settings.json 훅과 CLAUDE.md를 읽고, Codex 확장은 CLI와 같은 설정 계층을 써요 ③ 데스크톱 앱. 어느 방식이든 이 저장소 폴더를 작업 폴더로 열어야 하네스가 적용돼요.",
		En: "AI agents run in three ways: ① claude or codex in a terminal ② the VS Code extensions — the Claude Code extension reads the project's .claude/settings.json hooks and CLAUDE.md, and the Codex extension shares the CLI's configuration layers ③ the desktop apps. In every case, open this repository folder as the workspace so the harness applies.",
		Zh: "AI 代理有三种运行方式：① 在终端运行 claude 或 codex ② VS Code 扩展 — Claude Code 扩展会读取项目的 .claude/settings.json 钩子和 CLAUDE.md，Codex 扩展与 CLI 共享同一配置层 ③ 桌面应用。无论哪种方式，都要把本仓库文件夹作为工作区打开，harness 才会生效。",
	}
	GuideMacShell = Text{
		Ko: "macOS 기본 셸은 zsh예요. 바꾸려면 Terminal → 설정 → 일반 → '셸 열기 방식'에서 '명령(전체 경로)'을 고르고 셸 경로를 입력하세요.",
		En: "The default macOS shell is zsh. To change it: Terminal → Settings → General → next to 'Shells open with', choose 'Command (complete path)' and enter the shell path.",
		Zh: "macOS 默认 shell 是 zsh。如需更改：终端 → 设置 → 通用 → 在“Shell 打开方式”中选择“命令（完整路径）”并输入 shell 路径。",
	}
	GuideLinuxInstall = Text{
		Ko: "패키지 설치 명령 형식: %s",
		En: "Package install command: %s",
		Zh: "软件包安装命令：%s",
	}
	GuideLinuxTerminal = Text{
		Ko: "기본 터미널 바꾸기(Debian·Ubuntu): sudo update-alternatives --config x-terminal-emulator 를 실행하고 번호를 고르세요. 다른 배포판은 데스크톱 환경의 기본 앱 설정에서 바꿔요.",
		En: "Change the default terminal (Debian, Ubuntu): run sudo update-alternatives --config x-terminal-emulator and pick a number. Other distributions change it in the desktop environment's default applications settings.",
		Zh: "更改默认终端（Debian、Ubuntu）：运行 sudo update-alternatives --config x-terminal-emulator 并选择编号。其他发行版请在桌面环境的默认应用设置中更改。",
	}
	GuideWSLTerminal = Text{
		Ko: "WSL 배포판은 Windows Terminal에 프로필이 자동으로 생겨요. Windows Terminal → 설정 → 시작 → '기본 프로필'을 그 배포판으로 바꾸면 열자마자 리눅스 셸이에요.",
		En: "Windows Terminal creates a profile for each WSL distribution automatically. Set Windows Terminal → Settings → Startup → 'Default profile' to that distribution to open straight into the Linux shell.",
		Zh: "Windows Terminal 会自动为每个 WSL 发行版创建配置文件。将 Windows Terminal → 设置 → 启动 → “默认配置文件”设为该发行版，即可直接打开 Linux shell。",
	}
	GuideWindowsReopen = Text{
		Ko: "새로 설치한 명령은 새로 연 Windows Terminal(pwsh) 창에서 인식돼요. 이 도우미도 그 창에서 다시 실행하면 돼요.",
		En: "Newly installed commands are found in a newly opened Windows Terminal (pwsh) window. You can rerun this wizard there.",
		Zh: "新安装的命令会在新打开的 Windows Terminal（pwsh）窗口中生效。可以在该窗口中重新运行本向导。",
	}
	GuideWindowsExecutionPolicy = Text{
		Ko: "npm으로 설치한 명령(.ps1)이 '스크립트를 실행할 수 없습니다'로 막히면 PowerShell에서 한 번 실행하세요: Set-ExecutionPolicy -Scope CurrentUser RemoteSigned",
		En: "If commands installed by npm (.ps1) fail with 'running scripts is disabled', run this once in PowerShell: Set-ExecutionPolicy -Scope CurrentUser RemoteSigned",
		Zh: "如果通过 npm 安装的命令（.ps1）提示“禁止运行脚本”，请在 PowerShell 中执行一次：Set-ExecutionPolicy -Scope CurrentUser RemoteSigned",
	}
	GuideGitLinux = Text{
		Ko: "Git이 없어요. 설치 명령: %s",
		En: "Git is missing. Install command: %s",
		Zh: "未安装 Git。安装命令：%s",
	}
	GuideGitMac = Text{
		Ko: "Git이 없어요. 'xcode-select --install' 또는 'brew install git'으로 설치하세요.",
		En: "Git is missing. Install it with 'xcode-select --install' or 'brew install git'.",
		Zh: "未安装 Git。请使用 'xcode-select --install' 或 'brew install git' 安装。",
	}
	GuideNode = Text{
		Ko: "Node.js가 없어요. 생성되는 훅·커밋 가드·서브 에이전트 실행기는 Node.js로 실행되므로 꼭 필요해요.",
		En: "Node.js is missing. The generated hooks, commit guard, and subagent launcher run on Node.js, so it is required.",
		Zh: "未安装 Node.js。生成的钩子、提交守卫和子代理启动器都运行在 Node.js 上，因此必须安装。",
	}
	GuideWSL = Text{
		Ko: "WSL에서 실행 중이에요. 리눅스 도구로 작업할 파일은 /mnt/c 아래보다 리눅스 홈(~/) 아래에 두는 것이 빨라요.",
		En: "You are inside WSL. Files you work on with Linux tools are faster under your Linux home (~/) than under /mnt/c.",
		Zh: "当前运行在 WSL 中。使用 Linux 工具处理的文件放在 Linux 主目录（~/）下比放在 /mnt/c 下更快。",
	}
	GuideMacBrew = Text{
		Ko: "Homebrew가 없어요. https://brew.sh 의 설치 명령을 먼저 실행하면 도구를 자동으로 설치할 수 있어요.",
		En: "Homebrew is missing. Run the install command from https://brew.sh first to enable automatic tool installation.",
		Zh: "未安装 Homebrew。请先运行 https://brew.sh 上的安装命令，以便自动安装工具。",
	}

	ProjectNameQuestion = Text{
		Ko: "프로젝트 이름 (영문자·숫자·.·_·-, 결과 폴더 이름이 돼요)",
		En: "Project name (letters, digits, . _ -; it becomes the output folder name)",
		Zh: "项目名称（字母、数字、. _ -；将作为输出文件夹名）",
	}
	ProjectNameInvalid = Text{
		Ko: "영문자나 숫자로 시작하고 영문자·숫자·.·_·-만 쓸 수 있어요.",
		En: "Start with a letter or digit and use only letters, digits, '.', '_' and '-'.",
		Zh: "必须以字母或数字开头，且只能包含字母、数字、'.'、'_' 和 '-'。",
	}
	ProjectSummaryQuestion = Text{
		Ko: "이 프로젝트가 무엇을 만드는지 한 문장 (Enter = 나중에 컨벤션 C02에서 채움)",
		En: "One sentence on what this project builds (Enter = fill it later in convention C02)",
		Zh: "用一句话说明本项目要构建什么（回车 = 稍后在约定 C02 中填写）",
	}
	ProjectSummaryPlaceholder = Text{
		Ko: "%s가 무엇을 만드는 프로젝트인지 한 문장으로 정의한다",
		En: "Define in one sentence what %s builds",
		Zh: "用一句话定义 %s 要构建什么",
	}
	SkillPrefixQuestion = Text{
		Ko: "스킬 이름 앞에 붙일 짧은 접두사 (소문자·숫자·-)",
		En: "Short prefix for skill names (lowercase letters, digits, -)",
		Zh: "技能名称的短前缀（小写字母、数字、-）",
	}
	SkillPrefixInvalid = Text{
		Ko: "소문자·숫자·-만 쓸 수 있어요.",
		En: "Use only lowercase letters, digits, and '-'.",
		Zh: "只能使用小写字母、数字和 '-'。",
	}
	KindQuestion = Text{
		Ko: "어떤 종류의 프로젝트인가요? (추천 도구가 달라져요)",
		En: "What kind of project is it? (this changes the recommended tools)",
		Zh: "这是哪类项目？（会影响推荐工具）",
	}
	ProjectLanguageQuestion = Text{
		Ko: "프로젝트의 대표 언어는 무엇인가요? (스킬 예시·검사 명령·추천 도구가 이 언어로 맞춰져요)",
		En: "What is the project's primary language? (skill examples, checks, and recommended tools follow it)",
		Zh: "项目的主要语言是什么？（技能示例、检查命令和推荐工具都会以此为准）",
	}
	Recommended = Text{Ko: "추천", En: "recommended", Zh: "推荐"}

	MainAgentQuestion = Text{
		Ko: "대표(메인) 에이전트를 고르세요. 훅·스킬·규칙이 이 에이전트에 연결돼요.",
		En: "Choose the main agent. Hooks, skills, and rules are connected to it.",
		Zh: "选择主代理。钩子、技能和规则都会连接到它。",
	}
	SubagentQuestion = Text{
		Ko: "서브 에이전트를 고르세요 (여러 개 가능, 없으면 Enter). 메인 에이전트가 조사·초안 같은 작은 일을 맡기고 결과를 검증해서 써요.",
		En: "Choose subagents (several allowed; Enter for none). The main agent delegates small tasks such as research or drafts and verifies the result before using it.",
		Zh: "选择子代理（可多选；回车表示不选）。主代理会把调研、草稿等小任务交给子代理，并在使用前验证结果。",
	}
	ProviderRegionQuestion = Text{
		Ko: "%s — 사용할 지역(엔드포인트)을 고르세요",
		En: "%s — choose the region (endpoint)",
		Zh: "%s — 选择区域（端点）",
	}
	ProviderModelQuestion = Text{
		Ko: "%s — 사용할 모델 이름 (모델 목록: %s)",
		En: "%s — model name (model list: %s)",
		Zh: "%s — 模型名称（模型列表：%s）",
	}
	ProviderModelInvalid = Text{
		Ko: "모델 이름에는 공백을 쓸 수 없어요.",
		En: "A model name cannot contain spaces.",
		Zh: "模型名称不能包含空格。",
	}
	ProviderTokenQuestion = Text{
		Ko: "%s — API 키를 담을 환경 변수 이름",
		En: "%s — environment variable that holds the API key",
		Zh: "%s — 存放 API 密钥的环境变量名",
	}
	EnvironmentVariableInvalid = Text{
		Ko: "대문자·숫자·_만 쓸 수 있고 숫자로 시작할 수 없어요.",
		En: "Use only uppercase letters, digits, and '_', not starting with a digit.",
		Zh: "只能使用大写字母、数字和 '_'，且不能以数字开头。",
	}

	RulesIntro = Text{
		Ko: "AI가 자주 만드는 인라인 코드와 명명 위반을 막는 규칙이에요. 기본으로 모두 켜져 있어요.\n[훅]은 훅이 해당 편집을 차단하고, [스킬]은 스킬·리뷰 절차와 컨벤션 조항으로 이행해요.",
		En: "These rules stop the inline code and naming violations AI agents produce most often. All are on by default.\n[hook] rules are blocked by the hook at edit time; [skill] rules are carried out by skill, review, and convention clauses.",
		Zh: "这些规则用于阻止 AI 最常产生的内联代码和命名违规，默认全部开启。\n[钩子] 规则在编辑时由钩子拦截；[技能] 规则通过技能、评审流程和约定条款执行。",
	}
	RuleHookTag  = Text{Ko: "[훅]", En: "[hook]", Zh: "[钩子]"}
	RuleSkillTag = Text{Ko: "[스킬]", En: "[skill]", Zh: "[技能]"}

	SummaryHeading  = Text{Ko: "선택한 내용", En: "Your choices", Zh: "已选内容"}
	SummaryLanguage = Text{Ko: "출력 언어: %s", En: "Output language: %s", Zh: "输出语言：%s"}
	SummaryProject  = Text{Ko: "프로젝트: %s — %s, %s", En: "Project: %s — %s, %s", Zh: "项目：%s — %s，%s"}
	SummaryPrefix   = Text{Ko: "스킬 접두사: %s", En: "Skill prefix: %s", Zh: "技能前缀：%s"}
	SummaryMain     = Text{Ko: "대표 에이전트: %s", En: "Main agent: %s", Zh: "主代理：%s"}
	SummarySubs     = Text{Ko: "서브 에이전트: %s", En: "Subagents: %s", Zh: "子代理：%s"}
	SummaryRules    = Text{Ko: "규칙: %s", En: "Rules: %s", Zh: "规则：%s"}
	SummaryOutput   = Text{Ko: "결과 폴더: %s", En: "Output folder: %s", Zh: "输出文件夹：%s"}
	None            = Text{Ko: "없음", En: "none", Zh: "无"}
	ConfirmGenerate = Text{
		Ko: "이대로 생성할까요?",
		En: "Generate with these choices?",
		Zh: "按以上选择生成吗？",
	}
	Cancelled = Text{
		Ko: "아무것도 만들지 않고 끝냈어요.",
		En: "Finished without writing anything.",
		Zh: "已结束，未写入任何内容。",
	}
	OutputExists = Text{
		Ko: "%s 폴더가 이미 있어요. 기존 결과를 지키기 위해 아무것도 쓰지 않았어요. 그 폴더를 옮기거나 다른 프로젝트 이름을 쓰세요.",
		En: "The folder %s already exists. Nothing was written, to preserve the existing result. Move that folder or use another project name.",
		Zh: "文件夹 %s 已存在。为保护已有结果，未写入任何内容。请移走该文件夹或使用其他项目名称。",
	}
	Generated = Text{
		Ko: "%d개 파일을 만들었어요: %s",
		En: "Wrote %d files to %s",
		Zh: "已写入 %d 个文件：%s",
	}
	GenerateFailed = Text{
		Ko: "생성에 실패했어요: %v",
		En: "Generation failed: %v",
		Zh: "生成失败：%v",
	}

	ToolsIntro = Text{
		Ko: "이 프로젝트와 선택한 에이전트에 맞는 도구예요. ✓ 설치됨, ✗ 없음.",
		En: "Tools matched to this project and the chosen agents. ✓ installed, ✗ missing.",
		Zh: "与本项目及所选代理匹配的工具。✓ 已安装，✗ 未安装。",
	}
	ToolsAllInstalled = Text{
		Ko: "필요한 도구가 모두 설치돼 있어요.",
		En: "Every recommended tool is already installed.",
		Zh: "所有推荐工具均已安装。",
	}
	ToolsInstallQuestion = Text{
		Ko: "없는 도구 %d개를 %s로 지금 설치할까요?",
		En: "Install the %d missing tools with %s now?",
		Zh: "现在用 %[2]s 安装缺少的 %[1]d 个工具吗？",
	}
	ToolsNoManager = Text{
		Ko: "자동 설치에 쓸 패키지 관리자가 없어요. 아래 명령·주소로 직접 설치하세요.",
		En: "No package manager is available for automatic installation. Install the tools with the commands and links below.",
		Zh: "没有可用于自动安装的包管理器。请使用下面的命令和链接手动安装。",
	}
	ToolsSkipped = Text{
		Ko: "설치를 건너뛰었어요. 아래 명령으로 언제든 설치할 수 있어요.",
		En: "Installation skipped. You can install them any time with the commands below.",
		Zh: "已跳过安装。之后可随时使用以下命令安装。",
	}
	ToolInstalling = Text{Ko: "설치 중: %s", En: "Installing: %s", Zh: "正在安装：%s"}
	ToolInstalled  = Text{Ko: "설치 완료: %s", En: "Installed: %s", Zh: "安装完成：%s"}
	ToolFailed     = Text{
		Ko: "설치 실패: %s (%v) — 위 출력에서 원인을 확인하세요.",
		En: "Installation failed: %s (%v) — check the output above for the cause.",
		Zh: "安装失败：%s（%v）— 请查看上方输出了解原因。",
	}
	ToolManual = Text{
		Ko: "직접 설치: %s → %s",
		En: "Install manually: %s → %s",
		Zh: "手动安装：%s → %s",
	}
	ToolsRestartShell = Text{
		Ko: "새로 설치한 명령은 터미널을 새로 열어야 인식돼요.",
		En: "Open a new terminal so newly installed commands are found.",
		Zh: "请打开新的终端，以便识别新安装的命令。",
	}

	NextCopy = Text{
		Ko: "%s 폴더를 열고, 안의 모든 파일·폴더(.으로 시작하는 폴더 포함)를 선택해 저장소 루트로 끌어다 놓으세요.",
		En: "Open %s, select every file and folder inside it (including folders starting with '.'), and drag them into your repository root.",
		Zh: "打开 %s，选中其中所有文件和文件夹（包括以 '.' 开头的文件夹），拖到仓库根目录。",
	}
	NextGitHooks = Text{
		Ko: "저장소 루트에서 실행: git config core.hooksPath .githooks (커밋 가드 연결)",
		En: "Run in the repository root: git config core.hooksPath .githooks (connects the commit guard)",
		Zh: "在仓库根目录运行：git config core.hooksPath .githooks（连接提交守卫）",
	}
	NextGitignore = Text{
		Ko: ".gitignore에 %s/ 를 추가하세요 (훅 상태·임시 산출물 폴더)",
		En: "Add %s/ to .gitignore (hook state and temporary output folder)",
		Zh: "将 %s/ 加入 .gitignore（钩子状态和临时产物文件夹）",
	}
	NextConventions = Text{
		Ko: "%s 에서 '(프로젝트 정의' 로 시작하는 행을 채우세요. 이 행들은 사용자가 정하는 결정이에요.",
		En: "Fill the rows starting with '(project definition' in %s. Those rows are decisions that belong to you.",
		Zh: "填写 %s 中以“（项目定义”开头的行。这些行由你来决定。",
	}
	NextClaude = Text{
		Ko: "Claude Code를 다시 시작하고 /hooks 로 SessionStart·UserPromptSubmit·PreToolUse·PostToolUse·PostToolUseFailure·Stop 훅이 보이는지 확인하세요.",
		En: "Restart Claude Code and check with /hooks that the SessionStart, UserPromptSubmit, PreToolUse, PostToolUse, PostToolUseFailure, and Stop hooks are listed.",
		Zh: "重启 Claude Code，并用 /hooks 确认 SessionStart、UserPromptSubmit、PreToolUse、PostToolUse、PostToolUseFailure、Stop 钩子均已列出。",
	}
	NextCodex = Text{
		Ko: "Codex에서 이 프로젝트를 신뢰(trust)하고 /hooks 에서 프로젝트 훅을 검토·승인하세요. 신뢰하지 않은 프로젝트의 .codex/ 설정과 훅은 읽히지 않아요.",
		En: "Trust this project in Codex and review the project hooks in /hooks. Codex skips the .codex/ layer, including hooks, of an untrusted project.",
		Zh: "在 Codex 中信任本项目，并在 /hooks 中审查项目钩子。未受信任项目的 .codex/ 配置和钩子不会被加载。",
	}
	NextSubagentKeys = Text{
		Ko: "서브 에이전트 API 키를 환경 변수로 설정하세요 (새 터미널부터 적용):",
		En: "Set the subagent API keys as environment variables (applies to new terminals):",
		Zh: "将子代理 API 密钥设置为环境变量（对新终端生效）：",
	}
	NextSubagentRun = Text{
		Ko: "서브 에이전트 실행: 작업 내용을 %s/ 안의 파일에 쓰고 node .harness/subagents/run_subagent.mjs <id> <작업 파일>",
		En: "Run a subagent: write the task to a file under %s/ and run node .harness/subagents/run_subagent.mjs <id> <task file>",
		Zh: "运行子代理：把任务写入 %s/ 下的文件，然后运行 node .harness/subagents/run_subagent.mjs <id> <任务文件>",
	}
	NextStructure = Text{
		Ko: "구조: AGENTS.md = AI 통제 뇌 · 훅 = 제어자 · 스킬 = 팔다리(명세·논리·좋은/나쁜 예시) · 컨벤션 = 프로젝트 방향·목적 선언 · 트래킹 = 진행 기록과 프로젝트 노드 관리",
		En: "Structure: AGENTS.md = the brain that controls the AI · hooks = the controller · skills = the limbs (specs, logic, good and bad examples) · conventions = the declared project direction and purpose · tracking = progress records and project node management",
		Zh: "结构：AGENTS.md = 控制 AI 的大脑 · 钩子 = 控制器 · 技能 = 四肢（规范、逻辑、正反示例）· 约定 = 项目方向与目的的声明 · 跟踪 = 进度记录与项目节点管理",
	}
	SetupGuideTitle = Text{
		Ko: "%s 하네스 설정 안내",
		En: "%s harness setup guide",
		Zh: "%s harness 设置指南",
	}
	NextGuideSaved = Text{
		Ko: "이 안내는 %s 에도 저장했어요.",
		En: "This guide is also saved to %s.",
		Zh: "本指南也已保存到 %s。",
	}
)
