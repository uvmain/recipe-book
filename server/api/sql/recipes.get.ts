import process from 'node:process'
import path from 'node:path'
import sqlite3 from 'sqlite3'
import type { Recipe } from '~/utils/recipe'

export default defineEventHandler(async (): Promise<Recipe[]> => {
  const sql = 'SELECT * FROM recipes'

  const filePath = path.join(process.cwd(), 'public/database', 'db.sqlite')

  const getRecipes = () => {
    return new Promise<Recipe[]>((resolve, reject) => {
      const db = new sqlite3.Database(filePath, (err) => {
        if (err) {
          console.error(err.message)
          reject(err)
        }
        else
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
    return recipes
  }
  catch (error) {
    const typedError = error as Error
    console.error(typedError.message)
    return []
  }
})
