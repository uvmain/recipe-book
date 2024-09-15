<script setup lang="ts">
const router = useRouter()

const currentPath = computed(() => {
  return router.currentRoute.value.path
})

const countOfRecipes = ref(0)
const selectedCourses = ref<string[]>([])
const selectedVegetarian = ref(false)
const selectedCalories = ref(1000)
const selectedCountry = ref()

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

function resetFilters() {
  selectedCourses.value = []
  selectedVegetarian.value = false
  selectedCalories.value = 1000
  selectedCountry.value = undefined
}

watch(selectedCourses, async () => {
  useState('selectedCourses').value = selectedCourses.value
  await navigateTo('/')
})

watch(selectedVegetarian, async () => {
  useState('selectedVegetarian').value = selectedVegetarian.value
  await navigateTo('/')
})

watch(selectedCalories, async () => {
  useState('selectedCalories').value = selectedCalories.value
  await navigateTo('/')
})

watch(selectedCountry, async () => {
  useState('selectedCountry').value = selectedCountry.value
  await navigateTo('/')
})

onBeforeMount(() => {
  getRecipeCount()
  setSidebarStatus()
})
</script>

<template>
  <div
    v-if="isSidebarOpen"
    class="pt-4 lg:pt-0 flex flex-col gap-4 fixed top-0 left-0 h-screen md:my-6 md:px-8 bg-blue-gray-100 z-10 w-screen lg:w-64"
  >
    <button
      type="button"
      class="mx-4 w-auto flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200 active:bg-blue-gray-300"
      @click="navToHome()"
    >
      <Icon name="carbon:home" class="size-6" />
      <span class="text-xl">Home</span>
    </button>
    <button
      type="button"
      class="mx-4 w-auto flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200 active:bg-blue-gray-300"
      @click="navToRandomRecipe()"
    >
      <Icon name="carbon:shuffle" class="size-6" />
      <span class="text-xl">Random</span>
    </button>
    <button
      type="button"
      class="mx-4 w-auto flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200 active:bg-blue-gray-300"
      @click="scrollToTop"
    >
      <Icon name="carbon:arrow-up" class="size-6" />
      <span class="text-xl">Scroll to top</span>
    </button>
    <SearchBar
      class="mx-4 w-auto p-0 flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200 outline-none text-xl placeholder-dark"
      :recipe-count="countOfRecipes ?? 0"
    />
    <button
      type="button"
      class="mx-4 w-auto flex gap-4 justify-center items-center h-12 text-center bg-white rounded-lg shadow-md hover:shadow-xl transition-shadow duration-100 border border-solid border-gray-200 active:bg-blue-gray-300"
      @click="resetFilters()"
    >
      <Icon name="carbon:reset" class="size-6" />
      <span class="text-xl">Reset filters</span>
    </button>
    
    <hr class="w-80% my-4 border-gray-300 border-1 border-solid">
    
    <div class="mx-8 grid grid-cols-2 gap-1">
      <div v-for="(course, index) in allowedCourses" :key="index" class="flex gap-1">
        <input
          :id="course"
          v-model="selectedCourses"
          type="checkbox"
          :value="course"
          class="accent-gray-600"
        >
        <label :for="course" class="font-light">
          {{ course }}
        </label>
      </div>
    </div>

    <hr class="w-80% my-2 border-gray-300 border-1 border-solid">

    <div class="mx-8 grid grid-cols-2 gap-1">
      <div class="flex flex-row gap-x-1">
        <input
          id="vegetarian"
          v-model="selectedVegetarian"
          type="checkbox"
          class="accent-gray-600"
        >
        <Icon name="lucide:vegan" class="size-6"/>
        <label for="vegetarian" class="font-light">
          vegetarian
        </label>
      </div>
    </div>

    <hr class="w-80% my-2 border-gray-300 border-1 border-solid">

    <div class="mx-8 flex flex-col gap-1">
      <label for="calories">
        Calories per serving: {{ selectedCalories == 1000 ? "Unlimited" : selectedCalories }}
      </label>
      <input
        id="calories"
        v-model="selectedCalories"
        type="range"
        min="100"
        max="1000"
        class="w-full accent-gray-600"
      >
    </div>

    <hr class="w-80% my-2 border-gray-300 border-1 border-solid">

    <div class="mx-8  flex flex-col gap-1">
      <label for="country" class="flex justify-between">
        <div class="flex flex-row gap-2">
          Country: <Icon v-if="selectedCountry" :name="countryFlags.filter((country) => country.name == selectedCountry)[0].icon" class="size-6" />
        </div>
        <button
          type="button"
          class="justify-center items-center size-8 bg-white rounded-lg border border-solid border-gray-200 active:bg-blue-gray-300 -translate-y-1"
          @click="selectedCountry = undefined"
        >
          <Icon name="carbon:reset" class="size-6 text-gray-600" />
        </button>
      </label>
      <select id="country" v-model="selectedCountry" class="bg-white rounded-lg h-6 border-solid border border-gray-200 text-sm text-gray-600">
        <option disabled value="">Please select one</option>
        <option v-for="(country, index) in countryFlags" :key=index>
          {{ country.name }}
        </option>
      </select>
    </div>

    <hr class="w-80% my-2 border-gray-300 border-1 border-solid">

    <Footer class="bottom-0 left-0"/>
  </div>
</template>
