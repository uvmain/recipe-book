const baseUrl = process.env.DIRECTUS_URL

export default defineEventHandler(async () => {
  const url = `${baseUrl}/items/Recipe`
  try {
    const response = await fetch(url)

    if (!response.ok) {
      throw new Error(`Failed to fetch data: ${response.statusText}`)
    }

    const { data } = await response.json()

    return data
  }
 catch (error) {
    return { error: 'Failed to fetch data from CMS' }
  }
})
