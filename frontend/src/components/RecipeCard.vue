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

const imageUrl = computed(() => {
  return `/api/images/${props.recipeCard.imageFilename}`
})

function onImageError(event: Event) {
  const target = event.target as HTMLImageElement
  target.onerror = null
  target.src = '/default.webp'
}
</script>

<template>
  <a
    class="text-center border text rounded block tracking-wider no-underline overflow-hidden hover:cursor-pointer transition-transform duration-200 ease-in-out hover:scale-102"
    :href="linkTarget"
  >
    <div id="card-header" class="flex flex-col h-60">
      <img :src="imageUrl" :alt="recipeCard.name" loading="lazy" :width="recipeCard.imageWidth" :height="recipeCard.imageHeight" class="w-full object-cover h-full" @error="onImageError">
    </div>
    <div>
      <h2 class="text-xl font-bold mx-1 titleText">
        {{ recipeCard.name }}
      </h2>
      <div class="flex items-center gap-3 mb-4 align-middle">
        <hr class="border flex-grow">
        <RecipeIcons :recipe="recipeCard" colour="text" />
        <hr class="border flex-grow">
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
