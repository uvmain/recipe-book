<script setup lang="ts">
defineProps({
  recipe: { type: Object, required: true },
})
</script>

<template>
  <div class="p-4 border rounded shadow-md">
    <img :src="recipe.imageUrl" :alt="recipe.title" class="w-full h-32 object-cover mb-4 rounded">

    <h3 class="text-xl font-bold mb-2">
      {{ recipe.title }}
    </h3>

    <div class="mb-2">
      <strong>Ingredients:</strong>
      <ul class="list-disc ml-4">
        <li v-for="(ingredient, index) in recipe.ingredients" :key="index">
          {{ ingredient }}
        </li>
      </ul>
    </div>

    <div>
      <strong>Instructions:</strong>
      <p>{{ recipe.instructions }}</p>
    </div>
  </div>
</template>
