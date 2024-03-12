<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

function getRouterLink() {
  return `/recipe/${props.recipe.slug}`
}

const caloriesPerServing = computed(() => {
  return Math.floor(props.recipe.calories / props.recipe.servings)
})
</script>

<template>
  <NuxtLink :to="getRouterLink()" class="text-white shadow-md border px-4 bg-blue-gray-600 rounded no-underline overflow-auto pb-4">
    <div class="flex items-baseline justify-between md:min-h-5rem">
      <h3 class="text-xl mb-2 font-bold">
        {{ recipe.name }}
      </h3>
    </div>
    <NuxtImg placeholder="/recipe-images/default.webp" :src="recipe.image" :alt="recipe.name" class="w-full rounded mb-2 h-32 object-cover" />
    <div>
      <div class="flex items-center gap-3 mx-2">
        <hr class="opacity-30 bg-gray-600 grow">
        <RecipeIcons :recipe="recipe" class="shrink" />
      </div>

      <p v-if="recipe.author">
        <strong>Author:</strong>
        {{ recipe.author }}
      </p>
      <p v-if="recipe.source" class="truncate">
        <strong>Source:</strong>
        {{ recipe.source }}
      </p>
      <span v-if="recipe.prepTime" class="text-sm">
        <strong>Prep time:</strong>
        {{ recipe.prepTime }}
      </span>
      <br>
      <span v-if="recipe.cookingTime" class="text-sm">
        <strong>Cooking time:</strong>
        {{ recipe.cookingTime }}
      </span>
      <br>
      <span v-if="recipe.calories && recipe.servings" class="text-sm">
        <strong>Calories:</strong>
        {{ caloriesPerServing }} ({{ recipe.calories }}/{{ recipe.servings }})
      </span>
    </div>
  </NuxtLink>
</template>
