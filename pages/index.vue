<script setup lang="ts">
import { useIntersectionObserver } from '@vueuse/core'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const page = ref(0)
const limit = 5
const latestRecipes = ref<any[]>([])

const loader = ref(null)
const loaderVisbible = ref(false)
const loaderStatus = ref('idle')

async function loadData() {
  if (loaderStatus.value === 'idle') {
    loaderStatus.value = 'loading'
    await queryContent('/recipes')
      .sort({ dateAdded: -1 })
      .skip(page.value * limit)
      .limit(limit)
      .find()
      .then((results) => {
        latestRecipes.value = latestRecipes.value.concat(results)
        page.value++
        loaderStatus.value = results.length < limit ? 'no-more' : 'idle'
      })

    if (loaderVisbible.value) {
      loadData()
    }
  }
}

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
  <div>
    <main v-if="latestRecipes.length" class="mx-auto w-19/20 md:w-4/5 mt-3 md:mt-8">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 lg:grid-cols-5 md:gap-8">
        <RecipeCard v-for="recipe in latestRecipes" :key="recipe.name" :recipe="recipe" />
      </div>
      <div ref="loader">
        <Icon v-if="latestRecipes.length && loaderStatus !== 'no-more'" name="svg-spinners:3-dots-move" class="scale-400 mt-4 mx-auto w-full" />
      </div>
    </main>
  </div>
</template>
