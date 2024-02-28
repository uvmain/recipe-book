<script setup lang="ts">
useHead({
  titleTemplate: 'RecipeBook: All',
})

const { data: recipes } = await useAsyncData('allrecipes', () =>
  queryContent('/recipes').find())
</script>

<template>
  <div>
    <main v-if="recipes" class="w-19/20 mx-auto mt-3 md:mt-8 md:w-4/5">
      <div class="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-5 md:gap-8">
        <RecipeCard v-for="recipe in recipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
  </div>
</template>
