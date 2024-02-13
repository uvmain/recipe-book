<script setup lang="ts">

useHead({
  titleTemplate: 'RecipeBook: All',
})

const { data: recipes } = await useAsyncData('recipes', () =>
  queryContent('/recipes').find())

</script>

<template>
  <div>
    <main v-if="recipes" class="w-80% mx-auto mt-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-8">
        <RecipeCard v-for="recipe in recipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
    <div v-else class="i-svg-spinners:3-dots-move w-1em h-1em" />
  </div>
</template>
