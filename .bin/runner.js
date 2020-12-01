import chalk from 'chalk';
import cp from 'child_process';
import { Command } from 'commander';
import fs from 'fs';
import path from 'path';

const today = new Date();
const program = new Command();

program
  .option('-y, --year <type>', 'year', parseFloat, today.getFullYear())
  .option('-d, --day <type>', 'day', parseFloat, today.getDate());

program.parse(process.argv);

const dayDir = `${program.year}/day${program.day}`;

const files = await fs.promises.readdir(dayDir);
const indexFileName = files.find(fileName => fileName.startsWith('index'));
const extension = path.extname(indexFileName);

let runner;

switch (extension) {
  case '.js': {
    runner = 'node';
    break;
  }
  case '.rb': {
    runner = 'ruby';
    break;
  }
  default: {
    console.log(chalk.red(`Unsupported extension: ${extension}`));
    process.exit(1);
  }
}

const child = cp.exec(`${runner} ./${indexFileName}`, {
  cwd: dayDir,
});

child.stdout.on('data', data => {
  console.log(data);
});

child.on('exit', code => {
  if (code !== 0) {
    throw new Error(`exit code - ${code}`);
  }
});
