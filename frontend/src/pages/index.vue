<script setup lang="ts">
import type { RecipeCard } from '../types/recipes'
import { useDebounceFn, useSessionStorage } from '@vueuse/core'
import { useHead } from '@vueuse/head'
import { backendFetchRequest } from '../composables/fetchFromBackend'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const allRecipeCards = ref<RecipeCard[]>([])

const searchInput = useSessionStorage<string>('searchInput', '')
const selectedCourses = useSessionStorage<string[]>('selectedCourses', [])
const selectedVegetarian = useSessionStorage<boolean>('selectedVegetarian', false)
const selectedCountry = useSessionStorage<string>('selectedCountry', '')
const selectedCalories = useSessionStorage<number>('selectedCalories', 1000)

async function getRecipes() {
  let url = searchInput.value ? `recipecards?filter=${searchInput.value}` : 'recipecards'
  url = selectedCourses.value?.length ? `${url}&courses=${selectedCourses.value}` : url
  url = selectedVegetarian.value === true ? `${url}&vegetarian=true` : url
  url = selectedCountry.value?.length > 0 ? `${url}&country=${selectedCountry.value}` : url

  try {
    const response = await backendFetchRequest(url)
    const jsonData = await response.json() as RecipeCard[]
    allRecipeCards.value.push(...jsonData.filter((recipe) => {
      return selectedCalories.value === 1000 || (Number.parseInt(recipe.calories) / Number.parseInt(recipe.servings)) < selectedCalories.value
    }))
  }
  catch (error) {
    console.error('Failed to fetch recipes:', error)
  }
}

const search = useDebounceFn(async () => {
  allRecipeCards.value = []
  await getRecipes()
}, 700)

watch([searchInput, selectedCourses, selectedVegetarian, selectedCalories, selectedCountry], () => {
  search()
})

onMounted(async () => {
  await getRecipes()
})
</script>

<template>
  <div class="background min-h-screen flex flex-col">
    <div class="grid gap-4 p-4 lg:p-6 grid-cols-1 grid-rows-2 lg:grid-cols-4 md:grid-cols-3 xl:grid-cols-5 md:gap-6">
      <RecipeCard v-for="recipeCard in allRecipeCards" :key="recipeCard.name" :recipe-card="recipeCard" class="recipeCardBackground flex-1" />
    </div>
  </div>
</template>
