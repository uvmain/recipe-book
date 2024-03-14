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
  await navigateTo(`/recipe/${randomSlugContent[0].slug}`)
}
</script>

<template>
  <div class="flex w-full text-white justify-center bg-gray-200">
    <header class="justify-center gap-2 grid grid-cols-6 m-4">
      <button
        type="button"
        class="header-button"
        :class="{ 'header-button-selected': currentPath === '/' }"
        @click="navigateTo('/')"
      >
        <Icon name="carbon:home" />
      </button>
      <button
        type="button"
        class="header-button"
        :class="{ 'header-button-selected': currentPath.startsWith('/recipe/') }"
        @click="navToRandomRecipe"
      >
        <Icon name="carbon:shuffle" />
      </button>
      <button
        type="button"
        class="header-button"
        :class="{ 'header-button-selected': currentPath === '/add-recipe' }"
        @click="navigateTo('/add-recipe')"
      >
        <Icon name="carbon:document-add" />
      </button>
      <SearchBar class="text-white text-xl rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 focus:bg-blue-gray-400 col-span-3 bg-white" />
    </header>
  </div>
</template>
