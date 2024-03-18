<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

const sourceTag = computed(() => {
  return `${props.recipe.source}`.startsWith('http') ? 'a' : 'span'
})

const caloriesPerServing = computed(() => {
  if (Number.isInteger(Number(props.recipe.servings)) && Number(props.recipe.servings) > 1 && Number.isInteger(Number(props.recipe.calories)))
    return `${Math.floor(props.recipe.calories / props.recipe.servings)} per serving (${props.recipe.calories}/${props.recipe.servings})`
  else if (Number.isInteger(Number(props.recipe.calories)))
    return `${props.recipe.calories} total`
  else return null
})
</script>

<template>
  <div class="w-full">
    <h2 class="text-center font-bold mb-4 text-4xl text-zinc-600">
      {{ recipe.name }}
    </h2>
    <div class="mb-4 flex items-center gap-3 md:mb-6 mx-auto md:w-4/5">
      <div class="w-full h-0.5 bg-gradient-to-r to-zinc-500 from-primarybg-300" />
      <RecipeIcons :recipe="recipe" />
      <div class="w-full h-0.5 to-zinc-500 from-primarybg-300 bg-gradient-to-l" />
    </div>

    <div class="justify-center flex flex-col md:flex-row gap-4">
      <div class="min-w-1/3 lg:max-w-1/2">
        <div class="mb-4 flex flex-col md:flex-row gap-4 p-4 bg-blue-gray-500 rounded-md items-baseline justify-between pt-0">
          <div class="text-center mx-auto md:text-left md:mx-0 md:max-w-3/4">
            <p v-if="recipe.author">
              <strong>Author:</strong>
              {{ recipe.author }}
            </p>
            <strong v-if="recipe.source">Source: </strong>
            <component :is="sourceTag" v-if="recipe.source" :href="recipe.source" target="_blank" class="text-white underline-none break-all">
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
        <img :src="recipe.image" :alt="recipe.name" class="w-full rounded-lg shadow-md h-auto md:mb-4">
      </div>
      <div class="lg:max-w-5/9">
        <div class="grid gap-4">
          <MarkdownBlock v-if="recipe.ingredients" label="Ingredients" :markdown-string="recipe.ingredients" props-class="bg-blue-gray-600" class="min-w-1/2" />
          <MarkdownBlock v-if="recipe.instructions" label="Instructions" :markdown-string="recipe.instructions" props-class="bg-gray-600" class="min-w-1/2" />
        </div>
      </div>
    </div>
  </div>
</template>
