<script setup lang="ts">
const router = useRouter()

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

const countOfRecipes = ref(0)

async function getRecipeCount() {
  const response = await $fetch<RecipesApiResponse>('/api/recipes')
  countOfRecipes.value =  await response.data.length
}

async function navToRandomRecipe() {
  const randomIndex = Math.floor(Math.random() * Number(countOfRecipes.value))
  
  const response = await $fetch('/api/recipes') as { data : Recipe[]}
  const randomRecipe = response.data[randomIndex]
  console.warn(randomRecipe.slug)

  if (currentPath.value === `/recipe/${randomRecipe.slug}`) {
    navToRandomRecipe()
  }
  if (useState<string>('searchInput').value.length) {
    useState<string>('searchInput').value = ''
  }
  await navigateTo(`/recipe/${randomRecipe.slug}`)
}

async function navToHome() {
  if (useState<string>('searchInput').value.length) {
    useState<string>('searchInput').value = ''
  }
  await navigateTo('/')
}

onBeforeMount(() => {
  getRecipeCount()
})
</script>

<template>
  <div class="text-white w-full bg-gradient-to-b to-primary_bg-300 from-zinc-400">
    <header class="flex flex-wrap gap-2 justify-center p-4">
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
        @click="navToRandomRecipe"
      >
        <Icon name="carbon:shuffle" />
      </button>
      <SearchBar class="search-bar" :recipe-count="countOfRecipes ?? 0" />
    </header>
  </div>
</template>
