<script setup lang="ts">
import type { Recipe } from '../../types/recipes'
import { ref } from 'vue'
import RecipeForm from '../../components/RecipeEditor.vue'
import { backendFetchRequest } from '../../composables/fetchFromBackend'
import { base64ToBlob } from '../../composables/imaging'

const route = useRoute()
const router = useRouter()
const recipeSlug = computed(() => `${route.params.slug}`)
const recipe = ref<Recipe>({} as Recipe)
const deleting = ref(false)

async function getRecipe() {
  try {
    const response = await backendFetchRequest(`recipes/${recipeSlug.value}`)
    const jsonData = await response.json() as Recipe
    recipe.value = jsonData
  }
  catch (error) {
    console.error('Failed to fetch recipe:', error)
  }
}

async function handleSave(recipe: Recipe, imageBase64?: string) {
  const imageFilename = `${recipe.slug}.webp`
  const originalFilename = recipe.imageFilename
  if (imageBase64 && imageBase64.length > 0) {
    recipe.imageFilename = imageFilename
    console.log(imageFilename)
    await patchImage(imageBase64, originalFilename, imageFilename)
  }

  const response = await backendFetchRequest(`recipes/${recipeSlug.value}`, {
    body: JSON.stringify(recipe),
    method: 'PATCH',
  })
  await response
  router.push(`/recipe/${recipe.slug}`)
}

async function patchImage(imageBase64: string, originalFilename: string, imageFilename: string) {
  try {
    const imageBlob = await base64ToBlob(imageBase64)
    const formData = new FormData()
    formData.append('file', imageBlob)
    formData.append('filename', imageFilename)

    const response = await backendFetchRequest('images', {
      body: formData,
      method: 'PATCH',
    })

    if (!await response.ok) {
      const errorText = await response.body
      console.error(`Image patch failed: ${errorText}`)
    }
  }
  catch (error) {
    console.error(`Image patch failed: ${error}`)
  }
}

function handleCancel() {
  router.push('/')
}

function toggleDeleting() {
  deleting.value = !deleting.value
}

async function deleteThisRecipe() {
  if (recipe.value.imageFilename && recipe.value.imageFilename.length > 0) {
    await backendFetchRequest(`images/${recipe.value.imageFilename}`, {
      method: 'DELETE',
    })
  }
  if (recipe.value.slug) {
    await backendFetchRequest(`recipes/${recipe.value.slug}`, {
      method: 'DELETE',
    })
  }
  router.push('/')
}

onBeforeMount(async () => {
  await getRecipe()
})
</script>

<template>
  <div>
    <RecipeForm v-if="recipe" :can-delete="true" :recipe="recipe" @cancel="handleCancel" @click-delete="toggleDeleting()" @save="handleSave" />
    <div v-if="deleting" class="flex justify-center gap-4 mx-auto lg:w-2/3 mb-8">
      <button aria-label="cancel" class="w-full textButton" @click="toggleDeleting">
        Cancel
      </button>
      <button aria-label="save" class="textButton w-full" @click="deleteThisRecipe()">
        Delete
      </button>
    </div>
  </div>
</template>
