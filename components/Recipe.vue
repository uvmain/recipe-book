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
    <RecipeIcons :recipe="recipe" class="mb-4 flex justify-center flex-auto" />
    <div class="h-0.5 mb-4 mx-auto bg-gradient-to-r via-zinc-500 from-primarybg-300 to-primarybg-300" />
    <div class="justify-center grid auto-cols-auto gap-4 md:gap-8 lg:grid-flow-col">
      <div class="">
        <div class="mb-4 flex bg-blue-gray-500 items-baseline rounded-md justify-between p-4 pt-0">
          <div>
            <p v-if="recipe.author">
              <strong>Author:</strong>
              {{ recipe.author }}
            </p>
            <strong>Source: </strong>
            <component :is="sourceTag" v-if="recipe.source" :href="recipe.source" target="_blank" class="text-white break-all underline-none">
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
