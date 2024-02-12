<script setup lang="ts">
import dayjs from 'dayjs'

useHead({
  titleTemplate: 'RecipeBook: Home',
})

const isLatestView = ref(true)

const toggleView = () => {
  isLatestView.value = !isLatestView.value
}

const { data: recipes } = await useAsyncData('recipes', () =>
  queryContent('/recipes').find())

const latestRecipes = computed(() => {
  if (recipes.value)
    return recipes.value.slice().sort((a, b) => dayjs(b.dateAdded).diff(dayjs(a.dateAdded))).slice(0, 10)
})

const allRecipes = computed(() => {
  return recipes.value
})

const displayedRecipes = computed(() => {
  if (recipes.value) {
    return isLatestView.value ? latestRecipes.value : allRecipes.value
  }
})
</script>

<template>
  <div>
    <main v-if="recipes" class="w-80% mx-auto mt-8">
      <h2 class="text-3xl font-bold">
          {{ isLatestView ? 'Latest Recipes' : 'All Recipes' }}
        </h2>
        <button @click="toggleView" class="text-blue-500 hover:underline">
          {{ isLatestView ? 'All Recipes' : 'Latest Recipes' }}
        </button>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-8">
        <RecipeCard v-for="recipe in displayedRecipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
    <div v-else class="i-svg-spinners:3-dots-move w-1em h-1em" />
  </div>
</template>
