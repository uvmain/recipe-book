<script setup lang="ts">
const router = useRouter()

const inputText = useState<string>('searchInput', () => '')

async function navigateHomeIfNeeded() {
  if (inputText.value.length && router.currentRoute.value.path !== '/') {
    navigateTo('/')
  }
}

const debouncedNav = useDebounceFn(async () => {
  await navigateHomeIfNeeded()
}, 700)

watch(inputText, () => {
  debouncedNav()
})
</script>

<template>
  <input
    id="search-input"
    v-model="inputText"
    placeholder="Search..."
  >
</template>
