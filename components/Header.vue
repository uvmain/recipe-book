<script setup lang="ts">
defineProps({
  modelValue: { type: String, default: '' },
})
const emit = defineEmits(['update:modelValue', 'inputChanged'])

const router = useRouter()

const searchInput = ref()

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

function handleInput(e: Event) {
  const value = (e.target as HTMLInputElement).value
  emit('update:modelValue', value)
  emit('inputChanged', value)
}

function searchClicked() {
  router.push('/')
}

function resetClicked() {
  emit('update:modelValue', '')
  emit('inputChanged', '')
}
</script>

<template>
  <div class="flex w-full justify-center bg-gray-100">
    <div class="text-dark" />
    <header class="flex justify-center flex-wrap m-4">
      <div>
        <button
          type="button"
          class="text-white text-xl bg-blue-gray-500 rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 me-2 mb-2"
          :class="{ 'opacity-50': currentPath !== '/' }"
          @click="$router.push('/')"
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
          :class="{ 'opacity-50': currentPath !== '/add-recipe' }"
          @click="$router.push('/add-recipe')"
        >
          <Icon name="carbon:add-alt" />
        </button>
        <input
          id="search-input"
          ref="searchInput"
          :value="modelValue"
          placeholder="Search..."
          class="text-white text-xl rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 mb-2 md:me-2 focus:bg-blue-gray-400 bg-blue-gray-200"
          @click="searchClicked"
          @input="handleInput"
        >
        <button
          type="button"
          class="text-white bg-blue-gray-500 rounded-lg text-center md:text-3xl px-5 py-2.5 mb-2 md:me-2 opacity-50 text-base"
          @click="resetClicked"
        >
          <Icon name="carbon:skip-back" />
        </button>
      </div>
    </header>
  </div>
</template>
