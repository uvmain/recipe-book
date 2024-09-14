<script setup lang="ts">
const router = useRouter()
const sidebarOpen = useState('sidebarOpen')

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

const countOfRecipes = ref(0)

const links = [
  "one",
  "two",
  "three"
]

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
  await navigateTo(`/recipe/${randomRecipe.slug}`)
}

async function navToHome() {
  useState<string>('searchInput').value = ''
  await navigateTo('/')
}

function toggleSidebar() {
  if (sidebarOpen.value == false ) {
    sidebarOpen.value = true
  }
  else {
    sidebarOpen.value = false
  }
}

function scrollToTop() {
  window.scrollTo(0, 0)
}

const isSidebarOpen = computed(() => {
  return sidebarOpen.value
})

function setSidebarStatus() {
  if (window.innerWidth > 800) {
    sidebarOpen.value = true
  }
  else {
    sidebarOpen.value = false
  }
}

onBeforeMount(() => {
  getRecipeCount()
  setSidebarStatus()
})
</script>

<template>
  <div class="fixed top-0 left-0 h-screen p-4 md:p-6 lg:p-8">
    <div
      v-if="isSidebarOpen"
      class="flex flex-col gap-4"
    >
      <button
        type="button"
        class="w-full flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200"
        @click="navToHome()"
      >
        <Icon name="carbon:home" class="size-6" />
        <span class="text-xl">Home</span>
      </button>
      <button
        type="button"
        class="w-full flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200"
        @click="navToRandomRecipe()"
      >
        <Icon name="carbon:shuffle" class="size-6" />
        <span class="text-xl">Random</span>
      </button>
      <button
        type="button"
        class="w-full flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200"
        @click="scrollToTop"
      >
        <Icon name="carbon:arrow-up" class="size-6" />
        <span class="text-xl">Scroll to top</span>
      </button>
      <SearchBar
        class="w-full p-0 flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200 outline-none text-xl placeholder-dark"
        :recipe-count="countOfRecipes ?? 0"
      />
      
      <hr class="w-80% my-4 border-gray-300 border-1 border-solid">
      
      <ul>
        <li v-for="(link, index) in links" :key="index">
          {{ link }}
        </li>
      </ul>

      <hr class="w-80% my-4 border-gray-300 border-1 border-solid">
      
      <Footer class="bottom-0 left-0"/>
    </div>

    <div>
      <div class="fixed bottom-20 left-0">
      <button
        type="button"
        class="items-center h-12 text-center bg-white rounded-r-lg rounded-l-none shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-400"
        @click="toggleSidebar()"
      >
        <Icon v-if="isSidebarOpen" name="gg:chevron-double-left" class="size-6"/>
        <Icon v-if="!isSidebarOpen" name="gg:chevron-double-right" class="size-6"/>
      </button>
    </div>
    </div>
  </div>
</template>
