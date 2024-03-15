<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

const sourceTag = computed(() => {
  return `${props.recipe.source}`.startsWith('http') ? 'a' : 'span'
})

const caloriesPerServing = computed(() => {
  return Math.floor(props.recipe.calories / props.recipe.servings)
})
</script>

<template>
  <div class="w-full">
    <h2 class="font-bold text-center mb-4 text-4xl text-zinc-600">
      {{ recipe.name }}
    </h2>
    <div class="flex items-center mb-4 gap-3 md:mb-6 mx-auto md:w-4/5">
      <div class="w-full h-0.5 bg-gradient-to-r to-zinc-500 from-primarybg-300" />
      <RecipeIcons :recipe="recipe" />
      <div class="w-full h-0.5 to-zinc-500 from-primarybg-300 bg-gradient-to-l" />
    </div>
    <div class="justify-center grid auto-cols-auto gap-4 md:gap-8 lg:grid-flow-col">
      <div class="">
        <div class="mb-4 flex rounded-md bg-blue-gray-500 items-baseline justify-between p-4 pt-0">
          <div>
            <p v-if="recipe.author">
              <strong>Author:</strong>
              {{ recipe.author }}
            </p>
            <strong v-if="recipe.source">Source: </strong>
            <component :is="sourceTag" v-if="recipe.source" :href="recipe.source" target="_blank" class="text-white underline-none break-all">
              {{ recipe.source }}
            </component>
          </div>
          <div class="grid text-sm text-right opacity-80 min-w-1/5">
            <span v-if="recipe.servings">{{ recipe.servings }} Servings</span>
            <span v-if="recipe.prepTime">Prep: {{ recipe.prepTime }}</span>
            <span v-if="recipe.cookingTime">Cook: {{ recipe.cookingTime }}</span>
            <span v-if="recipe.servings && recipe.calories">{{ caloriesPerServing }} kcal each, {{ recipe.calories }} total</span>
          </div>
        </div>
        <img :src="recipe.image" :alt="recipe.name" class="rounded-lg w-full shadow-md h-auto md:mb-4">
      </div>
      <div>
        <div class="grid gap-4">
          <MarkdownBlock v-if="recipe.ingredients" label="Ingredients" :markdown-string="recipe.ingredients" props-class="bg-blue-gray-600" class="min-w-1/2" />
          <MarkdownBlock v-if="recipe.instructions" label="Instructions" :markdown-string="recipe.instructions" props-class="bg-gray-600" class="min-w-1/2" />
        </div>
      </div>
    </div>
  </div>
</template>
