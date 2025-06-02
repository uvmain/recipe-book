import type { RecipeCard } from '../types/recipes'
import { useSessionStorage } from '@vueuse/core'
import { backendFetchRequest } from '../composables/fetchFromBackend'

const searchInput = useSessionStorage<string>('searchInput', '')
const selectedCourses = useSessionStorage<string[]>('selectedCourses', [])
const selectedVegetarian = useSessionStorage<boolean>('selectedVegetarian', false)
const selectedCountry = useSessionStorage<string>('selectedCountry', '')
const selectedCalories = useSessionStorage<number>('selectedCalories', 1000)

export async function getFilteredRecipeCards(): Promise<RecipeCard[]> {
  let url = searchInput.value || selectedCourses.value?.length || selectedVegetarian.value === true || selectedCountry.value?.length > 0 || selectedCalories.value !== 1000 ? 'recipecards?' : 'recipecards'
  url = searchInput.value ? `${url}&filter=${searchInput.value}` : url
  url = selectedCourses.value?.length ? `${url}&courses=${selectedCourses.value.join(',')}` : url
  url = selectedVegetarian.value === true ? `${url}&vegetarian=true` : url
  url = selectedCountry.value?.length > 0 ? `${url}&country=${selectedCountry.value}` : url
  url = selectedCalories.value !== 1000 ? `${url}&calories=${selectedCalories.value}` : url

  try {
    const response = await backendFetchRequest(url)
    const jsonData = await response.json() as RecipeCard[]
    return jsonData
  }
  catch (error) {
    console.error('Failed to fetch recipes:', error)
    return [] as RecipeCard[]
  }
}

export async function getRandomFilteredRecipeCard(): Promise<RecipeCard> {
  const recipeCards = await getFilteredRecipeCards()
  const recipeCard = recipeCards[Math.floor(Math.random() * recipeCards.length)]
  return recipeCard
}

export async function getRecipeCount(): Promise<number> {
  const response = await backendFetchRequest('recipe-count')
  const count = await response.json() as number
  return count
}
