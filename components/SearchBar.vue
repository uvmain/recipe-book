<script setup lang="ts">
const props = defineProps({
  recipeCount: { type: Number, default: 0 },
})

const router = useRouter()

const inputText = useState<string>('searchInput')

async function navigateHomeIfNeeded() {
  if (inputText.value.length && router.currentRoute.value.path !== '/') {
    navigateTo('/')
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
  <input
    id="search-input"
    v-model="inputText"
    :placeholder="placeHolder"
    type="text"
  >
</template>
