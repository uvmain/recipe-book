import process from 'node:process'
import path from 'node:path'
import sqlite3 from 'sqlite3'

export default defineEventHandler(async () => {
  const sql = `INSERT INTO recipes (
    slug, 
    dateAdded, 
    name, 
    author, 
    source, 
    course, 
    country, 
    vegetarian, 
    prepTime, 
    cookingTime, 
    calories, 
    servings, 
    ingredients, 
    instructions, 
    image
  ) 
  VALUES (
    'army-stew', 
    '2024-02-12', 
    'Budae Jjigae (Army Stew)', 
    'My Korean Kitchen', 
    'https://mykoreankitchen.com/army-stew-budae-jjigae/', 
    'mains', 
    'South Korea', 
    0, 
    '20 minutes', 
    '10 minutes', 
    '2612', 
    '4', 
    '**Main**\n- 1 ltr chicken stock\n- 200g SPAM thinly sliced\n- 4 cocktail Frankfurt sausages (150g)\n- thinly & diagonally sliced\n- 250g tofu sliced (about 1.5cm, 1/2 inch thickness)\n- 200g enoki mushrooms base stem removed & stems separated\n- 200g king oyster mushrooms thinly sliced length ways\n- 100g shiitake mushroom caps thinly sliced\n- 1/2 cup aged Kimchi , cut into bite sized pieces\n- 110g instant ramen noodles\n- 50g Korean rice cakes for soup soaked in cold water for 15 mins if it was frozen\n- 30g green onion thinly & diagonally sliced\n- 1 cheese slice\n\n**Sauce** (mix in a bowl)\n- 2 tbsp Korean chili flakes (Gochugaru)\n- 2 tbsp rice wine (mirin)\n- 1 tbsp soy sauce\n- 1 tbsp minced garlic\n- 1/2 tbsp sugar\n- 1/2 tbsp Korean chilli paste (Gochujang)\n- ground black pepper', 
    '- Assemble the main ingredients (except for instant ramen noodles, rice cakes, green onion and cheese) in a shallow pot.\n- Add the sauce in the middle. Pour the stock in the corner of the pot. Close the lid and boil it on medium high heat until the stock starts to boil (about 8 mins).\n- Add the remaining ingredients - instant ramen noodles, rice cakes, green onion and cheese on top of the pot and boil uncovered until the noodles are cooked (about 2 to 3 mins).\n- Reduce the heat to low (if you''re cooking on a portable burner and sharing the food at the dinning table).\n- Start dishing out soup, protein and vegetables onto your own soup bowl. Serve with steamed rice (& with other Korean side dishes).', 
    '/recipe-images/army-stew.webp'
  );`

  const filePath = path.join(process.cwd(), 'public/database', 'db.sqlite')

  const db = new sqlite3.Database(filePath, (err) => {
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
