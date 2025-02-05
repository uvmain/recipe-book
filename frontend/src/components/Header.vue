<script setup lang="ts">
import type { Recipe } from '../types/recipes'
import { useDark, useSessionStorage, useToggle } from '@vueuse/core'
import { backendFetchRequest } from '../composables/fetchFromBackend'
import { resetFilters, resetSearch } from '../composables/resets'

const isDark = useDark()
const toggleDark = useToggle(isDark)
const isModalOpened = ref(false)
const route = useRoute()
const router = useRouter()

const filtered = useSessionStorage<boolean>('filtered', false)
const showFilters = useSessionStorage<boolean>('showFilters', false)
const userLoginState = useSessionStorage('untrustedLoginState', false)

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
  const randomRecipe = await response.json() as Recipe

  if (currentPath.value === `/recipe/${randomRecipe.slug}`) {
    navToRandomRecipe()
  }
  resetSearch()
  await router.push(`/recipe/${randomRecipe.slug}`)
}

async function navToHome() {
  resetSearch()
  resetFilters()
  await router.push('/')
}

async function navToNew() {
  await router.push('/new')
}

async function navToEdit() {
  await router.push(`/edit/${route.params.slug}`)
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
      <div class="group">
        <button
          type="button"
          aria-label="navigate to homepage"
          class="headerButton"
          @click="navToHome()"
        >
          <icon-lucide-home class="headerButtonIcon" />
        </button>
        <span class="group-hover:opacity-90 dark:border-neutral-200 border-neutral-800 border-1 border-solid rounded px-2 py-1 text-sm text background invisible absolute group-hover:visible ml-2">
          Home
        </span>
      </div>
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
        v-if="userLoginState && route.path.startsWith('/recipe/')"
        type="button"
        aria-label="enable editing mode"
        class="headerButton"
        @click="navToEdit()"
      >
        <icon-lucide-edit class="headerButtonIcon" />
      </button>
      <button
        v-if="userLoginState"
        type="button"
        aria-label="create new recipe"
        class="headerButton"
        @click="navToNew()"
      >
        <icon-lucide-square-plus class="headerButtonIcon" />
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

<style>
@keyframes emptyanim {}

.fade:hover {
    animation:emptyanim;
}
</style>
