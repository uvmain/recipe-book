<script setup lang="ts">
import type { Recipe } from '../types/recipes'
import { ref } from 'vue'
import RecipeForm from '../components/RecipeEditor.vue'

import { backendFetchRequest } from '../composables/fetchFromBackend'
import { base64ToBlob } from '../composables/imaging'

const router = useRouter()
const newRecipe = ref<Recipe>({
  slug: '',
  dateCreated: '',
  name: '',
  author: '',
  source: '',
  course: '',
  country: '',
  vegetarian: false,
  prepTime: '',
  cookingTime: '',
  calories: '',
  servings: '',
  ingredients: '',
  instructions: '',
  imageFilename: '',
  imageWidth: 0,
  imageHeight: 0,
})

async function handleSave(recipe: Recipe, imageBase64?: string) {
  const imageFilename = `${recipe.slug}.webp`
  if (imageBase64 && imageBase64.length > 0) {
    recipe.imageFilename = imageFilename
    await postImage(imageBase64, imageFilename)
  }

  const response = await backendFetchRequest('recipes', {
    body: JSON.stringify(recipe),
    method: 'POST',
  })

  await response

  router.push(`/recipe/${recipe.slug}`)
}

async function postImage(imageBase64: string, imageFilename: string) {
  const imageBlob = await base64ToBlob(imageBase64)
  const formData = new FormData()
  formData.append('file', imageBlob)
  formData.append('filename', imageFilename)
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

function handleCancel() {
  router.push('/')
}
</script>

<template>
  <RecipeForm :can-delete="false" :recipe="newRecipe" @cancel="handleCancel" @save="handleSave" />
</template>
