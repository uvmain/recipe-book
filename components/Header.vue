<script setup lang="ts">
const props = defineProps({
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
  await navigateTo(`/recipe/${randomSlugContent[0].slug}`)
}

async function handleInput(e: Event) {
  const value = (e.target as HTMLInputElement).value
  if (value.trim().length && currentPath.value !== '/') {
    await navigateTo({
      path: '/',
      query: {
        searchInput: value,
      },
    })
  }
  emit('update:modelValue', value)
  emit('inputChanged', value)
}

onMounted(() => {
  if (currentPath.value === '/' && props.modelValue.length) {
    searchInput.value.focus()
  }
})
</script>

<template>
  <div class="flex w-full text-white justify-center bg-gray-100">
    <header class="gap-2 justify-center grid grid-cols-6 m-4">
      <button
        type="button"
        class="text-white text-xl rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 bg-blue-gray-500"
        :class="{ 'opacity-50': currentPath !== '/' }"
        @click="navigateTo('/')"
      >
        <Icon name="carbon:home" />
      </button>
      <button
        type="button"
        class="text-white text-xl bg-blue-gray-500 rounded-lg text-center font-medium md:text-3xl px-5 py-2.5"
        :class="{ 'opacity-50': !currentPath.startsWith('/recipe/') }"
        @click="navToRandomRecipe"
      >
        <Icon name="carbon:shuffle" />
      </button>
      <button
        type="button"
        class="text-white text-xl bg-blue-gray-500 rounded-lg text-center font-medium md:text-3xl px-5 py-2.5"
        :class="{ 'opacity-50': currentPath !== '/add-recipe' }"
        @click="navigateTo('/add-recipe')"
      >
        <Icon name="carbon:add-alt" />
      </button>
      <input
        id="search-input"
        ref="searchInput"
        :value="modelValue"
        placeholder="Search..."
        class="text-white text-xl rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 focus:bg-blue-gray-400 bg-blue-gray-200 col-span-3"
        @input="handleInput"
      >
    </header>
  </div>
</template>
