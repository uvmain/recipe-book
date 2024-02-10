<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

function getRouterLink() {
  return `/recipe/${props.recipe.id}`
}
</script>

<template>
  <NuxtLink :to="getRouterLink()" class="p-4 border rounded shadow-md bg-blue-gray-600 no-underline text-white">
    <h3 class="text-xl font-bold mb-2">
      {{ recipe.name }}
    </h3>

    <img :src="recipe.image" :alt="recipe.name" class="w-full h-32 object-cover mb-4 rounded">

    <div class="mb-2">
      <strong>Ingredients:</strong>
      <ul class="list-disc">
        <li v-for="(ingredient, index) in recipe.ingredients" :key="index">
          {{ ingredient }}
        </li>
      </ul>
    </div>

    <div>
      <strong>Instructions:</strong>
      <p>{{ recipe.instructions }}</p>
    </div>
  </NuxtLink>
</template>
