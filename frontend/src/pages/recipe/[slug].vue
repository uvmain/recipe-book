<script setup lang="ts">
import type { Recipe } from '../../types/recipes'
import { backendFetchRequest } from '../../logic/fetchFromBackend'

const route = useRoute('/recipe/[slug]')
const recipe = ref()

const recipeSlug = computed(() => `${route.params.slug}`)

async function getRecipe() {
  try {
    const response = await backendFetchRequest(`recipes/${recipeSlug.value}`)
    const jsonData = await response.json() as Recipe
    recipe.value = jsonData
    document.title = recipe.value
      ? `RecipeBook: ${recipe.value.name}`
      : 'RecipeBook'
  }
  catch (error) {
    console.error('Failed to fetch recipe:', error)
  }
}

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
