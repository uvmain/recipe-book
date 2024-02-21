let minutesArray: number[] = []

function clearMinutesArray() {
  minutesArray = []
}

function addToMinutesArray(amount: string, word: string) {
  if (Number(amount))
    minutesArray.push(word.includes('hour') ? Number(amount) * 60 : Number(amount))
}

export function timerParse(recipeInstructions: string): number[] {
  clearMinutesArray()

  const recipesWords = recipeInstructions.toLowerCase().split(' ')

  recipesWords.forEach((word: string, index) => {
    if (word.includes('min') || word.includes('hour')) {
      const prevWord = `${recipesWords[index - 1]}`.replaceAll('½', '.5').replaceAll('¾', '.75')
      if (prevWord === 'few') {
        addToMinutesArray('3', word)
      }
      else if (prevWord.startsWith('~')) {
        addToMinutesArray(prevWord.replace('~', ''), word)
      }
      else if (prevWord === 'of' && recipesWords[index - 2] === 'couple') {
        addToMinutesArray('2', word)
      }
      else if (prevWord.includes('-')) {
        addToMinutesArray(prevWord.split('-').slice(0)[0], word)
        addToMinutesArray(prevWord.split('-').slice(-1)[0], word)
      }
      else if (prevWord === 'further' && !word.endsWith('s')) {
        addToMinutesArray('1', word)
      }
      else if (Number(prevWord) && ['to', 'or'].includes(recipesWords[index - 2])) {
        addToMinutesArray(recipesWords[index - 3], word)
        addToMinutesArray(prevWord, word)
      }
      else {
        addToMinutesArray(prevWord, word)
      }
    }
  })

  return minutesArray
}
