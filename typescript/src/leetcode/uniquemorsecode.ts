/* eslint-disable no-multi-spaces */

/**
 * International Morse Code defines a standard encoding where each
 * letter is mapped to a series of dots and dashes, as follows:
 * 'a' maps to ".-",
 * 'b' maps to "-...",
 * 'c' maps to "-.-.", and so on.
 *
 * Given an array of strings words where each word can be written
 * as a concatenation of the Morse code of each letter.
 *
 * Return the number of different transformations among all words we have.
 * @example
 * Input: words = ["gin","zen","gig","msg"]
 * Output: 2
 * Explanation: The transformation of each word is:
 * "gin" -> "--...-."
 * "zen" -> "--...-."
 * "gig" -> "--...--."
 * "msg" -> "--...--."
 * There are 2 different transformations: "--...-." and "--...--.".
 */
export default class UniqueMorseCode {
  private static readonly MORSE_CODE_ARRAY: string[] = [
    '.-',     // a
    '-...',   // b
    '-.-.',   // c
    '-..',    // d
    '.',      // e
    '..-.',   // f
    '--.',    // g
    '....',   // h
    '..',     // i
    '.---',   // j
    '-.-',    // k
    '.-..',   // l
    '--',     // m
    '-.',     // n
    '---',    // o
    '.--.',   // p
    '--.-',   // q
    '.-.',    // r
    '...',    // s
    '-',      // t
    '..-',    // u
    '...-',   // v
    '.--',    // w
    '-..-',   // x
    '-.--',   // y
    '--..',   // z
  ];

  private static readonly ENGLISH_ALPHABETS: string = 'abcdefghijklmnopqrstuvwxyz';

  /**
   * @param words input consist of lowercase English letters
   * @returns the number of unique morsed words
   */
  public static Solver(words: string[]): number {
    const morseTransformationArray: string[] = [];
    for (let  i = 0; i < words.length; i += 1) {
      let asciiInterpolation = '';
      for (let j = 0; j < words[i].length; j += 1) {
        asciiInterpolation += this.MORSE_CODE_ARRAY[words[i][j].charCodeAt(0) - 97];
      }
      morseTransformationArray.push(asciiInterpolation);
    }

    // Don't forget to always use Set for filtering unique elements;
    return new Set(morseTransformationArray).size;
  }

  /**
   * @param words input consist of lowercase English letters
   * @returns the number of unique morsed words
   */
  public static MapSolver(words: string[]): number {
    const morsedWord: string[] = words.map(this.stringToMorseHelperFunction);
    const uniqMorsedWord: Set<string> = new Set(morsedWord);

    return uniqMorsedWord.size;
  }

  /**
   * @helper
   * @param word
   */
  private static stringToMorseHelperFunction(word: string) {
    const MORSE_CODE_ARRAY: string[] = [
      '.-',     // a
      '-...',   // b
      '-.-.',   // c
      '-..',    // d
      '.',      // e
      '..-.',   // f
      '--.',    // g
      '....',   // h
      '..',     // i
      '.---',   // j
      '-.-',    // k
      '.-..',   // l
      '--',     // m
      '-.',     // n
      '---',    // o
      '.--.',   // p
      '--.-',   // q
      '.-.',    // r
      '...',    // s
      '-',      // t
      '..-',    // u
      '...-',   // v
      '.--',    // w
      '-..-',   // x
      '-.--',   // y
      '--..',   // z
    ];

    let morsedWord: string = '';
    for (let i = 0; i < word.length; i += 1) {
      morsedWord += MORSE_CODE_ARRAY[word[i].charCodeAt(0) - 97];
    }

    return morsedWord;
  }

  public static UsingNewMapSolver(words: string[]) {
    const mapMorse = new Map(this.MORSE_CODE_ARRAY.map(
      (code, index) => [this.ENGLISH_ALPHABETS[index], code],
    ));

    const transformations: Set<string> = new Set();
    for (let i = 0; i < words.length; i += 1) {
      transformations.add(
        words[i]
          .split('')
          .map((char) => mapMorse.get(char)).join(''),
      );
    }

    return transformations.size;
  }
}
