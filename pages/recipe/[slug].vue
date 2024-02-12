<script setup lang="ts">
const route = useRoute()
const recipe = ref()

const recipeSlug = `${route.params.slug}`

async function getRecipe() {
  const content = await queryContent('/recipes').where({ slug: { $eq: recipeSlug } }).findOne()
  recipe.value = content
}

useHead({
  titleTemplate: recipe.value ? `RecipeBook: ${recipe.value.name}` : 'RecipeBook'
})

getRecipe()
</script>

<template>
  <div>
    <main class="container mx-auto mt-8">
      <Recipe v-if="recipe" :recipe="recipe" />
      <div v-else>
        <p>Loading...</p>
      </div>
    </main>
  </div>
</template>
