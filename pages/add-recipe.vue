<script setup lang="ts">
const addFilesInput = ref()
const recipe = ref < Recipe >({} as Recipe)
const imageBase64 = ref<string>()

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

const canSave = computed(() => {
  if (recipe.value.name
    && recipe.value.author
    && recipe.value.source
    && recipe.value.ingredients
    && recipe.value.instructions
    && recipe.value.course
  ) {
    return true
  }
  return false
})

const recipeSlug = computed(() => {
  return recipe.value.name.length ? convertToAlphanumeric(recipe.value.name) : ''
})

function handleFileChange() {
  const file = addFilesInput.value.files?.[0]

  if (file) {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = () => {
      if (typeof reader.result === 'string') {
        imageBase64.value = reader.result
        recipe.value.image = reader.result
      }
    }
  }
}

function resetImage() {
  imageBase64.value = ''
  recipe.value.image = ''
}

async function saveRecipeImageAsWebp() {
  if (imageBase64.value) {
    const image = new Image()
    image.src = imageBase64.value

    image.onload = async function () {
      const canvas = document.createElement('canvas')
      const ctx = canvas.getContext('2d')

      // Calculate new width and height while preserving aspect ratio
      let newWidth = image.width
      let newHeight = image.height
      if (image.width > 800 || image.height > 800) {
        if (image.width > image.height) {
          newWidth = 800
          newHeight = Math.round((800 / image.width) * image.height)
        }
        else {
          newHeight = 800
          newWidth = Math.round((800 / image.height) * image.width)
        }
      }

      // Resize the canvas
      canvas.width = newWidth
      canvas.height = newHeight
      ctx?.drawImage(image, 0, 0, newWidth, newHeight)

      // Convert canvas to WebP format with 80% quality
      const webpData = canvas.toDataURL('image/webp', 0.8)

      // Convert base64 to blob
      const blob = await fetch(webpData).then(res => res.blob())

      // Create a download link
      const link = document.createElement('a')
      link.href = URL.createObjectURL(blob)
      const slug = recipeSlug.value || 'recipeImage'
      link.download = `${slug}.webp`
      link.click()
    }
  }
}

async function downloadRecipe() {
  if (!canSave.value) {
    return
  }
  saveRecipeImageAsWebp()

  recipe.value.slug = recipeSlug.value
  recipe.value.dateAdded = new Date().toISOString().split('T')[0]
  recipe.value.image = `/recipe-images/${recipe.value.slug}.webp`
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
  <div class="mt-10">
    <div class="grid gap-4 grid-cols-3">
      <div class="mx-auto text-zinc-800 text-left">
        <p>Recipe name:</p>
        <input id="nameInput" v-model="recipe.name" label="" type="text" class="add-form-component">

        <p>Author:</p>
        <input id="authorInput" v-model="recipe.author" type="text" class="add-form-component">

        <p>Source:</p>
        <input id="sourceInput" v-model="recipe.source" type="text" class="add-form-component">

        <FormDropdown id="courseInput" v-model="recipe.course" label="Course" :options="courseOptions" class="text-left w-full flex gap-4 my-4 flex-auto md:w-1/2 grid-rows-1" />
        <FormDropdown id="countryInput" v-model="recipe.country" label="Country" :options="countryOptions" class="text-left w-full flex flex-auto gap-4 mb-4 md:w-1/2 grid-rows-1" />
        <FormCheckbox id="vegetarianInput" v-model="recipe.vegetarian" label="Vegetarian?" class="text-left w-full md:w-1/2 grid gap-2 grid-rows-1 mb-4 flex flex-auto grid-cols-2" />

        <p>Prep Time:</p>
        <input id="prepTimeInput" v-model="recipe.prepTime" type="text" class="add-form-component">

        <p>Cooking Time:</p>
        <input id="cookingTimeInput" v-model="recipe.cookingTime" type="text" class="add-form-component">

        <p>Calories:</p>
        <input id="caloriesInput" v-model="recipe.calories" type="number" class="add-form-component">

        <p>Servings:</p>
        <input id="servingsInput" v-model="recipe.servings" type="number21" class="add-form-component">

        <div class="flex flex-row items-baseline mt-2">
          Ingredients:
          <div class="px-2 py-1 rounded text-sm m-1 bg-gray-200">
            '** ' = <Icon name="carbon:text-bold" />
          </div>
          <div class="m-1 px-2 py-1 rounded bg-gray-200 text-sm">
            '- ' = <Icon name="tabler:point-filled" />
          </div>
        </div>
        <textarea id="ingredientsInput" v-model="recipe.ingredients" class="add-form-component h-40" />

        <div class="flex flex-row items-baseline mt-2">
          Instructions:
          <div class="m-1 px-2 py-1 rounded bg-gray-200 text-sm">
            '** ' = <Icon name="carbon:text-bold" />
          </div>
          <div class="m-1 px-2 py-1 rounded bg-gray-200 text-sm">
            '- ' = <Icon name="tabler:point-filled" />
          </div>
        </div>
        <textarea id="instructionsInput" v-model="recipe.instructions" class="add-form-component h-40" />

        <p>Recipe Image:</p>
        <div class="flex flex-row gap-2">
          <input
            id="addFiles"
            ref="addFilesInput"
            type="file"
            accept="image/*"
            multiple="false"
            class="add-form-component"
            @change="handleFileChange"
          >
          <button
            class="text-white rounded-md px-4 focus:outline-none py-2 bg-gray-500 hover:bg-blue-600 focus:bg-blue-600 text-3xl"
            @click="resetImage"
          >
            <Icon name="carbon:reset" />
          </button>
        </div>
        <div>
          <div class="flex gap-4 ml-4 mt-8">
            <button
              class="text-white rounded-md px-4 focus:outline-none py-2 bg-gray-500 hover:bg-blue-600 focus:bg-blue-600 text-3xl"
              :class="{ 'bg-red hover:bg-red focus:bg-red': !(canSave) }"
              @click="downloadRecipe"
            >
              Download Recipe
            </button>
          </div>
        </div>
      </div>
      <Recipe :recipe="recipe" class="col-span-2" />
    </div>
  </div>
  <FloatingScrollToTop />
</template>
