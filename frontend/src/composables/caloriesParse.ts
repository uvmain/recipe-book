export function calculateCaloriesPerServing(caloriesIn: string, servingsIn: string): string | null {
  const servings = Number.parseInt(servingsIn)
  const calories = Number.parseInt(caloriesIn)
  if (servings > 0 && calories > 0)
    return `${Math.floor(calories / servings)} per serving (${caloriesIn}/${servingsIn})`
  else if (calories > 0)
    return `${caloriesIn} total`
  else return null
}
