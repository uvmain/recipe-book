<script setup lang="ts">
import { useDebounceFn } from '@vueuse/core'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const allRecipes = ref<Recipe[]>([])

const input = useState<string>('searchInput')

async function loadData() {
  const url = input.value ? `/api/recipes?filter=${input.value}` : '/api/recipes'
  try {
    const response = await $fetch<RecipesApiResponse>(url)
    .catch((error) => {
      console.error(`Failed to fetch data: ${JSON.stringify(error.data)}`)
    });

    allRecipes.value = response ? response.data : []
  }
  catch (error) {
    console.error('Failed to fetch recipes:', error)
  }

}

const search = useDebounceFn(async () => {
  allRecipes.value = []
  await loadData()
}, 700)

watch(input, () => {
  search()
})

onMounted(async () => {
  await loadData()
})
</script>

<template>
  <main v-if="allRecipes" class="mx-auto w-11/12 p-4">
    <div class="grid gap-6 md:gap-10 grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5">
      <RecipeCard v-for="recipe in allRecipes" :key="recipe.name" :recipe="recipe" class="flex-1" />
    </div>
    <FloatingScrollToTop />
  </main>
</template>
