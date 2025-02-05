<script setup lang="ts">
import type { Recipe } from '../types/recipes'
import { ref } from 'vue'
import RecipeForm from '../components/RecipeEditor.vue'

const router = useRouter()
const initialRecipe = ref<Recipe>({
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
})

function handleCancel() {
  router.push('/')
}

function handleSave(recipe: Recipe) {
  router.push(`/recipe/${recipe.slug}`)
}
</script>

<template>
  <RecipeForm :recipe="initialRecipe" @cancel="handleCancel" @save="handleSave" />
</template>
