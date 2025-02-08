<script setup lang="ts">
import { calculateCaloriesPerServing } from '../composables/caloriesParse'

const props = defineProps({
  recipeCard: { type: Object, required: true },
})

const router = useRouter()

function navigateToRecipe() {
  router.push(`/recipe/${props.recipeCard.slug}`)
}

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
  <div class="tracking-wider text-center border text rounded no-underline block overflow-hidden shadow-md hover:shadow-xl transition-shadow duration-100 dark:hover:shadow-gray-600 hover:cursor-pointer" @click="navigateToRecipe()">
    <div id="card-header" class="flex flex-col h-60">
      <img :src="imageAddress" :alt="recipeCard.name" loading="lazy" :width="recipeCard.imageWidth" :height="recipeCard.imageHeight" class="w-full object-cover h-full">
    </div>
    <div>
      <h2 class="text-xl font-bold titleText mx-1">
        {{ recipeCard.name }}
      </h2>
      <div class="flex items-center gap-3 mb-4 align-middle">
        <hr class="flex-grow border-gray-300">
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
  </div>
</template>
