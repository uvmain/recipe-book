<script setup lang="ts">
import { useSessionStorage } from '@vueuse/core'
import { countries } from '../composables/countries'
import { resetFilters } from '../composables/resets'
import { allowedCourses } from '../types/recipes'

const showFilters = useSessionStorage<boolean>('showFilters', false)
const selectedCourses = useSessionStorage<string[]>('selectedCourses', [])
const selectedVegetarian = useSessionStorage<boolean>('selectedVegetarian', false)
const selectedCalories = useSessionStorage<number>('selectedCalories', 1000)
const selectedCountry = useSessionStorage<string>('selectedCountry', '')
const filtered = useSessionStorage<boolean>('filtered', false)

const computedFiltered = computed(() => {
  return selectedCourses.value.length > 0 || selectedVegetarian.value === true || selectedCalories.value < 1000 || selectedCountry.value !== undefined
})

watch(computedFiltered, () => {
  filtered.value = computedFiltered.value
})

function closeFilters() {
  showFilters.value = false
}
</script>

<template>
  <div class="flex flex-col mx-auto p-2 md:p-4 lg:p-6 text">
    <div class="mx-auto grid grid-cols-2 lg:grid-cols-3 gap-y-1 gap-x-3">
      <div v-for="(course, index) in allowedCourses" :key="index" class="flex gap-1">
        <input
          :id="course"
          v-model="selectedCourses"
          type="checkbox"
          :value="course"
          class="accent-gray-600 dark:accent-gray-400"
        >
        <label :for="course" class="font-light">
          {{ course }}
        </label>
      </div>
      <div class="flex flex-row gap-x-1 items-center">
        <input
          id="vegetarian"
          v-model="selectedVegetarian"
          type="checkbox"
          class="accent-gray-600 dark:accent-gray-400"
        >
        <icon-lucide-leafy-green class="text-lg text-green-600" />
        <label for="vegetarian" class="font-light">
          vegetarian
        </label>
      </div>
    </div>

    <hr class="filterHr">

    <div class="mx-auto flex flex-col gap-1">
      <label for="calories">
        Calories per serving: {{ selectedCalories === 1000 ? "Unlimited" : selectedCalories }}
      </label>
      <input
        id="calories"
        v-model="selectedCalories"
        type="range"
        min="100"
        max="1000"
        class="dark:accent-gray-400 w-full accent-gray-400"
      >
    </div>

    <hr class="filterHr">

    <div class="flex flex-col gap-1 mx-8">
      <label for="country" class="flex justify-between">
        <div class="flex flex-row items-center gap-2">
          Country: <Flag :country="selectedCountry" class="text-lg" />
        </div>
        <button
          type="button"
          aria-label="select country"
          class="items-center flex items-center justify-center size-8 border-solid border-gray-200 active:bg-blue-gray-300 recipeCardBackground border"
          @click="selectedCountry = ''"
        >
          <icon-lucide-eraser class="text-lg headerButtonIcon" />
        </button>
      </label>
      <select id="country" v-model="selectedCountry" class="border text recipeCardBackground h-8 text-base">
        <option disabled value="">
          Please select one
        </option>
        <option v-for="(country, index) in countries" :key="index">
          {{ country }}
        </option>
      </select>
    </div>

    <hr class="filterHr">

    <div class="flex flex-row gap-2 mx-auto">
      <button
        type="button"
        aria-label="reset filters"
        class="flex flex-row gap-2 items-center w-auto p-4 headerButton"
        @click="resetFilters()"
      >
        <icon-lucide-eraser class="headerButtonIcon" />
        <p class="text text-lg">
          Reset filters
        </p>
      </button>

      <button
        type="button"
        aria-label="close filters"
        class="headerButton w-auto flex flex-row gap-2 items-center p-4"
        @click="closeFilters()"
      >
        <icon-lucide-save class="headerButtonIcon" />
        <p class="text text-lg">
          Close filters
        </p>
      </button>
    </div>
  </div>
</template>
