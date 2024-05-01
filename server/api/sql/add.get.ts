import sqlite3 from 'sqlite3'

export default defineEventHandler(async () => {
  const sql = `INSERT INTO recipes (slug, dateAdded, name, author, source, course, country, vegetarian, prepTime, cookingTime, calories, servings, ingredients, instructions, image)
  VALUES (
    "salsa-negra",
    "2024-02-14",
    "Salsa Negra",
    "Honest Food",
    "https://honest-food.net/salsa-negra/",
    "sauces",
    "Mexico",
    0,
    "20 minutes",
    "5 minutes",
    "1790",
    "10",
    "- 180ml cup sunflower, canola or vegetable oil\n- 2 tsp lime juice\n- 1 tsp smoked salt\n- 1 tbsp dry roasted peanuts\n- 1 tbsp coriander seeds, toasted\n- 1 pasilla or guajillo chile, stemmed and seeded and torn up\n- 1 ancho chile, stemmed and seeded and torn up\n- 1 tbsp cumin\n- 1 tsp chipotle powder\n- 2 tbsp sesame seeds, toasted\n- 1/2 tsp garlic powder\n- 3 cloves black garlic",
    "- Puree all the ingredients on high in a blender until smooth.\n- Store in an airtight container in the fridge indefinitely.",
    "/recipe-images/salsa-negra.webp"
  );`

  const db = new sqlite3.Database('./public/database/db.sqlite', (err) => {
    if (err) {
      console.error(err.message)
    }
    else
      console.info('Connected to the recipes database.')
  })

  db.all(sql, [], (err) => {
    if (err) {
      console.warn(`Failed to add recipe to recipes table; ${err.message}`)
      return { outcome: err }
    }
    console.info('Added recipe to recipes table')
    return { outcome: 'success' }
  })

  db.close()
})
