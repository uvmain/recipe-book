<script setup>
const input = ref('')

const { data: recipes } = await useAsyncData('allrecipes', () =>
  queryContent('/recipes').find())

function flattenRecipe(jsonObj) {
  const result = []
  function traverse(obj) {
    for (const key in obj) {
      if (typeof obj[key] === 'object' && obj[key] !== null) {
        traverse(obj[key])
      }
      else if (key === 'vegetarian' && obj[key] === true) {
        result.push('vegetarian')
      }
      else if (typeof obj[key] === 'string') {
        result.push(obj[key].replace(/\W/g, ' '))
      }
    }
  }
  traverse(jsonObj)
  return result
}

function onlyUnique(value, index, array) {
  return array.indexOf(value) === index
}

const slugKeywords = computed(() => {
  if (!recipes.value)
    return []
  const slugKeywords = []
  for (const recipe of recipes.value) {
    const flatRecipe = flattenRecipe(recipe)
    slugKeywords.push({
      [recipe.slug]: flatRecipe.join(' ').split(' ').filter(word => word.length > 2).filter(onlyUnique),
    })
  }
  return slugKeywords
})

function filteredList() {
  const filteredSlugs = []
  const filteredRecipeList = slugKeywords.value.filter(recipe =>
    Object.values(recipe)[0].some(value =>
      typeof value === 'string' && value.includes(input.value.toLowerCase()),
    ),
  )
  for (const recipe of filteredRecipeList) {
    filteredSlugs.push(Object.keys(recipe)[0])
  }
  searchedSlugs.value = filteredSlugs
  return filteredSlugs
}
</script>

<template>
  <div v-if="recipes" class="flex items-center grid grid-cols-2 gap-5 min-w-1/2">
    <input
      v-model="input"
      type="text"
      placeholder="Search recipes..."
      class="block w-full px-3 py-3 text-base font-normal text-dark bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0"
      @click="$router.push('/search')"
      @input="filteredList"
    >
    <div v-if="input && !filteredList().length" class="text-red font-semibold">
      <span>No results found!</span>
    </div>
  </div>
</template>
