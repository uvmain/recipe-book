import sqlite3 from 'sqlite3'
import type { Recipe } from '~/utils/recipe'

interface RecipesResponse {
  recipes?: Recipe[]
  error?: string
}

export default defineEventHandler(async (): Promise<RecipesResponse> => {
  const sql = 'SELECT * FROM recipes'

  const getRecipes = () => {
    return new Promise<Recipe[]>((resolve, reject) => {
      const db = new sqlite3.Database('../../recipes.db', (err) => {
        if (err) {
          console.error(err.message)
          reject(err)
        }
        console.debug('Connected to the recipes database.')
      })

      db.all(sql, [], (err, rows: Recipe[]) => {
        db.close()

        if (err) {
          console.warn(`Failed to select from recipes table; ${err.message}`)
          reject(err)
        }

        console.info('Returning rows from recipes table')
        console.log(JSON.stringify(rows))
        resolve(rows)
      })
    })
  }

  try {
    const recipes = await getRecipes()
    return { recipes }
  }
  catch (error) {
    const typedError = error as Error
    return { error: typedError.message }
  }
})