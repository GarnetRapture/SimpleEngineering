package catalog

import "simpleharness/src/internal/locale"

type Agent struct {
	ID        string
	Label     string
	Detail    locale.Text
	SkillRoot string
}

var MainAgents = []Agent{
	{
		ID: "claude", Label: "Claude Code", SkillRoot: ".claude/skills",
		Detail: locale.Text{
			Ko: "Anthropic — 훅 .claude/settings.json, 스킬 .claude/skills, 진입 CLAUDE.md",
			En: "Anthropic — hooks .claude/settings.json, skills .claude/skills, entry CLAUDE.md",
			Zh: "Anthropic — 钩子 .claude/settings.json，技能 .claude/skills，入口 CLAUDE.md",
		},
	},
	{
		ID: "codex", Label: "Codex", SkillRoot: ".agents/skills",
		Detail: locale.Text{
			Ko: "OpenAI — 훅 .codex/hooks.json, 스킬 .agents/skills, 진입 AGENTS.md",
			En: "OpenAI — hooks .codex/hooks.json, skills .agents/skills, entry AGENTS.md",
			Zh: "OpenAI — 钩子 .codex/hooks.json，技能 .agents/skills，入口 AGENTS.md",
		},
	},
}

func AgentByID(id string) Agent {
	for _, agent := range MainAgents {
		if agent.ID == id {
			return agent
		}
	}
	panic("catalog: unknown main agent " + id)
}

type Region struct {
	ID      string
	Label   locale.Text
	BaseURL string
}

type Provider struct {
	ID           string
	Label        string
	Detail       locale.Text
	Runtime      string
	MainAgents   []string
	Regions      []Region
	AuthVariable string
	TokenEnv     string
	TokenValue   string
	Model        string
	ModelsURL    string
}

const (
	RuntimeClaude = "claude"
	RuntimeCodex  = "codex"
	RuntimeOllama = "ollama"
)

const (
	authToken = "ANTHROPIC_AUTH_TOKEN"
	authKey   = "ANTHROPIC_API_KEY"
)

var global = locale.Text{Ko: "글로벌(국제)", En: "Global (international)", Zh: "国际"}
var china = locale.Text{Ko: "중국 본토", En: "Mainland China", Zh: "中国大陆"}

var Providers = []Provider{
	{
		ID: "codex", Label: "Codex CLI", Runtime: RuntimeCodex, MainAgents: []string{"claude"},
		Detail: locale.Text{
			Ko: "codex exec — 읽기 전용 샌드박스로 조사·검토를 맡김",
			En: "codex exec — delegate research and review in a read-only sandbox",
			Zh: "codex exec — 在只读沙箱中委派调研与审查",
		},
	},
	{
		ID: "claude", Label: "Claude Code CLI", Runtime: RuntimeClaude, MainAgents: []string{"codex"},
		Detail: locale.Text{
			Ko: "claude -p — 편집 도구를 뺀 비대화 모드로 조사·검토를 맡김",
			En: "claude -p — delegate research and review in print mode without edit tools",
			Zh: "claude -p — 以去除编辑工具的非交互模式委派调研与审查",
		},
	},
	{
		ID: "ollama-local", Label: "Ollama", Runtime: RuntimeOllama, MainAgents: []string{"claude", "codex"},
		Regions:   []Region{{ID: "local", Label: locale.Text{Ko: "이 PC", En: "This PC", Zh: "本机"}, BaseURL: "http://localhost:11434"}},
		Model:     "qwen3.5",
		ModelsURL: "https://ollama.com/search",
		Detail: locale.Text{
			Ko: "로컬 Ollama — 내려받은 모델과 :cloud 모델(ollama 로그인 필요)",
			En: "Local Ollama — downloaded models and :cloud models (requires signing in to Ollama)",
			Zh: "本地 Ollama — 已下载模型以及 :cloud 模型（需登录 Ollama）",
		},
	},
	{
		ID: "ollama-cloud", Label: "Ollama Cloud", Runtime: RuntimeClaude, MainAgents: []string{"claude", "codex"},
		Regions:      []Region{{ID: "global", Label: global, BaseURL: "https://ollama.com"}},
		AuthVariable: authToken, TokenEnv: "OLLAMA_API_KEY",
		Model:     "glm-5.3-flash",
		ModelsURL: "https://ollama.com/search?c=cloud",
		Detail: locale.Text{
			Ko: "Ollama 설치 없이 클라우드 API 직접 사용 (API 키)",
			En: "Use the cloud API directly without installing Ollama (API key)",
			Zh: "无需安装 Ollama，直接使用云端 API（API 密钥）",
		},
	},
	{
		ID: "alibaba-coding-plan", Label: "Alibaba Cloud Model Studio — Coding Plan", Runtime: RuntimeClaude, MainAgents: []string{"claude", "codex"},
		Regions: []Region{
			{ID: "intl", Label: global, BaseURL: "https://coding-intl.dashscope.aliyuncs.com/apps/anthropic"},
			{ID: "cn", Label: china, BaseURL: "https://coding.dashscope.aliyuncs.com/apps/anthropic"},
		},
		AuthVariable: authToken, TokenEnv: "DASHSCOPE_CODING_PLAN_API_KEY",
		Model:     "qwen3.7-plus",
		ModelsURL: "https://www.alibabacloud.com/help/en/model-studio/coding-plan",
		Detail: locale.Text{
			Ko: "Qwen 등 코딩 요금제 (Coding Plan 전용 sk-sp- 키)",
			En: "Coding subscription for Qwen and other models (Coding Plan sk-sp- key)",
			Zh: "Qwen 等编码订阅套餐（Coding Plan 专用 sk-sp- 密钥）",
		},
	},
	{
		ID: "alibaba-token-plan", Label: "Alibaba Cloud Model Studio — Token Plan", Runtime: RuntimeClaude, MainAgents: []string{"claude", "codex"},
		Regions:      []Region{{ID: "intl", Label: global, BaseURL: "https://token-plan.ap-southeast-1.maas.aliyuncs.com/apps/anthropic"}},
		AuthVariable: authToken, TokenEnv: "DASHSCOPE_TOKEN_PLAN_API_KEY",
		ModelsURL: "https://www.alibabacloud.com/help/en/model-studio/claude-code",
		Detail: locale.Text{
			Ko: "토큰 단위 요금제 (Token Plan 키)",
			En: "Token-based plan (Token Plan key)",
			Zh: "按 Token 计费套餐（Token Plan 密钥）",
		},
	},
	{
		ID: "kimi", Label: "Moonshot Kimi", Runtime: RuntimeClaude, MainAgents: []string{"claude", "codex"},
		Regions: []Region{
			{ID: "global", Label: global, BaseURL: "https://api.moonshot.ai/anthropic"},
			{ID: "cn", Label: china, BaseURL: "https://api.moonshot.cn/anthropic"},
		},
		AuthVariable: authToken, TokenEnv: "MOONSHOT_API_KEY",
		Model:     "kimi-k3[1m]",
		ModelsURL: "https://platform.kimi.ai/docs/guide/claude-code-kimi",
		Detail: locale.Text{
			Ko: "Kimi 공식 API",
			En: "Official Kimi API",
			Zh: "Kimi 官方 API",
		},
	},
	{
		ID: "zai-glm", Label: "Z.ai GLM Coding Plan", Runtime: RuntimeClaude, MainAgents: []string{"claude", "codex"},
		Regions:      []Region{{ID: "global", Label: global, BaseURL: "https://api.z.ai/api/anthropic"}},
		AuthVariable: authToken, TokenEnv: "ZAI_API_KEY",
		Model:     "glm-5.3[1m]",
		ModelsURL: "https://docs.z.ai/devpack/tool/claude",
		Detail: locale.Text{
			Ko: "GLM 코딩 요금제",
			En: "GLM coding plan",
			Zh: "GLM 编码套餐",
		},
	},
	{
		ID: "deepseek", Label: "DeepSeek", Runtime: RuntimeClaude, MainAgents: []string{"claude", "codex"},
		Regions:      []Region{{ID: "global", Label: global, BaseURL: "https://api.deepseek.com/anthropic"}},
		AuthVariable: authKey, TokenEnv: "DEEPSEEK_API_KEY",
		Model:     "deepseek-flash",
		ModelsURL: "https://api-docs.deepseek.com/guides/anthropic_api",
		Detail: locale.Text{
			Ko: "DeepSeek Anthropic 호환 API",
			En: "DeepSeek Anthropic-compatible API",
			Zh: "DeepSeek Anthropic 兼容 API",
		},
	},
	{
		ID: "minimax", Label: "MiniMax", Runtime: RuntimeClaude, MainAgents: []string{"claude", "codex"},
		Regions: []Region{
			{ID: "global", Label: global, BaseURL: "https://api.minimax.io/anthropic"},
			{ID: "cn", Label: china, BaseURL: "https://api.minimaxi.com/anthropic"},
		},
		AuthVariable: authToken, TokenEnv: "MINIMAX_API_KEY",
		Model:     "MiniMax-M3",
		ModelsURL: "https://platform.minimax.io/docs/guides/text-ai-coding-tools",
		Detail: locale.Text{
			Ko: "MiniMax 공식 API (토큰 요금제 포함)",
			En: "Official MiniMax API (token plans included)",
			Zh: "MiniMax 官方 API（含 Token 套餐）",
		},
	},
	{
		ID: "openrouter", Label: "OpenRouter", Runtime: RuntimeClaude, MainAgents: []string{"claude", "codex"},
		Regions:      []Region{{ID: "global", Label: global, BaseURL: "https://openrouter.ai/api"}},
		AuthVariable: authToken, TokenEnv: "OPENROUTER_API_KEY",
		ModelsURL: "https://openrouter.ai/models",
		Detail: locale.Text{
			Ko: "여러 제공자의 모델을 키 하나로 사용",
			En: "Many providers' models behind one key",
			Zh: "一个密钥使用多家提供商的模型",
		},
	},
}

func ProvidersFor(mainAgent string) []Provider {
	var selected []Provider
	for _, provider := range Providers {
		for _, agent := range provider.MainAgents {
			if agent == mainAgent {
				selected = append(selected, provider)
				break
			}
		}
	}
	return selected
}

func ProviderByID(id string) Provider {
	for _, provider := range Providers {
		if provider.ID == id {
			return provider
		}
	}
	panic("catalog: unknown subagent provider " + id)
}
