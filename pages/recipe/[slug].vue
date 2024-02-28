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
    <main class="mx-auto mt-8 max-w-full md:max-w-4/5">
      <Recipe v-if="recipe" :recipe="recipe" class="ml-2 md:ml" />
    </main>
  </div>
</template>
