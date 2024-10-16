import { getQuery } from 'h3'

const baseUrl = process.env.DIRECTUS_URL as string
const token = process.env.DIRECTUS_TOKEN as string

interface Query {
  filter?: string;
  limit?: string;
  offset?: string;
  courses?: string;
  vegetarian?: string;
  country?: string;
}

export interface Recipe {
  slug: string;
  date_created: string;
  name: string;
  author: string;
  source: string;
  course: string;
  vegetarian: boolean;
  prep_time: string;
  cooking_time: string;
  calories: string;
  servings: string;
  ingredients: string;
  instructions: string;
  image: string;
  country: string;
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
let cache: { [key: string]: any } = {}

const SIXTY_SECONDS_IN_MILLISECONDS = 60000
setInterval(clearCache, SIXTY_SECONDS_IN_MILLISECONDS)

function clearCache() {
  cache = {}
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
    _and.push({ _or })
  })

  if (query.courses?.length) {
    _and.push({ "course": { "_in": query.courses.split(",") }})
  }

  if (query.vegetarian?.length) {
    _and.push({ "vegetarian": { "_eq": "true" }})
  }

  if (query.country?.length) {
    _and.push({ "country": { "_eq": query.country }})
  }
  
  return { _and }
}

export default defineEventHandler(async (event) => {
  if (baseUrl.length < 1 || token.length < 1) {
    return {
      error: 'CMS connection variables are not defined'
    }
  }
  
  const url = `${baseUrl}/items/Recipe`
  const unparsedQuery = getQuery(event) as Query

  const cacheKey = JSON.stringify(unparsedQuery)

  if (unparsedQuery.filter === 'cacheclear') {
    cache = {}
    return { warning: 'Recipe cache cleared' }
  }

  if (cache[cacheKey]) {
    return cache[cacheKey]
  }

  try {
    const options = {
      headers: {
        Authorization: `Bearer ${token}`
      },
      query: {
        filter: generateFilter(unparsedQuery),
        sort: '-date_created',
        limit: unparsedQuery.limit,
        offset: unparsedQuery.offset,
      },
    }
    const response = await $fetch<{ data: Recipe[] }>(url, options).catch((error) => {
      throw new Error(`Failed to fetch data: ${JSON.stringify(error.data)}`)
    })

    cache[cacheKey] = response

    return response
  }
 catch (error) {
    return { error: `Failed to fetch data from CMS: ${error}` }
  }
})
