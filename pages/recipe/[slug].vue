<script setup lang="ts">
const route = useRoute()
const recipe = ref()

const recipeSlug = `${route.params.slug}`

async function getRecipe() {
  const content = await queryContent('/recipes').where({ slug: { $eq: recipeSlug } }).find()
  recipe.value = content[0]
}

useHead({
  titleTemplate: recipe.value ? `RecipeBook: ${recipe.value.name}` : 'RecipeBook'
})

getRecipe()
</script>

<template>
  <div>
    <main class="container mx-auto mt-8">
      <div v-if="recipe">
        <h2 class="text-3xl font-bold mb-4">
          {{ recipe.name }}
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div>
            <img :src="recipe.image" :alt="recipe.name" class="w-full h-auto mb-4 rounded">
          </div>
          <div>
            <h3 class="text-xl font-bold mb-2">
              Ingredients:
            </h3>
            <ul class="list-disc ml-4">
              <li v-for="(ingredient, index) in recipe.ingredients" :key="index">
                {{ ingredient }}
              </li>
            </ul>
            <h3 class="text-xl font-bold mt-4">
              Instructions:
            </h3>
            <p>{{ recipe.instructions }}</p>
          </div>
        </div>
      </div>
      <div v-else>
        <p>Loading...</p>
      </div>
    </main>
  </div>
</template>
