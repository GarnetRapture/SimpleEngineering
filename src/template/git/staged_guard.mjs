import { execFileSync } from 'node:child_process';
import fs from 'node:fs';
import path from 'node:path';

const git = (args, cwd, input) => execFileSync('git', args, { cwd, input, encoding: 'utf8', maxBuffer: 64 * 1024 * 1024 });
const root = git(['rev-parse', '--show-toplevel'], process.cwd()).trim();
const config = JSON.parse(fs.readFileSync(path.join(root, 'harness.config.json'), 'utf8'));
const blockedExtensions = new Set(config.commitGuard.blockedExtensions.map(extension => extension.toLowerCase()));

const staged = git(['diff', '--cached', '--name-only', '-z', '--diff-filter=ACMR'], root)
  .split('\0')
  .filter(entry => entry.length > 0);

if (staged.length === 0) {
  process.exit(0);
}

const sizes = git(['cat-file', '--batch-check=%(objectsize)'], root, staged.map(file => ':' + file).join('\n') + '\n')
  .split('\n')
  .filter(line => line.length > 0);

const toMegabytes = bytes => (bytes / (1024 * 1024)).toFixed(2);
const violations = [];
let commitBytes = 0;

staged.forEach((file, index) => {
  const extension = file.includes('.') ? file.slice(file.lastIndexOf('.') + 1).toLowerCase() : '';
  const size = Number.parseInt(sizes[index] ?? '', 10);
  if (Number.isNaN(size)) {
    violations.push(file + '  the staged blob size is unreadable: git cat-file --batch-check returned no object size for it');
    return;
  }
  commitBytes += size;
  if (blockedExtensions.has(extension)) {
    violations.push(file + '  blocked extension .' + extension + ' listed in harness.config.json commitGuard.blockedExtensions (' + toMegabytes(size) + ' MB)');
  }
});

if (violations.length > 0) {
  process.stderr.write('C16 commit guard (.githooks/pre-commit): the commit is stopped because these staged files violate the C02 repository row (large files stay out of git tracking):\n');
  for (const violation of violations) {
    process.stderr.write('  ' + violation + '\n');
  }
  process.stderr.write('Keep these files out of git with .gitignore or move them to a local path, then commit the same message again.\n');
  process.exit(1);
}

process.stdout.write('staged ' + staged.length + ' files, ' + toMegabytes(commitBytes) + ' MB\n');
