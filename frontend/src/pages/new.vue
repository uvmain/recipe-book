<script setup lang="ts">
import type { Recipe } from '../types/recipes'
import { useHead } from '@vueuse/head'
import { checkIfLoggedIn } from '../composables/auth'
import { backendFetchRequest } from '../composables/fetchFromBackend'

const router = useRouter()
const recipe = ref<Recipe>({} as Recipe)

async function cancel() {
  router.push('/')
}

async function save() {
  recipe.value.slug = recipe.value.slug || ''
  recipe.value.date_created = new Date().toISOString()
  recipe.value.name = recipe.value.name || ''

  const response = await backendFetchRequest('recipes', {
    body: JSON.stringify(recipe.value),
    method: 'POST',
  })
  const jsonData = await response.json()
  console.log(jsonData)
}

watch(
  () => recipe.value.name,
  (newName) => {
    recipe.value.slug = newName
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
  <div class="background text min-h-screen flex flex-col gap-4 mx-auto lg:max-w-18/19">
    {{ recipe }}
    <div class="">
      <TextInput v-model="recipe.slug" class="py-2 text-xl" label="slug" />
      <TextInput v-model="recipe.name" class="py-2 text-xl" label="name" />
      <div class="flex justify-center gap-4 mt-4">
        <button aria-label="cancel" class="headerButton w-full" @click="cancel()">
          Cancel
        </button>
        <button aria-label="save" class="headerButton w-full" @click="save()">
          Save
        </button>
      </div>
    </div>
  </div>
</template>
