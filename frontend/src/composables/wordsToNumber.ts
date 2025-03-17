const numericWords: Record<string, number> = {
  zero: 0,
  one: 1,
  two: 2,
  three: 3,
  four: 4,
  five: 5,
  six: 6,
  seven: 7,
  eight: 8,
  nine: 9,
  ten: 10,
  eleven: 11,
  twelve: 12,
  thirteen: 13,
  fourteen: 14,
  fifteen: 15,
  sixteen: 16,
  seventeen: 17,
  eighteen: 18,
  nineteen: 19,
  twenty: 20,
  thirty: 30,
  forty: 40,
  fifty: 50,
  sixty: 60,
  seventy: 70,
  eighty: 80,
  ninety: 90,
  onehundred: 100,
} as const

export function convertWordToNum(word: string): number {
  const lowercaseWord = word.toLowerCase()

  // Direct match (e.g., "fourteen")
  if (lowercaseWord in numericWords) {
    return numericWords[lowercaseWord]
  }

  // Handle concatenated words (e.g., "twentyone" -> 21)
  for (const key of Object.keys(numericWords)) {
    if (lowercaseWord.startsWith(key)) {
      const firstValue = numericWords[key]
      const remaining = lowercaseWord.slice(key.length)
      if (remaining in numericWords) {
        return firstValue + numericWords[remaining]
      }
    }
  }

  // If no match, return 0 as a fallback
  return 0
}
