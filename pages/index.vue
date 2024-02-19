<script setup lang="ts">
import dayjs from 'dayjs'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const { data: recipes } = await useAsyncData('allrecipes', () =>
  queryContent('/recipes').find())

const latestRecipes = computed(() => {
  if (recipes.value)
    return recipes.value.slice().sort((a, b) => dayjs(b.dateAdded).diff(dayjs(a.dateAdded))).slice(0, 10)
})
</script>

<template>
  <div>
    <main v-if="recipes" class="mx-auto w-19/20 md:w-4/5 mt-3 md:mt-8">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 lg:grid-cols-5 md:gap-8">
        <RecipeCard v-for="recipe in latestRecipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
  </div>
</template>
