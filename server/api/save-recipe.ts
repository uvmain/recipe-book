import path from 'node:path'
import fs from 'node:fs'
import process from 'node:process'
import type { Recipe } from '../../utils/recipe'

export default defineEventHandler(async (event) => {
  const body: Recipe = await readBody(event)
  try {
    const slug = body?.slug
    const jsonObject = JSON.stringify(body, null, 2)
    const filepath = path.join(process.cwd(), '/content/recipes', `${slug}.json`)
    fs.writeFileSync(filepath, jsonObject)
    console.log(`saving ${slug}.json`)
    return 200
  }
  catch (err) {
    console.error(err)
    return 500
  }
})
