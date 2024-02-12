<script setup lang="ts">
useHead({
  titleTemplate: 'RecipeBook: Home',
})

const { data: recipes } = await useAsyncData('recipes', () =>
  queryContent('/recipes').find())

const latestRecipes = computed(() => {
  if (recipes.value)
    return recipes.value.slice().sort((a, b) => b.id - a.id).slice(0, 10)
})
</script>

<template>
  <div>
    <main v-if="recipes && latestRecipes" class="w-80% mx-auto mt-8">
      <h2 class="text-3xl font-bold mb-4">
        Latest Recipes
      </h2>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-8">
        <RecipeCard v-for="recipe in latestRecipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
    <div v-else class="i-svg-spinners:3-dots-move w-1em h-1em" />
  </div>
</template>
