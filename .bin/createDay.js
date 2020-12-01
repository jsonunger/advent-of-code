import fs from 'fs';
import chalk from 'chalk';
import { Command } from 'commander';

const today = new Date();
const program = new Command();

program
  .option('-e, --extension <type>', 'choose extension', 'js')
  .option('-y, --year <type>', 'year', parseFloat, today.getFullYear())
  .option('-d, --day <type>', 'day', parseFloat, today.getDate());

program.parse(process.argv);

const dayDir = `${program.year}/day${program.day}`;

console.log(chalk.cyan(`Checking for ${dayDir}`));

let exists = true;

try {
  await fs.promises.access(dayDir);
} catch {
  exists = false;
}

if (exists) {
  console.log(chalk.yellow(`${dayDir} already exists. Exiting.`));
  process.exit(0);
}

try {
  await fs.promises.mkdir(dayDir, { recursive: true });
  const file = await fs.promises.open(
    `${dayDir}/index.${program.extension}`,
    'a'
  );
  fs.closeSync(file.fd);
} catch (e) {
  console.log(chalk.red(`Error creating ${dayDir}`));
  console.error(e);
  process.exit(1);
}

console.log(chalk.green(`Successfully created ${dayDir}!`));
