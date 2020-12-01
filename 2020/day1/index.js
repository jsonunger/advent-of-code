import fs from 'fs';

const testCase = [1721, 979, 366, 299, 675, 1456];

const input = fs
  .readFileSync('./input.txt', 'utf8')
  .split('\n')
  .filter(x => !!x)
  .map(Number);

const sum2020 = arr => {
  for (let i = 0; i < arr.length - 1; i += 1) {
    for (let j = i + 1; j < arr.length; j += 1) {
      if (arr[i] + arr[j] === 2020) {
        console.log(arr[i] * arr[j]);
      }
    }
  }
};

console.log('Test Case for two numbers:');
sum2020(testCase);
console.log('Input for two numbers:');
sum2020(input);

const sum2020With3 = arr => {
  for (let i = 0; i < arr.length - 2; i += 1) {
    for (let j = i + 1; j < arr.length - 1; j += 1) {
      for (let k = j + 1; k < arr.length; k += 1) {
        if (arr[i] + arr[j] + arr[k] === 2020) {
          console.log(arr[i] * arr[j] * arr[k]);
        }
      }
    }
  }
};

console.log('Test Case for three numbers:');
sum2020With3(testCase);
console.log('Input for three numbers:');
sum2020With3(input);
