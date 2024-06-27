const baseUrl = process.env.DIRECTUS_URL
const token = process.env.DIRECTUS_TOKEN as string

export default defineEventHandler(async (event) => {
  const slug = event.context.params?.slug
  const url = `${baseUrl}/items/Recipe/${slug}`

  try {
    const response = await $fetch(url, {
      headers: {
        Authorization: `Bearer ${token}`
      },
    })
    .catch((error) => {
      throw new Error(`Failed to fetch data: ${JSON.stringify(error.data)}`)
    });

    return response
  }
 catch (error) {
    return { error: `Failed to fetch data from CMS: ${error}` }
  }
})
