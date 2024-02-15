/* eslint-disable ts/no-unsafe-assignment */
export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  return { body }
})
