const baseUrl = process.env.DIRECTUS_URL as string
const token = process.env.DIRECTUS_TOKEN as string

interface Query {
  filter?: string
}

function generateFilter(query: Query) {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const _and: any[] = []

  query.filter?.split(' ').forEach(filterString => {
    const _or = []
    _or.push({ "slug": { "_icontains": filterString }})
    _or.push({ "name": { "_icontains": filterString }})
    _or.push({ "author": { "_icontains": filterString }})
    _or.push({ "source": { "_icontains": filterString }})
    _or.push({ "course": { "_icontains": filterString }})
    _or.push({ "country": { "_icontains": filterString }})
    _or.push({ "ingredients": { "_icontains": filterString }})
    _or.push({ "instructions": { "_icontains": filterString }})
    _and.push({_or })
  })
  return { _and }
}

export default defineEventHandler(async (event) => {
  const url = `${baseUrl}/items/Recipe`
  const unparsedQuery = getQuery(event)
  try {
    const response = await $fetch(url, {
      headers: {
        Authorization: `Bearer ${token}`
      },
      query: {
        filter: generateFilter(unparsedQuery)
      }
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