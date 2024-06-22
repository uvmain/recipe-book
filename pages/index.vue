<script setup lang="ts">
import { useDebounceFn } from '@vueuse/core'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const allRecipes = ref<Recipe[]>([])

const input = useState<string>('searchInput')

// const whereClauses = computed(() => {
//   const whereClauseArray: unknown[] = []
//   const inputs = input.value.split(' ').filter(word => word.trim().length)
//   inputs.forEach((word) => {
//     if (word.toLowerCase() === 'vegetarian') {
//       whereClauseArray.push({ vegetarian: true })
//     }
//     else {
//       const wordArray = []
//       if (word.trim().length) {
//         wordArray.push({ slug: { $icontains: word } })
//         wordArray.push({ name: { $icontains: word } })
//         wordArray.push({ author: { $icontains: word } })
//         wordArray.push({ source: { $icontains: word } })
//         wordArray.push({ course: { $icontains: word } })
//         wordArray.push({ country: { $icontains: word } })
//         wordArray.push({ ingredients: { $icontains: word } })
//         wordArray.push({ instructions: { $icontains: word } })
//       }
//       whereClauseArray.push({ $or: wordArray })
//     }
//   })
//   return inputs.length === 1 ? whereClauseArray : inputs.length > 1 ? { $and: whereClauseArray } : {}
// })

async function loadData() {
  try {
    const response = await fetch('/api/recipes')

    if (!response.ok) {
      throw new Error('Failed to fetch recipes')
    }

    const data = await response.json()
    
    allRecipes.value = data
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
  <main v-if="allRecipes" class="mx-auto w-19/20">
    <div class="grid gap-3 grid-cols-1 md:grid-cols-3 lg:grid-cols-6">
      <RecipeCard v-for="recipe in allRecipes" :key="recipe.name" :recipe="recipe" class="flex-1" />
    </div>
    <FloatingScrollToTop />
  </main>
</template>
