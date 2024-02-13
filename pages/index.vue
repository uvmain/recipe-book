<script setup lang="ts">
import dayjs from 'dayjs'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const { data: recipes } = await useAsyncData('recipes', () =>
  queryContent('/recipes').find())

const latestRecipes = computed(() => {
  if (recipes.value)
    return recipes.value.slice().sort((a, b) => dayjs(b.dateAdded).diff(dayjs(a.dateAdded))).slice(0, 10)
})
</script>

<template>
  <div>
    <main v-if="recipes" class=" w-95% md:w-80% mx-auto mt-3 md:mt-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4 md:gap-8">
        <RecipeCard v-for="recipe in latestRecipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
    <div v-else class="i-svg-spinners:3-dots-move w-1em h-1em" />
  </div>
</template>
