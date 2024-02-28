import path from 'node:path'
import process from 'node:process'
import buffer from 'node:buffer'
import sharp from 'sharp'

export default defineEventHandler(async (event) => {
  const fileName = `${getRouterParam(event, 'name')}`
  try {
    const base64string = await readRawBody(event, 'utf-8')
    if (base64string) {
      const base64ImageData = base64string.replace(/^data:image\/\w+;base64,/, '')
      const imageBuffer = buffer.Buffer.from(base64ImageData, 'base64')
      const filepath = path.join(process.cwd(), '/public/recipe-images', `${fileName}.webp`)

      await sharp(imageBuffer)
        .toFormat('webp')
        .toFile(filepath)

      console.log(`Saved ${fileName}`)

      return { statusCode: 200, body: 'Image saved successfully' }
    }
    else {
      console.error('Base64 string is empty or undefined')
      return { statusCode: 400, body: 'Base64 string is empty or undefined' }
    }
  }
  catch (err) {
    console.error(`Failed to save ${fileName}`)
    console.error(err)
    return { statusCode: 500, body: 'Failed to save image' }
  }
})
