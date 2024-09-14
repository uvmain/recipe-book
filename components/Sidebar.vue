<script setup lang="ts">
const router = useRouter()

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

const countOfRecipes = ref(0)
const selectedCourses = ref<string[]>([])

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
  useState<string>('searchInput').value = ''
  setSidebarStatus()
  await navigateTo(`/recipe/${randomRecipe.slug}`)
}

async function navToHome() {
  useState<string>('searchInput').value = ''
  setSidebarStatus()
  await navigateTo('/')
}

function scrollToTop() {
  window.scrollTo(0, 0)
  setSidebarStatus()
}

function setSidebarStatus() {
  if (window.innerWidth > 800) {
    isSidebarOpen.value = true
  }
  else {
    isSidebarOpen.value = false
  }
}

watch(selectedCourses, () => {
  useState('selectedCourses').value = selectedCourses.value
})

onBeforeMount(() => {
  getRecipeCount()
  setSidebarStatus()
})
</script>

<template>
  <div
    v-if="isSidebarOpen"
    class="flex flex-col gap-4 fixed top-0 left-0 h-screen p-4 md:p-6 lg:p-8 bg-blue-gray-100 z-10 w-full lg:w-64"
  >
    <button
      type="button"
      class="w-auto flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200 mr-8 md:mr-0"
      @click="navToHome()"
    >
      <Icon name="carbon:home" class="size-6" />
      <span class="text-xl">Home</span>
    </button>
    <button
      type="button"
      class="w-auto flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200 mr-8 md:mr-0"
      @click="navToRandomRecipe()"
    >
      <Icon name="carbon:shuffle" class="size-6" />
      <span class="text-xl">Random</span>
    </button>
    <button
      type="button"
      class="w-auto flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200 mr-8 md:mr-0"
      @click="scrollToTop"
    >
      <Icon name="carbon:arrow-up" class="size-6" />
      <span class="text-xl">Scroll to top</span>
    </button>
    <SearchBar
      class="w-auto p-0 flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200 outline-none text-xl placeholder-dark mr-8 md:mr-0"
      :recipe-count="countOfRecipes ?? 0"
    />
    
    <hr class="w-80% my-4 border-gray-300 border-1 border-solid">
    
    <div class="grid grid-cols-2 gap-1">
      <div v-for="(course, index) in allowedCourses" :key="index" class="flex gap-1">
        <input
          v-model="selectedCourses"
          type="checkbox"
          :value="course"
        >
        <label class="font-light">
          {{ course }}
        </label>
      </div>
    </div>

    <hr class="w-80% my-4 border-gray-300 border-1 border-solid">
    
    <Footer class="bottom-0 left-0"/>
  </div>
</template>
