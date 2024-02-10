<script>
export default {
  data() {
    return {
      recipes: [
        { name: 'Delicious Pasta', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!' },
      ],
    }
  },
}
</script>

<template>
  <div>
    <main class="container mx-auto mt-8">
      <h2 class="text-3xl font-bold mb-4">
        Latest Recipes
      </h2>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <Recipe v-for="recipe in recipes" :key="recipe.name" :recipe="recipe" />
      </div>
    </main>
  </div>
</template>
