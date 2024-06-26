<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

function getRouterLink() {
  return `/recipe/${props.recipe.slug}`
}

const caloriesPerServing = computed(() => {
  return calculateCaloriesPerServing(props.recipe.calories, props.recipe.servings)
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
      <img :src="imageAddress" :alt="recipe.name" class="rounded flex-1 object-cover truncate shadow-md">
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
        <span v-if="recipe.prep_time" class="text-sm">
          <strong>Prep time:</strong>
          {{ recipe.prep_time }}
        </span>
        <br v-if="recipe.prep_time">
        <span v-if="recipe.cooking_time" class="text-sm">
          <strong>Cooking time:</strong>
          {{ recipe.cooking_time }}
        </span>
        <br v-if="recipe.cooking_time">
        <span v-if="caloriesPerServing" class="text-sm">
          <strong>Calories:</strong>
          {{ caloriesPerServing }}
        </span>
      </div>
    </div>
  </NuxtLink>
</template>
