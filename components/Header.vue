<script setup lang="ts">
const router = useRouter()

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

const countOfRecipes = ref(0)

async function getRecipeCount() {
  const response = await $fetch<{ count: number }>('/api/recipes/count')
  countOfRecipes.value = response.count
}

async function navToRandomRecipe() {
  const randomIndex = Math.floor(Math.random() * Number(countOfRecipes.value))
  
  const response = await $fetch(`/api/recipes?offset=${randomIndex}&limit=1`) as { data : Recipe[]}
  const randomRecipe = response.data[0]

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
  <div class="bg-gray-400 w-full">
    <header class="flex flex-row gap-1 md:gap-4 justify-center items-center p-6">
      <button
        type="button"
        class="flex justify-center items-center h-12 text-zinc-600 text-xl hover:text-white hover:bg-zinc-500 rounded-lg text-center font-medium md:text-3xl px-5 border-1 border-solid border-zinc-500 bg-blue-gray-100"
        @click="navToHome"
      >
        <Icon name="carbon:home" class="w-6 h-6" />
      </button>
      <button
        type="button"
        class="flex justify-center items-center h-12 text-zinc-600 text-xl hover:text-white hover:bg-zinc-500 rounded-lg text-center font-medium md:text-3xl px-5 border-1 border-solid border-zinc-500 bg-blue-gray-100"
        @click="navToRandomRecipe"
      >
        <Icon name="carbon:shuffle" class="w-6 h-6" />
      </button>
      <SearchBar class="h-12 text-zinc-600 text-xl bg-blue-gray-100 hover:bg-zinc-300 focus:bg-zinc-200 rounded-lg text-center font-medium md:text-2xl px-5 border-1 border-solid border-zinc-500 outline-none w-1/2 md:w-1/5" :recipe-count="countOfRecipes ?? 0" />
    </header>
  </div>
</template>


