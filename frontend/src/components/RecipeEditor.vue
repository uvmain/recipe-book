<script setup lang="ts">
import type { Recipe } from '../types/recipes'
import { useHead } from '@vueuse/head'
import { checkIfLoggedIn } from '../composables/auth'
import { countries } from '../composables/countries'
import { slugify } from '../composables/slugify'
import { allowedCourses } from '../types/recipes'
import Dropdown from './Dropdown.vue'
import ImageSelector from './ImageSelector.vue'
import Vulgars from './Vulgars.vue'

const props = defineProps<{
  recipe: Recipe
  canDelete: boolean
}>()

const emit = defineEmits(['cancel', 'clickDelete', 'save'])

const localRecipe = ref<Recipe>({ ...props.recipe })
const resizedImage = ref<string>('')

watch(
  () => props.recipe,
  (newRecipe) => {
    localRecipe.value = { ...newRecipe }
  },
  { deep: true },
)

watch(
  () => localRecipe.value.name,
  (newName) => {
    localRecipe.value.slug = slugify(newName)
  },
)

function handleDelete() {
  if (localRecipe.value.slug && localRecipe.value.slug.length > 0) {
    emit('clickDelete')
  }
}

function cancel() {
  emit('cancel')
}

function save() {
  if (localRecipe.value.slug && localRecipe.value.slug.length > 0) {
    emit('save', localRecipe.value, resizedImage.value)
  }
}

useHead({
  titleTemplate: 'RecipeBook: New',
})

onBeforeMount(async () => {
  const loggedIn = await checkIfLoggedIn()
  if (!loggedIn) {
    emit('cancel')
  }
})
</script>

<template>
  <div class="background text min-h-screen flex flex-col gap-4 mx-auto lg:max-w-200 min-w-1/3 my-8">
    <TextInput v-model="localRecipe.slug" type="text" label="Slug" />
    <TextInput v-model="localRecipe.name" type="text" label="Name" />
    <TextInput v-model="localRecipe.author" type="text" label="Author" />
    <TextInput v-model="localRecipe.source" type="text" label="Source" />
    <Dropdown v-model="localRecipe.course" type="select" label="Course" :options="allowedCourses" placeholder="Select a course" />
    <Dropdown v-model="localRecipe.country" type="select" label="Country" :options="countries" placeholder="Select a country" />
    <Checkbox v-model="localRecipe.vegetarian" label="Vegetarian?" name="vegetarian" />
    <TextInput v-model="localRecipe.prepTime" type="text" label="PrepTime" />
    <TextInput v-model="localRecipe.cookingTime" type="text" label="CookingTime" />
    <TextInput v-model="localRecipe.calories" type="text" label="Calories" />
    <TextInput v-model="localRecipe.servings" type="text" label="Servings" />
    <TextInput v-model="localRecipe.ingredients" type="textarea" label="Ingredients" />
    <Vulgars />
    <TextInput v-model="localRecipe.instructions" type="textarea" label="Instructions" />
    <ImageSelector v-model:resized-image="resizedImage" v-model:resized-image-width="localRecipe.imageWidth" v-model:resized-image-height="localRecipe.imageHeight" label="Image" :recipe="recipe" />

    <div class="flex justify-center gap-4">
      <button aria-label="delete" class="w-full textButton" @click="handleDelete()">
        Delete
      </button>
      <button aria-label="cancel" class="w-full textButton" @click="cancel()">
        Cancel
      </button>
      <button aria-label="save" class="textButton w-full" @click="save()">
        Save
      </button>
    </div>
  </div>
</template>
