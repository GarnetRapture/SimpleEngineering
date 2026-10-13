package catalog

import (
	"simpleharness/src/internal/harness"
	"simpleharness/src/internal/locale"
)

const gateWrapper = `\[(?:목표달성완주근거|증명불가미완료)\]`

var Responses = map[locale.Language]harness.Response{
	locale.Korean: {
		Language:        "ko",
		Fields:          []string{"규칙", "무엇", "어떻게", "어디서", "언제", "왜", "어떤 것 때문에", "왜 존재하는지", "어떻게 존재해야 하는지"},
		BasisLabels:     []string{"수학적 근거", "기술적 근거", "과학적 근거"},
		BasisStatement:  `기반으로[\s\S]*이행[\s\S]*근거는\s*다음과\s*같`,
		NumberedBasis:   `근거\s*1[\s\S]*따라서`,
		BannedJudgment:  `멈출게|멈추겠|멈춥니다|중단할게|중단하겠|미룰게|미루겠|나중에|보류할|보류하겠|건너뛸|건너뛰겠|좁힐게|좁히겠|일부만|축소할|축소하겠|생략할|생략하겠|후순위|범위 밖|원래 없|오판|A안|B안|지금은 예외`,
		Acknowledgment:  `지적[^\n.]{0,12}맞|(?:^|[\n.!?]\s*)(?:네[,.]?\s*)?(?:맞아요|맞습니다)|보고해요|보고할게요|보고합니다|보고드`,
		ReportKeyword:   "보고해",
		StopInstruction: `멈춰|멈추라|중지|중단|종료해|종료하`,
		GuessedJudgment: `추측으로|일반적으로`,
		GateWrapper:     gateWrapper,
		SummaryHeading:  "### 컨벤션 요약",
		TrackingHeading: "세션 GOAL 요약",
		TodoHeading:     "## 3. 누락 TODO",
		OpenStatus:      "열림",
		NodeHeading:     "## 프로젝트 노드",
	},
	locale.English: {
		Language:        "en",
		Fields:          []string{"Rule", "What", "How", "Where", "When", "Why", "Because of", "Why it exists", "How it must exist"},
		BasisLabels:     []string{"Mathematical basis", "Technical basis", "Scientific basis"},
		BasisStatement:  `Based on[\s\S]*carried out[\s\S]*evidence for this action is as follows`,
		NumberedBasis:   `Evidence\s*1[\s\S]*Therefore`,
		BannedJudgment:  `\b(?:I(?:'ll| will) stop|stopping here|I(?:'ll| will) defer|deferring|for later|on hold|I(?:'ll| will) skip|skipping|narrow(?:ing)? the scope|only part of|partially done|out of scope|originally absent|a misjudgment|plan A|plan B|an exception for now)\b`,
		Acknowledgment:  `\b(?:you(?:'re| are) right|that is correct|I(?:'ll| will) report|reporting back)\b`,
		ReportKeyword:   "report",
		StopInstruction: `\b(?:stop|halt|pause|cancel|abort|terminate)\b`,
		GuessedJudgment: `\b(?:by guess|guessing|generally|in general|usually)\b`,
		GateWrapper:     gateWrapper,
		SummaryHeading:  "### Conventions summary",
		TrackingHeading: "Session GOAL summary",
		TodoHeading:     "## 3. Open TODO",
		OpenStatus:      "open",
		NodeHeading:     "## Project nodes",
	},
	locale.Chinese: {
		Language:        "zh",
		Fields:          []string{"规则", "什么", "如何", "在哪里", "何时", "为什么", "因为什么", "为何存在", "应如何存在"},
		BasisLabels:     []string{"数学依据", "技术依据", "科学依据"},
		BasisStatement:  `基于[\s\S]*执行了[\s\S]*依据如下`,
		NumberedBasis:   `依据\s*1[\s\S]*因此`,
		BannedJudgment:  `我(?:先|会|将)?停止|我(?:先|会|将)?暂停|推迟到|稍后再做|以后再做|先搁置|跳过这|缩小范围|只做一部分|部分完成|超出范围|原本没有|误判|方案A|方案B|暂时例外`,
		Acknowledgment:  `你说得对|您说得对|没错|我会汇报|汇报如下`,
		ReportKeyword:   "汇报",
		StopInstruction: `停止|暂停|中止|终止|取消`,
		GuessedJudgment: `推测|一般来说|通常`,
		GateWrapper:     gateWrapper,
		SummaryHeading:  "### 约定摘要",
		TrackingHeading: "会话 GOAL 摘要",
		TodoHeading:     "## 3. 未完成 TODO",
		OpenStatus:      "未完成",
		NodeHeading:     "## 项目节点",
	},
}
