<script setup lang="ts">
import markdownit from 'markdown-it'
import { calculateCaloriesPerServing } from '../composables/caloriesParse'
import { getTimers } from '../composables/timerParse'

const props = defineProps({
  recipe: { type: Object, required: true },
})

const md = markdownit()

const sourceTag = computed(() => {
  return `${props.recipe.source}`.startsWith('http') ? 'a' : 'span'
})

const caloriesPerServing = computed(() => {
  return calculateCaloriesPerServing(props.recipe.calories, props.recipe.servings)
})

const imageAddress = computed(() => {
  return props.recipe.imageFilename ? `/api/images/${props.recipe.imageFilename}` : '/default.webp'
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

const details = ref<HTMLDivElement>()
const image = ref<HTMLDivElement>()
const ingredients = ref<HTMLDivElement>()
const instructions = ref<HTMLDivElement>()

const parsedSource = computed(() => {
  if (props.recipe.source.startsWith('http')) {
    return new URL(props.recipe.source).host
  }
  else {
    return props.recipe.source
  }
})
</script>

<template>
  <div class="mx-2 flex gap-4 lg:gap-6 flex-col sm:max-w-90vw sm:max-w-60vw lg:max-w-40vw">
    <h2 class="text-center mb-4 text-4xl titleText">
      {{ recipe.name }}
    </h2>
    <div class="flex items-center mx-auto gap-4 w-4/5">
      <div class="w-full h-0.5 bg-gradient-to-r to-zinc-500 from-gray-400 dark:from-zinc-500 dark:to-gray-400" />
      <RecipeIcons :recipe="recipe" />
      <div class="w-full h-0.5 to-zinc-500 from-gray-400 dark:from-zinc-500 dark:to-gray-400 bg-gradient-to-l" />
    </div>
    <div class="flex flex-col gap-4 p-4 recipeCardBackground text justify-between border-solid md:flex-row border-1 border-gray-400 rounded">
      <div ref="details" class="text-center mx-auto md:text-left md:mx-0 md:max-w-3/4">
        <div v-if="recipe.author">
          <strong>Author:</strong>
          {{ recipe.author }}
        </div>
        <strong v-if="recipe.source">Source: </strong>
        <component :is="sourceTag" v-if="recipe.source" :href="recipe.source" target="_blank" class="text underline-none break-all text-wrap">
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
      <img :src="imageAddress" :alt="recipe.name" loading="lazy" :width="recipe.imageWidth" :height="recipe.imageHeight" class="w-full border-1 border-solid border-gray-400 rounded object-cover max-h-60vh">
      <!-- <img :src="imageAddress" :alt="recipe.name" loading="lazy" :width="recipe.imageWidth" :height="recipe.imageHeight" class="h-auto w-auto max-h-60vh max-w-95vw border-1 border-gray-400 border-solid box-border" /> -->
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
      <div v-if="timers.length" class="flex gap-2 flex-wrap md:gap-4 justify-center">
        <Timer v-for="(timer, index) of timers" :key="index" :minutes="timer" />
      </div>
    </div>
  </div>
</template>
