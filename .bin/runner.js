import * as x from 'child_process';
import getTodayDir from './getTodayDir.js';

const dayDir = getTodayDir();

const process = x.fork('./index.js', [], {
  cwd: dayDir,
});

process.on('exit', code => {
  if (code !== 0) {
    throw new Error(`exit code - ${code}`);
  }
});
