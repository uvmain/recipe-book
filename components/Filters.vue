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
  <div class="flex flex-col w-full md:w-2/3 lg:w-1/2 mx-auto p-2 md:p-4 lg:p-6">
    <div class="mx-8 grid grid-cols-2 gap-1">
      <div v-for="(course, index) in allowedCourses" :key="index" class="flex gap-1">
        <input
          :id="course"
          v-model="selectedCourses"
          type="checkbox"
          :value="course"
          class="accent-gray-600"
        >
        <label :for="course" class="font-light">
          {{ course }}
        </label>
      </div>
    </div>

    <hr class="w-80% my-2 border-gray-300 border-1 border-solid">

    <div class="mx-8 grid grid-cols-2">
      <div class="flex flex-row gap-x-1 items-center">
        <input
          id="vegetarian"
          v-model="selectedVegetarian"
          type="checkbox"
          class="accent-gray-600"
        >
        <Icon name="lucide:leafy-green" class="text-lg text-green-600"/>
        <label for="vegetarian" class="font-light">
          vegetarian
        </label>
      </div>
    </div>

    <hr class="w-80% my-2 border-gray-300 border-1 border-solid">

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
        class="w-full accent-gray-600"
      >
    </div>

    <hr class="w-80% my-2 border-gray-300 border-1 border-solid">

    <div class="mx-8  flex flex-col gap-1">
      <label for="country" class="flex justify-between">
        <div class="flex flex-row gap-2 items-center">
          Country: <Icon v-if="selectedCountry" :name="countryFlags.filter((country) => country.name == selectedCountry)[0].icon" class="text-lg" />
        </div>
        <button
          type="button"
          class="justify-center items-center size-8 bg-white rounded-lg border border-solid border-gray-200 active:bg-blue-gray-300 flex items-center"
          @click="selectedCountry = undefined"
        >
          <Icon name="lucide:eraser" class="text-lg text-gray-600" />
        </button>
      </label>
      <select id="country" v-model="selectedCountry" class="bg-white rounded-lg h-8 border-solid border border-gray-200 text-base text-gray-600">
        <option disabled value="">Please select one</option>
        <option v-for="(country, index) in countryFlags" :key=index>
          {{ country.name }}
        </option>
      </select>
    </div>

    <hr class="w-80% mt-2 mb-4 border-gray-300 border-1 border-solid">

    <div class="flex flex-row gap-2 mx-auto">
      <button
        type="button"
        class="h-12 bg-white rounded-lg overflow-hidden shadow-md hover:shadow-lg transition-shadow duration-300 border-1 border-solid border-gray-200 active:bg-blue-gray-300 flex flex-row items-center gap-x-2 text-lg px-2"
        @click="resetFilters()"
      >
        <Icon name="lucide:eraser" class="text-2xl align-middle text-center text-gray-800" />
        <p>Reset filters</p>
      </button>

      <button
        type="button"
        class="h-12 bg-white rounded-lg overflow-hidden shadow-md hover:shadow-lg transition-shadow duration-300 border-1 border-solid border-gray-200 active:bg-blue-gray-300 flex flex-row items-center gap-x-2 text-lg px-2"
        @click="closeFilters()"
      >
        <Icon name="lucide:save" class="text-2xl align-middle text-center text-gray-800" />
        <p>Close filters</p>
      </button>
    </div>
  </div>
</template>