<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

function getRouterLink() {
  return `/recipe/${props.recipe.slug}`
}
</script>

<template>
  <NuxtLink :to="getRouterLink()" class="p-4 border rounded shadow-md bg-blue-gray-600 no-underline text-white overflow-auto">
    <div class="flex items-baseline justify-between">
      <h3 class="text-xl font-bold mb-2">
        {{ recipe.name }}
      </h3>
      <div class="group flex relative">
        <Icon v-if="recipe.vegetarian" name="lucide:vegan" class="text-green-500" />
        <span class="group-hover:opacity-100 transition-opacity bg-gray-800 px-1 text-sm text-gray-100 rounded-md absolute left-1/2 -translate-x-1/2 translate-y-full opacity-0 m-4 -ml-10">
          Vegetarian
        </span>
    </div>
    </div>
    <img :src="recipe.image" :alt="recipe.name" class="w-full h-32 object-cover mb-4 rounded">

    <div>
      <p v-if="recipe.course">
        <strong>Course:</strong>
        {{ recipe.course }}
      </p>
      <p v-if="recipe.prepTime">
        <strong>Prep time:</strong>
        {{ recipe.prepTime }}
      </p>
      <p v-if="recipe.cookingTime">
        <strong>Cooking time:</strong>
        {{ recipe.cookingTime }}
      </p>
      <p v-if="recipe.author">
        <strong>Author:</strong>
        {{ recipe.author }}
      </p>
      <p v-if="recipe.source" class="truncate">
        <strong>Source:</strong>
        {{ recipe.source }}
      </p>
    </div>
  </NuxtLink>
</template>
