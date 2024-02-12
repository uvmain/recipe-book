<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

const sourceTag = computed(() => {
  return `${props.recipe.source}`.startsWith('http') ? 'a' : 'p'
})

const caloriesPerServing = computed(() => {
  return props.recipe.calories / props.recipe.servings
})
</script>

<template>
  <div>
    <h2 class="text-3xl font-bold mb-4">
      {{ recipe.name }}
    </h2>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div>
        <div class="mb-4 bg-blue-gray-500 p-4 pt-0 rounded-md flex items-baseline justify-between">
          <div class="">
            <p v-if="recipe.author">
              <strong>Author:</strong>
              {{ recipe.author }}
            </p>
            <component v-if="recipe.source" :is="sourceTag" :href="recipe.source" target="_blank" class="break-all underline-none"><strong>Source: </strong>{{ recipe.source }}</component>
          </div>
          <div class="grid grid-rows-3 text-sm text-right">
            <span v-if="recipe.servings">{{ recipe.servings }} Servings</span>
            <span v-if="recipe.servings">Prep: {{ recipe.prepTime }}</span>
            <span v-if="recipe.servings">Cook: {{ recipe.cookingTime }}</span>
            <span v-if="recipe.servings">{{ recipe.calories }} kcal, {{ caloriesPerServing }} ea.</span>
          </div>
        </div>
        <img :src="recipe.image" :alt="recipe.name" class="w-full h-auto mb-4 rounded-lg shadow-md">
      </div>
      <div>
        <div>
          <h3 class="text-xl font-bold mb-2">
            Ingredients:
          </h3>
          <ul class="list-disc ml-2">
            <li v-for="(ingredient, index) in recipe.ingredients" :key="index" :class="{ 'opacity-0': ingredient === '', 'list-none': ingredient.startsWith(' ') }">
              {{ ingredient }}
            </li>
          </ul>
        </div>
        <div>
          <h3 class="text-xl font-bold">
            Instructions:
          </h3>
          <ul class="list-disc ml-2">
            <li v-for="(instruction, index) in recipe.instructions" :key="index" :class="{ 'opacity-0': instruction === '' }">
              {{ instruction }}
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>