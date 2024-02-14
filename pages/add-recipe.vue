<script setup lang="ts">
import { ref } from 'vue'

export interface Recipe {
  slug: string
  dateAdded: string
  name: string
  author: string
  source: string
  course: string
  vegetarian: false
  prepTime: string
  cookingTime: string
  calories: string
  servings: string
  ingredients: string[]
  instructions: string[]
  image: string
}
const recipe = ref<Recipe>({} as Recipe)

const ingredients = ref()
const instructions = ref()

const courseOptions = computed(() => {
  const courses: any[] = []
  allowedCourses.forEach((course) => {
    courses.push({
      title: course,
      value: course,
    })
  })
  return courses
})

async function saveRecipe() {
  recipe.value.dateAdded = new Date().toISOString().split('T')[0]
  recipe.value.image = `/recipe-images/${recipe.value.slug}.webp`

  recipe.value.ingredients = ingredients.value.split('\n')
  recipe.value.instructions = instructions.value.split('\n')

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
  <div class="grid gap-4 grid-cols-1 place-items-center mt-4">
    <form class="pl md:pl-5 md:w-1/3 w-3/4" @submit.prevent="saveRecipe">
      <FormInput id="slug" v-model="recipe.slug" label="Slug" type="text" />
      <FormInput id="name" v-model="recipe.name" label="Recipe Name" type="text" />
      <FormInput id="author" v-model="recipe.author" label="Author" type="text" />
      <FormInput id="source" v-model="recipe.source" label="Source" type="text" />
      <FormDropdown v-model="recipe.course" label="Course" :options="courseOptions" class="w-full md:w-1/2" />
      <FormCheckbox id="vegetarian" v-model="recipe.vegetarian" label="Vegetarian?" />
      <FormInput id="prepTime" v-model="recipe.prepTime" label="Prep Time" type="text" />
      <FormInput id="cookingTime" v-model="recipe.cookingTime" label="Cooking Time" type="text" />
      <FormInput id="calories" v-model="recipe.calories" label="Total Calories" type="number" />
      <FormInput id="servings" v-model="recipe.servings" label="Servings" type="number" />
      <FormTextarea id="ingredients" v-model="ingredients" label="Ingredients" />
      <FormTextarea id="Ingredients" v-model="instructions" label="Instructions" />

      <div class="pt-8 flex justify-center">
        <button type="submit" class="block w-full px-3 py-3 text-base font-normal text-dark bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 lg:w-1/2">
          Save Recipe
        </button>
      </div>
    </form>
  </div>
</template>
