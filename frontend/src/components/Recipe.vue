<script setup lang="ts">
import { useMediaQuery, useWindowSize } from '@vueuse/core'
import markdownit from 'markdown-it'
import { calculateCaloriesPerServing } from '../composables/caloriesParse'
import { getTimers } from '../composables/timerParse'

const props = defineProps({
  recipe: { type: Object, required: true },
})
const { width } = useWindowSize()
const isLargeScreen = useMediaQuery('(min-width: 1201px)')

const md = markdownit()

const sourceTag = computed(() => {
  return `${props.recipe.source}`.startsWith('http') ? 'a' : 'span'
})

const caloriesPerServing = computed(() => {
  return calculateCaloriesPerServing(props.recipe.calories, props.recipe.servings)
})

const imageAddress = computed(() => {
  return props.recipe.image ? `http://localhost:3001/api/images/${props.recipe.image}` : '/default.webp'
})

const ingredientsMarkdown = computed(() => {
  return md.render(props.recipe.ingredients)
})

const instructionsMarkdown = computed(() => {
  return md.render(props.recipe.instructions)
})

const timers = computed(() => {
  const fetchedTimers = getTimers(props.recipe.instructions)
  return [...new Set(fetchedTimers)]
})

const details = ref<HTMLDivElement>()
const image = ref<HTMLDivElement>()
const ingredients = ref<HTMLDivElement>()
const instructions = ref<HTMLDivElement>()
const placement = ref()

const left = ref()
const right = ref()

const flexClass = computed(() => {
  return isLargeScreen.value ? 'flex-row' : 'flex-col'
})

const widthClass = computed(() => {
  return isLargeScreen.value ? 'max-w-1/2' : ''
})

async function setTimerPlacement() {
  if (details.value && image.value && ingredients.value && instructions.value) {
    if (width.value < 1024) {
      placement.value = 'right'
    }
    else {
      left.value = (details.value.offsetHeight + image.value.offsetHeight)
      right.value = (ingredients.value.offsetHeight + instructions.value.offsetHeight)
      placement.value = (left.value + 40) <= right.value ? 'left' : 'right'
    }
  }
  else {
    placement.value = 'right'
  }
}
</script>

<template>
  <div class="mx-2">
    <h2 class="text-center mb-4 text-4xl titleText">
      {{ recipe.name }}
    </h2>
    <div class="mb-4 flex items-center mx-auto gap-3 md:mb-6 md:w-4/5">
      <div class="w-full h-0.5 bg-gradient-to-r to-zinc-500 from-gray-400 dark:from-zinc-500 dark:to-gray-400" />
      <RecipeIcons :recipe="recipe" />
      <div class="w-full h-0.5 to-zinc-500 from-gray-400 dark:from-zinc-500 dark:to-gray-400 bg-gradient-to-l" />
    </div>
    <div class="justify-center flex gap-4" :class="flexClass">
      <div class="flex flex-col gap-4 min-w-1/3" :class="widthClass">
        <div class="flex flex-col gap-4 p-4 recipeCardBackground text justify-between border-solid md:flex-row border-1 border-gray-400 rounded">
          <div ref="details" class="text-center mx-auto md:text-left md:mx-0 md:max-w-3/4">
            <div v-if="recipe.author">
              <strong>Author:</strong>
              {{ recipe.author }}
            </div>
            <strong v-if="recipe.source">Source: </strong>
            <component :is="sourceTag" v-if="recipe.source" :href="recipe.source" target="_blank" class="text underline-none break-all">
              {{ recipe.source }}
            </component>
          </div>
          <div class="text-center mx-auto grid text-sm md:text-right opacity-80 min-w-1/5 lg:mx-0">
            <span v-if="recipe.servings">Servings: {{ recipe.servings }}</span>
            <span v-if="recipe.prep_time">Prep: {{ recipe.prep_time }}</span>
            <span v-if="recipe.cooking_time">Cook: {{ recipe.cooking_time }}</span>
            <span v-if="caloriesPerServing">Calories: {{ caloriesPerServing }}</span>
          </div>
        </div>
        <div ref="image" class="flex">
          <img :src="imageAddress" :alt="recipe.name" class="w-full border-1 border-solid border-gray-400 rounded object-cover max-h-200 md:mb-4" @load="setTimerPlacement">
        </div>
        <div v-if="timers.length && placement === 'left'" class="flex gap-2 flex-wrap md:gap-4 justify-end sm:mt-4 md:-mt-4">
          <Timer v-for="(timer, index) of timers" :key="index" :minutes="timer" />
        </div>
      </div>
      <div class="" :class="widthClass">
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
            <div v-html="instructionsMarkdown" />
          </div>
          <div v-if="timers.length && placement === 'right'" class="flex flex-wrap gap-2 lg:gap-x-4">
            <Timer v-for="(timer, index) of timers" :key="index" :minutes="timer" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
