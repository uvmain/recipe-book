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
  <div class="flex relative">
    <input
      id="search-input"
      v-model="inputText"
      :placeholder="placeHolder"
      type="text"
      class="headerButton grow headerSearch"
    >
    <span v-if="inputText.length > 0" class="absolute inset-y-0 right-2 flex items-center pl-3" @click="inputText = ''">
      <icon-lucide-eraser class="headerButtonIcon" />
    </span>
  </div>
</template>
