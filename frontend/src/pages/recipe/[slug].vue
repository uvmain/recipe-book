<script setup lang="ts">
import type { RecipesApiResponse } from '../../types/recipes'
import { useHead } from '@vueuse/head'
import { backendFetchRequest } from '../../composables/fetchFromBackend'

const route = useRoute()
const recipe = ref()

const recipeSlug = `${route.params.slug}`

async function getRecipe() {
  try {
    const response = await backendFetchRequest(`recipes/${recipeSlug}`)
    const json = await response.json() as RecipesApiResponse
    recipe.value = json ? json.data : []
  }
  catch (error) {
    console.error('Failed to fetch recipes:', error)
  }
}

const computedHead = computed(() => {
  return recipe.value ? `RecipeBook: ${recipe.value.name}` : 'RecipeBook'
})

useHead({
  titleTemplate: computedHead,
})

onBeforeMount(() => {
  getRecipe()
})
</script>

<template>
  <div class="background min-h-screen flex flex-col">
    <div class="mx-auto lg:max-w-18/19">
      <Recipe v-if="recipe" :recipe="recipe" />
    </div>
  </div>
</template>
