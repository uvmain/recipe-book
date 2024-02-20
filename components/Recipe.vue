<script setup lang="ts">
import markdownit from 'markdown-it'

const props = defineProps({
  recipe: { type: Object, required: true },
})

const md = markdownit()

const sourceTag = computed(() => {
  return `${props.recipe.source}`.startsWith('http') ? 'a' : 'span'
})

const caloriesPerServing = computed(() => {
  return props.recipe.calories / props.recipe.servings
})
</script>

<template>
  <div class="mr-2 md:mr">
    <h2 class="text-center font-bold text-4xl mb-4">
      {{ recipe.name }}
    </h2>
    <RecipeIcons :recipe="recipe" class="flex justify-center mb-4 flex-auto" />
    <hr class="mb-4 opacity-30">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div>
        <div class="mb-4 flex bg-blue-gray-500 items-baseline justify-between p-4 rounded-md pt-0">
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
            <span v-if="recipe.servings && recipe.calories">{{ recipe.calories }} kcal, {{ caloriesPerServing }} ea.</span>
          </div>
        </div>
        <NuxtImg placeholder="/recipe-images/default.webp" :src="recipe.image" :alt="recipe.name" class="w-full rounded-lg shadow-md h-auto md:mb-4" />
      </div>
      <div class="grid grid-cols-1 gap-4 auto-rows-min">
        <div class="rounded-lg bg-blue-gray-600 p-2 pt-1">
          <h3 class="font-bold text-xl mb-2">
            Ingredients:
          </h3>
          <div class="pl-2 md:pl-4 leading-relaxed" v-html="md.render(recipe.ingredients)" />
        </div>
        <div class="rounded-lg p-2 pt-1 bg-gray-600">
          <h3 class="text-xl font-bold">
            Instructions:
          </h3>
          <div class="pl-2 md:pl-4 leading-relaxed" v-html="md.render(recipe.instructions)" />
        </div>
      </div>
    </div>
  </div>
</template>
