<script setup lang="ts">
import { calculateCaloriesPerServing } from '../composables/caloriesParse'

const props = defineProps({
  recipeCard: { type: Object, required: true },
})

const linkTarget = computed(() => {
  return `/recipe/${props.recipeCard.slug}`
})

const caloriesPerServing = computed(() => {
  return calculateCaloriesPerServing(props.recipeCard.calories, props.recipeCard.servings)
})

const parsedSource = computed(() => {
  if (props.recipeCard.source.startsWith('http')) {
    return new URL(props.recipeCard.source).host
  }
  else {
    return props.recipeCard.source
  }
})

const showAuthor = computed(() => {
  return props.recipeCard.author && `${props.recipeCard.author}`.toLowerCase() !== 'unknown'
})

const showSource = computed(() => {
  return props.recipeCard.source && `${props.recipeCard.source}`.toLowerCase() !== 'unknown'
})

const imageAddress = computed(() => {
  return props.recipeCard.imageFilename ? `/api/images/${props.recipeCard.imageFilename}` : '/default.webp'
})
</script>

<template>
  <a
    class="text-center border text rounded block shadow-md tracking-wider no-underline overflow-hidden hover:shadow-xl transition-shadow duration-100 dark:hover:shadow-gray-600 hover:cursor-pointer"
    :href="linkTarget"
  >
    <div id="card-header" class="flex flex-col h-60">
      <img :src="imageAddress" :alt="recipeCard.name" loading="lazy" :width="recipeCard.imageWidth" :height="recipeCard.imageHeight" class="w-full object-cover h-full">
    </div>
    <div>
      <h2 class="text-xl font-bold mx-1 titleText">
        {{ recipeCard.name }}
      </h2>
      <div class="flex items-center gap-3 mb-4 align-middle">
        <hr class="border-gray-300 flex-grow">
        <RecipeIcons :recipe="recipeCard" colour="text" />
        <hr class="flex-grow border-gray-300">
      </div>
      <div class="mb-0">
        <p v-if="showAuthor" class="text-sm recipeCardText">
          <strong>Author:</strong> {{ recipeCard.author }}
        </p>
        <p v-if="showSource" class="text-sm recipeCardText">
          <strong>Source:</strong> {{ parsedSource }}
        </p>
      </div>
      <div class="text-sm recipeCardText">
        <p v-if="recipeCard.prepTime">
          <strong>Prep time:</strong> {{ recipeCard.prepTime }}
        </p>
        <p v-if="recipeCard.cookingTime">
          <strong>Cooking time:</strong> {{ recipeCard.cookingTime }}
        </p>
        <p v-if="caloriesPerServing">
          <strong>Calories:</strong> {{ caloriesPerServing }}
        </p>
      </div>
    </div>
  </a>
</template>
