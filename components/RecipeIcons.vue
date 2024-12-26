<script setup lang="ts">
const props = defineProps({
  recipe: { type: Object, required: true },
  colour: { type: String, default: 'text' },
})

const courseIcon = computed(() => {
  if (props.recipe.course === 'baking')
    return 'lucide:croissant'
  else if (props.recipe.course === 'cocktails')
    return 'lucide-lab:cocktail'
  else if (props.recipe.course === 'desserts')
    return 'lucide:dessert'
  else if (props.recipe.course === 'fish')
    return 'lucide:fish-symbol'
  else if (props.recipe.course === 'mains')
    return 'lucide:cooking-pot'
  else if (props.recipe.course === 'salads')
    return 'lucide:salad'
  else if (props.recipe.course === 'sauces')
    return 'lucide:droplets'
  else if (props.recipe.course === 'sides')
    return 'lucide:bean'
  else if (props.recipe.course in ['soups', 'stews'])
    return 'lucide:soup'
  else if (props.recipe.course === 'seasonings')
    return 'lucide-lab:garlic'
  else return 'lucide:cooking-pot'
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
    <div v-if="recipe.vegetarian" id="vegetarianIcon" class="flex relative group/veg">
      <Icon name="lucide:leafy-green" class="text-green-600" />
      <span class="text-sm bg-zinc-700 text-gray-100 transition-opacity py-1 absolute -translate-y-full -translate-x-1/2 opacity-0 px-2 group-hover/veg:opacity-90">
        Vegetarian
      </span>
    </div>
    <div v-if="recipe.country" id="countryIcon" class="flex relative group/country">
      <Icon v-if="flagIcon" :name="flagIcon" />
      <span class="text-sm transition-opacity bg-zinc-700 py-1 text-gray-100 absolute -translate-y-full -translate-x-1/2 opacity-0 px-2 group-hover/country:opacity-90">
        {{ recipe.country }}
      </span>
    </div>
    <div id="courseIcon" class="flex relative group/course">
      <Icon :name="courseIcon" :class="colour" />
      <span class="text-sm transition-opacity bg-zinc-700 py-1 text-gray-100 absolute -translate-y-full -translate-x-1/2 opacity-0 px-2 group-hover/course:opacity-90">
        {{ recipe.course }}
      </span>
    </div>
  </div>
</template>
