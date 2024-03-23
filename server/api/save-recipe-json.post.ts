import { writeFile } from 'node:fs/promises'
import { join } from 'node:path'
import process from 'node:process'

interface Body {
  fileName: string
  fileContent: string
}

export default defineEventHandler(async (event) => {
  const body: Body = await readBody(event)
  const fileName = body.fileName
  const fileContent = body.fileContent

  if (fileName && fileContent) {
    const filePath = join(process.cwd(), 'content/recipes', fileName)
    await writeFile(filePath, fileContent)
    return `${fileName} saved`
  }
  else {
    if (!fileName)
      console.log('missing fileName')
    if (!fileContent)
      console.log('missing fileContent')
  }
})
