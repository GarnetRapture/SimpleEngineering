import fs from 'node:fs';
import path from 'node:path';
import { spawn } from 'node:child_process';
import { fileURLToPath } from 'node:url';

const root = path.resolve(fileURLToPath(new URL('../..', import.meta.url)));
const usage = 'C12 subagent row: usage is node .harness/subagents/run_subagent.mjs <subagent-id> <task-file>, where the task file holds the bounded task text inside the scratch directory (harness.config.json paths.scratch).';
const claudeReadOnlyTools = ['Edit', 'Write', 'NotebookEdit'];
const npmShimScript = /"%dp0%\\([^"]+\.js)"/i;

function readConfig() {
  return JSON.parse(fs.readFileSync(path.join(root, 'harness.config.json'), 'utf8'));
}

function isExecutableFile(candidate) {
  try {
    const stat = fs.statSync(candidate);
    return stat.isFile() && (stat.mode & 0o111) !== 0;
  } catch {
    return false;
  }
}

function resolveExecutable(name) {
  const directories = (process.env.PATH ?? '').split(path.delimiter).filter(Boolean);
  if (process.platform !== 'win32') {
    const found = directories.map(directory => path.join(directory, name)).find(isExecutableFile);
    return found ? { command: found, prefix: [] } : null;
  }
  const executable = directories.map(directory => path.join(directory, name + '.exe')).find(candidate => fs.existsSync(candidate));
  if (executable) return { command: executable, prefix: [] };
  for (const directory of directories) {
    const shim = path.join(directory, name + '.cmd');
    if (!fs.existsSync(shim)) continue;
    const script = npmShimScript.exec(fs.readFileSync(shim, 'utf8'));
    if (script) return { command: process.execPath, prefix: [path.join(directory, script[1])] };
  }
  return null;
}

function providerEnvironment(subagent, token) {
  if (!subagent.baseUrl) return null;
  const environment = { ANTHROPIC_BASE_URL: subagent.baseUrl };
  environment[subagent.authVariable] = token;
  if (subagent.authVariable === 'ANTHROPIC_AUTH_TOKEN') environment.ANTHROPIC_API_KEY = '';
  if (subagent.model) {
    environment.ANTHROPIC_MODEL = subagent.model;
    environment.ANTHROPIC_DEFAULT_OPUS_MODEL = subagent.model;
    environment.ANTHROPIC_DEFAULT_SONNET_MODEL = subagent.model;
    environment.ANTHROPIC_DEFAULT_HAIKU_MODEL = subagent.model;
    environment.CLAUDE_CODE_SUBAGENT_MODEL = subagent.model;
  }
  return environment;
}

function runProcess(executable, args, logFile) {
  return new Promise((resolve, reject) => {
    const child = spawn(executable.command, [...executable.prefix, ...args], { cwd: root, stdio: ['ignore', 'pipe', 'pipe'], windowsHide: true });
    const stdout = [];
    const stderr = [];
    child.stdout.on('data', chunk => stdout.push(chunk));
    child.stderr.on('data', chunk => stderr.push(chunk));
    child.on('error', reject);
    child.on('close', code => {
      const output = Buffer.concat(stdout).toString('utf8');
      const errors = Buffer.concat(stderr).toString('utf8');
      fs.writeFileSync(logFile, JSON.stringify({ command: executable.command, args: [...executable.prefix, ...args], exit: code, stdout: output, stderr: errors }, null, 2));
      resolve({ exit: code ?? 1, stdout: output, stderr: errors });
    });
  });
}

async function runClaude(subagent, task, runDirectory, stamp) {
  const executable = resolveExecutable('claude');
  if (!executable) throw new Error('C12 subagent row: the claude command is not on PATH; install Claude Code (SimpleHarness recommended tools) before delegating to ' + subagent.id + '.');
  const args = ['-p', task, '--output-format', 'json', '--no-session-persistence', '--setting-sources', 'user'];
  let settingsFile = null;
  const token = subagent.tokenEnv ? process.env[subagent.tokenEnv] : subagent.tokenValue;
  const environment = providerEnvironment(subagent, token);
  if (environment) {
    settingsFile = path.join(runDirectory, stamp + '.settings.json');
    fs.writeFileSync(settingsFile, JSON.stringify({ env: environment }), { mode: 0o600, flag: 'wx' });
    args.push('--settings', settingsFile);
  }
  if (subagent.model) args.push('--model', subagent.model);
  args.push('--disallowedTools', ...claudeReadOnlyTools);
  try {
    const result = await runProcess(executable, args, path.join(runDirectory, stamp + '.log'));
    if (result.exit !== 0) return result;
    const parsed = JSON.parse(result.stdout);
    return { ...result, answer: String(parsed.result ?? '') };
  } finally {
    if (settingsFile) fs.rmSync(settingsFile, { force: true });
  }
}

async function runCodex(subagent, task, runDirectory, stamp) {
  const executable = resolveExecutable('codex');
  if (!executable) throw new Error('C12 subagent row: the codex command is not on PATH; install Codex (SimpleHarness recommended tools) before delegating to ' + subagent.id + '.');
  const answerFile = path.join(runDirectory, stamp + '.answer.md');
  const args = ['exec', '--sandbox', 'read-only', '-c', 'features.hooks=false', '--json', '-o', answerFile];
  if (subagent.localProvider) args.push('--oss', '--local-provider', subagent.localProvider);
  if (subagent.model) args.push('-m', subagent.model);
  args.push(task);
  const result = await runProcess(executable, args, path.join(runDirectory, stamp + '.log'));
  if (result.exit !== 0) return result;
  return { ...result, answer: fs.readFileSync(answerFile, 'utf8') };
}

async function main(argv) {
  if (argv.length !== 2) throw new Error(usage);
  const [id, taskArgument] = argv;
  const config = readConfig();
  const subagent = config.subagents.find(entry => entry.id === id);
  if (!subagent) throw new Error(`C12 subagent row (harness.config.json subagents): no subagent has the id "${id}"; registered ids are ${config.subagents.map(entry => entry.id).join(', ') || 'none'}.`);
  const scratch = path.join(root, config.paths.scratch);
  const taskFile = path.resolve(process.cwd(), taskArgument);
  const relativeTask = path.relative(scratch, taskFile);
  if (relativeTask.startsWith('..') || path.isAbsolute(relativeTask)) throw new Error(usage + ' The given file is outside ' + config.paths.scratch + '/: ' + taskFile);
  const task = fs.readFileSync(taskFile, 'utf8').trim();
  if (task.length === 0) throw new Error(usage + ' The task file is empty: ' + taskFile);
  if (subagent.tokenEnv && !process.env[subagent.tokenEnv]) {
    throw new Error(`C12 subagent row (harness.config.json subagents ${id}.tokenEnv): the environment variable ${subagent.tokenEnv} is empty; set the provider API key in it before delegating.`);
  }
  const runDirectory = path.join(scratch, 'subagent', id);
  fs.mkdirSync(runDirectory, { recursive: true });
  const stamp = new Date().toISOString().replace(/[:.]/g, '-');
  const result = subagent.runtime === 'codex'
    ? await runCodex(subagent, task, runDirectory, stamp)
    : await runClaude(subagent, task, runDirectory, stamp);
  const logFile = path.relative(root, path.join(runDirectory, stamp + '.log')).split(path.sep).join('/');
  if (result.exit !== 0) {
    process.stderr.write(`C12 subagent row: ${id} exited with ${result.exit}; the full output is in ${logFile}.\n${result.stderr.slice(-4000)}\n`);
    return result.exit;
  }
  process.stdout.write(`[subagent ${id} — unverified result; verify it against the sources before use (C12 subagent row); log ${logFile}]\n${result.answer}\n`);
  return 0;
}

try {
  process.exitCode = await main(process.argv.slice(2));
} catch (error) {
  process.stderr.write(`${error.message ?? error}\n`);
  process.exitCode = 1;
}
