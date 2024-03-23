import { writeFile } from 'node:fs/promises'
import { join } from 'node:path'
import process from 'node:process'
import { Buffer } from 'node:buffer'

interface Body {
  fileName: string
  imageString: string
}

export default defineEventHandler(async (event) => {
  const body: Body = await readBody(event)
  const fileName = body.fileName
  const fileContentBase64 = body.imageString
  const base64Data = fileContentBase64.replace(/^data:image\/\w+;base64,/, '')
  const buffer = Buffer.from(base64Data, 'base64')

  if (fileName && fileContentBase64) {
    const filePath = join(process.cwd(), 'public/recipe-images', fileName)
    await writeFile(filePath, buffer)
    return `${fileName} saved`
  }
  else {
    if (!fileName)
      console.log('missing fileName')
    if (!fileContentBase64)
      console.log('missing fileContent')
  }
})
