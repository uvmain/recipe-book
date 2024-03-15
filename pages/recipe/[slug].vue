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
    <main class="mx-auto pt-1 px-2">
      <Recipe v-if="recipe" :recipe="recipe" class="lg:mx-auto lg:max-w-18/20" />
    </main>
  </div>
  <FloatingScrollToTop />
</template>
