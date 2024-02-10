<script setup lang="ts">
interface Recipe {
  id: number
  name: string
  ingredients: string[]
  instructions: string
  image: string
}

const recipes = ref<Recipe[]>([
  { id: 1, name: 'Delicious Pasta', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 2, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 3, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 4, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 5, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 6, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 7, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 8, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 9, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 10, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 11, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 12, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 13, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
])

const latestRecipes = computed(() => {
  // Sort recipes by id in descending order and take the first 10
  return recipes.value.slice().sort((a, b) => b.id - a.id).slice(0, 10)
})
</script>

<template>
  <div>
    <main class="w-80% mx-auto mt-8">
      <h2 class="text-3xl font-bold mb-4">
        Latest Recipes
      </h2>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-8">
        <RecipeCard v-for="recipe in latestRecipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
  </div>
</template>
