<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
})

const courseIcon = computed(() => {
  if (props.recipe.course === 'baking')
    return 'mingcute:bread-line'
  else if (props.recipe.course === 'mains')
    return 'icon-park-outline:cook'
  else if (props.recipe.course === 'soups')
    return 'lucide:soup'
  else if (props.recipe.course === 'desserts')
    return 'ep:dessert'
  else if (props.recipe.course === 'cocktails')
    return 'la:cocktail'
  else if (props.recipe.course === 'sides')
    return 'mingcute:fries-line'
  else if (props.recipe.course === 'sauces')
    return 'icon-park-outline:bottle-two'
  else if (props.recipe.course === 'salads')
    return 'lucide:salad'
  else if (props.recipe.course === 'spices')
    return 'tabler:salt'
  else return 'icon-park-outline:cook'
})

const flagIcon = computed(() => {
  if (props.recipe.country) {
    const countryFlag = countryFlags.find(country => country.name === props.recipe.country)
    return countryFlag?.icon
  }
  else
    return null
})
</script>

<template>
  <div class="flex gap-2 flex-row">
    <div id="vegetarianIcon" class="flex relative group/veg">
      <Icon v-if="recipe.vegetarian" name="lucide:vegan" class="text-zinc-700" />
      <span class="text-sm rounded-md transition-opacity bg-zinc-800 px-1 text-gray-100 absolute left-1/2 -translate-x-1/2 translate-y-full opacity-0 -ml-10 group-hover/veg:opacity-100">
        Vegetarian
      </span>
    </div>
    <div id="countryIcon" class="flex relative group/country">
      <Icon v-if="flagIcon" :name="flagIcon" />
      <span v-if="recipe.country" class="text-sm rounded-md transition-opacity bg-zinc-800 px-1 text-gray-100 absolute left-1/2 -translate-x-1/2 translate-y-full opacity-0 -ml-10 group-hover/country:opacity-100">
        {{ recipe.country }}
      </span>
    </div>
    <div id="courseIcon" class="flex relative group/course">
      <Icon :name="courseIcon" class="text-zinc-700" />
      <span class="text-sm rounded-md transition-opacity bg-zinc-800 px-1 text-gray-100 absolute left-1/2 -translate-x-1/2 translate-y-full opacity-0 -ml-10 group-hover/course:opacity-100">
        {{ recipe.course }}
      </span>
    </div>
  </div>
</template>
