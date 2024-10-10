<script setup lang="ts">
import { useDebounceFn, useIntersectionObserver } from '@vueuse/core'

useHead({
  titleTemplate: 'RecipeBook: Latest',
})

const allRecipes = ref<Recipe[]>([])
const input = useState<string>('searchInput')
const loading = ref(false)
const page = ref(1)
const hasMore = ref(true)
const observerTarget = ref(null)
const limit = ref(10)

const offset = computed(() => {
  return (page.value * 10) - 10
})

async function loadData() {
  if (loading.value || !hasMore.value) return
  loading.value = true

  let url = input.value ? `/api/recipes?filter=${input.value}&limit=${limit.value}&offset=${offset.value}` : `/api/recipes?limit=${limit.value}&offset=${offset.value}`
  
  if (useState<string[]>('selectedCourses').value?.length) {
    url = `${url}&courses=${useState('selectedCourses').value}`
  }

  if (useState<boolean>('selectedVegetarian').value == true) {
    url = `${url}&vegetarian=true`
  }

  if (useState<string>('selectedCountry').value?.length > 0) {
    url = `${url}&country=${useState('selectedCountry').value}`
  }

  // useState('selectedCountry').value
  
  try {
    const response = await $fetch<RecipesApiResponse>(url)
    .catch((error) => {
      console.error(`Failed to fetch data: ${JSON.stringify(error.data)}`)
    })

    if (response) {
      if (response.data.length > 0) {
        allRecipes.value.push(...response.data.filter((recipe) => {
          const calories = useState<number>('selectedCalories').value || 1000
          return calories == 1000 || (parseInt(recipe.calories) / parseInt(recipe.servings)) < calories
        }))
        page.value++
      }
      else {
        hasMore.value = false
      }
    }
  }
  catch (error) {
    console.error('Failed to fetch recipes:', error)
  }
  finally {
    loading.value = false
    console.log(url)
    useIntersectionObserver(
      observerTarget,
      ([{ isIntersecting }]) => {
        if (isIntersecting) {
          loadData()
        }
      },
    )
  }
}

const search = useDebounceFn(async () => {
  allRecipes.value = []
  page.value = 1
  hasMore.value = true
  await loadData()
  },
  700
)

watch(input, () => {
  search()
})

watch(useState('selectedCourses'), () => {
  search()
})

watch(useState('selectedVegetarian'), () => {
  search()
})

watch(useState('selectedCalories'), () => {
  search()
})

watch(useState('selectedCountry'), () => {
  search()
})

onMounted(async () => {
  await loadData()
})
</script>

<template>
  <div class="grid gap-4 md:gap-6 grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 p-4 lg:p-6 grid-rows-2">
    <RecipeCard v-for="recipe in allRecipes" :key="recipe.name" :recipe="recipe" class="flex-1" />
    <div ref="observerTarget" />
  </div>
</template>