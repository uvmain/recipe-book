export function getTimer(sentence: string): number {
  let timerDetails: number = 0

  const recipesWords = convertToAlphanumeric(sentence).split('-')

  recipesWords.forEach((word: string, index) => {
    if (word.includes('min') || word.includes('hour')) {
      const prevWord = `${recipesWords[index - 1]}`.replaceAll('½', '.5').replaceAll('¾', '.75')
      if (prevWord === 'few') {
        timerDetails = word.startsWith('h') ? 180 : 3
      }
      else if (prevWord === 'of' && recipesWords[index - 2] === 'couple') {
        timerDetails = word.startsWith('h') ? 120 : 2
      }
      else if (prevWord === 'further' && !word.endsWith('s')) {
        timerDetails = word.startsWith('h') ? 60 : 1
      }
      else {
        timerDetails = word.startsWith('h') ? Number(prevWord) * 60 : Number(prevWord)
      }
    }
  })

  return timerDetails
}
