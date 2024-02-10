<script setup lang="ts">
interface Recipe {
  id: number
  name: string
  ingredients: string[]
  instructions: string
  image: string
}

const recipes: Recipe[] = [
  { id: 1, name: 'Delicious Pasta', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 2, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
]
</script>

<template>
  <div>
    <main class="container mx-auto mt-8">
      <h2 class="text-3xl font-bold mb-4">
        Latest Recipes
      </h2>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <RecipeCard v-for="recipe in recipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
  </div>
</template>
