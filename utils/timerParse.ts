export function getTimers(paragraph: string): number[] {
  const timerDetails: number[] = []

  const recipesWords = convertToAlphanumeric(paragraph).split('-')

  recipesWords.forEach((word: string, index) => {
    if (word.includes('min') || word.includes('hour')) {
      const prevWord = recipesWords[index - 1]
      if (prevWord === 'few') {
        timerDetails.push(word.startsWith('h') ? 180 : 3)
      }
      else if (prevWord === 'of' && recipesWords[index - 2] === 'couple' && word.endsWith('s')) {
        timerDetails.push(word.startsWith('h') ? 120 : 2)
      }
      else if (Number(prevWord) === 5 && Number(recipesWords[index - 2])) {
        const fraction = Number(`${recipesWords[index - 2]}.${prevWord}`)
        timerDetails.push(word.startsWith('h') ? 60 * fraction : fraction)
      }
      else if (Number(prevWord) && Number(recipesWords[index - 2])) {
        timerDetails.push(word.startsWith('h') ? 60 * Number(recipesWords[index - 2]) : Number(recipesWords[index - 2]))
      }
      else if (prevWord === 'more' && !Number.isNaN(recipesWords[index - 2])) {
        timerDetails.push(word.startsWith('h') ? 60 * Number(recipesWords[index - 2]) : Number(recipesWords[index - 2]))
      }
      else if (Number.isInteger(Number(prevWord)) && !word.endsWith('s')) {
        timerDetails.push(word.startsWith('h') ? 60 : 1)
      }
      else if (prevWord === 'an' || prevWord === 'a') {
        timerDetails.push(word.startsWith('h') ? 60 : 1)
      }
      else if (Number.isInteger(Number(prevWord))) {
        timerDetails.push(word.startsWith('h') ? Number(prevWord) * 60 : Number(prevWord))
      }
    }
  })

  return timerDetails.filter((time) => {
    return Number.isInteger(Number(time))
  })
}
