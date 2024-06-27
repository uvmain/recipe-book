const baseUrl = process.env.DIRECTUS_URL as string
const token = process.env.DIRECTUS_TOKEN as string

interface Query {
  filter?: string
}

export interface Recipe {
  slug: string
  date_created: string
  name: string
  author: string
  source: string
  course: string
  vegetarian: boolean
  prep_time: string
  cooking_time: string
  calories: string
  servings: string
  ingredients: string
  instructions: string
  image: string
  country: string
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
  if (baseUrl.length < 1 || token.length < 1) {
    return {
      error: 'CMS connection variables are not defined'
    }
  }
  
  const url = `${baseUrl}/items/Recipe`
  const unparsedQuery = getQuery(event)
  try {
    const options = {
      headers: {
        Authorization: `Bearer ${token}`
      },
      query: {
        filter: generateFilter(unparsedQuery),
        sort: '-date_created',
      },
    }
    const response = await $fetch<{ data: Recipe[]}>(url, options)
    .catch((error) => {
      throw new Error(`Failed to fetch data: ${JSON.stringify(error.data)}`)
    });

    return response
  }
 catch (error) {
    return { error: `Failed to fetch data from CMS: ${error}` }
  }
})