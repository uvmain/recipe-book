<script setup lang="ts">
const showFilters = useState('showFilters') || false
const selectedCourses = useState<string[]>('selectedCourses', () => [])
const selectedVegetarian = useState<boolean>('selectedVegetarian')
const selectedCalories = useState<number>('selectedCalories', () => 1000)
const selectedCountry = useState('selectedCountry')

function closeFilters() {
  showFilters.value = false
}
</script>

<template>
  <div class="flex flex-col w-full md:w-2/3 lg:w-1/2 mx-auto p-2 md:p-4 lg:p-6 text">
    <div class="mx-8 grid grid-cols-2 gap-1">
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
    </div>

    <hr class="filterHr">

    <div class="mx-8 grid grid-cols-2">
      <div class="flex flex-row gap-x-1 items-center">
        <input
          id="vegetarian"
          v-model="selectedVegetarian"
          type="checkbox"
          class="accent-gray-600 dark:accent-gray-400"
        >
        <Icon name="lucide:leafy-green" class="text-lg text-green-600"/>
        <label for="vegetarian" class="font-light">
          vegetarian
        </label>
      </div>
    </div>

    <hr class="filterHr">

    <div class="mx-8 flex flex-col gap-1">
      <label for="calories">
        Calories per serving: {{ selectedCalories == 1000 ? "Unlimited" : selectedCalories }}
      </label>
      <input
        id="calories"
        v-model="selectedCalories"
        type="range"
        min="100"
        max="1000"
        class="w-full accent-gray-400 dark:accent-gray-400"
      >
    </div>

    <hr class="filterHr">

    <div class="mx-8  flex flex-col gap-1">
      <label for="country" class="flex justify-between">
        <div class="flex flex-row gap-2 items-center">
          Country: <Icon v-if="selectedCountry" :name="countryFlags.filter((country) => country.name == selectedCountry)[0].icon" class="text-lg" />
        </div>
        <button
          type="button"
          class="justify-center items-center size-8 recipeCardBackground border border-solid border-gray-200 active:bg-blue-gray-300 flex items-center"
          @click="selectedCountry = undefined"
        >
          <Icon name="lucide:eraser" class="text-lg headerButtonIcon" />
        </button>
      </label>
      <select id="country" v-model="selectedCountry" class="h-8 border text-base text recipeCardBackground">
        <option disabled value="">Please select one</option>
        <option v-for="(country, index) in countryFlags" :key=index>
          {{ country.name }}
        </option>
      </select>
    </div>

    <hr class="filterHr">

    <div class="flex flex-row gap-2 mx-auto">
      <button
        type="button"
        class="headerButton w-auto flex flex-row gap-2 items-center p-4"
        @click="resetFilters()"
      >
        <Icon name="lucide:eraser" class="headerButtonIcon" />
        <p class="text text-lg">Reset filters</p>
      </button>

      <button
        type="button"
        class="headerButton w-auto flex flex-row gap-2 items-center p-4"
        @click="closeFilters()"
      >
        <Icon name="lucide:save" class="headerButtonIcon" />
        <p class="text text-lg">Close filters</p>
      </button>
    </div>
  </div>
</template>