package catalog

import (
	"simpleharness/src/internal/harness"
	"simpleharness/src/internal/locale"
)

type Rule struct {
	ID         string
	Clause     string
	Label      locale.Text
	Definition locale.Text
	Languages  []string
	Hook       bool
	Inline     []harness.InlineRule
}

const RuleNoInlineComment = "no-inline-comment"

var webMarkupExtensions = []string{".html", ".vue", ".svelte", ".php"}

var Rules = []Rule{
	{
		ID: "domain-structure", Clause: "C04",
		Label: locale.Text{
			Ko: "도메인 기반 구조 — {도메인}/{클래스}/{대상}/{이름}",
			En: "Domain-based structure — {domain}/{class}/{target}/{name}",
			Zh: "基于领域的结构 — {领域}/{类别}/{对象}/{名称}",
		},
		Definition: locale.Text{
			Ko: "소스는 `{도메인}/{클래스}/{대상}/{이름}` 계층으로 둔다. 각 계층은 책임·기능·대상을 구분하며 깊이를 채우기 위한 폴더를 만들지 않는다. 공유 개념(타입·상수·로직·스타일)의 소유자는 자기 도메인 안에 하나다",
			En: "Source lives in the `{domain}/{class}/{target}/{name}` hierarchy. Each level separates responsibility, function, and target, and no folder exists only to fill depth. Every shared concept (type, constant, logic, style) has exactly one owner inside its domain",
			Zh: "源码按 `{领域}/{类别}/{对象}/{名称}` 层级组织。每一层区分职责、功能与对象，不为凑层级而建文件夹。每个共享概念（类型、常量、逻辑、样式）在其领域内只有一个所有者",
		},
	},
	{
		ID: "domain-naming", Clause: "C05",
		Label: locale.Text{
			Ko: "도메인 기반 명명 — 도메인 용어+동작+대상, 범용 이름 금지",
			En: "Domain-based naming — domain term + action + target, no generic names",
			Zh: "基于领域的命名 — 领域术语+动作+对象，禁止通用名称",
		},
		Definition: locale.Text{
			Ko: "이름은 도메인 용어와 동작·대상을 담고 C05 표의 언어별 대소문자 규칙을 따른다. `util`·`helper`·`common`·`manager`·`misc`·`temp`·`new`·`data`·`info`처럼 책임을 숨기는 범용 이름을 쓰지 않는다",
			En: "Names carry the domain term with the action and target and follow the per-language casing of the C05 table. Generic names that hide responsibility such as `util`, `helper`, `common`, `manager`, `misc`, `temp`, `new`, `data`, and `info` are not used",
			Zh: "名称包含领域术语以及动作与对象，并遵循 C05 表中各语言的大小写规则。不使用 `util`、`helper`、`common`、`manager`、`misc`、`temp`、`new`、`data`、`info` 等隐藏职责的通用名称",
		},
	},
	{
		ID: RuleNoInlineComment, Clause: "C04", Hook: true,
		Label: locale.Text{
			Ko: "인라인 주석 금지 — 의도는 이름과 구조로",
			En: "No inline comments — intent lives in names and structure",
			Zh: "禁止内联注释 — 意图由名称与结构表达",
		},
		Definition: locale.Text{
			Ko: "자체 작성 코드에 줄·블록·문서 주석을 쓰지 않는다. 의도는 이름과 구조로 드러낸다. 훅이 주석 수가 늘어나는 편집을 차단한다",
			En: "Authored code carries no line, block, or documentation comments; intent is expressed through names and structure. The hook blocks an edit that increases the comment count",
			Zh: "自编代码中不写行注释、块注释或文档注释；意图通过名称与结构表达。钩子会拦截使注释数量增加的编辑",
		},
	},
	{
		ID: "no-inline-type", Clause: "C04", Hook: true, Languages: []string{"typescript"},
		Label: locale.Text{
			Ko: "인라인 타입 금지 — 타입은 도메인의 types 소유 파일에만",
			En: "No inline types — types live only in the domain's types owner files",
			Zh: "禁止内联类型 — 类型只放在领域的 types 所属文件中",
		},
		Definition: locale.Text{
			Ko: "`type`·`interface`·`enum` 선언은 자기 도메인의 `types/` 폴더 또는 `*.types.ts` 파일에만 둔다. 컴포넌트·로직 파일 안의 즉석 타입 선언은 훅이 차단한다",
			En: "`type`, `interface`, and `enum` declarations live only in the domain's `types/` folder or `*.types.ts` files. The hook blocks an ad hoc type declaration inside a component or logic file",
			Zh: "`type`、`interface`、`enum` 声明只放在所属领域的 `types/` 文件夹或 `*.types.ts` 文件中。钩子会拦截在组件或逻辑文件中临时声明类型",
		},
		Inline: []harness.InlineRule{
			{ID: "no-inline-type", Clause: "C04", Extensions: []string{".ts", ".tsx", ".mts", ".cts"}, Pattern: `^[ \t]*(?:export[ \t]+)?(?:declare[ \t]+)?(?:type|interface|enum)[ \t]+[A-Za-z_$][\w$]*`, AllowPath: `(^|/)types/|\.types\.(ts|tsx|mts|cts)$|\.d\.ts$`, Reason: "C04 no-inline-type: declare types in the domain's types/ folder or *.types.ts owner file"},
		},
	},
	{
		ID: "no-inline-style", Clause: "C04", Hook: true, Languages: []string{"typescript", "php"},
		Label: locale.Text{
			Ko: "인라인 스타일 금지 — 스타일은 도메인 스타일 소유 파일에",
			En: "No inline styles — styles live in the domain's style owner files",
			Zh: "禁止内联样式 — 样式放在领域的样式所属文件中",
		},
		Definition: locale.Text{
			Ko: "JSX `style={{…}}`, HTML `style=\"…\"` 속성, `<style>` 블록을 쓰지 않는다. 스타일은 자기 도메인의 스타일 소유 파일에 이름 있는 클래스로 둔다",
			En: "No JSX `style={{…}}`, HTML `style=\"…\"` attributes, or `<style>` blocks. Styles live as named classes in the domain's style owner file",
			Zh: "不使用 JSX `style={{…}}`、HTML `style=\"…\"` 属性或 `<style>` 块。样式以具名类的形式放在所属领域的样式文件中",
		},
		Inline: []harness.InlineRule{
			{ID: "no-inline-style", Clause: "C04", Extensions: []string{".tsx", ".jsx"}, Pattern: `\bstyle=\{\{`, Reason: "C04 no-inline-style: move the style to a named class in the domain's style owner file"},
			{ID: "no-inline-style", Clause: "C04", Extensions: webMarkupExtensions, Pattern: `\sstyle\s*=\s*["']|<style[\s>]`, Reason: "C04 no-inline-style: move the style to a named class in the domain's style owner file"},
		},
	},
	{
		ID: "no-inline-handler", Clause: "C05", Hook: true, Languages: []string{"typescript", "php"},
		Label: locale.Text{
			Ko: "인라인 이벤트 핸들러 금지 — 이름 있는 핸들러 함수로",
			En: "No inline event handlers — use named handler functions",
			Zh: "禁止内联事件处理器 — 使用具名处理函数",
		},
		Definition: locale.Text{
			Ko: "JSX `onClick={() => …}` 같은 즉석 화살표 함수와 HTML `onclick=\"…\"` 속성을 쓰지 않는다. 핸들러는 동작+대상 이름을 가진 함수로 정의해 참조한다",
			En: "No ad hoc arrow functions such as JSX `onClick={() => …}` and no HTML `onclick=\"…\"` attributes. Handlers are functions named by action and target and are referenced by name",
			Zh: "不使用 JSX `onClick={() => …}` 之类的临时箭头函数，也不使用 HTML `onclick=\"…\"` 属性。处理器定义为以“动作+对象”命名的函数并按名称引用",
		},
		Inline: []harness.InlineRule{
			{ID: "no-inline-handler", Clause: "C05", Extensions: []string{".tsx", ".jsx"}, Pattern: `\bon[A-Z]\w*=\{\s*(?:async\s*)?(?:\([^)]*\)|[A-Za-z_$][\w$]*)\s*=>|\bon[A-Z]\w*=\{\s*function\b`, Reason: "C05 no-inline-handler: define a handler function named by action and target and pass it by name"},
			{ID: "no-inline-handler", Clause: "C05", Extensions: webMarkupExtensions, Pattern: `\son[a-z]+\s*=\s*["']`, Reason: "C05 no-inline-handler: bind a named handler function from the domain's script owner"},
		},
	},
	{
		ID: "no-magic-literal", Clause: "C04",
		Label: locale.Text{
			Ko: "매직 리터럴 금지 — 숫자·문자열·URL·경로·SQL·정규식은 이름 있는 상수로",
			En: "No magic literals — numbers, strings, URLs, paths, SQL, and regexes become named constants",
			Zh: "禁止魔法字面量 — 数字、字符串、URL、路径、SQL、正则都定义为具名常量",
		},
		Definition: locale.Text{
			Ko: "의미를 가진 숫자·문자열·URL·경로·포트·SQL·정규식을 코드 중간에 직접 쓰지 않는다. 자기 도메인의 상수 소유 파일에 이름 있는 상수로 두고 참조한다",
			En: "Meaningful numbers, strings, URLs, paths, ports, SQL, and regular expressions are not written inline. They are named constants in the domain's constant owner file and referenced by name",
			Zh: "有含义的数字、字符串、URL、路径、端口、SQL 和正则表达式不直接写在代码中，而是在所属领域的常量文件中定义为具名常量并按名称引用",
		},
	},
}

func (rule Rule) AppliesTo(language string) bool {
	if len(rule.Languages) == 0 {
		return true
	}
	for _, candidate := range rule.Languages {
		if candidate == language {
			return true
		}
	}
	return false
}

func RuleByID(id string) Rule {
	for _, rule := range Rules {
		if rule.ID == id {
			return rule
		}
	}
	panic("catalog: unknown rule " + id)
}
