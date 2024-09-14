<script setup lang="ts">
import { useDebounceFn, useIntersectionObserver } from '@vueuse/core'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const allRecipes = ref<Recipe[]>([])
const input = useState<string>('searchInput')
const loading = ref(false)
const page = ref(1)
const hasMore = ref(true)
const observerTarget = ref(null)
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
    useIntersectionObserver(
      observerTarget,
      ([{ isIntersecting }]) => {
        if (isIntersecting) {
          loadData()
        }
      },
    )
  }
}

const search = useDebounceFn(async () => {
  allRecipes.value = []
  page.value = 1
  hasMore.value = true
  await loadData()
  if (window.innerWidth < 800) {
    isSidebarOpen.value = false
  }
}, 700)

watch(input, () => {
  search()
})

watch(useState('selectedCourses'), () => {
  search()
})

onMounted(async () => {
  await loadData()
})
</script>

<template>
  <div class="grid gap-6 md:gap-10 grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 p-2 md:p-4 lg:p-6">
    <RecipeCard v-for="recipe in allRecipes" :key="recipe.name" :recipe="recipe" class="flex-1" />
    <div ref="observerTarget" />
  </div>
</template>