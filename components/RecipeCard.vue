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

const parsedSource = computed(() => {
  if (props.recipe.source.startsWith('http')) {
    return new URL(props.recipe.source).hostname
  }
  else return props.recipe.source
})

const showAuthor = computed(() => {
  return props.recipe.author && `${props.recipe.author}`.toLowerCase() !== 'unknown'
})

const showSource = computed(() => {
  return props.recipe.source && `${props.recipe.source}`.toLowerCase() !== 'unknown'
})
</script>

<template>
  <NuxtLink :to="getRouterLink()" class="text-white px-4 bg-blue-gray-500 rounded no-underline overflow-auto pb-4 text-center">
    <div id="card-header" class="h-60 flex flex-col gap-2 mb-2 w-full">
      <p class="text-xl font-bold mb-2">
        {{ recipe.name }}
      </p>
      <img :src="recipe.image" :alt="recipe.name" class="rounded flex-1 object-cover truncate">
    </div>
    <div>
      <div class="flex items-center gap-3">
        <hr class="grow border border-solid border-blue-gray-400">
        <RecipeIcons :recipe="recipe" colour="text-white" />
        <hr class="grow border border-solid border-blue-gray-400">
      </div>
      <div class="my-2 flex flex-col gap-2">
        <span v-if="showAuthor">
          <strong>Author:</strong>
          {{ recipe.author }}
        </span>
        <span v-if="showSource">
          <strong>Source:</strong>
          {{ parsedSource }}
        </span>
      </div>
      <div class="text-gray-100">
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
    </div>
  </NuxtLink>
</template>
