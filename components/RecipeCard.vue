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
  <NuxtLink :to="getRouterLink()" class="text-white border-1 border-zinc-200 border-solid px-4 bg-blue-gray-600 rounded no-underline overflow-auto pb-4">
    <p class="text-xl font-bold mb-2 min-h-3.5rem align-middle">
      {{ recipe.name }}
    </p>
    <NuxtImg :src="recipe.image" :alt="recipe.name" class="rounded mb-2 w-full h-32 object-cover" />
    <div>
      <div class="flex items-center gap-3 mx-2">
        <hr class="opacity-30 grow bg-gray-600">
        <RecipeIcons :recipe="recipe" class="shrink" colour="text-white" />
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
