<script setup lang="ts">
import { useDark, useMediaQuery, useSessionStorage } from '@vueuse/core'
import { getRecipeCount } from '../composables/fetches'
import { resetFilters, resetSearch } from '../composables/resets'

const isDark = useDark()
const isModalOpened = ref(false)
const route = useRoute()
const router = useRouter()
const countOfRecipes = ref(0)

const filtered = useSessionStorage<boolean>('filtered', false)
const showFilters = useSessionStorage<boolean>('showFilters', false)
const userLoginState = useSessionStorage('untrustedLoginState', false)
const lastShuffled = useSessionStorage<string>('lastShuffled', '')
const isLargeScreen = useMediaQuery('(min-width: 1024px)')

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

function toggleDark() {
  if (!document.startViewTransition) {
    // Fallback for browsers that don't support view transitions
    isDark.value = !isDark.value
    return
  }

  document.startViewTransition(() => {
    isDark.value = !isDark.value
  })
}

async function randomiseRecipes() {
  lastShuffled.value = new Date().toISOString()
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
  const slug = (route.params as { slug?: string }).slug
  if (slug) {
    await router.push(`/edit/${slug}`)
  }
}

function toggleFilters() {
  showFilters.value = !showFilters.value
}

const showFiltersButton = computed(() => {
  return currentPath.value === '/'
})

const filterBorderClass = computed(() => {
  return filtered.value ? 'border-zinc-400 dark:border-zinc-500' : ''
})

function openModal() {
  isModalOpened.value = true
}

function closeModal() {
  isModalOpened.value = false
}

onBeforeMount(async () => {
  countOfRecipes.value = await getRecipeCount()
})
</script>

<template>
  <div class="background w-full">
    <header class="flex flex-col gap-1 justify-center p-3 lg:p-6 lg:gap-4 items-center">
      <div class="flex flex-wrap gap-1 justify-center">
        <div>
          <button
            type="button"
            aria-label="navigate to homepage"
            class="headerButton"
            @click="navToHome()"
          >
            <icon-lucide-home class="headerButtonIcon" />
          </button>
        </div>
        <div>
          <button
            type="button"
            aria-label="navigate to random recipe"
            class="headerButton"
            @click="randomiseRecipes()"
          >
            <icon-lucide-shuffle class="headerButtonIcon" />
          </button>
        </div>
        <SearchBar v-if="isLargeScreen" :recipe-count="countOfRecipes ?? 0" />
        <div v-if="showFiltersButton">
          <button
            type="button"
            aria-label="toggle filters"
            class="headerButton"
            :class="filterBorderClass"
            @click="toggleFilters"
          >
            <icon-lucide-filter class="headerButtonIcon" />
          </button>
        </div>
        <div v-if="userLoginState && route.path.startsWith('/recipe/')">
          <button
            type="button"
            aria-label="enable editing mode"
            class="headerButton"
            @click="navToEdit()"
          >
            <icon-lucide-edit class="headerButtonIcon" />
          </button>
        </div>
        <div v-if="userLoginState">
          <button
            type="button"
            aria-label="create new recipe"
            class="headerButton"
            @click="navToNew()"
          >
            <icon-lucide-square-plus class="headerButtonIcon" />
          </button>
        </div>
        <div>
          <button
            type="button"
            aria-label="toggle dark mode"
            class="headerButton"
            @click="toggleDark()"
          >
            <icon-lucide-sun v-if="isDark" class="headerButtonIcon" />
            <icon-lucide-sun-moon v-else class="headerButtonIcon" />
          </button>
        </div>
        <div>
          <button
            type="button"
            aria-label="toggle dark mode"
            class="headerButton"
            @click="openModal()"
          >
            <icon-lucide-user class="headerButtonIcon" />
          </button>
        </div>
      </div>
      <SearchBar v-if="!isLargeScreen" :recipe-count="countOfRecipes ?? 0" />
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
