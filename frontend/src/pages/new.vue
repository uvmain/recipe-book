<script setup lang="ts">
import type { Recipe } from '../types/recipes'
import { useHead } from '@vueuse/head'
import Dropdown from '../components/Dropdown.vue'
import ImageSelector from '../components/ImageSelector.vue'
import { checkIfLoggedIn } from '../composables/auth'
import { countries } from '../composables/countries'
import { backendFetchRequest } from '../composables/fetchFromBackend'
import { base64ToBlob } from '../composables/imaging'
import { slugify } from '../composables/slugify'
import { allowedCourses } from '../types/recipes'

const router = useRouter()
const resizedImage = ref<string>('')
const recipe = ref<Recipe>({
  slug: '',
  date_created: '',
  name: '',
  author: '',
  source: '',
  course: '',
  country: '',
  vegetarian: false,
  prep_time: '',
  cooking_time: '',
  calories: '',
  servings: '',
  ingredients: '',
  instructions: '',
  imageFilename: '',
  imageWidth: 0,
  imageHeight: 0,
} as Recipe)

async function cancel() {
  router.push('/')
}

async function save() {
  await postImage()
  recipe.value.slug = slugify(recipe.value.slug)
  recipe.value.imageFilename = `${recipe.value.slug}.webp`

  const response = await backendFetchRequest('recipes', {
    body: JSON.stringify(recipe.value),
    method: 'POST',
  })
  await response.body
  await router.push(`/recipe/${recipe.value.slug}`)
}

async function postImage() {
  const imageBlob = await base64ToBlob(resizedImage.value)
  const formData = new FormData()
  formData.append('file', imageBlob)
  formData.append('filename', `${recipe.value.slug}.webp`)
  try {
    const response = await backendFetchRequest('images', {
      body: formData,
      method: 'POST',
    })

    if (!response.ok) {
      const errorText = await response.body
      console.error(`Upload failed: ${errorText}`)
    }
  }
  catch (error) {
    console.error(`Upload failed: ${error}`)
  }
}

watch(
  () => recipe.value.name,
  (newName) => {
    recipe.value.slug = slugify(newName)
  },
)

useHead({
  titleTemplate: 'RecipeBook: New',
})

onBeforeMount(async () => {
  const loggedIn = await checkIfLoggedIn()
  if (!loggedIn) {
    router.push('/')
  }
})
</script>

<template>
  <div class="background text min-h-screen min-w-1/3 flex flex-col gap-4 mx-auto lg:max-w-18/19">
    <TextInput v-model="recipe.slug" type="text" label="Slug" />
    <TextInput v-model="recipe.name" type="text" label="Name" />
    <TextInput v-model="recipe.author" type="text" label="Author" />
    <TextInput v-model="recipe.source" type="text" label="Source" />
    <Dropdown v-model="recipe.course" type="select" label="Course" :options="allowedCourses" placeholder="Select a course" />
    <Dropdown v-model="recipe.country" type="select" label="Country" :options="countries" placeholder="Select a country" />
    <Checkbox v-model="recipe.vegetarian" label="Vegetarian?" name="vegetarian" />
    <TextInput v-model="recipe.prep_time" type="text" label="PrepTime" />
    <TextInput v-model="recipe.cooking_time" type="text" label="CookingTime" />
    <TextInput v-model="recipe.calories" type="text" label="Calories" />
    <TextInput v-model="recipe.servings" type="text" label="Servings" />
    <TextInput v-model="recipe.ingredients" type="textarea" label="Ingredients" />
    <TextInput v-model="recipe.instructions" type="textarea" label="Instructions" />
    <ImageSelector v-model:resized-image="resizedImage" v-model:resized-image-width="recipe.imageWidth" v-model:resized-image-height="recipe.imageHeight" label="Image" />

    <div class="flex justify-center gap-4 mt-4">
      <button aria-label="cancel" class="w-full headerButton" @click="cancel()">
        Cancel
      </button>
      <button aria-label="save" class="headerButton w-full" @click="save()">
        Save
      </button>
    </div>
  </div>
</template>
