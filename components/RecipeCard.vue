<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

function getRouterLink() {
  return `/recipe/${props.recipe.slug}`
}

const caloriesPerServing = computed(() => {
  if (Number.isInteger(Number(props.recipe.servings)) && Number(props.recipe.servings) > 1 && Number.isInteger(Number(props.recipe.calories)))
    return `${Math.floor(props.recipe.calories / props.recipe.servings)} per serving (${props.recipe.calories}/${props.recipe.servings})`
  else if (Number.isInteger(Number(props.recipe.calories)))
    return `${props.recipe.calories} total`
  else return null
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
  <NuxtLink :to="getRouterLink()" class="text-white text-center bg-blue-gray-500 px-4 rounded no-underline overflow-auto pb-4">
    <div id="card-header" class="flex w-full flex-col gap-2 mb-2 h-60">
      <p class="text-xl font-bold mb-2">
        {{ recipe.name }}
      </p>
      <img :src="recipe.image" :alt="recipe.name" class="rounded flex-1 object-cover truncate">
    </div>
    <div>
      <div class="flex items-center gap-3">
        <hr class="border-solid grow border border-blue-gray-400">
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
        <br v-if="recipe.prepTime">
        <span v-if="recipe.cookingTime" class="text-sm">
          <strong>Cooking time:</strong>
          {{ recipe.cookingTime }}
        </span>
        <br v-if="recipe.cookingTime">
        <span v-if="caloriesPerServing" class="text-sm">
          <strong>Calories:</strong>
          {{ caloriesPerServing }}
        </span>
      </div>
    </div>
  </NuxtLink>
</template>
