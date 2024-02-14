<script setup lang="ts">
const route = useRoute()
const recipe = ref()

const recipeSlug = `${route.params.slug}`

async function getRecipe() {
  const content = await queryContent('/recipes').where({ slug: { $eq: recipeSlug } }).findOne()
  recipe.value = content
}

const computedHead = computed(() => {
  return recipe.value ? `RecipeBook: ${recipe.value.name}` : 'RecipeBook'
})

useHead({
  titleTemplate: computedHead,
})

getRecipe()
</script>

<template>
  <div>
    <main class="mx-auto container mt-8">
      <Recipe v-if="recipe" :recipe="recipe" class="ml-2 md:ml" />
      <div v-else class="i-svg-spinners:3-dots-move w-1em h-1em" />
    </main>
  </div>
</template>
