<script setup lang="ts">
import { useDebounceFn, useIntersectionObserver } from '@vueuse/core'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const page = ref(0)
const limit = 5
const latestRecipes = ref<any[]>([])

const loader = ref(null)
const loaderVisbible = ref(false)
const loaderStatus = ref('idle')

const input = ref<string>('')

const whereClauses = computed(() => {
  const whereClauseArray: any[] = []
  const inputs = input.value.split(' ').filter(word => word.trim().length)
  inputs.forEach((word) => {
    if (word.toLowerCase().startsWith('vege')) {
      whereClauseArray.push({ vegetarian: true })
    }
    else {
      const wordArray = []
      if (word.trim().length) {
        wordArray.push({ slug: { $icontains: word } })
        wordArray.push({ name: { $icontains: word } })
        wordArray.push({ author: { $icontains: word } })
        wordArray.push({ source: { $icontains: word } })
        wordArray.push({ course: { $icontains: word } })
        wordArray.push({ country: { $icontains: word } })
        wordArray.push({ ingredients: { $icontains: word } })
        wordArray.push({ instructions: { $icontains: word } })
      }
      whereClauseArray.push({ $or: wordArray })
    }
  })
  return inputs.length === 1 ? whereClauseArray : inputs.length > 1 ? { $and: whereClauseArray } : {}
})

async function loadData() {
  if (loaderStatus.value === 'idle') {
    loaderStatus.value = 'loading'
    await queryContent('/recipes')
      .where(whereClauses.value as any)
      .sort({ dateAdded: -1 })
      .skip(page.value * limit)
      .limit(limit)
      .find()
      .then((results) => {
        latestRecipes.value = latestRecipes.value.concat(results)
        page.value++
        loaderStatus.value = results.length < limit ? 'no-more' : 'idle'
      })

    if (loaderVisbible.value && loaderStatus.value !== 'no-more') {
      loadData()
    }
  }
}

const search = useDebounceFn(async () => {
  latestRecipes.value = []
  page.value = 0
  loaderStatus.value = 'idle'
  await loadData()
}, 700)

onMounted(async () => {
  await loadData()

  useIntersectionObserver(
    loader,
    // eslint-disable-next-line unused-imports/no-unused-vars
    ([{ isIntersecting }], observerElement) => {
      loaderVisbible.value = isIntersecting
      if (isIntersecting)
        loadData()
    },
  )
})
</script>

<template>
  <Header v-model="input" @input-changed="search" />
  <div>
    <main v-if="latestRecipes.length" class="mx-auto w-19/20 md:w-4/5 mt-3 md:mt-8">
      <div class="grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-5 md:gap-8">
        <RecipeCard v-for="recipe in latestRecipes" :key="recipe.name" :recipe="recipe" />
      </div>
      <div ref="loader">
        <Icon v-if="latestRecipes.length && loaderStatus !== 'no-more'" name="svg-spinners:3-dots-move" class="mx-auto mt-4 w-full scale-400" />
      </div>
    </main>
  </div>
</template>
