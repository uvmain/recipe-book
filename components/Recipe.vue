<script setup lang="ts">
import markdownit from 'markdown-it'
import { useWindowSize } from '@vueuse/core'
const { width } = useWindowSize()

const md = markdownit()

const props = defineProps({
  recipe: { type: Object, required: true },
})

const sourceTag = computed(() => {
  return `${props.recipe.source}`.startsWith('http') ? 'a' : 'span'
})

const caloriesPerServing = computed(() => {
  return calculateCaloriesPerServing(props.recipe.calories, props.recipe.servings)
})

const imageAddress = computed(() => {
  return props.recipe.image ? `/api/images/${props.recipe.image}` : '/default.webp'
})

const ingredientsMarkdown = computed(() => {
  return md.render(props.recipe.ingredients);
})

const instructionsMarkdown = computed(() => {
  return md.render(props.recipe.instructions);
})

const timers = computed(() => {
  const fetchedTimers = getTimers(props.recipe.instructions)
  return [... new Set(fetchedTimers)]
})

const details = ref<HTMLDivElement>()
const image = ref<HTMLDivElement>()
const ingredients = ref<HTMLDivElement>()
const instructions = ref<HTMLDivElement>()
const placement = ref()

const left = ref()
const right = ref()

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
  else placement.value = 'right'
}
</script>

<template>
  <div class="w-full">
    <h2 class="text-center font-bold mb-4 text-4xl text-zinc-600">
      {{ recipe.name }}
    </h2>
    <div class="mb-4 flex items-center gap-3 md:mb-6 mx-auto md:w-4/5">
      <div class="w-full h-0.5 bg-gradient-to-r to-zinc-500 from-primary_bg-300" />
      <RecipeIcons :recipe="recipe" />
      <div class="w-full h-0.5 to-zinc-500 from-primary_bg-300 bg-gradient-to-l" />
    </div>
    <div class="justify-center flex flex-col md:flex-row gap-4">      
      <div class="min-w-1/3 lg:max-w-1/2">
        <div class="mb-4 flex flex-col md:flex-row gap-4 p-4 bg-blue-gray-100 text-dark rounded-md items-baseline justify-between pt-0 border-1 border-solid border-gray-400">
          <div ref="details" class="text-center mx-auto md:text-left md:mx-0 md:max-w-3/4">
            <p v-if="recipe.author">
              <strong>Author:</strong>
              {{ recipe.author }}
            </p>
            <strong v-if="recipe.source">Source: </strong>
            <component :is="sourceTag" v-if="recipe.source" :href="recipe.source" target="_blank" class="text-dark underline-none break-all">
              {{ recipe.source }}
            </component>
          </div>
          <div class="text-center mx-auto grid text-sm md:text-right opacity-80 min-w-1/5 lg:mx-0">
            <span v-if="recipe.servings">Servings: {{ recipe.servings }}</span>
            <span v-if="recipe.prepTime">Prep: {{ recipe.prepTime }}</span>
            <span v-if="recipe.cookingTime">Cook: {{ recipe.cookingTime }}</span>
            <span v-if="caloriesPerServing">Calories: {{ caloriesPerServing }}</span>
          </div>
        </div>
        <div ref="image">
          <img :src="imageAddress" :alt="recipe.name" class="w-full rounded-lg h-auto md:mb-4 border-1 border-solid border-gray-400" @load="setTimerPlacement">
        </div>
        <div v-if="timers.length && placement === 'left'" class="flex flex-wrap gap-x-4 justify-end sm:mt-4 md:mt-0" >
            <Timer v-for="(timer, index) of timers" :key="index" :minutes="timer" />
        </div>
      </div>
      <div ref="secondaryDiv" class="lg:max-w-5/9">
        <div class="grid gap-4">
          <!-- ingredients -->
          <div ref="ingredients" class="rounded-lg pt-1 p-2 pr-4 bg-blue-gray-200 text-dark border-1 border-solid border-gray-400">
            <h3 class="font-bold text-xl pl-2">
              Ingredients
            </h3>
            <div v-html="ingredientsMarkdown" />
          </div>
          <!-- instructions -->
          <div ref="instructions" class="rounded-lg pt-1 p-2 pr-4 bg-blue-gray-300 text-dark border-1 border-solid border-gray-400">
            <h3 class="font-bold text-xl pl-2">
              Instructions
            </h3>
            <div v-html="instructionsMarkdown" />
          </div>
          <div v-if="timers.length && placement === 'right'" class="flex flex-wrap gap-4" >
            <Timer v-for="(timer, index) of timers" :key="index" :minutes="timer" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
