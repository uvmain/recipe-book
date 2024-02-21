<script setup lang="ts">
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'

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

const countryOptions = computed(() => {
  const countries: any[] = []
  countryFlags.forEach((country) => {
    countries.push({
      title: country.name,
      value: country.name,
    })
  })
  return countries
})

async function saveRecipe() {
  recipe.value.slug = recipe.value.name.toLowerCase().replaceAll(' ', '-')
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
      <FormInput id="name" v-model="recipe.name" label="Recipe Name" type="text" />
      <FormInput id="author" v-model="recipe.author" label="Author" type="text" />
      <FormInput id="source" v-model="recipe.source" label="Source" type="text" />
      <FormDropdown v-model="recipe.course" label="Course" :options="courseOptions" class="w-full md:w-1/2" />
      <FormDropdown v-model="recipe.country" label="Country" :options="countryOptions" class="w-full md:w-1/2" />
      <FormCheckbox id="vegetarian" v-model="recipe.vegetarian" label="Vegetarian?" />
      <FormInput id="prepTime" v-model="recipe.prepTime" label="Prep Time" type="text" />
      <FormInput id="cookingTime" v-model="recipe.cookingTime" label="Cooking Time" type="text" />
      <FormInput id="calories" v-model="recipe.calories" label="Total Calories" type="number" />
      <FormInput id="servings" v-model="recipe.servings" label="Servings" type="number" />

      <label class="text-white m-2 block">
        Ingredients
      </label>
      <MdEditor v-model="ingredients" editor-id="ingredients" class="add-form-component" language="en-US" />

      <label class="m-2 block text-white">
        Instructions
      </label>
      <MdEditor v-model="instructions" class="add-form-component" language="en-US" editor-id="instructions" />

      <div class="flex justify-center pt-8">
        <button type="submit" class="w-full block border rounded px-3 py-3 text-base font-normal text-dark bg-white bg-clip-padding border-solid border-gray-300 transition ease-in-out m-0 lg:w-1/2">
          Download Recipe
        </button>
      </div>
    </form>
  </div>
</template>
