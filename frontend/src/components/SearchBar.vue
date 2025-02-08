<script setup lang="ts">
import { useDebounceFn, useSessionStorage } from '@vueuse/core'

const props = defineProps({
  recipeCount: { type: Number, default: 0 },
})

const router = useRouter()

const inputText = useSessionStorage<string>('searchInput', '')

async function navigateHomeIfNeeded() {
  if (inputText.value.length && router.currentRoute.value.path !== '/') {
    router.push('/')
  }
}

const debouncedNav = useDebounceFn(async () => {
  await navigateHomeIfNeeded()
}, 700)

const placeHolder = computed(() => {
  return props.recipeCount > 0 ? `Search ${props.recipeCount} recipes..` : 'Search..'
})

watch(inputText, () => {
  debouncedNav()
})
</script>

<template>
  <div class="flex shadow-md hover:shadow-lg">
    <input
      id="search-input"
      v-model="inputText"
      :placeholder="placeHolder"
      type="text"
      class="headerButton border-r-none rounded-r-none grow headerSearch"
    >
    <button
      type="button"
      aria-label="clear search"
      class="headerButton border-l-none rounded-l-none"
      @click="inputText = ''"
    >
      <icon-lucide-eraser class="headerButtonIcon" />
    </button>
  </div>
</template>
