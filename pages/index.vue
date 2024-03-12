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

async function loadData() {
  if (loaderStatus.value === 'idle') {
    loaderStatus.value = 'loading'
    await queryContent('/recipes')
      .where({
        $or: [
          { slug: { $icontains: input.value } },
          { name: { $icontains: input.value } },
          { author: { $icontains: input.value } },
          { source: { $icontains: input.value } },
          { course: { $icontains: input.value } },
          { country: { $icontains: input.value } },
          { ingredients: { $icontains: input.value } },
          { instructions: { $icontains: input.value } },
        ],
      })
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
    {{ input }}
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
