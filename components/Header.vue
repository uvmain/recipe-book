<script setup lang="ts">
const router = useRouter()

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

const { data: countOfRecipes } = await useAsyncData('recipeCount', () =>
  queryContent('/recipes').count())

async function navToRandomRecipe() {
  const randomIndex = Math.floor(Math.random() * Number(countOfRecipes.value))
  const randomSlugContent = await queryContent('/recipes').skip(randomIndex).limit(1).only('slug').find()
  if (currentPath.value === `/recipe/${randomSlugContent[0].slug}`) {
    navToRandomRecipe()
  }
  router.push(`/recipe/${randomSlugContent[0].slug}`)
}

const targetLatest: string = '/'
const targetAdd: string = '/add-recipe'
</script>

<template>
  <div class="flex w-full justify-center bg-gray-100">
    <header class="flex justify-center flex-wrap m-4">
      <div>
        <button
          type="button"
          class="text-white text-xl bg-blue-gray-500 rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 me-2 mb-2"
          :class="{ 'opacity-50': currentPath !== targetLatest }"
          @click="$router.push(targetLatest)"
        >
          <Icon name="carbon:home" />
        </button>
        <button
          type="button"
          class="text-white text-xl bg-blue-gray-500 rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 me-2 mb-2"
          :class="{ 'opacity-50': !currentPath.startsWith('/recipe/') }"
          @click="navToRandomRecipe"
        >
          <Icon name="carbon:shuffle" />
        </button>
        <button
          type="button"
          class="text-white text-xl bg-blue-gray-500 rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 mb-2 md:me-2"
          :class="{ 'opacity-50': currentPath !== targetAdd }"
          @click="$router.push(targetAdd)"
        >
          <Icon name="carbon:add-alt" />
        </button>
      </div>
      <SearchBar class="" />
    </header>
  </div>
</template>
