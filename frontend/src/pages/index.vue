<script setup lang="ts">
import type { RecipeCard } from '../types/recipes'
import { useSessionStorage } from '@vueuse/core'
import { useHead } from '@vueuse/head'
import { getFilteredRecipeCards } from '../composables/fetches'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const allRecipeCards = ref<RecipeCard[]>([])

const searchInput = useSessionStorage<string>('searchInput', '')
const selectedCourses = useSessionStorage<string[]>('selectedCourses', [])
const selectedVegetarian = useSessionStorage<boolean>('selectedVegetarian', false)
const selectedCountry = useSessionStorage<string>('selectedCountry', '')
const selectedCalories = useSessionStorage<number>('selectedCalories', 1000)

watch([searchInput, selectedCourses, selectedVegetarian, selectedCalories, selectedCountry], async () => {
  allRecipeCards.value = await getFilteredRecipeCards()
})

onMounted(async () => {
  allRecipeCards.value = await getFilteredRecipeCards()
})
</script>

<template>
  <div class="background min-h-screen flex flex-col">
    <div class="grid gap-4 p-4 lg:p-6 grid-cols-1 grid-rows-2 lg:grid-cols-4 md:grid-cols-3 xl:grid-cols-5 md:gap-6">
      <RecipeCard
        v-for="(recipeCard, index) in allRecipeCards"
        :key="recipeCard.slug"
        :recipe-card="recipeCard"
        :index="index"
        class="recipeCardBackground flex-1"
      />
    </div>
  </div>
</template>
