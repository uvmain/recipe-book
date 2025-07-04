<script setup lang="ts">
import markdownit from 'markdown-it'
import { calculateCaloriesPerServing } from '../composables/caloriesParse'
import { getTimers } from '../composables/timerParse'

const props = defineProps({
  recipe: { type: Object, required: true },
})

const details = ref<HTMLDivElement>()
const image = ref<HTMLDivElement>()
const ingredients = ref<HTMLDivElement>()
const instructions = ref<HTMLDivElement>()

const md = markdownit()

const sourceTag = computed(() => {
  return `${props.recipe.source}`.startsWith('http') ? 'a' : 'span'
})

const caloriesPerServing = computed(() => {
  return calculateCaloriesPerServing(props.recipe.calories, props.recipe.servings)
})

const ingredientsMarkdown = computed(() => {
  return md.render(props.recipe.ingredients || '')
})

const instructionsMarkdown = computed(() => {
  return md.render(props.recipe.instructions || '')
})

const timers = computed(() => {
  const fetchedTimers = getTimers(props.recipe.instructions)
  return [...new Set(fetchedTimers)]
})

const parsedSource = computed(() => {
  if (props.recipe.source.startsWith('http')) {
    return new URL(props.recipe.source).host
  }
  else {
    return props.recipe.source
  }
})

const imageUrl = computed(() => {
  return `/api/images/${props.recipe.imageFilename}`
})

function onImageError(event: Event) {
  const target = event.target as HTMLImageElement
  target.onerror = null
  target.src = '/default.webp'
}
</script>

<template>
  <div class="flex gap-4 flex-col tracking-wider mx-2 lg:gap-6 sm:max-w-90vw sm:max-w-60vw lg:max-w-200">
    <div class="text-center titleText text-4xl font-semibold my-4">
      {{ recipe.name }}
    </div>
    <div class="flex items-center mx-auto gap-4 w-4/5">
      <div class="w-full h-0.5 bg-gradient-to-r to-zinc-500 from-gray-400 dark:from-zinc-500 dark:to-gray-400" />
      <RecipeIcons :recipe="recipe" />
      <div class="w-full h-0.5 to-zinc-500 from-gray-400 dark:from-zinc-500 dark:to-gray-400 bg-gradient-to-l" />
    </div>
    <div class="flex flex-col gap-4 p-4 recipeCardBackground text justify-between border-solid border-1 rounded md:flex-row border-gray-400">
      <div ref="details" class="text-center mx-auto md:text-left md:mx-0 md:max-w-3/4">
        <div v-if="recipe.author">
          <strong>Author:</strong>
          {{ recipe.author }}
        </div>
        <strong v-if="recipe.source">Source: </strong>
        <component :is="sourceTag" v-if="recipe.source" :href="recipe.source" target="_blank" class="text underline-none text-wrap">
          {{ parsedSource }}
        </component>
      </div>
      <div class="text-center mx-auto grid text-sm md:text-right opacity-80 min-w-1/5 lg:mx-0">
        <span v-if="recipe.servings">Servings: {{ recipe.servings }}</span>
        <span v-if="recipe.prepTime">Prep: {{ recipe.prepTime }}</span>
        <span v-if="recipe.cookingTime">Cook: {{ recipe.cookingTime }}</span>
        <span v-if="caloriesPerServing">Calories: {{ caloriesPerServing }}</span>
      </div>
    </div>
    <div ref="image" class="flex">
      <img :src="imageUrl" :alt="recipe.name" loading="eager" :width="recipe.imageWidth" :height="recipe.imageHeight" class="w-full border-1 border-solid border-gray-400 rounded object-cover max-h-60vh" @error="onImageError">
    </div>
    <div class="grid gap-4 mb-4">
      <!-- ingredients -->
      <div ref="ingredients" class="recipeCardBackground text border-1 border-solid border-gray-400 rounded pt-1 px-4">
        <h3 class="font-bold text-xl pl-2">
          Ingredients
        </h3>
        <div v-html="ingredientsMarkdown" />
      </div>
      <!-- instructions -->
      <div ref="instructions" class="pt-1 px-4 recipeCardBackground text border-1 border-solid border-gray-400 rounded">
        <h3 class="font-bold text-xl pl-2">
          Instructions
        </h3>
        <div class="text-wrap" v-html="instructionsMarkdown" />
      </div>
      <div v-if="timers.length" class="flex gap-2 justify-center flex-wrap md:gap-4">
        <Timer v-for="(timer, index) of timers" :key="index" :minutes="timer" />
      </div>
    </div>
  </div>
</template>
