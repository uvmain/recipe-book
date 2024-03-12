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
  <div class="mr-2 md:mr">
    <h2 class="text-center font-bold mb-4 text-4xl">
      {{ recipe.name }}
    </h2>
    <RecipeIcons :recipe="recipe" class="mb-4 flex justify-center flex-auto" />
    <hr class="mb-4 opacity-30">
    <div class="justify-center grid md:grid-flow-col auto-cols-auto gap-8">
      <div class="md:min-w-170">
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
            <span v-if="recipe.servings && recipe.calories">{{ recipe.calories }} kcal, {{ caloriesPerServing }} ea.</span>
          </div>
        </div>
        <NuxtImg :src="recipe.image" :alt="recipe.name" class="w-full rounded-lg shadow-md h-auto md:mb-4" />
      </div>
      <div>
        <div class="grid gap-4">
          <MarkdownBlock label="Ingredients" :markdown-string="recipe.ingredients" props-class="bg-blue-gray-600" class="min-w-1/2" />
          <MarkdownBlock label="Instructions" :markdown-string="recipe.instructions" props-class="bg-gray-600" class="min-w-1/2" />
        </div>
      </div>
    </div>
  </div>
</template>
