<script setup lang="ts">
const route = useRoute()
const recipe = ref()

const recipeSlug = `${route.params.slug}`

async function getRecipe() {
  try {
    const response = await $fetch<RecipesApiResponse>((`/api/recipes/${recipeSlug}`))
    .catch((error) => {
      console.error(`Failed to fetch data: ${JSON.stringify(error.data)}`)
    })

    
    recipe.value = response ? response.data : []
  }
  catch (error) {
    console.error('Failed to fetch recipes:', error)
  }

}

const computedHead = computed(() => {
  return recipe.value ? `RecipeBook: ${recipe.value.name}` : 'RecipeBook'
})

useHead({
  titleTemplate: computedHead,
})

onBeforeMount(() => {
  getRecipe()
})
</script>

<template>
  <div class="text-black mx-auto md:max-w-18/19">
    <Recipe v-if="recipe" :recipe="recipe"/>
  </div>
</template>
