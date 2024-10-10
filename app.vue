<script setup>
const router = useRouter()

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

const showFilters = useState('showFilters') || false

const showFiltersComponent = computed(() => {
  return currentPath.value === '/' && showFilters.value
})

useHead({
  htmlAttrs: {
    lang: 'en',
  },
  charset: 'utf-8',
  title: 'RecipeBook',
})

</script>

<template>
  <div id="app" class="bg-blue-gray-100 min-h-screen flex flex-col">
    <NuxtLoadingIndicator />
    <Header />
    <Filters v-if="showFiltersComponent" />
    <NuxtPage />
  </div>
</template>

<style>
html, body, #__nuxt {
  margin: 0;
  padding: 0;
  border: 0;
  font-family: 'Poppins', sans-serif;
}
</style>