<script setup lang="ts">
import { useDebounceFn, useInfiniteScroll } from '@vueuse/core'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const allRecipes = ref<Recipe[]>([])
const input = useState<string>('searchInput')
const loading = ref(false)
const page = ref(1)
const hasMore = ref(true)

const limit = ref(10)

const offset = computed(() => {
  return (page.value * 10) - 10
})

async function loadData() {
  if (loading.value || !hasMore.value) return
  loading.value = true

  const url = input.value ? `/api/recipes?filter=${input.value}&limit=${limit.value}&offset=${offset.value}` : `/api/recipes?limit=${limit.value}&offset=${offset.value}`
  try {
    const response = await $fetch<RecipesApiResponse>(url)
    .catch((error) => {
      console.error(`Failed to fetch data: ${JSON.stringify(error.data)}`)
    });

    if (response) {
      if (response.data.length > 0) {
        allRecipes.value.push(...response.data)
        page.value++
      }
      else {
        hasMore.value = false
      }
    }
  }
  catch (error) {
    console.error('Failed to fetch recipes:', error)
  }
  finally {
    loading.value = false
  }
}

const search = useDebounceFn(async () => {
  allRecipes.value = []
  page.value = 1
  hasMore.value = true
  await loadData()
}, 700)

watch(input, () => {
  search()
})

onMounted(async () => {
  await loadData()
})

const scrollContainer = ref<HTMLElement | null>(null)

useInfiniteScroll(scrollContainer, loadData, {
  distance: 200, // Adjust the distance to trigger loading more items
})
</script>

<template>
  <main ref="scrollContainer" class="mx-auto w-11/12 p-4 overflow-y-auto h-75vh scroll-smooth">
    <div class="grid gap-6 md:gap-10 grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5">
      <RecipeCard v-for="recipe in allRecipes" :key="recipe.name" :recipe="recipe" class="flex-1" />
    </div>
    <FloatingScrollToTop />
  </main>
</template>

<style>
body, html {
  height: 100%;
  margin: 0;
  /* overflow: hidden; /* Prevent scrolling on body */
}

::-webkit-scrollbar {
    width: 10px;
}
 
/* Track */
::-webkit-scrollbar-track {
    background: #c0c0c0;
    border-radius: 5px;
}
 
/* Handle */
::-webkit-scrollbar-thumb {
    background: #ffffff;
    border-radius: 5px;
}
 
/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
    background: #555;
}
</style>