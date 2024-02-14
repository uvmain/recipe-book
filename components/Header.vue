<script setup lang="ts">
const router = useRouter()

const { data: countOfRecipes } = await useAsyncData('recipeCount', () =>
  queryContent('/recipes').count())

async function navToRandomRecipe() {
  const randomIndex = Math.floor(Math.random() * Number(countOfRecipes.value))
  const randomSlugContent = await queryContent('/recipes').skip(randomIndex).limit(1).only('slug').find()
  if (router.currentRoute.value.path === `/recipe/${randomSlugContent[0].slug}`) {
    navToRandomRecipe()
  }
  router.push(`/recipe/${randomSlugContent[0].slug}`)
}
</script>

<template>
  <div class="flex w-full justify-center bg-gray-100">
    <header class="flex w-full m-4 md:w-2/3 justify-around md:justify-around w-2/3">
      <button
        type="button"
        class="text-white text-xl bg-blue-gray-500 rounded-lg text-center font-medium md:text-3xl px-5 py-2.5 me-2 mb-2"
        @click="$router.push('/')"
      >
        Latest
      </button>
      <button
        type="button"
        class="text-white bg-blue-gray-500 font-medium rounded-lg text-xl md:text-3xl px-5 py-2.5 text-center me-2 mb-2"
        @click="$router.push('/all-recipes')"
      >
        All
      </button>
      <button
        type="button"
        class="text-white bg-blue-gray-500 font-medium rounded-lg text-xl md:text-3xl px-5 py-2.5 text-center me-2 mb-2"
        @click="navToRandomRecipe"
      >
        Random
      </button>
      <button
        type="button"
        class="text-white bg-blue-gray-500 font-medium rounded-lg text-xl md:text-3xl px-5 py-2.5 text-center me-2 mb-2"
        @click="$router.push('/add-recipe')"
      >
        <Icon name="lucide:plus" class="scale-140" />
      </button>
    </header>
  </div>
</template>
