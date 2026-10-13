package harness

type Config struct {
	Project        Project         `json:"project"`
	Output         Output          `json:"output"`
	Agents         Agents          `json:"agents"`
	Skills         Skills          `json:"skills"`
	Paths          Paths           `json:"paths"`
	Time           Time            `json:"time"`
	Version        Version         `json:"version"`
	Stacks         []Stack         `json:"stacks"`
	Rules          []string        `json:"rules"`
	InlineRules    []InlineRule    `json:"inlineRules"`
	Observations   []Observation   `json:"observations"`
	ProtectedPaths []ProtectedPath `json:"protectedPaths"`
	ForbiddenPaths []PatternReason `json:"forbiddenPaths"`
	DeniedCommands []PatternReason `json:"deniedCommands"`
	MCP            MCP             `json:"mcp"`
	Subagents      []Subagent      `json:"subagents"`
	Claude         ClaudeSettings  `json:"claude"`
	CommitGuard    CommitGuard     `json:"commitGuard"`
	Response       Response        `json:"response"`
}

type Project struct {
	Name     string `json:"name"`
	Summary  string `json:"summary"`
	Kind     string `json:"kind"`
	Language string `json:"language"`
}

type Output struct {
	Language string `json:"language"`
}

type Agents struct {
	Main string `json:"main"`
}

type Skills struct {
	Prefix string        `json:"prefix"`
	Root   string        `json:"root"`
	Domain []DomainSkill `json:"domain"`
}

type DomainSkill struct {
	Name         string   `json:"name"`
	Trigger      string   `json:"trigger"`
	Paths        []string `json:"paths"`
	Commands     []string `json:"commands"`
	ToolPrefixes []string `json:"toolPrefixes"`
	Extensions   []string `json:"extensions"`
}

type Paths struct {
	Conventions      string `json:"conventions"`
	Patterns         string `json:"patterns"`
	Tracking         string `json:"tracking"`
	Audit            string `json:"audit"`
	Feedback         string `json:"feedback"`
	Scratch          string `json:"scratch"`
	CommitGuard      string `json:"commitGuard"`
	Hook             string `json:"hook"`
	SubagentLauncher string `json:"subagentLauncher"`
}

type Time struct {
	Label   string `json:"label"`
	Command string `json:"command"`
}

type Version struct {
	Source  string `json:"source"`
	Command string `json:"command"`
}

type Stack struct {
	Name       string      `json:"name"`
	Markers    []string    `json:"markers"`
	Extensions []string    `json:"extensions"`
	Checks     []Check     `json:"checks"`
	Forbidden  []Forbidden `json:"forbidden"`
}

type Check struct {
	Kind       string     `json:"kind"`
	Command    string     `json:"command"`
	Acceptance Acceptance `json:"acceptance"`
}

type Acceptance struct {
	Exit       int      `json:"exit"`
	Line       string   `json:"line,omitempty"`
	ForbidLine string   `json:"forbidLine,omitempty"`
	Counts     []string `json:"counts,omitempty"`
}

type Forbidden struct {
	Command string `json:"command"`
	Reason  string `json:"reason"`
}

type InlineRule struct {
	ID         string   `json:"id"`
	Clause     string   `json:"clause"`
	Extensions []string `json:"extensions"`
	Pattern    string   `json:"pattern"`
	AllowPath  string   `json:"allowPath,omitempty"`
	Reason     string   `json:"reason"`
}

type Observation struct {
	Kind           string   `json:"kind"`
	ChangeTypes    string   `json:"changeTypes"`
	EditExtensions []string `json:"editExtensions"`
	Route          string   `json:"route"`
}

type ProtectedPath struct {
	Pattern       string   `json:"pattern"`
	ShellSegments []string `json:"shellSegments"`
	Target        string   `json:"target"`
	Reason        string   `json:"reason"`
	Route         string   `json:"route"`
}

type PatternReason struct {
	Pattern string `json:"pattern"`
	Reason  string `json:"reason"`
}

type MCP struct {
	QueryRoute            string      `json:"queryRoute"`
	ContextTool           *string     `json:"contextTool"`
	InventoryTool         *string     `json:"inventoryTool"`
	BatchTool             *string     `json:"batchTool"`
	ToolRunner            *ToolRunner `json:"toolRunner"`
	ExecutingToolPrefixes []string    `json:"executingToolPrefixes"`
	HumanInputTools       []string    `json:"humanInputTools"`
	ObservationTools      []string    `json:"observationTools"`
	Servers               []MCPServer `json:"servers"`
}

type ToolRunner struct {
	Tool         string            `json:"tool"`
	Capabilities map[string]string `json:"capabilities"`
}

type MCPServer struct {
	Name      string   `json:"name"`
	Command   string   `json:"command,omitempty"`
	Args      []string `json:"args,omitempty"`
	URL       string   `json:"url,omitempty"`
	Harnesses []string `json:"harnesses"`
}

type Subagent struct {
	ID            string `json:"id"`
	Provider      string `json:"provider"`
	Runtime       string `json:"runtime"`
	BaseURL       string `json:"baseUrl,omitempty"`
	AuthVariable  string `json:"authVariable,omitempty"`
	TokenEnv      string `json:"tokenEnv,omitempty"`
	TokenValue    string `json:"tokenValue,omitempty"`
	Model         string `json:"model,omitempty"`
	LocalProvider string `json:"localProvider,omitempty"`
}

type ClaudeSettings struct {
	PermissionsAllow []string `json:"permissionsAllow"`
}

type CommitGuard struct {
	BlockedExtensions []string `json:"blockedExtensions"`
}

type Response struct {
	Language        string   `json:"language"`
	Fields          []string `json:"fields"`
	BasisLabels     []string `json:"basisLabels"`
	BasisStatement  string   `json:"basisStatement"`
	NumberedBasis   string   `json:"numberedBasis"`
	BannedJudgment  string   `json:"bannedJudgment"`
	Acknowledgment  string   `json:"acknowledgment"`
	ReportKeyword   string   `json:"reportKeyword"`
	StopInstruction string   `json:"stopInstruction"`
	GuessedJudgment string   `json:"guessedJudgment"`
	GateWrapper     string   `json:"gateWrapper"`
	SummaryHeading  string   `json:"summaryHeading"`
	TrackingHeading string   `json:"trackingHeading"`
	TodoHeading     string   `json:"todoHeading"`
	OpenStatus      string   `json:"openStatus"`
	NodeHeading     string   `json:"nodeHeading"`
}
