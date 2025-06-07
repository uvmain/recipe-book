<script setup lang="ts">
defineProps({
  recipe: { type: Object, required: true },
  colour: { type: String, default: 'text' },
})
</script>

<template>
  <div class="flex gap-2 flex-row">
    <div v-if="recipe.vegetarian" id="vegetarianIcon" class="flex relative group/veg">
      <icon-lucide-leafy-green class="text-zinc-600 dark:text-zinc-400" />
      <span class="text-sm bg-zinc-800 text-zinc-100 transition-opacity py-1 absolute -translate-y-full -translate-x-1/2 opacity-0 px-2 group-hover/veg:opacity-90">
        Vegetarian
      </span>
    </div>
    <div v-if="recipe.country" class="flex relative group/country">
      <Flag :country="recipe.country" />
      <span class="text-sm transition-opacity bg-zinc-800 text-zinc-100 py-1 absolute -translate-y-full -translate-x-1/2 opacity-0 px-2 group-hover/country:opacity-90">
        {{ recipe.country }}
      </span>
    </div>
    <div id="courseIcon" class="flex relative group/course">
      <icon-lucide-croissant v-if="recipe.course === 'baking'" :class="colour" />
      <icon-lucide-lab-cocktail v-else-if="recipe.course === 'cocktails'" :class="colour" />
      <icon-lucide-dessert v-else-if="recipe.course === 'desserts'" :class="colour" />
      <icon-lucide-fish-symbol v-else-if="recipe.course === 'fish'" :class="colour" />
      <icon-lucide-cooking-pot v-else-if="recipe.course === 'mains'" :class="colour" />
      <icon-lucide-salad v-else-if="recipe.course === 'salads'" :class="colour" />
      <icon-lucide-bean v-else-if="recipe.course === 'sides'" :class="colour" />
      <icon-lucide-soup v-else-if="recipe.course in ['soups', 'stews']" :class="colour" />
      <icon-lucide-lab-garlic v-else-if="recipe.course === 'seasonings'" :class="colour" />
      <icon-lucide-cooking-pot v-else :class="colour" />
      <span class="text-sm transition-opacity bg-zinc-800 text-zinc-100 py-1 absolute -translate-y-full -translate-x-1/2 opacity-0 px-2 group-hover/course:opacity-90">
        {{ recipe.course }}
      </span>
    </div>
  </div>
</template>
