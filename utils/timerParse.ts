let timerArray: number[] = []

function addToTimerArray(amount: string, word: string) {
  if (Number(amount))
    timerArray.push(word.includes('hour') ? Number(amount) * 60 : Number(amount))
}

export function timerParse(recipeInstructions: string): number[] {
  timerArray = []
  const recipesWords = recipeInstructions.toLowerCase().split(' ')

  recipesWords.forEach((word: string, index) => {
    if (word.includes('min') || word.includes('hour')) {
      const prevWord = `${recipesWords[index - 1]}`.replaceAll('½', '.5').replaceAll('¾', '.75')
      if (prevWord === 'few') {
        addToTimerArray('3', word)
      }
      else if (prevWord.startsWith('~')) {
        addToTimerArray(prevWord.replace('~', ''), word)
      }
      else if (prevWord === 'of' && recipesWords[index - 2] === 'couple') {
        addToTimerArray('2', word)
      }
      else if (prevWord.includes('-')) {
        addToTimerArray(prevWord.split('-').slice(0)[0], word)
        addToTimerArray(prevWord.split('-').slice(-1)[0], word)
      }
      else if (prevWord === 'further' && !word.endsWith('s')) {
        addToTimerArray('1', word)
      }
      else if (Number(prevWord) && ['to', 'or'].includes(recipesWords[index - 2])) {
        addToTimerArray(recipesWords[index - 3], word)
        addToTimerArray(prevWord, word)
      }
      else {
        addToTimerArray(prevWord, word)
      }
    }
  })

  return timerArray
}
