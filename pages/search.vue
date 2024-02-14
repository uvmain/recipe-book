<script setup lang="ts">
useHead({
  titleTemplate: 'RecipeBook: Search',
})

const recipes = ref()
</script>

<template>
  <div>
    {{ searchedSlugs }}
    <main v-if="recipes" class="w-95% md:w-80% mx-auto mt-3 md:mt-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4 md:gap-8">
        <RecipeCard v-for="recipe in recipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
    <div v-else class="i-svg-spinners:3-dots-move w-1em h-1em" />
  </div>
</template>
