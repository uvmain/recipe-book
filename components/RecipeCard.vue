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
  <NuxtLink :to="getRouterLink()" class="no-underline text-center block rounded-lg overflow-hidden shadow-md hover:shadow-xl transition-shadow duration-300 border text">
    <div id="card-header" class="flex flex-col h-60">
      <img :src="imageAddress" :alt="recipe.name" class="h-full w-full object-cover">
    </div>
    <div class="">
      <h2 class="text-xl font-bold mx-1 titleText">
        {{ recipe.name }}
      </h2>
      <div class="flex items-center gap-3 mb-4">
        <hr class="flex-grow border-gray-300">
        <RecipeIcons :recipe="recipe" colour="text" />
        <hr class="flex-grow border-gray-300">
      </div>
      <div class="mb-0">
        <p v-if="showAuthor" class="text-sm recipeCardText">
          <strong>Author:</strong> {{ recipe.author }}
        </p>
        <p v-if="showSource" class="text-sm recipeCardText">
          <strong>Source:</strong> {{ parsedSource }}
        </p>
      </div>
      <div class="text-sm recipeCardText">
        <p v-if="recipe.prep_time">
          <strong>Prep time:</strong> {{ recipe.prep_time }}
        </p>
        <p v-if="recipe.cooking_time">
          <strong>Cooking time:</strong> {{ recipe.cooking_time }}
        </p>
        <p v-if="caloriesPerServing">
          <strong>Calories:</strong> {{ caloriesPerServing }}
        </p>
      </div>
    </div>
  </NuxtLink>
</template>


