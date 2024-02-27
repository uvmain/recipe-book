import { writeFile } from 'node:fs/promises'

export default defineEventHandler(async (event): any => {
  const test = await readBody(event)
  // const slug = recipeJson.slug
  return test
})

// export async function saveRecipe(recipeJson: Recipe) {
//   try {
//     const slug = recipeJson.slug
//     await writeFile(`content/recipes/${slug}.json`, JSON.stringify(recipeJson))
//   }
//   catch (error) {
//     console.error('Error saving JSON data:', error)
//   }
// }
