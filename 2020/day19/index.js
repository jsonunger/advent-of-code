import fs from 'fs';

const INPUT_FILE = 'input.txt';
const END_OF_LINE = '\n';

const ruleNumOverrides = {
  8: [[42], [42, 8]],
  11: [
    [42, 31],
    [42, 11, 31],
  ],
};

// Load input data
const inputData = (function () {
  // read input file as UTF 8
  const inputText = fs.readFileSync(INPUT_FILE, 'utf-8');

  // Split into major sections
  const [ruleSectionText, messageSectionText] = inputText
    .split(`${END_OF_LINE}${END_OF_LINE}`)
    .map(section => section.split(END_OF_LINE));

  // extract rules
  const rules = ruleSectionText
    .map(ruleText => {
      // split rule line into sections
      const [ruleNumText, ruleDefText] = ruleText.trim().split(': ');

      // create rule object
      return {
        // get rule number
        ruleNum: parseInt(ruleNumText, 10),

        // get rule options
        options: ruleDefText.split(' | ').map(option =>
          // get option sequence
          option.split(' ').map(token => {
            if (token.startsWith('"')) {
              // token is a string character to match
              return token.substring(1, token.length - 1);
            }
            // token is a reference to another rule
            return parseInt(token, 10);
          })
        ),
      };
    })
    .map(rule => ({
      ...rule,
      options: ruleNumOverrides[rule.ruleNum] || rule.options,
    }));

  // trim messages and split into chars
  const messages = messageSectionText.map(message =>
    Array.from(message.trim())
  );

  /** @type {InputData} */
  const data = { rules, messages };
  return data;
})();

function findRule(ruleNum) {
  const rule = inputData.rules.find(r => r.ruleNum === ruleNum);
  if (!rule) throw new Error(`No rule found for ${ruleNum}`);
  return rule;
}

function testMessageAgainstRule(message, startRule) {
  const NO_MATCH = 0;

  /** @param {number} offset @param {RuleSequence} sequence @param {number} seqIndex @param {number} loopCount @returns {number[]} */
  function testAt(offset, sequence, seqIndex, loopCount) {
    // if we overflow the end, then fail
    if (offset >= message.length) {
      return [NO_MATCH];
    }

    // if we recurse too many times without a match, then fail
    if (loopCount > message.length) {
      return [NO_MATCH];
    }

    // get current token
    const token = sequence[seqIndex];
    const nextSeqIndex = seqIndex + 1;
    const isEndOfSequence = nextSeqIndex >= sequence.length;

    /** @type {number[]} */
    let nextOffsets;
    if (typeof token === 'number') {
      // token is rule reference
      const nextRule = findRule(token);

      // test each option
      nextOffsets = nextRule.options
        .map(nextSeq => testAt(offset, nextSeq, 0, loopCount + 1))
        .flat();
    } else if (message[offset] === token) {
      nextOffsets = [offset + 1];
    } else {
      nextOffsets = [NO_MATCH];
    }

    // if this is NOT the end of the sequence, then check next
    if (!isEndOfSequence) {
      nextOffsets = nextOffsets
        .filter(consumed => consumed !== NO_MATCH)
        .map(off => testAt(off, sequence, nextSeqIndex, 0))
        .flat();
    }

    // if this is end of the sequence, then return all the offsets that should be checked next
    return nextOffsets.filter(consumed => consumed !== NO_MATCH);
  }

  // Check if any of the matches used exactly all the characters from the input
  return startRule.options
    .map(sequence => testAt(0, sequence, 0, 0))
    .flat()
    .filter(off => off !== NO_MATCH)
    .some(off => off === message.length);
}

const rule0 = findRule(0);

const messagesMatchingRule0 = inputData.messages.filter(message =>
  testMessageAgainstRule(message, rule0)
);
console.log(
  `Number of messages matching rule 0 = ${messagesMatchingRule0.length}`
);
