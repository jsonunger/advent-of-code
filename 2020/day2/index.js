import fs from 'fs';

const testCase = ['1-3 a: abcde', '1-3 b: cdefg', '2-9 c: ccccccccc'];

const input = fs.readFileSync('./input.txt', 'utf8').split('\n');

const numValid = inputs => {
  return inputs.filter(i => {
    const { min, max, letter, password } = i.match(
      /(?<min>\d+)-(?<max>\d+) (?<letter>\w): (?<password>\w+)/
    ).groups;
    const numLetters = password.replace(new RegExp(`[^${letter}]`, 'g'), '')
      .length;
    return numLetters >= min && numLetters <= max;
  }).length;
};

console.log('Test Case:', numValid(testCase));
console.log('Input:', numValid(input));

const numValid2 = inputs => {
  return inputs.filter(i => {
    const { first, second, letter, password } = i.match(
      /(?<first>\d+)-(?<second>\d+) (?<letter>\w): (?<password>\w+)/
    ).groups;
    const firstMatches = password[first - 1] === letter ? 1 : 0;
    const secondMatches = password[second - 1] === letter ? 1 : 0;
    return firstMatches + secondMatches === 1;
  }).length;
};

console.log('Test Case:', numValid2(testCase));
console.log('Input:', numValid2(input));
