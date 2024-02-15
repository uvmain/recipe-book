<script setup>
const router = useRouter()
const input = ref('')

const { data: recipes } = await useAsyncData('allrecipes', () =>
  queryContent('/recipes').find())

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

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
  const inputWords = input.value.toLowerCase().split(' ')

  const filteredRecipeList = slugKeywords.value.filter(recipe =>
    inputWords.every(word =>
      Object.values(recipe)[0].some(value =>
        typeof value === 'string' && value.toLowerCase().includes(word),
      ),
    ),
  )

  for (const recipe of filteredRecipeList) {
    filteredSlugs.push(Object.keys(recipe)[0])
  }

  searchedSlugs.value = filteredSlugs
  return filteredSlugs
}

onMounted(() => {
  filteredList()
})
</script>

<template>
  <div v-if="recipes" class="flex items-center grid grid-cols-2 gap-3 ml-2">
    <input
      v-model="input"
      type="text"
      placeholder="Search recipes..."
      class="text-white text-xl bg-blue-gray-500 rounded-lg font-medium md:text-3xl px-5 py-2.5 me-2 mb-2 placeholder-gray-400"
      :class="{ 'opacity-50': currentPath !== '/search' }"
      @click="$router.push('/search')"
      @input="filteredList"
    >
    <div v-if="input && !filteredList().length" class="text-red-500 font-semibold">
      <span>No results found!</span>
    </div>
  </div>
</template>
