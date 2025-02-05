<script setup lang="ts">
import type { Recipe } from '../../types/recipes'
import { useHead } from '@vueuse/head'
import { backendFetchRequest } from '../../composables/fetchFromBackend'

const route = useRoute()
const recipe = ref()

const recipeSlug = computed(() => `${route.params.slug}`)

async function getRecipe() {
  try {
    const response = await backendFetchRequest(`recipes/${recipeSlug.value}`)
    const jsonData = await response.json() as Recipe
    recipe.value = jsonData
  }
  catch (error) {
    console.error('Failed to fetch recipe:', error)
  }
}

const computedHead = computed(() => {
  return recipe.value ? `RecipeBook: ${recipe.value.name}` : 'RecipeBook'
})

useHead({
  titleTemplate: computedHead,
})

watch(recipeSlug, () => {
  getRecipe()
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
