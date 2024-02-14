<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

function getRouterLink() {
  return `/recipe/${props.recipe.slug}`
}

const caloriesPerServing = computed(() => {
  return props.recipe.calories / props.recipe.servings
})

const courseIcon = computed(() => {
  if (props.recipe.course === 'main')
    return 'icon-park-outline:cook'
  if (props.recipe.course === 'soup')
    return 'lucide:soup'
  if (props.recipe.course === 'dessert')
    return 'ep:dessert'
  if (props.recipe.course === 'cocktail')
    return 'la:cocktail'
  if (props.recipe.course === 'sides')
    return 'mingcute:fries-line'
  if (props.recipe.course === 'sauces')
    return 'icon-park-outline:bottle-two'
  else return 'icon-park-outline:cook'
})
</script>

<template>
  <NuxtLink :to="getRouterLink()" class="p-4 border rounded shadow-md text-white bg-blue-gray-600 no-underline overflow-auto">
    <div class="flex items-baseline justify-between">
      <h3 class="text-xl font-bold mb-2">
        {{ recipe.name }}
      </h3>
      <div class="flex group relative">
        <Icon v-if="recipe.vegetarian" name="lucide:vegan" class="text-green-300" />
        <span class="text-sm rounded-md group-hover:opacity-100 transition-opacity bg-gray-800 px-1 text-gray-100 absolute left-1/2 -translate-x-1/2 translate-y-full opacity-0 -ml-10">
          Vegetarian
        </span>
      </div>
    </div>
    <img :src="recipe.image" :alt="recipe.name" class="w-full rounded mb-2 h-32 object-cover" loading="lazy">
    <div>
      <div class="flex items-baseline justify-between">
        <hr class="ml-2 opacity-30 bg-gray-600 grow">
        <div class="group relative text-white text-xl items-start shrink ml-4 mr-2">
          <Icon :name="courseIcon" />
          <span class="text-sm rounded-md group-hover:opacity-100 transition-opacity bg-gray-800 px-1 text-gray-100 absolute left-1/2 -translate-x-1/2 translate-y-full opacity-0 -ml-10">
            {{ recipe.course }}
          </span>
        </div>
      </div>

      <p v-if="recipe.author">
        <strong>Author:</strong>
        {{ recipe.author }}
      </p>
      <p v-if="recipe.source" class="truncate">
        <strong>Source:</strong>
        {{ recipe.source }}
      </p>
      <span v-if="recipe.prepTime" class="text-sm">
        <strong>Prep time:</strong>
        {{ recipe.prepTime }}
      </span>
      <br>
      <span v-if="recipe.cookingTime" class="text-sm">
        <strong>Cooking time:</strong>
        {{ recipe.cookingTime }}
      </span>
      <br>
      <span v-if="recipe.calories && recipe.servings" class="text-sm">
        <strong>Calories:</strong>
        {{ caloriesPerServing }} ({{ recipe.calories }}/{{ recipe.servings }})
      </span>
    </div>
  </NuxtLink>
</template>
