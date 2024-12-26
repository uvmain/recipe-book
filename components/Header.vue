<script setup lang="ts">
import { useDark, useToggle } from '@vueuse/core'

const isDark = useDark()
const toggleDark = useToggle(isDark)
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
  <div class="background w-full h-18 lg:h-20">
    <header class="flex flex-row gap-1 lg:gap-4 justify-center items-center p-4 lg:p-6">
      <button
        type="button"
        class="headerButton"
        @click="navToHome()"
      >
        <Icon name="lucide:home" class="headerButtonIcon" />
      </button>
      <button
        type="button"
        class="headerButton"
        @click="navToRandomRecipe()"
      >
        <Icon name="lucide:shuffle" class="headerButtonIcon" />
      </button>
      <SearchBar class="headerButton headerSearch" :recipe-count="countOfRecipes ?? 0" />
      <button
        v-if="showFiltersButton"
        type="button"
        class="headerButton"
        @click="toggleFilters"
      >
        <Icon name="lucide:filter" class="headerButtonIcon" />
      </button>
      <button
        type="button"
        class="headerButton"
        @click="toggleDark()"
      >
        <Icon v-if="isDark" name="lucide:sun" class="headerButtonIcon" />
        <Icon v-else name="lucide:sun-moon" class="headerButtonIcon" />
      </button>
    </header>
  </div>
</template>

