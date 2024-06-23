<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

function getRouterLink() {
  return `/recipe/${props.recipe.slug}`
}

const caloriesPerServing = computed(() => {
  if (parseInt(props.recipe.servings) > 0 && parseInt(props.recipe.calories) > 1)
    return `${Math.floor(props.recipe.calories / props.recipe.servings)} per serving (${props.recipe.calories}/${props.recipe.servings})`
  else if (parseInt(props.recipe.calories) > 1)
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

const imageAddress = computed(() => {
  return props.recipe.image ? `/api/thumbnail/${props.recipe.image}` : '/default.webp'
})
</script>

<template>
  <NuxtLink :to="getRouterLink()" class="text-dark-2 text-center bg-blue-gray-100 rounded no-underline overflow-auto p-3 border-gray-400 border-1 border-solid">
    <div id="card-header" class="flex w-full flex-col h-60 gap-2 mb-2">
      <span class="text-xl font-bold">
        {{ recipe.name }}
      </span>
      <img :src="imageAddress" :alt="recipe.name" class="rounded flex-1 object-cover truncate shadow-md" :onerror="imageError = true">
    </div>
    <div>
      <div class="flex items-center gap-3">
        <hr class="border-solid grow border border-gray-400">
        <RecipeIcons :recipe="recipe" colour="text-dark" />
        <hr class="grow border border-solid border-gray-400">
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
      <div class="text-dark">
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
