import fs from 'node:fs';
import path from 'node:path';
import { createHash, randomUUID } from 'node:crypto';
import { fileURLToPath } from 'node:url';

const root = path.resolve(fileURLToPath(new URL('../..', import.meta.url)));
const harnessNames = new Set(['claude', 'codex']);
const harness = process.argv.slice(1).find(argument => harnessNames.has(argument)) ?? null;
const events = new Set(['SessionStart', 'UserPromptSubmit', 'PreToolUse', 'PostToolUse', 'PostToolUseFailure', 'Stop']);
const claudeTools = {
  edit: new Set(['Edit', 'Write', 'NotebookEdit']),
  read: new Set(['Read', 'Grep', 'Glob']),
  shell: new Set(['PowerShell', 'Bash']),
  question: new Set(['AskUserQuestion', 'EnterPlanMode', 'ExitPlanMode']),
};
const codexTools = { shell: 'Bash', patch: 'apply_patch' };
const processSkillNames = ['workflow', 'fix-process', 'verify'];
const baseExecutables = ['npm', 'npx', 'node', 'python', 'python3', 'py'];
const contextCharacterLimit = 10000;
const responseLanguageNames = { ko: 'Korean', en: 'English', zh: 'Simplified Chinese' };
const evidenceToken = /\bC\d{2}\b|\d|https?:\S+|[\w.@-]+[\\/][\w.@-]+/;
const harnessScratch = /appdata[\\/]local[\\/]temp[\\/]|\$env:(?:temp|tmp)\b|[\\/]\.claude[\\/]projects[\\/][^\\/]+[\\/][^\\/]+[\\/]tool-results(?:[\\/]|$)/i;
const jobStop = /\b(?:stop-process|taskkill(?:\.exe)?|kill)\b/i;
const codexExit = /^Process exited with code (-?\d+)\s*$/m;
const patchHeader = /^\*\*\* (Add File|Update File|Delete File|Move to): (.+)$/;
const mutatingCommand = /\b(?:set-content|add-content|out-file|new-item|copy-item|move-item|rename-item|remove-item|clear-content)\b/i;
const movingCommand = /\b(?:move-item|rename-item|git\s+mv)\b/i;
const foregroundWait = /\bStart-Sleep\b|\bGet-Content\b[^;|]*\s-Wait\b/i;
const cmdRoute = new RegExp('(?:^|[\\s;&|])(cmd(?:\\.exe)?|npm\\.cmd|[^\\s;|"\']+\\.bat)(?=\\s|$|["\'])', 'i');
const tmpWrite = new RegExp('(?:set-content|add-content|out-file|new-item|copy-item|move-item|rename-item)\\b[^;|]*["\']/tmp/', 'i');
const bashLiteral = new RegExp("'[^']*'|\"(?:[^\"\\\\]|\\\\[\\s\\S])*\"", 'g');
const powershellLiteral = new RegExp("@'\\r?\\n[\\s\\S]*?\\r?\\n'@|@\"\\r?\\n[\\s\\S]*?\\r?\\n\"@|'(?:[^']|'')*'|\"(?:[^\"\\u0060]|\\u0060[\\s\\S]|\"\")*\"", 'g');
const bashSubstitution = new RegExp('\\$\\(|\\u0060');
const powershellSubstitution = new RegExp('\\$\\(');

function readConfig() {
  return JSON.parse(fs.readFileSync(path.join(root, 'harness.config.json'), 'utf8'));
}

function escapeRegex(text) {
  return text.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}

function comparable(text) {
  return process.platform === 'win32' ? text.toLowerCase() : text;
}

function tokenPattern(command) {
  return new RegExp('(?:^|[\\s;&|(])' + escapeRegex(command) + '(?=\\s|$|[;&|)])', 'i');
}

function segmentPattern(command) {
  return new RegExp('(?:^|[;&|]\\s*)' + escapeRegex(command) + '(?=\\s|$|[;&|)])', 'i');
}

function commandTokens(command) {
  return command.split(/[\s;&|()]+/).filter(token => token.length > 0).map(token => comparable(token.replace(/\.exe$/i, '')));
}

function activeStacks(config) {
  return config.stacks.filter(stack => stack.markers.every(marker => fs.existsSync(path.join(root, marker))));
}

function compileOptional(pattern, flags) {
  return pattern ? new RegExp(pattern, flags) : null;
}

function createRules(config) {
  const skill = name => config.skills.prefix + '-' + name;
  const skillRoot = config.skills.root;
  const skillFile = name => skillRoot + '/' + name + '/SKILL.md';
  const mandatory = processSkillNames.map(skill);
  const stacks = activeStacks(config);
  const conventions = config.paths.conventions;
  const scratch = config.paths.scratch;
  const response = config.response;
  const fields = response.fields;
  const languageName = responseLanguageNames[config.output.language];
  const domain = config.skills.domain.map(entry => ({ ...entry, skill: skill(entry.name), extensionSet: new Set(entry.extensions.map(comparable)) }));
  const knownSkills = [...mandatory, skill('structure'), skill('stack'), skill('mcp'), ...domain.map(entry => entry.skill)];
  const domainText = domain.map(entry => entry.skill + ' when ' + entry.trigger).join('; ');
  const loadMethod = harness === 'codex'
    ? 'read the SKILL.md file of each skill under ' + skillRoot + '/ (for example ' + skillFile(mandatory[0]) + ')'
    : 'load each skill through the Skill tool';
  const contextSource = config.mcp.contextTool
    ? 'the ' + config.mcp.contextTool + ' summaries and source identities'
    : 'the ' + conventions + ' summary table and the AGENTS behavior summary with their source identities';
  const basisLabels = response.basisLabels.map(label => label + ':').join(', ');
  const toolBasis = `C00 tool-and-approach row of ${conventions} (C15, C16): use a tool and approach only as the conventions specify. Before every tool call, the same message names the convention clause ID and row that define this exact tool and approach and states the confirmed basis as ${basisLabels} with its source path, measurement, or observation. Queries and measurements run through ${config.mcp.queryRoute}; a temporary script is allowed only when no registered tool has that capability (C12). Before running project code or a tool on it, judge every executed logic through the ${skill('verify')} skill (C14).`;
  const skillLoad = `AGENTS section 3 skill row (C01): ${loadMethod}. Mandatory before any change, run, or final response: ${mandatory.join(', ')}. Situational: ${skill('structure')} when a file is added, moved, or renamed; ${skill('stack')} when a file of an active stack changes or a stack check runs${domainText ? '; ' + domainText : ''}. Reason: only the skill holds the method of the behavior AGENTS selects, and an unloaded skill leaves the model to fill the procedure with its own guess (C01). The project hook denies a change or run whose required skills are not loaded.`;
  const responseText = `C15 response row of ${conventions} (harness.config.json response): every final and progress response, in polite ${languageName}, states each labelled field ${fields.map(field => field + ':').join(' ')}, where ${fields[0]}: names the applicable convention clause IDs and every field carries the source path, measurement, or observation that settles it; then the basis statement, the numbered evidence lines, and the conclusion the C15 response row defines. A guess or a generality is no basis. A change, however small, follows the ${skill('fix-process')} FIX steps and an analysis its diagnostic steps.`;
  const shellSegments = config.protectedPaths.flatMap(entry => entry.shellSegments).map(escapeRegex);
  return {
    config,
    conventions,
    scratch,
    fields,
    mandatory,
    domain,
    knownSkills,
    skillFile,
    structureSkill: skill('structure'),
    stackSkill: skill('stack'),
    fixSkill: skill('fix-process'),
    stackExtensions: new Set(stacks.flatMap(stack => stack.extensions.map(comparable))),
    checks: stacks.flatMap(stack => stack.checks.map(check => ({ ...check, id: stack.name + ':' + check.kind, pattern: segmentPattern(check.command) }))),
    forbiddenCommands: stacks.flatMap(stack => stack.forbidden.map(entry => ({ ...entry, stack: stack.name, pattern: tokenPattern(entry.command) }))),
    observations: config.observations,
    executables: new Set([...baseExecutables, ...config.stacks.flatMap(stack => [...stack.checks, ...stack.forbidden].map(entry => entry.command.split(/\s+/)[0]))].map(comparable)),
    executingToolPrefixes: config.mcp.executingToolPrefixes,
    humanInputTools: new Set(config.mcp.humanInputTools),
    observationTools: new Set(config.mcp.observationTools),
    toolRunner: config.mcp.toolRunner,
    protectedPaths: config.protectedPaths.map(entry => ({ reason: entry.reason, regex: new RegExp(entry.pattern, process.platform === 'win32' ? 'i' : '') })),
    protectedShell: shellSegments.length > 0
      ? new RegExp('(?:set-content|add-content|out-file|new-item|copy-item|move-item|rename-item|remove-item|clear-content)\\b[^;|]*(?:["\'\\s]|[\\\\/])(?:' + shellSegments.join('|') + ')[\\\\/]', 'i')
      : null,
    forbiddenPaths: config.forbiddenPaths.map(entry => ({ reason: entry.reason, regex: new RegExp(entry.pattern, 'i') })),
    deniedCommands: config.deniedCommands.map(entry => ({ reason: entry.reason, regex: new RegExp(entry.pattern, 'i') })),
    commentRule: config.rules.includes('no-inline-comment'),
    inlineRules: config.inlineRules.map(entry => ({ ...entry, regex: new RegExp(entry.pattern, 'gm'), extensionSet: new Set(entry.extensions.map(comparable)), allow: compileOptional(entry.allowPath, 'i') })),
    documents: {
      summaryHeading: response.summaryHeading,
      trackingHeading: new RegExp('^## \\d{4}-\\d{2}-\\d{2} ' + escapeRegex(response.trackingHeading)),
      todoHeading: response.todoHeading,
      openStatus: response.openStatus,
      nodeHeading: response.nodeHeading,
    },
    response: {
      basisLabel: response.basisLabels.map(label => escapeRegex(label).replace(/ /g, '\\s*')).join('|'),
      basisStatement: new RegExp(response.basisStatement, 'iu'),
      numberedBasis: new RegExp(response.numberedBasis, 'iu'),
      bannedJudgment: new RegExp(response.bannedJudgment, 'iu'),
      acknowledgment: new RegExp(response.acknowledgment, 'iu'),
      reportKeyword: response.reportKeyword,
      stopInstruction: new RegExp(response.stopInstruction, 'iu'),
      guessedJudgment: new RegExp(response.guessedJudgment, 'iu'),
      gateWrapper: new RegExp(response.gateWrapper, 'u'),
    },
    text: {
      toolBasis,
      skillLoad,
      response: responseText,
      newSession: `C15 session-start row of ${conventions}: before any other action read ${conventions} in full, every line, in bounded portions that return every line, then AGENTS.md in full; summaries, searches, and partial reads never replace that reading. Then consume ${contextSource}, the newest tracking summary and the project node table (${config.paths.tracking}), the open TODO rows (${config.paths.audit}), and the applicable skill details. Follow AGENTS section 2. Responses and progress begin with the applicable convention IDs and use polite ${languageName}. ${toolBasis} ${skillLoad} ${responseText}`,
      resume: `C15 session-start row of ${conventions} (resume): do not reread every document; check only that the clauses, procedure rows, and owner files tied to the recent work still hold, retain valid evidence and unfinished requirements, and continue. Follow AGENTS section 2. Responses and progress begin with the applicable convention IDs and use polite ${languageName}. ${toolBasis} ${skillLoad} ${responseText}`,
      prompt: `C00 instruction row and C15 instruction-receipt row of ${conventions} (AGENTS 2.2): retain each atomic requirement verbatim in internal working state; merge compatible unfinished work and replace only contradicted actions. A correction requires the next authorized action in this turn. Do not create another ledger, scope, approval flow, or stopping condition. ${toolBasis} ${skillLoad} ${responseText}`,
      scratch: `C12 scratch row of ${conventions} (harness.config.json paths.scratch): task output is read and written only in the project ${scratch}/ directory; the harness scratch (AppData/Local/Temp, $env:TEMP, ~/.claude/projects/*/tool-results) is not a task location. Bound the query so its result returns inline.`,
    },
  };
}

function readDocument(relativePath) {
  const bytes = fs.readFileSync(path.join(root, relativePath));
  return { path: relativePath, sha256: createHash('sha256').update(bytes).digest('hex'), lines: bytes.toString('utf8').split(/\r?\n/) };
}

function headingLevel(line) {
  const match = /^(#+) /.exec(line);
  return match ? match[1].length : 0;
}

function section(document, heading) {
  const start = document.lines.findIndex(line => heading.test(line));
  if (start < 0) return null;
  const level = headingLevel(document.lines[start]);
  let end = start + 1;
  while (end < document.lines.length && !(headingLevel(document.lines[end]) > 0 && headingLevel(document.lines[end]) <= level)) end += 1;
  return { startLine: start + 1, endLine: end, lines: document.lines.slice(start, end) };
}

function tableRows(lines) {
  return lines
    .map(line => line.trim())
    .filter(line => line.startsWith('|') && line.endsWith('|') && !/^\|\s*-/.test(line))
    .map(line => line.slice(1, -1).split('|').map(cell => cell.trim()));
}

function exactHeading(text) {
  return new RegExp('^' + escapeRegex(text) + '$');
}

function requireSection(document, heading) {
  const found = section(document, exactHeading(heading));
  if (!found) throw new Error(`C16 compaction context: ${document.path} has no "${heading}" section, which the project hook reads after compaction (harness.config.json response headings).`);
  return found;
}

function compactContext(rules) {
  const headings = rules.documents;
  const conventions = readDocument(rules.conventions);
  const summary = requireSection(conventions, headings.summaryHeading);
  const clauses = tableRows(summary.lines).slice(1).filter(row => /^C\d{2}$/.test(row[0]));
  const agents = readDocument('AGENTS.md');
  const behavior = requireSection(agents, '### Behavior summary');
  const tracking = readDocument(rules.config.paths.tracking);
  const latest = section(tracking, headings.trackingHeading);
  const nodes = section(tracking, exactHeading(headings.nodeHeading));
  const nodeRows = nodes ? tableRows(nodes.lines).slice(1) : [];
  const audit = readDocument(rules.config.paths.audit);
  const todo = requireSection(audit, headings.todoHeading);
  const openRows = tableRows(todo.lines).slice(1).filter(row => row[1] === headings.openStatus);
  const text = [
    'C15 session-start row and C16 hook row (compaction): context summary the project hook read from the canonical documents. Do not reread every document; check only that the clauses, procedure rows, and owner files tied to the recent work still hold. A summary omission never exempts a clause.',
    `Conventions ${conventions.path} lines ${summary.startLine}-${summary.endLine} sha256 ${conventions.sha256}`,
    ...clauses.map(row => row[0] + ' ' + row[1]),
    `Behavior ${agents.path} lines ${behavior.startLine}-${behavior.endLine} sha256 ${agents.sha256}`,
    latest
      ? `Tracking ${tracking.path} lines ${latest.startLine}-${latest.endLine} sha256 ${tracking.sha256}`
      : `Tracking ${tracking.path} sha256 ${tracking.sha256}: no accepted session summary recorded`,
    ...(latest ? latest.lines : []),
    `Project nodes ${tracking.path} rows ${nodeRows.length}`,
    ...nodeRows.map(row => row.join(' | ')),
    `Open TODOs ${audit.path} lines ${todo.startLine}-${todo.endLine} sha256 ${audit.sha256} rows ${openRows.length}`,
    ...openRows.map(row => row.join(' | ')),
    rules.text.toolBasis,
    rules.text.skillLoad,
    rules.text.response,
  ].join('\n');
  if (text.length > contextCharacterLimit) throw new Error(`C16 compaction context: ${text.length} characters exceed the ${contextCharacterLimit}-character context limit; shorten the conventions summary table, the tracking summary, or the project node table.`);
  return text;
}

function sessionStartContext(rules, source) {
  if (source === 'startup' || source === 'clear') return rules.text.newSession;
  if (source === 'compact') return compactContext(rules);
  if (source === 'resume' || source === 'fork') return rules.text.resume;
  throw new Error(`C16 project hook: unsupported SessionStart source "${source}"; the supported sources are startup, clear, compact, resume, and fork.`);
}

function responseLabel(label) {
  return new RegExp('(?<![\\p{L}])(?:' + label + ')\\*{0,2}\\s*[:：]', 'u');
}

function fieldPattern(field) {
  return escapeRegex(field).replace(/ /g, '\\s*');
}

function missingResponseFields(rules, message) {
  return rules.fields.filter(field => !responseLabel(fieldPattern(field)).test(message));
}

function responseLabelValues(rules, message) {
  const labels = [...rules.fields.map(field => [field, fieldPattern(field)]), ['basis', rules.response.basisLabel]]
    .map(([name, label]) => ({ name, match: responseLabel(label).exec(message) }))
    .filter(label => label.match !== null)
    .sort((a, b) => a.match.index - b.match.index);
  return new Map(labels.map((label, index) => {
    const end = index + 1 < labels.length ? labels[index + 1].match.index : message.length;
    return [label.name, message.slice(label.match.index + label.match[0].length, end).trim()];
  }));
}

function citedSources(message) {
  const local = message.replace(/https?:\S+/g, ' ');
  return [...local.matchAll(/((?:[A-Za-z]:[\\/])?(?:[\w.@-]+[\\/])+[\w.@-]+\.[A-Za-z0-9]+)(?:#L(\d+)(?:-L(\d+))?|:(\d+)(?:-(\d+))?)?/g)]
    .map(match => ({ target: match[1], line: Math.max(...match.slice(2).map(value => Number(value ?? 0))) }));
}

function missingSources(message) {
  return citedSources(message).filter(({ target, line }) => {
    const file = path.resolve(root, target);
    if (!fs.existsSync(file) || !fs.statSync(file).isFile()) return true;
    return line > fs.readFileSync(file, 'utf8').split(/\r?\n/).length;
  }).map(({ target, line }) => (line > 0 ? target + ':' + line : target));
}

function userPrompts(state) {
  return fs.readdirSync(state)
    .filter(name => name.endsWith('-UserPromptSubmit.json'))
    .sort()
    .map(name => String(JSON.parse(fs.readFileSync(path.join(state, name), 'utf8')).prompt ?? ''));
}

let outputWritten = false;

function context(event, message) {
  outputWritten = true;
  process.stdout.write(JSON.stringify({ hookSpecificOutput: { hookEventName: event, additionalContext: message } }));
}

function record(state, kind, value) {
  const file = `${Date.now()}-${randomUUID()}-${kind}.json`;
  fs.writeFileSync(path.join(state, file), JSON.stringify(value), { flag: 'wx' });
}

function deny(state, message) {
  record(state, 'denial', { message });
  process.stderr.write(message);
  process.exitCode = 2;
}

function block(state, message) {
  record(state, 'block', { message });
  process.stderr.write(message);
  process.exitCode = 2;
}

function relative(raw, base) {
  if (typeof raw !== 'string' || raw.length === 0) throw new Error('C12 project hook: the tool input names no target path, so the protected-path and skill checks cannot resolve it.');
  const missing = [];
  let ancestor = path.resolve(base, raw);
  while (!fs.existsSync(ancestor)) {
    const parent = path.dirname(ancestor);
    if (parent === ancestor) throw new Error(`C12 project hook: no existing ancestor directory for the target path ${raw}.`);
    missing.unshift(path.basename(ancestor));
    ancestor = parent;
  }
  const absolute = path.resolve(fs.realpathSync(ancestor), ...missing);
  const rel = path.relative(fs.realpathSync(root), absolute);
  return rel === '..' || rel.startsWith('..' + path.sep) || path.isAbsolute(rel)
    ? 'OUTSIDE:' + absolute
    : rel.split(path.sep).join('/');
}

function inputBase(input) {
  return typeof input.cwd === 'string' && input.cwd.length > 0 ? input.cwd : root;
}

function claudeEdits(tool, params) {
  return [{
    target: params.file_path ?? params.notebook_path,
    operation: tool === 'Write' ? 'write' : 'update',
    before: tool === 'Write' ? '' : String(params.old_string ?? ''),
    after: String(tool === 'Write' ? params.content ?? '' : params.new_string ?? ''),
  }];
}

function patchText(command) {
  if (typeof command === 'string') return command;
  if (Array.isArray(command)) return command.find(part => typeof part === 'string' && part.includes('*** Begin Patch')) ?? '';
  return '';
}

function patchEdits(text) {
  const edits = [];
  let current = null;
  for (const line of text.split(/\r?\n/)) {
    const header = patchHeader.exec(line);
    if (header && header[1] === 'Move to') {
      if (current) current.moveTo = header[2].trim();
      continue;
    }
    if (header) {
      const operation = header[1] === 'Add File' ? 'add' : header[1] === 'Delete File' ? 'delete' : 'update';
      current = { target: header[2].trim(), operation, removed: [], added: [] };
      edits.push(current);
      continue;
    }
    if (!current || line.startsWith('***')) continue;
    if (line.startsWith('+')) current.added.push(line.slice(1));
    else if (line.startsWith('-')) current.removed.push(line.slice(1));
  }
  return edits.map(edit => ({ target: edit.target, operation: edit.operation, moveTo: edit.moveTo, before: edit.removed.join('\n'), after: edit.added.join('\n') }));
}

function shellText(command) {
  if (Array.isArray(command)) return command.map(String).join(' ');
  return String(command ?? '');
}

function decodeTool(input) {
  const tool = input.tool_name;
  const params = input.tool_input ?? {};
  if (harness === 'codex') {
    if (tool === codexTools.patch) return { kind: 'edit', edits: patchEdits(patchText(params.command)) };
    if (tool === codexTools.shell) return { kind: 'shell', command: shellText(params.command), background: false };
    return { kind: 'other' };
  }
  if (claudeTools.edit.has(tool)) return { kind: 'edit', edits: claudeEdits(tool, params) };
  if (claudeTools.shell.has(tool)) return { kind: 'shell', command: String(params.command ?? ''), background: params.run_in_background === true };
  if (claudeTools.read.has(tool)) return { kind: 'read', target: params.file_path ?? params.path };
  if (claudeTools.question.has(tool)) return { kind: 'question' };
  if (tool === 'TaskStop') return { kind: 'stop-task' };
  return { kind: 'other' };
}

function protectedReason(rules, rel) {
  if (rel.startsWith('OUTSIDE:')) return 'C12 file editing row: the target is outside the project root';
  const match = rules.protectedPaths.find(entry => entry.regex.test(rel));
  return match ? match.reason + ' (harness.config.json protectedPaths)' : '';
}

function countMatches(regex, text) {
  regex.lastIndex = 0;
  return (text.match(regex) ?? []).length;
}

function inlineViolation(rules, rel, edit) {
  const ext = comparable(path.extname(rel));
  return rules.inlineRules.find(rule => rule.extensionSet.has(ext)
    && !(rule.allow && rule.allow.test(rel))
    && countMatches(rule.regex, edit.after) > countMatches(rule.regex, edit.before)) ?? null;
}

function commentCount(code, ext) {
  const cLike = new Set(['ts', 'tsx', 'mts', 'cts', 'js', 'jsx', 'mjs', 'cjs', 'java', 'kt', 'cs', 'go', 'rs', 'swift', 'dart', 'php', 'glsl', 'hlsl', 'hlsli', 'c', 'cc', 'cpp', 'cxx', 'ixx', 'hpp', 'hh', 'h', 'cu', 'cuh']);
  const hash = new Set(['py', 'sh', 'rb', 'ps1', 'psm1']);
  if (!cLike.has(ext) && !hash.has(ext) && ext !== 'lua') return 0;
  let count = 0;
  let quote = '';
  for (let i = 0; i < code.length; i += 1) {
    const c = code[i];
    if (quote) {
      if (c === '\\') { i += 1; continue; }
      if (c === quote || (c === '\n' && quote !== '\u0060')) quote = '';
      continue;
    }
    if (c === '"' || c === '\'' || (c === '\u0060' && cLike.has(ext))) { quote = c; continue; }
    if (cLike.has(ext) && c === '/' && (code[i + 1] === '/' || code[i + 1] === '*')) { count += 1; i += 1; }
    if (hash.has(ext) && c === '#' && !(i === 0 && code[1] === '!')) count += 1;
    if (ext === 'lua' && c === '-' && code[i + 1] === '-') { count += 1; i += 1; }
  }
  return count;
}

function pendingNames(state, kind) {
  return fs.readdirSync(state).filter(name => name.startsWith('pending_' + kind + '_'));
}

function pendingAdd(state, kind) {
  fs.writeFileSync(path.join(state, 'pending_' + kind + '_' + randomUUID()), '', { flag: 'wx' });
}

function pendingKind(id) {
  return createHash('sha256').update(id).digest('hex').slice(0, 16);
}

function skillMarker(state, skill) {
  return path.join(state, 'skill_' + createHash('sha256').update(skill).digest('hex'));
}

function skillReset(state) {
  for (const name of fs.readdirSync(state).filter(entry => entry.startsWith('skill_'))) fs.unlinkSync(path.join(state, name));
}

function commandOutsideLiterals(tool, command) {
  const bash = tool === 'Bash' && process.platform !== 'win32';
  const literal = bash ? bashLiteral : powershellLiteral;
  const substitution = bash ? bashSubstitution : powershellSubstitution;
  return command.replace(literal, text => ((text.startsWith('"') || text.startsWith('@"')) && substitution.test(text) ? text : '\'\''));
}

function requiredSkills(rules, input, decoded, state) {
  const tool = input.tool_name;
  const base = inputBase(input);
  const shell = decoded.kind === 'shell';
  const command = shell ? commandOutsideLiterals(tool, decoded.command) : '';
  const slashCommand = command.replace(/\\/g, '/');
  const tokens = shell ? commandTokens(command) : [];
  const edits = decoded.kind === 'edit' ? decoded.edits : [];
  const targets = edits.flatMap(edit => [edit.target, edit.moveTo].filter(Boolean)).map(target => comparable(relative(target, base)));
  const mutates = edits.length > 0 || (shell && mutatingCommand.test(command));
  const executes = rules.executingToolPrefixes.some(prefix => tool.startsWith(prefix)) || tokens.some(token => rules.executables.has(token));
  const skills = mutates || executes ? [...rules.mandatory] : [];
  const scratchPrefix = comparable(rules.scratch + '/');
  const created = edits.some(edit => {
    const creates = edit.operation === 'add' || (edit.operation === 'write' && !fs.existsSync(path.resolve(base, edit.target)));
    return creates && !comparable(relative(edit.target, base)).startsWith(scratchPrefix);
  });
  const moved = edits.some(edit => edit.moveTo) || (shell && movingCommand.test(command));
  if (created || moved) skills.push(rules.structureSkill);
  const stackEdit = targets.some(target => rules.stackExtensions.has(path.extname(target)));
  const stackCheck = shell && rules.checks.some(check => check.pattern.test(command));
  if (stackEdit || stackCheck) skills.push(rules.stackSkill);
  for (const entry of rules.domain) {
    const byPath = targets.some(target => entry.paths.some(prefix => target.startsWith(comparable(prefix))));
    const byExtension = targets.some(target => entry.extensionSet.has(path.extname(target)));
    const byCommand = shell && entry.commands.some(item => tokenPattern(item).test(slashCommand));
    const byTool = entry.toolPrefixes.some(prefix => tool.startsWith(prefix));
    if (byPath || byExtension || byCommand || byTool) skills.push(entry.skill);
  }
  return [...new Set(skills)].filter(name => !fs.existsSync(skillMarker(state, name)));
}

function codexReceipt(text) {
  const exit = codexExit.exec(text);
  if (!exit) return null;
  const marker = text.indexOf('\nOutput:\n');
  return { exit: Number(exit[1]), stdout: marker >= 0 ? text.slice(marker + '\nOutput:\n'.length) : text, stderr: '', raw: {} };
}

function receipt(response) {
  if (typeof response === 'string') return codexReceipt(response);
  if (!response || typeof response !== 'object') return null;
  if (response.isError === true) return null;
  if (typeof response.output === 'string' && codexExit.test(response.output)) return codexReceipt(response.output);
  if (response.structuredContent) return receipt(response.structuredContent);
  if (Array.isArray(response.content)) {
    for (const item of response.content) {
      if (item.type !== 'text') continue;
      try {
        const parsed = receipt(JSON.parse(item.text));
        if (parsed) return parsed;
      } catch (error) {
        if (!(error instanceof SyntaxError)) throw error;
      }
    }
  }
  const exit = response.exit ?? response.exitCode ?? response.exit_code;
  if (!Number.isInteger(exit)) return null;
  return { exit, stdout: response.stdout ?? response.output ?? '', stderr: response.stderr ?? '', raw: response };
}

function checkFor(rules, input, decoded) {
  const runner = rules.toolRunner;
  if (runner && input.tool_name === runner.tool) {
    const id = runner.capabilities[String(input.tool_input?.capability ?? '')];
    return rules.checks.find(check => check.id === id) ?? null;
  }
  if (decoded.kind !== 'shell') return null;
  const command = commandOutsideLiterals(input.tool_name, decoded.command);
  const matched = rules.checks.filter(check => check.pattern.test(command));
  return matched.length === 1 ? matched[0] : null;
}

function checkAccepted(check, result) {
  if (!result || result.exit !== check.acceptance.exit || result.raw.interrupted === true || result.raw.ok === false) return false;
  const lines = (String(result.stdout) + '\n' + String(result.stderr)).split(/\r?\n/);
  if (typeof check.acceptance.line === 'string' && !lines.some(line => line.trim() === check.acceptance.line)) return false;
  if (typeof check.acceptance.forbidLine === 'string') {
    const forbidden = new RegExp(check.acceptance.forbidLine, 'i');
    if (lines.some(line => forbidden.test(line))) return false;
  }
  if (Array.isArray(check.acceptance.counts) && !check.acceptance.counts.every(key => result.raw[key] === 0)) return false;
  return true;
}

function editViolation(rules, decoded, base) {
  const deleted = new Set(decoded.edits.filter(edit => edit.operation === 'delete').map(edit => comparable(relative(edit.target, base))));
  for (const edit of decoded.edits) {
    const rel = relative(edit.target, base);
    for (const candidate of [rel, edit.moveTo ? relative(edit.moveTo, base) : null].filter(Boolean)) {
      const reason = protectedReason(rules, candidate);
      if (reason) return reason + ': ' + candidate;
    }
    const exists = fs.existsSync(path.resolve(base, edit.target));
    if (edit.operation === 'write' && exists) {
      return 'C12 file editing row: an existing file changes in part through Edit; Write creates new files only: ' + rel;
    }
    if (edit.operation === 'add' && (exists || deleted.has(comparable(rel)))) {
      return 'C12 file editing row: Add File creates new files only, and deleting then re-adding a file overwrites it; change an existing file in place with an Update File hunk: ' + rel;
    }
    const ext = path.extname(rel).slice(1).toLowerCase();
    if (rules.commentRule && commentCount(edit.after, ext) > commentCount(edit.before, ext)) {
      return 'C04 no-inline-comment rule (harness.config.json rules): this edit adds a comment to authored code; express the intent through names and structure: ' + rel;
    }
    const inline = inlineViolation(rules, rel, edit);
    if (inline) return `${inline.reason} (harness.config.json inlineRules ${inline.id}): ${rel}`;
  }
  return '';
}

function shellViolation(rules, tool, decoded, state) {
  const shellCommand = decoded.command;
  const outside = commandOutsideLiterals(tool, shellCommand);
  if (jobStop.test(outside) && !rules.response.stopInstruction.test(userPrompts(state).at(-1) ?? '')) {
    return 'C12 background row (AGENTS section 7): a running job is followed through its handle and completion notification; stopping it requires a stop instruction in the latest user prompt (harness.config.json response.stopInstruction).';
  }
  if (process.platform === 'win32' && cmdRoute.test(shellCommand)) {
    return 'C12 shell row: CMD, npm.cmd, and batch files are not the project command route; run the command in PowerShell.';
  }
  const forbiddenPath = rules.forbiddenPaths.find(entry => entry.regex.test(shellCommand));
  if (forbiddenPath) return forbiddenPath.reason + ' (harness.config.json forbiddenPaths)';
  if (!decoded.background && foregroundWait.test(shellCommand)) {
    return 'C12 background row (AGENTS 2.4.7): follow a running job through its handle and completion notification; this command waits in the foreground.';
  }
  const forbiddenCommand = rules.forbiddenCommands.find(entry => entry.pattern.test(outside));
  if (forbiddenCommand) {
    return `C12 stack table (harness.config.json stacks ${forbiddenCommand.stack}): "${forbiddenCommand.command}" is not a verification route (${forbiddenCommand.reason}).`;
  }
  const denied = rules.deniedCommands.find(entry => entry.regex.test(outside));
  if (denied) return denied.reason + ' (harness.config.json deniedCommands)';
  if (harnessScratch.test(shellCommand)) return rules.text.scratch;
  if (tmpWrite.test(shellCommand)) {
    return `C04 and C12 scratch rows (harness.config.json paths.scratch): temporary output belongs in the project ${rules.scratch}/ directory.`;
  }
  if (rules.protectedShell && rules.protectedShell.test(shellCommand)) {
    return 'C13 protected-target table (harness.config.json protectedPaths shellSegments): protected sources and generated output change only through their owning route.';
  }
  return '';
}

function preTool(rules, input, state) {
  const tool = input.tool_name;
  const decoded = decodeTool(input);
  const base = inputBase(input);
  if (decoded.kind === 'question') {
    deny(state, 'C12 interactive question row (AGENTS section 5): interactive question and approval tools are prohibited; the existing instruction is the authorization.');
    return;
  }
  if (harness === 'claude' && tool === 'Bash' && process.platform === 'win32') {
    deny(state, 'C12 shell row: on this Windows host project commands run through the PowerShell tool; the Bash tool runs Git Bash, which is not the project command route.');
    return;
  }
  if (decoded.kind === 'stop-task' && !rules.response.stopInstruction.test(userPrompts(state).at(-1) ?? '')) {
    deny(state, 'C12 background row (AGENTS section 7): a running job is followed through its handle and completion notification; stopping it requires a stop instruction in the latest user prompt (harness.config.json response.stopInstruction).');
    return;
  }
  if (decoded.kind === 'edit') {
    const violation = editViolation(rules, decoded, base);
    if (violation) { deny(state, violation); return; }
  }
  if (decoded.kind === 'read' && typeof decoded.target === 'string') {
    const forbidden = rules.forbiddenPaths.find(entry => entry.regex.test(decoded.target));
    if (forbidden) { deny(state, forbidden.reason + ' (harness.config.json forbiddenPaths)'); return; }
    if (harnessScratch.test(decoded.target)) { deny(state, rules.text.scratch + ' ' + decoded.target); return; }
  }
  if (decoded.kind === 'shell') {
    const violation = shellViolation(rules, tool, decoded, state);
    if (violation) { deny(state, violation); return; }
  }
  const unloaded = requiredSkills(rules, input, decoded, state);
  if (unloaded.length > 0) {
    deny(state, rules.text.skillLoad + ' Not loaded for this ' + tool + ' call: ' + unloaded.join(', ') + '.');
    return;
  }
  const check = checkFor(rules, input, decoded);
  if (check && typeof input.tool_use_id === 'string') {
    const token = createHash('sha256').update(input.tool_use_id).digest('hex');
    fs.writeFileSync(path.join(state, 'check_' + token + '.json'), JSON.stringify({ id: check.id, pending: pendingNames(state, pendingKind(check.id)) }));
  }
  const messages = [rules.text.toolBasis];
  if (rules.humanInputTools.has(tool)) {
    messages.push('C12 observation row: send no input while a human operates; preserve the running session and observe before and after authorized input.');
  }
  context('PreToolUse', messages.join(' '));
}

function markReadSkills(rules, input, state) {
  const serialized = JSON.stringify(input.tool_input ?? {}).replace(/\\\\/g, '/');
  for (const name of rules.knownSkills) {
    if (serialized.includes(rules.skillFile(name))) fs.writeFileSync(skillMarker(state, name), '');
  }
}

function postTool(rules, input, state) {
  if (input.hook_event_name === 'PostToolUseFailure') {
    context('PostToolUseFailure', `C06 FIX row of ${rules.conventions}: retain this failed tool input, cwd, full diagnostic, and exit; correct its demonstrated owner inside scope and continue the same goal. No verification was accepted. Apply the ${rules.fixSkill} skill before the next file change (AGENTS FIX behavior 3).`);
    return;
  }
  const decoded = decodeTool(input);
  const result = receipt(input.tool_response);
  const succeeded = input.tool_response?.isError !== true && (result === null || result.exit === 0);
  if (input.tool_name === 'Skill' && typeof input.tool_input?.skill === 'string') {
    fs.writeFileSync(skillMarker(state, input.tool_input.skill), '');
  }
  if (succeeded) markReadSkills(rules, input, state);
  if (decoded.kind === 'edit' && input.tool_response?.isError !== true) {
    for (const edit of decoded.edits) {
      for (const target of [edit.target, edit.moveTo].filter(Boolean)) {
        const ext = comparable(path.extname(String(target)));
        if (rules.stackExtensions.has(ext)) {
          for (const check of rules.checks) pendingAdd(state, pendingKind(check.id));
        }
        for (const observation of rules.observations) {
          if (observation.editExtensions.map(comparable).includes(ext)) pendingAdd(state, pendingKind('observation:' + observation.kind));
        }
      }
    }
  }
  const check = checkFor(rules, input, decoded);
  if (check && typeof input.tool_use_id === 'string') {
    const token = createHash('sha256').update(input.tool_use_id).digest('hex');
    const file = path.join(state, 'check_' + token + '.json');
    const snapshot = fs.existsSync(file) ? JSON.parse(fs.readFileSync(file, 'utf8')) : null;
    const accepted = snapshot?.id === check.id && checkAccepted(check, result);
    if (accepted) {
      for (const name of snapshot.pending) {
        if (pendingNames(state, pendingKind(check.id)).includes(name)) fs.unlinkSync(path.join(state, name));
      }
    }
    record(state, 'check_result', { tool_use_id: input.tool_use_id, id: check.id, accepted, result });
    context('PostToolUse', accepted
      ? `C07 and C14 fixed-check rows: ${check.id} (harness.config.json stacks checks) is accepted only for the edits that preceded this invocation; it proves neither runtime behavior nor whole-goal completion.`
      : `C06 and C14 rows: ${check.id} (harness.config.json stacks checks) is unaccepted because the invocation, exit, or acceptance condition is missing or the result failed. Retain every diagnostic; no pending check was cleared.`);
    return;
  }
  if (rules.observationTools.has(input.tool_name)) {
    context('PostToolUse', 'C12 observation row and C14: an observation call alone proves neither the affected path nor the same-build result. Retain the build, session, input, and frame identity and the actual consumption before acceptance.');
  }
}

function stop(rules, input, state) {
  const message = String(input.last_assistant_message ?? '').trim();
  const response = rules.response;
  const unloaded = rules.mandatory.filter(skill => !fs.existsSync(skillMarker(state, skill)));
  if (unloaded.length > 0) {
    block(state, rules.text.skillLoad + ' Not loaded: ' + unloaded.join(', ') + '.');
    return;
  }
  const banned = response.bannedJudgment.exec(message);
  if (banned) {
    block(state, `C00 forbidden-judgment and executor-judgment rows of ${rules.conventions} (harness.config.json response.bannedJudgment): "${banned[0]}". Execute the instructed correction instead of stopping, deferring, narrowing, or excusing.`);
    return;
  }
  const wrapper = response.gateWrapper.exec(message);
  if (wrapper) {
    block(state, `C16 user-designated block (AGENTS 2.6.4; harness.config.json response.gateWrapper): "${wrapper[0]}" is a user-level gate marker, not a project response form, and ending with an incomplete proof is deferral (C09). Execute the remaining requirement; report only result, evidence, unverified boundary, and remaining requirements.`);
    return;
  }
  const acknowledged = response.acknowledgment.exec(message);
  if (acknowledged && !(userPrompts(state).at(-1) ?? '').includes(response.reportKeyword)) {
    block(state, `C00 and C07 correction rows (harness.config.json response.acknowledgment): "${acknowledged[0].trim()}" answers a correction, but the latest user instruction does not contain "${response.reportKeyword}". Stop acknowledging or reporting and correct the pointed-out target now.`);
    return;
  }
  const conventions = fs.readFileSync(path.join(root, rules.conventions), 'utf8');
  const clauses = new Set([...conventions.matchAll(/<a id="(C\d+)"><\/a>/g)].map(match => match[1]));
  const cited = [...message.matchAll(/\bC\d+\b/g)].map(match => match[0]);
  if (cited.length === 0 || cited.some(clause => !clauses.has(clause))) {
    block(state, `C15 response row of ${rules.conventions}: report in the polite output language and state the actual applicable convention clause IDs that exist as <a id="Cnn"></a> anchors. Cite evidence only within its observed boundary.`);
    return;
  }
  const missing = missingResponseFields(rules, message);
  if (missing.length > 0) {
    block(state, rules.text.response + ' Missing: ' + missing.join(', ') + '.');
    return;
  }
  if (!responseLabel(response.basisLabel).test(message)) {
    block(state, rules.text.toolBasis + ' Missing: one of the basis labels (harness.config.json response.basisLabels).');
    return;
  }
  const values = responseLabelValues(rules, message);
  const unsupported = [...values].filter(([, value]) => !evidenceToken.test(value)).map(([name]) => name);
  if (unsupported.length > 0) {
    block(state, rules.text.response + ' Unsupported: ' + unsupported.join(', ') + '. Each label carries the clause ID, source path, measurement, or observation that settles it.');
    return;
  }
  if (!/\bC\d{2}\b/.test(values.get(rules.fields[0]) ?? '')) {
    block(state, `C15 response row: ${rules.fields[0]}: names the applicable convention clause IDs.`);
    return;
  }
  const absent = missingSources(message);
  if (absent.length > 0) {
    block(state, 'C09 and C14 evidence rows: cited sources do not exist or lack the cited line: ' + absent.join(', ') + '. Cite only sources actually read or measured.');
    return;
  }
  const guessed = response.guessedJudgment.exec(message);
  if (guessed) {
    block(state, rules.text.response + ' "' + guessed[0] + '" is a guess or a generality, not a basis (harness.config.json response.guessedJudgment).');
    return;
  }
  if (!response.basisStatement.test(message) || !response.numberedBasis.test(message)) {
    block(state, rules.text.response + ' Missing: the basis statement, the numbered evidence lines, or the conclusion (harness.config.json response.basisStatement and response.numberedBasis).');
    return;
  }
  if (input.stop_hook_active === true) return;
  const pendingIds = [...rules.checks.map(check => check.id), ...rules.observations.map(observation => 'observation:' + observation.kind)];
  const pending = pendingIds.filter(id => pendingNames(state, pendingKind(id)).length > 0);
  if (pending.length > 0) {
    const reminder = 'C00, C07, C14, and C15 rows: retained unaccepted checks: ' + pending.join(', ') + '. Execute the next authorized action. Do not claim completion from a command name, document, build, or observation call. Only an observed external dependency or governing constraint can prevent the exact remaining action; finish independent authorized work and report that boundary.';
    if (harness === 'codex') block(state, reminder);
    else context('Stop', reminder);
  }
}

async function readInput() {
  let raw = '';
  process.stdin.setEncoding('utf8');
  for await (const chunk of process.stdin) raw += chunk.toString();
  return raw;
}

function main(rules, raw) {
  const input = JSON.parse(raw);
  if (!events.has(input.hook_event_name)) throw new Error(`C16 project hook: unsupported hook event ${input.hook_event_name}; the registration lists SessionStart, UserPromptSubmit, PreToolUse, PostToolUse, PostToolUseFailure, and Stop.`);
  if (typeof input.session_id !== 'string' || input.session_id.length === 0) throw new Error('C16 project hook: the hook input has no session_id, so the session state directory cannot be resolved.');
  if (input.agent_id) return;
  const id = /^[a-zA-Z0-9_-]+$/.test(input.session_id) ? input.session_id : createHash('sha256').update(input.session_id).digest('hex');
  const state = path.join(root, rules.scratch, 'hook', harness, id);
  fs.mkdirSync(state, { recursive: true });
  record(state, input.hook_event_name, input);
  switch (input.hook_event_name) {
    case 'SessionStart':
      if (input.source === 'compact') skillReset(state);
      context('SessionStart', sessionStartContext(rules, input.source));
      break;
    case 'UserPromptSubmit':
      context('UserPromptSubmit', rules.text.prompt);
      break;
    case 'PreToolUse': preTool(rules, input, state); break;
    case 'PostToolUse':
    case 'PostToolUseFailure': postTool(rules, input, state); break;
    case 'Stop': stop(rules, input, state); break;
    default: throw new Error(`C16 project hook: no handler for the hook event ${input.hook_event_name}.`);
  }
}

let raw = '';
let failureDirectory = '';
try {
  if (!harness) throw new Error('C16 project hook: the registration passes no harness name; the generated hook command ends with "claude" or "codex" (SimpleHarness CLI registration).');
  const config = readConfig();
  failureDirectory = path.join(root, config.paths.scratch, 'hook', 'failure');
  raw = await readInput();
  main(createRules(config), raw);
  if (harness === 'codex' && !outputWritten && process.exitCode !== 2) process.stdout.write('{}');
} catch (error) {
  process.exitCode = 1;
  process.stderr.write(`C16 project hook failure (record kept under ${failureDirectory || 'the scratch hook/failure directory'}): ${error.stack ?? error}\n`);
  if (failureDirectory) {
    try {
      fs.mkdirSync(failureDirectory, { recursive: true });
      record(failureDirectory, 'failure', { error: String(error.stack ?? error), raw });
    } catch (recordError) {
      process.stderr.write(`project hook failure record failed: ${recordError.stack ?? recordError}\n`);
    }
  }
}
