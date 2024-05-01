import sqlite3 from 'sqlite3'

const createTableSql = `CREATE TABLE recipes (
  slug TEXT PRIMARY KEY,
  dateAdded DATE,
  name TEXT,
  author TEXT,
  source TEXT,
  course TEXT,
  country TEXT,
  vegetarian INTEGER,
  prepTime TEXT,
  cookingTime TEXT,
  calories TEXT,
  servings TEXT,
  ingredients TEXT,
  instructions TEXT,
  image TEXT
);`

export default defineEventHandler(async () => {
  const db = new sqlite3.Database('../../recipes.db', (err) => {
    if (err) {
      console.error(err.message)
    }
    console.info('Connected to the recipes database.')
  })

  db.all(createTableSql, [], (err) => {
    if (err) {
      console.warn(`Failed to create recipes table; ${err.message}`)
      return { outcome: err }
    }
    console.info('Created recipes table.')
    return { outcome: 'success' }
  })
  db.close()
})
