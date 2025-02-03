<script setup lang="ts">
import type { Recipe } from '../types/recipes'
import { useDark, useSessionStorage, useToggle } from '@vueuse/core'
import { backendFetchRequest } from '../composables/fetchFromBackend'
import { resetFilters, resetSearch } from '../composables/resets'

const isDark = useDark()
const toggleDark = useToggle(isDark)
const isModalOpened = ref(false)
const router = useRouter()

const filtered = useSessionStorage<boolean>('filtered', false)
const showFilters = useSessionStorage<boolean>('showFilters', false)

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

const countOfRecipes = ref(0)

async function getRecipeCount() {
  const response = await backendFetchRequest('recipe-count')
  const count = await response.json() as number
  countOfRecipes.value = count
}

async function navToRandomRecipe() {
  const response = await backendFetchRequest('random-recipe')
  const randomRecipe = await response.json() as { data: Recipe }

  if (currentPath.value === `/recipe/${randomRecipe.data.slug}`) {
    navToRandomRecipe()
  }
  resetSearch()
  await router.push(`/recipe/${randomRecipe.data.slug}`)
}

async function navToHome() {
  resetSearch()
  resetFilters()
  await router.push('/')
}

function toggleFilters() {
  showFilters.value = !showFilters.value
}

const showFiltersButton = computed(() => {
  return currentPath.value === '/'
})

const filterBorderClass = computed(() => {
  return filtered.value ? 'border-green dark:border-green' : ''
})

function openModal() {
  isModalOpened.value = true
}

function closeModal() {
  isModalOpened.value = false
}

onBeforeMount(() => {
  getRecipeCount()
})
</script>

<template>
  <div class="background w-full h-18 lg:h-20">
    <header class="flex flex-row gap-1 justify-center items-center p-4 lg:p-6 lg:gap-4">
      <button
        type="button"
        aria-label="navigate to homepage"
        class="headerButton"
        @click="navToHome()"
      >
        <icon-lucide-home class="headerButtonIcon" />
      </button>
      <button
        type="button"
        aria-label="navigate to random recipe"
        class="headerButton"
        @click="navToRandomRecipe()"
      >
        <icon-lucide-shuffle class="headerButtonIcon" />
      </button>
      <SearchBar :recipe-count="countOfRecipes ?? 0" />
      <button
        v-if="showFiltersButton"
        type="button"
        aria-label="toggle filters"
        class="headerButton"
        :class="filterBorderClass"
        @click="toggleFilters"
      >
        <icon-lucide-filter class="headerButtonIcon" />
      </button>
      <button
        type="button"
        aria-label="toggle dark mode"
        class="headerButton"
        @click="toggleDark()"
      >
        <icon-lucide-sun v-if="isDark" class="headerButtonIcon" />
        <icon-lucide-sun-moon v-else class="headerButtonIcon" />
      </button>
      <button
        type="button"
        aria-label="toggle dark mode"
        class="headerButton"
        @click="openModal()"
      >
        <icon-lucide-user class="headerButtonIcon" />
      </button>
      <LoginModal :is-open="isModalOpened" @modal-close="closeModal" />
    </header>
  </div>
</template>
