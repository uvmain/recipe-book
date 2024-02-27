import path from 'node:path'
import fs from 'node:fs'
import process from 'node:process'

export default defineEventHandler(async (event) => {
  const name = getRouterParam(event, 'name')

  try {
    // eslint-disable-next-line ts/no-unsafe-assignment
    const body = await readBody(event)
    // console.warn(`body: ${body}`)
    // const extension = name?.split('.').pop() // Extract extension from the file name
    // const filepath = path.join(process.cwd(), '/public/recipe-images', `${name}.${extension}`)
    // // eslint-disable-next-line ts/no-unsafe-argument
    // fs.writeFileSync(filepath, body, 'binary') // Write binary data to the file
    // console.log(`saving ${name}`)
    return 200
  }
  catch (err) {
    console.error(`Failed to save ${name}`)
    console.error(err)
    return 500
  }
})
