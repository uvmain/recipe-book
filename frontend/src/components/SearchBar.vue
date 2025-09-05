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
  <div class="grid items-center justify-items-end">
    <input
      id="search-input"
      v-model="inputText"
      :placeholder="placeHolder"
      type="text"
      class="headerButton headerSearch col-span-full row-span-full z-1 w-auto"
    >
    <icon-lucide-eraser
      v-if="inputText.length > 0"
      class="headerButtonIcon col-span-full row-span-full z-2 px-3"
      @click="inputText = ''"
    />
  </div>
</template>
