<script setup lang="ts">
import { ref, watch } from 'vue'
import { searchedSlugs } from '../utils/searchedSlugs'

useHead({
  titleTemplate: 'RecipeBook: Search',
})

const contentQuery = ref(searchedSlugs)
const recipes = ref<any[]>([])

watch(contentQuery, async (newValue) => {
  recipes.value = await queryContent('/recipes').where({ slug: { $in: newValue } }).find()
})
</script>

<template>
  <div>
    <main v-if="recipes && recipes.length" class="w-95% md:w-80% mx-auto mt-3 md:mt-8">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 lg:grid-cols-5 md:gap-8">
        <RecipeCard v-for="recipe in recipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
    <div v-else class="i-svg-spinners:3-dots-move w-1em h-1em" />
  </div>
</template>
