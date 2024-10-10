<script setup lang="ts">
const router = useRouter()

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

const countOfRecipes = ref(0)
const showFilters = useState('showFilters') || false

async function getRecipeCount() {
  const response = await $fetch<{ count: number }>('/api/recipes/count')
  countOfRecipes.value = response.count
}

async function navToRandomRecipe() {
  const randomIndex = Math.floor(Math.random() * Number(countOfRecipes.value))
  
  const response = await $fetch(`/api/recipes?offset=${randomIndex}&limit=1`) as { data : Recipe[]}
  const randomRecipe = response.data[0]

  if (currentPath.value === `/recipe/${randomRecipe.slug}`) {
    navToRandomRecipe()
  }
  resetSearch()
  await navigateTo(`/recipe/${randomRecipe.slug}`)
}

async function navToHome() {
  resetSearch()
  resetFilters()
  await navigateTo('/')
}

function toggleFilters() {
  showFilters.value = !showFilters.value
}

const showFiltersButton = computed(() => {
  return currentPath.value === '/'
})

onBeforeMount(() => {
  getRecipeCount()
})
</script>

<template>
  <div class="bg-blue-gray-100 w-full h-18 lg:h-20">
    <header class="flex flex-row gap-1 lg:gap-4 justify-center items-center p-4 lg:p-6">
      <button
        type="button"
        class="size-12 bg-white rounded-lg overflow-hidden shadow-md hover:shadow-lg transition-shadow duration-300 border-1 border-solid border-gray-200 active:bg-blue-gray-300"
        @click="navToHome()"
      >
        <Icon name="lucide:home" class="text-2xl align-middle text-center text-gray-800" />
      </button>
      <button
        type="button"
        class="size-12 bg-white rounded-lg overflow-hidden shadow-md hover:shadow-lg transition-shadow duration-300 border-1 border-solid border-gray-200 active:bg-blue-gray-300"
        @click="navToRandomRecipe()"
      >
        <Icon name="lucide:shuffle" class="text-2xl align-middle text-center text-gray-800" />
      </button>
      <SearchBar class="h-11 text-gray-800 text-xl focus:placeholder-op-0 border-gray-200 focus:bg-blue-gray-300 rounded-lg text-center shadow-md hover:shadow-lg transition-shadow duration-300 md:text-2xl px-2 border-1 border-solid outline-none w-1/2 md:w-1/4" :recipe-count="countOfRecipes ?? 0" />
      <button
        v-if="showFiltersButton"
        type="button"
        class="size-12 bg-white rounded-lg overflow-hidden shadow-md hover:shadow-lg transition-shadow duration-300 border-1 border-solid border-gray-200 active:bg-blue-gray-300"
        @click="toggleFilters"
      >
        <Icon name="lucide:filter" class="text-2xl align-middle text-center text-gray-800" />
      </button>
    </header>
  </div>
</template>


