<script setup lang="ts">
import { ref } from 'vue'

const recipe = ref({
  slug: '',
  dateAdded: '',
  name: '',
  author: '',
  source: '',
  course: '',
  vegetarian: false,
  prepTime: '',
  cookingTime: '',
  calories: 0,
  servings: 0,
  ingredients: [],
  instructions: [],
  image: ''
})

async function saveRecipe() {
  recipe.value.dateAdded = new Date().toISOString().split("T")[0]
  const jsonString = JSON.stringify(recipe.value, null, 2)
  const blob = new Blob([jsonString], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `${recipe.value.slug}.json`
  document.body.appendChild(a)
  a.click()
  window.URL.revokeObjectURL(url)
}
</script>

<template>
  <div class="flex w-full">
    <form @submit.prevent="saveRecipe" class="grid grid-cols-1 w-full md:w-2/3 justify-center">
      <FormInput id="slug" v-model="recipe.slug" label="Slug" type="text" />
      <FormInput id="name" v-model="recipe.name" label="Recipe Name" type="text" />
      <FormInput id="author" v-model="recipe.author" label="Author" type="text" />
      <FormInput id="source" v-model="recipe.source" label="Source" type="text" />
      
      <button type="submit">Save Recipe</button>
    </form>
  </div>
</template>

