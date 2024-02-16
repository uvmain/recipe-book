<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

const courseIcon = computed(() => {
  if (props.recipe.course === 'mains')
    return 'icon-park-outline:cook'
  if (props.recipe.course === 'soups')
    return 'lucide:soup'
  if (props.recipe.course === 'desserts')
    return 'ep:dessert'
  if (props.recipe.course === 'cocktails')
    return 'la:cocktail'
  if (props.recipe.course === 'sides')
    return 'mingcute:fries-line'
  if (props.recipe.course === 'sauces')
    return 'icon-park-outline:bottle-two'
  else return 'icon-park-outline:cook'
})

const flagIcon = computed(() => {
  if (props.recipe.country) {
    const countryFlag = countryFlags.find(country => country.name === props.recipe.country)
    return countryFlag?.icon || 'fxemoji:europeafricaglobe'
  }
  else
    return 'fxemoji:europeafricaglobe'
})
</script>

<template>
  <div class="flex flex-row gap-2">
    <div id="vegetarianIcon" class="flex relative group/veg">
      <Icon v-if="recipe.vegetarian" name="lucide:vegan" class="text-green-300" />
      <span class="text-sm rounded-md transition-opacity bg-gray-800 px-1 text-gray-100 absolute left-1/2 -translate-x-1/2 translate-y-full opacity-0 -ml-10 group-hover/veg:opacity-100">
        Vegetarian
      </span>
    </div>
    <div id="countryIcon" class="flex group/country relative">
      <Icon :name="flagIcon" />
      <span v-if="recipe.country" class="text-sm rounded-md transition-opacity bg-gray-800 px-1 text-gray-100 absolute left-1/2 -translate-x-1/2 translate-y-full opacity-0 -ml-10 group-hover/country:opacity-100">
        {{ recipe.country }}
      </span>
    </div>
    <div id="courseIcon" class="flex group/course relative">
      <Icon :name="courseIcon" />
      <span class="text-sm rounded-md transition-opacity bg-gray-800 px-1 text-gray-100 absolute left-1/2 -translate-x-1/2 translate-y-full opacity-0 -ml-10 group-hover/course:opacity-100">
        {{ recipe.course }}
      </span>
    </div>
  </div>
</template>
