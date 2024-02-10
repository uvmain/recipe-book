<script setup lang="ts">
const recipes = [
  { id: 1, name: 'Delicious Pasta', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
  { id: 2, name: 'Puttanesca', ingredients: ['Pasta', 'Sauce', 'Cheese'], instructions: 'Cook pasta, add sauce, sprinkle cheese, enjoy!', image: '/recipe-images/Puttanesca-fd5810c.webp' },
]

const recipe = recipes[0]
</script>

<template>
  <Recipe :recipe="recipe" />
</template>
