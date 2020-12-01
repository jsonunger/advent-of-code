import fs from 'fs';
import chalk from 'chalk';
import getTodayDir from './getTodayDir.js';

const dayDir = getTodayDir();

console.log(chalk.cyan(`Checking for ${dayDir}`));

let exists = true;

try {
  await fs.promises.access(dayDir);
} catch {
  exists = false;
}

if (exists) {
  console.log(chalk.green(`${dayDir} already exists. Exiting.`));
  process.exit(0);
}

try {
  await fs.promises.mkdir(dayDir, { recursive: true });
  const file = await fs.promises.open(`${dayDir}/index.js`, 'a');
  fs.closeSync(file.fd);
} catch (e) {
  console.log(chalk.red(`Error creating ${dayDir}`));
  console.error(e);
  process.exit(1);
}

console.log(chalk.green(`Successfully created ${dayDir}!`));
