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
  if (useState<string>('searchInput').value.length) {
    useState<string>('searchInput').value = ''
  }
  await navigateTo(`/recipe/${randomSlugContent[0].slug}`)
}

async function navToHome() {
  if (useState<string>('searchInput').value.length) {
    useState<string>('searchInput').value = ''
  }
  await navigateTo('/')
}
</script>

<template>
  <div class="text-white w-full bg-gradient-to-b from-gray-500 to-gray-300">
    <header class="grid flex justify-center gap-2 grid-cols-6 m-4">
      <button
        type="button"
        class="header-button"
        :class="{ 'header-button-selected': currentPath === '/' }"
        @click="navToHome"
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
      <SearchBar class="search-bar" />
    </header>
  </div>
</template>
